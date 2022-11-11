package davinci

import (
	"context"
	// "fmt"
	// "log"
	"strconv"
	"time"

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
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"connector_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"company_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"customer_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_date": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"properties": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "Connection configuration",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"value": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"type": {
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"environment_id": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func dataSourceConnectionsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	resp, err := c.ReadConnections(&c.CompanyID, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	conns := make([]interface{}, len(resp), len(resp))
	for i, connItem := range resp {
		conn := make(map[string]interface{})
		conn = map[string]interface{}{
			"connection_id": connItem.ConnectionID,
			"connector_id":  connItem.ConnectorID,
			"name":          connItem.Name,
			"created_date":  connItem.CreatedDate,
			"company_id":    connItem.CompanyID,
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
			conn["properties"] = connProps
		}
		conns[i] = conn
	}

	if err := d.Set("connections", conns); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("environment_id", c.CompanyID); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}
