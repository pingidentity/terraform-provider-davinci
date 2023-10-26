package davinci

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

func DataSourceConnections() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConnectionsRead,
		Schema: map[string]*schema.Schema{
			"connections": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Returned set of connections matching environment and/or the filter criteria.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "connection_id for this connection.",
						},
						"connector_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "DaVinci internal connector type. Only found via API read response (e.g Http Connector is 'httpConnector')",
						},
						"company_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "PingOne environment id. Matches environment_id and will be deprecated in the future.",
						},
						"customer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Internal DaVinci id. Should not be set by user.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the connection displayed in UI. Also used for mapping id on flows between environments.",
						},
						"created_date": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Resource creation date as epoch.",
						},
						"property": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Connection properties. These are specific to the connector type. Get connection properties from connection API read response.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of the property.",
									},
									"value": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Value of the property as string. If the property is an array, use a comma separated string.",
									},
									"type": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Type of the property. This is used to cast the value to the correct type. Must be: string or boolean. Use 'string' for array",
									},
								},
							},
						},
					},
				},
			},
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PingOne environment id",
			},
			"connector_ids": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filters list of returned connections",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceConnectionsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	//TODO: clean this up
	var cIdsFilter []string
	if cIds, ok := d.GetOk("connector_ids"); ok {
		cIdsInterface := cIds.([]interface{})
		for _, v := range cIdsInterface {
			cIdsFilter = append(cIdsFilter, v.(string))
		}
	}

	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.ReadConnections(&c.CompanyID, nil)
	}, nil)

	if err != nil {
		return diag.FromErr(err)
	}

	res, ok := sdkRes.([]dv.Connection)
	if !ok {
		err = fmt.Errorf("Unable to parse connections response from Davinci API")
		return diag.FromErr(err)
	}
	resp := []dv.Connection{}
	if cIdsFilter != nil {
		for _, resItem := range res {
			cIdFound := contains(cIdsFilter, resItem.ConnectorID)
			if !cIdFound {
				continue
			}
			resp = append(resp, resItem)
		}
	} else {
		resp = res
	}
	conns := make([]interface{}, len(resp))
	for i, connItem := range resp {
		conn := map[string]interface{}{
			"id":           connItem.ConnectionID,
			"connector_id": connItem.ConnectorID,
			"name":         connItem.Name,
			"created_date": connItem.CreatedDate,
			"company_id":   connItem.CompanyID,
		}
		if connItem.Properties != nil {
			connProps := []map[string]interface{}{}
			for propi, propv := range connItem.Properties {
				pMap := propv.(map[string]interface{})
				if pMap == nil {
					return diag.Errorf("Unable to assert Property to map interface")
				}
				thisProp := map[string]interface{}{
					"name":  propi,
					"value": "",
				}
				// In some properties, if the value is blank in the UI, the value is not returned in the API response
				if _, ok := pMap["value"]; !ok {
					continue
				}
				if pType, ok := pMap["type"].(string); ok {
					thisProp["type"] = pType
					switch pType {
					case "string", "":
						if _, ok := pMap["value"].(string); ok {
							thisProp["value"] = pMap["value"].(string)
						}
					case "boolean":
						if pValue, ok := pMap["value"].(bool); ok {
							thisProp["value"] = strconv.FormatBool(pValue)
						}
					default:
						return diag.Errorf("For Connection '%v' and Property '%v': unable to identify value type, only string or boolean is currently supported", connItem.Name, thisProp["name"])
					}
				} else {
					switch pMap["value"].(type) {
					case string:
						if _, ok := pMap["value"].(string); ok {
							thisProp["value"] = pMap["value"].(string)
						}
					case bool:
						if pValue, ok := pMap["value"].(bool); ok {
							thisProp["value"] = strconv.FormatBool(pValue)
						}
					default:
						return diag.Errorf("For Connection '%v' and Property '%v': unable to identify value type, only string or boolean is currently supported", connItem.Name, thisProp["name"])
					}
				}
				connProps = append(connProps, thisProp)
			}
			conn["property"] = connProps
		}
		conns[i] = conn
	}

	if err := d.Set("connections", conns); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("environment_id", c.CompanyID); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(fmt.Sprintf("id-%s-connections", c.CompanyID))
	return diags
}

func contains(s []string, str string) bool {
	for _, a := range s {
		if a == str {
			return true
		}
	}
	return false
}
