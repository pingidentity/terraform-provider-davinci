package davinci

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

func DataSourceConnection() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConnectionRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "A string that specifies the ID of the connection to retrieve.  This field is deprecated for retrieving the connection and will be made read only in a future release, use `connection_id` instead.",
				ExactlyOneOf: []string{"id", "connection_id", "name"},
			},
			"connection_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "A string that specifies the ID of the connection to retrieve. Either `connection_id` or `name` must be specified.",
				ExactlyOneOf: []string{"id", "connection_id", "name"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "A string that specifies the name of the connection to retrieve. Either `id` or `name` must be specified.",
				ExactlyOneOf: []string{"id", "connection_id", "name"},
			},
			"connector_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The DaVinci internal connector type ID, which can be found in the [DaVinci Connection Definitions](../../resources/connection#davinci-connection-definitions) documentation.",
			},
			"environment_id": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The ID of the PingOne environment to retrieve a connection from. Must be a valid PingOne resource ID.",
				ValidateDiagFunc: validation.ToDiagFunc(verify.ValidP1ResourceID),
			},
			"customer_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An ID that represents the customer tenant.",
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
							Sensitive:   true,
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
		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(20 * time.Minute),
		},
	}
}

func dataSourceConnectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	var connId *string
	var err error

	environmentID := d.Get("environment_id").(string)

	// Make sure we have all connections propagated
	_, err = readAllConnections(ctx, c, environmentID, d.Timeout(schema.TimeoutRead))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	// prep case where name is provided
	connName, ok := d.GetOk("name")
	if ok {
		connId, err = getConnectionIdByName(ctx, c, environmentID, connName.(string))
		if err != nil {
			err = fmt.Errorf("Connection not found")
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}
	}

	// prep case where id is provided
	if v, ok := d.GetOk("id"); ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Use of `id` to select a connection is deprecated, please use `connection_id` instead",
			Detail:   "The use of `id` is deprecated and will be made a computed attribute in a future release, please use `connection_id` instead",
		})
		value := v.(string)
		connId = &value
	}

	// prep case where id is provided
	if v, ok := d.GetOk("connection_id"); ok {
		value := v.(string)
		connId = &value
	}

	// get connection by id and parse response
	sdkRes, err := sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.ReadConnectionWithResponse(environmentID, *connId)
		},
	)

	if err != nil {
		if dvError, ok := err.(dv.ErrorResponse); ok {
			if dvError.HttpResponseCode == http.StatusNotFound || dvError.Code == dv.DV_ERROR_CODE_CONNECTION_NOT_FOUND {
				err = fmt.Errorf("Connection not found")
				diags = append(diags, diag.FromErr(err)...)
				return diags
			}
		}

		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	res, ok := sdkRes.(*dv.Connection)
	if !ok {
		err = fmt.Errorf("Unable to parse response from Davinci API on connection id: %v", connId)
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId(*res.ConnectionID)

	if err := d.Set("name", res.Name); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	if err := d.Set("connector_id", res.ConnectorID); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	if err := d.Set("created_date", res.CreatedDate.UnixMilli()); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	if err := d.Set("environment_id", res.CompanyID); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	if err := d.Set("customer_id", res.CustomerID); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	props, err := flattenConnectionProperties(res.Properties)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	if err := d.Set("property", props); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	return diags
}

func getConnectionIdByName(ctx context.Context, c *dv.APIClient, environmentID, connName string) (*string, error) {

	sdkRes, err := sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.ReadConnectionsWithResponse(environmentID, nil)
		},
	)
	if err != nil {
		return nil, err
	}

	res, ok := sdkRes.([]dv.Connection)
	if !ok {
		return nil, fmt.Errorf("Unable to parse response from Davinci API on connection name: %v", connName)
	}

	for _, conn := range res {
		if *conn.Name == connName {
			return conn.ConnectionID, nil
		}
	}
	return nil, fmt.Errorf("Unable to find connection with name: %v", connName)
}
