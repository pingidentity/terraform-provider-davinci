package davinci

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

func DataSourceConnections() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConnectionsRead,
		Schema: map[string]*schema.Schema{
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the PingOne environment to retrieve connections for. Must be a valid PingOne resource ID.",

				ValidateDiagFunc: validation.ToDiagFunc(verify.ValidP1ResourceID),
			},
			"connector_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "A list of connector IDs to filter from the returned connections.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"connections": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "The returned set of connections matching the environment and the optional filter criteria.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID for this connection, otherwise known as the \"Connection ID\".",
						},
						"connector_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The DaVinci internal connector type ID, which can be found in the [DaVinci Connection Definitions](../../resources/connection#davinci-connection-definitions) documentation.",
						},
						"company_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Deprecated:  "This attribute is deprecated and will be removed in the next major release.",
							Description: "**Deprecation Notice** This attribute is deprecated and will be removed in the next major release.  The PingOne environment ID.",
						},
						"customer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An ID that represents the customer tenant.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the connection displayed in UI. Also used for mapping id on flows between environments.",
						},
						"created_date": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Resource creation date as epoch timestamp.",
						},
						"property": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Connection properties. These are specific to the connector type configured in `connector_id`. See the [DaVinci Connection Definitions](#davinci-connection-definitions) document to find the appropriate property name/value pairs for the connection.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the property.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The value of the property as string. If the property is an array, the value will be a comma separated string.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Type of the property. This is used to cast the value to the correct type. Will be either: `string` or `boolean`. `string` is used for array types.",
									},
								},
							},
						},
					},
				},
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(20 * time.Minute),
		},
	}
}

func dataSourceConnectionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	var cIdsFilter []string
	if cIds, ok := d.GetOk("connector_ids"); ok {
		cIdsInterface := cIds.([]interface{})
		for _, v := range cIdsInterface {
			cIdsFilter = append(cIdsFilter, v.(string))
		}
	}

	environmentID := d.Get("environment_id").(string)

	res, err := readAllConnections(ctx, c, environmentID, d.Timeout(schema.TimeoutRead))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	resp := []dv.Connection{}
	if cIdsFilter != nil {
		for _, resItem := range res {
			cIdFound := slices.Contains(cIdsFilter, resItem.ConnectorID)
			if !cIdFound {
				continue
			}
			resp = append(resp, resItem)
		}
	} else {
		resp = res
	}
	conns := make([]map[string]interface{}, 0)
	for _, connItem := range resp {
		conn := map[string]interface{}{
			"id":           connItem.ConnectionID,
			"connector_id": connItem.ConnectorID,
			"name":         connItem.Name,
			"created_date": connItem.CreatedDate,
			"company_id":   connItem.CompanyID,
			"customer_id":  connItem.CustomerID,
		}
		if v := connItem.Properties; v != nil {
			props, err := flattenConnectionProperties(&v)
			if err != nil {
				diags = append(diags, diag.FromErr(err)...)
				return diags
			}

			conn["property"] = props
		}
		conns = append(conns, conn)
	}

	if err := d.Set("connections", conns); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	if err := d.Set("environment_id", c.CompanyID); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId(fmt.Sprintf("id-%s-connections", c.CompanyID))
	return diags
}

func readAllConnections(ctx context.Context, c *dv.APIClient, environmentID string, timeout time.Duration) ([]dv.Connection, error) {

	connections := -1

	stateConf := &retry.StateChangeConf{
		Pending: []string{
			"false",
		},
		Target: []string{
			"true",
			"err",
		},
		Refresh: func() (interface{}, string, error) {

			// Run the API call
			sdkRes, err := sdk.DoRetryable(
				ctx,
				c,
				environmentID,
				func() (interface{}, *http.Response, error) {
					return c.ReadConnectionsWithResponse(&environmentID, nil)
				},
			)

			if err != nil {
				return nil, "err", err
			}

			res, ok := sdkRes.([]dv.Connection)
			if !ok {
				err = fmt.Errorf("Unable to parse connections response from Davinci API")
				return nil, "err", err
			}

			// If the number of connections has changed since last time, we need to keep waiting
			if len(res) != connections {
				connections = len(res)
				return res, "false", nil
			}

			return res, "true", nil
		},
		Timeout:                   timeout - time.Minute,
		Delay:                     2 * time.Second,
		MinTimeout:                2 * time.Second,
		ContinuousTargetOccurence: 5, // we want five consecutive successful reads of the same number of connections
	}
	sdkRes, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return nil, err
	}

	res, ok := sdkRes.([]dv.Connection)
	if !ok {
		err = fmt.Errorf("Unable to parse connections response from Davinci API")
		return nil, err
	}

	return res, nil

}
