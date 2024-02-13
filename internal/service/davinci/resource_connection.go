package davinci

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pingidentity/terraform-provider-davinci/internal/framework"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

func ResourceConnection() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceConnectionCreate,
		ReadContext:   resourceConnectionRead,
		UpdateContext: resourceConnectionUpdate,
		DeleteContext: resourceConnectionDelete,
		Schema: map[string]*schema.Schema{
			"connector_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The DaVinci connector type identifier. See the [DaVinci Connection Definitions](#davinci-connection-definitions) below to find the appropriate connector ID value. This field is immutable and will trigger a replace plan if changed.",
				ForceNew:    true,
			},
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the PingOne environment to create the DaVinci connection. Must be a valid PingOne resource ID. This field is immutable and will trigger a replace plan if changed.",

				ValidateDiagFunc: validation.ToDiagFunc(verify.ValidP1ResourceID),
				ForceNew:         true,
			},
			"customer_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An ID that represents the customer tenant.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the connection displayed in UI. Also used for mapping id on flows between environments. This field is immutable and will trigger a replace plan if changed.",
				ForceNew:    true,
			},
			"created_date": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Resource creation date as epoch timestamp.",
			},
			"property": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Connection properties. These are specific to the connector type configured in `connector_id`. See the [DaVinci Connection Definitions](#davinci-connection-definitions) below to find the appropriate property name/value pairs for the connection.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the property.",
						},
						"value": {
							Type:        schema.TypeString,
							Required:    true,
							Sensitive:   true,
							Description: "The value of the property as string. If the property is an array, use a comma separated string.",
						},
						"type": {
							Type:         schema.TypeString,
							Optional:     true,
							Description:  "Type of the property. This is used to cast the value to the correct type. Must be: `string` or `boolean`. Use `string` for array types.",
							ValidateFunc: validation.StringInSlice([]string{"string", "boolean"}, false),

							Default: "string",
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: resourceConnectionImport,
		},
	}
}

func resourceConnectionCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	connection := dv.Connection{
		ConnectorID: d.Get("connector_id").(string),
		Name:        d.Get("name").(string),
	}

	connection.Properties = makeProperties(d)

	environmentID := d.Get("environment_id").(string)

	sdkRes, err := sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.CreateInitializedConnectionWithResponse(environmentID, &connection)
		},
	)

	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	resp, ok := sdkRes.(*dv.Connection)
	if !ok || resp.Name == "" {
		err = fmt.Errorf("failed to cast created response to Connection")
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId(resp.ConnectionID)

	// Set properties based on incoming config after successful create
	// not using reponse itself because it may contain obfuscated values
	configProps := makePropsListMap(d)
	if err := d.Set("property", configProps); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	resourceConnectionRead(ctx, d, meta)

	return diags
}

// get properties from incoming config
func makePropsListMap(d *schema.ResourceData) []map[string]interface{} {
	propsList := d.Get("property").(*schema.Set).List()
	propsListMap := []map[string]interface{}{}
	for _, prop := range propsList {
		propMap := prop.(map[string]interface{})
		propsListMap = append(propsListMap, propMap)
	}
	return propsListMap
}

func resourceConnectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	connId := d.Id()

	environmentID := d.Get("environment_id").(string)

	sdkRes, err := sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.ReadConnectionWithResponse(environmentID, connId)
		},
	)
	if err != nil {
		if dvError, ok := err.(dv.ErrorResponse); ok {
			if dvError.HttpResponseCode == http.StatusNotFound || dvError.Code == dv.DV_ERROR_CODE_CONNECTION_NOT_FOUND {
				d.SetId("")
				return diags
			}
		}

		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	res, ok := sdkRes.(*dv.Connection)
	if !ok {
		err = fmt.Errorf("Unable to cast Connection type to response from Davinci API on connection id: %v", connId)
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId(res.ConnectionID)

	if err := d.Set("name", res.Name); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	if err := d.Set("connector_id", res.ConnectorID); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	if err := d.Set("created_date", res.CreatedDate); err != nil {
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
	// // override props with state props if obfuscated
	stateProps := makePropsListMap(d)
	for _, prop := range props {
		if prop["value"] == "******" {
			for _, stateProp := range stateProps {
				if prop["name"] == stateProp["name"] {
					prop["value"] = stateProp["value"]
				}
			}
		}
	}

	if err := d.Set("property", props); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	return diags
}

func resourceConnectionUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := meta.(*dv.APIClient)

	connId := d.Id()
	// API only allows property changes
	if d.HasChanges("property") {
		connection := dv.Connection{
			ConnectorID:  d.Get("connector_id").(string),
			Name:         d.Get("name").(string),
			ConnectionID: connId,
		}

		connection.Properties = makeProperties(d)

		environmentID := d.Get("environment_id").(string)

		sdkRes, err := sdk.DoRetryable(
			ctx,
			c,
			environmentID,
			func() (interface{}, *http.Response, error) {
				return c.UpdateConnectionWithResponse(environmentID, &connection)
			},
		)
		if err != nil {
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}

		res, ok := sdkRes.(*dv.Connection)
		if !ok || res.Name == "" {
			err = fmt.Errorf("Unable to parse update response from Davinci API on connection id: %v", connId)
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}

		// Set properties based on incoming config after successful create
		// not using reponse itself because it may contain obfuscated values
		configProps := makePropsListMap(d)
		if err := d.Set("property", configProps); err != nil {
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}
	}

	return resourceConnectionRead(ctx, d, meta)
}

func resourceConnectionDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	connId := d.Id()

	environmentID := d.Get("environment_id").(string)

	_, err := sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.DeleteConnectionWithResponse(environmentID, connId)
		},
	)
	if err != nil {
		if dvError, ok := err.(dv.ErrorResponse); ok {
			// Can indicate environment already deleted/missing
			if dvError.HttpResponseCode == http.StatusNotFound && dvError.Code == dv.DV_ERROR_CODE_CONNECTION_NOT_FOUND {
				return diags
			}
		}
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	return diags
}

func resourceConnectionImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	idComponents := []framework.ImportComponent{
		{
			Label:  "environment_id",
			Regexp: verify.P1ResourceIDRegexp,
		},
		{
			Label:     "davinci_connection_id",
			Regexp:    verify.P1DVResourceIDRegexp,
			PrimaryID: true,
		},
	}

	attributes, err := framework.ParseImportID(d.Id(), idComponents...)
	if err != nil {
		return nil, err
	}

	if err = d.Set("environment_id", attributes["environment_id"]); err != nil {
		return nil, err
	}

	d.SetId(attributes["davinci_connection_id"])

	resourceConnectionRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func flattenConnectionProperties(connectionProperties map[string]interface{}) ([]map[string]interface{}, error) {
	if connectionProperties == nil {
		return nil, nil
	}
	connProps := []map[string]interface{}{}
	for propName, propVal := range connectionProperties {
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
					thisProp["type"] = "string"
				}
			case "boolean":
				if pValue, ok := pMap["value"].(bool); ok {
					thisProp["value"] = strconv.FormatBool(pValue)
					thisProp["type"] = "boolean"
				}
			default:
				return nil, fmt.Errorf("For Property '%v': unable to identify value type, only string or boolean is currently supported", thisProp["name"])
			}
		} else {
			switch pMap["value"].(type) {
			case string:
				if _, ok := pMap["value"].(string); ok {
					thisProp["value"] = pMap["value"].(string)
					thisProp["type"] = "string"
				}
			case bool:
				if pValue, ok := pMap["value"].(bool); ok {
					thisProp["value"] = strconv.FormatBool(pValue)
					thisProp["type"] = "boolean"
				}
			default:
				return nil, fmt.Errorf("For Property '%v': unable to identify value type, only string or boolean is currently supported", thisProp["name"])
			}
		}
		connProps = append(connProps, thisProp)
	}
	return connProps, nil
}

func makeProperties(d *schema.ResourceData) map[string]interface{} {
	connProps := map[string]interface{}{}
	props := d.Get("property").(*schema.Set).List()
	for _, raw := range props {
		prop := raw.(map[string]interface{})
		connProps[prop["name"].(string)] = map[string]interface{}{
			"value": prop["value"].(string),
		}
	}
	return connProps
}
