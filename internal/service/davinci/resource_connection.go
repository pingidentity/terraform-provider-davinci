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

func ResourceConnection() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConnectionCreate,
		ReadContext:   resourceConnectionRead,
		UpdateContext: resourceConnectionUpdate,
		DeleteContext: resourceConnectionDelete,
		Schema: map[string]*schema.Schema{
			"connection_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DaVinci generated identifier for the connection.",
			},
			"connector_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "DaVinci internal connector type. Not found in UI. Look in API read response (e.g Http Connector is 'httpConnector'",
			},
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PingOne environment id",
			},
			"customer_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the connection displayed in UI. Also used for mapping connection_id on flows between environments.",
			},
			"created_date": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"properties": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Connection properties. These are specific to the connector type. Get connection properties from connection API read response.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Name of the property.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Value of the property as string. If the property is an array, use a comma separated string.",
						},
						"type": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Type of the property. This is used to cast the value to the correct type. Must be: string, boolean. Use 'string' for array",
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceConnectionCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	connection := dv.Connection{
		ConnectorID: d.Get("connector_id").(string),
		Name:        d.Get("name").(string),
	}

	connection.Properties = *makeProperties(d)

	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.CreateInitializedConnection(&c.CompanyID, &connection)
	}, nil)

	if err != nil {
		return diag.FromErr(err)
	}

	resp, ok := sdkRes.(*dv.Connection)
	if !ok || resp.Name == "" {
		err = fmt.Errorf("failed to cast created response to Connection")
		return diag.FromErr(err)
	}

	d.SetId(resp.ConnectionID)

	resourceConnectionRead(ctx, d, m)

	return diags
}

func resourceConnectionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	connId := d.Id()

	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.ReadConnection(&c.CompanyID, connId)
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	res, ok := sdkRes.(*dv.Connection)
	if !ok {
		err = fmt.Errorf("Unable to cast Connection type to response from Davinci API on connection id: %v", connId)
		return diag.FromErr(err)
	}

	if err := d.Set("name", res.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("connection_id", res.ConnectionID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("connector_id", res.ConnectorID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_date", res.CreatedDate); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("environment_id", res.CompanyID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("customer_id", res.CustomerID); err != nil {
		return diag.FromErr(err)
	}
	props, err := flattenConnectionProperties(&res.Properties)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("properties", props); err != nil {
		return diag.FromErr(err)
	}
	return diags
}

func resourceConnectionUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}
	connId := d.Id()
	if d.HasChanges("properties", "name") {
		connection := dv.Connection{
			ConnectorID:  d.Get("connector_id").(string),
			Name:         d.Get("name").(string),
			ConnectionID: connId,
		}

		connection.Properties = *makeProperties(d)

		sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
			return c.UpdateConnection(&c.CompanyID, &connection)
		}, nil)
		if err != nil {
			return diag.FromErr(err)
		}

		res, ok := sdkRes.(*dv.Connection)
		if !ok || res.Name == "" {
			err = fmt.Errorf("Unable to parse update response from Davinci API on connection id: %v", connId)
			return diag.FromErr(err)
		}
	}

	return resourceConnectionRead(ctx, d, m)
}

func resourceConnectionDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	connId := d.Id()

	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.DeleteConnection(&c.CompanyID, connId)
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	res, ok := sdkRes.(*dv.Message)
	if !ok || res.Message == "" {
		err = fmt.Errorf("Unable to parse update response from Davinci API on connection id: %v", connId)
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func flattenConnectionProperties(connectionProperties *dv.Properties) ([]map[string]interface{}, error) {
	if connectionProperties == nil {
		return nil, fmt.Errorf("no properties")
	}
	connProps := []map[string]interface{}{}
	for propName, propVal := range *connectionProperties {
		pMap := propVal.(map[string]interface{})
		if pMap == nil {
			return nil, fmt.Errorf("Unable to assert property values for %v\n", propName)
		}

		if _, ok := pMap["value"]; !ok {
			continue
		}

		thisProp := map[string]interface{}{
			"name":  propName,
			"value": "",
		}

		if propType, ok := pMap["type"].(string); ok {
			thisProp["type"] = propType
			switch propType {
			case "string", "":
				if _, ok := pMap["value"].(string); ok {
					thisProp["value"] = pMap["value"].(string)
				}
			case "boolean":
				if pValue, ok := pMap["value"].(bool); ok {
					thisProp["value"] = strconv.FormatBool(pValue)
				}
			default:
				return nil, fmt.Errorf("For Property '%v': unable to identify value type, only string or boolean is currently supported", thisProp["name"])
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
				return nil, fmt.Errorf("For Property '%v': unable to identify value type, only string or boolean is currently supported", thisProp["name"])
			}
		}
		connProps = append(connProps, thisProp)
	}
	return connProps, nil
}

func makeProperties(d *schema.ResourceData) *dv.Properties {
	connProps := dv.Properties{}
	props := d.Get("properties").(*schema.Set).List()
	for _, raw := range props {
		prop := raw.(map[string]interface{})
		connProps[prop["name"].(string)] = map[string]interface{}{
			"value": prop["value"].(string),
		}
	}
	return &connProps
}

func getCompanyId(c *dv.Client, d *schema.ResourceData) {
	if cId, ok := d.GetOk("environment_id"); ok {
		c.CompanyID = cId.(string)
	}
}
