package davinci

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

func ResourceVariable() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceVariableCreate,
		ReadContext:   resourceVariableRead,
		UpdateContext: resourceVariableUpdate,
		DeleteContext: resourceVariableDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the variable",
			},
			"context": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"company", "flowInstance", "user"}, false),
				Description:  "Must be one of: company, flowInstance, user",
			},
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PingOne environment id",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of variable",
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"string", "number", "boolean", "object"}, false),
				Description:  "Must be one of: string, number, boolean, object",
			},
			"mutable": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Value as string, type will be inferred",
				Sensitive:   true,
			},
			"min": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
				Default:     0,
			},
			"max": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "",
				Default:     2000,
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceVariableCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}
	variablePayload := getVariableAttributes(d)

	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.CreateVariable(&c.CompanyID, &variablePayload)
	}, nil)

	if err != nil {
		return diag.FromErr(err)
	}
	res, ok := sdkRes.(map[string]dv.Variable)
	if !ok {
		err = fmt.Errorf("Unable to parse response from Davinci API for variable")
		return diag.FromErr(err)
	}

	if err != nil {
		return diag.FromErr(err)
	}
	if len(res) != 1 {
		return diag.Errorf("Received error while attempting to create variable")
	}
	for i := range res {
		d.SetId(fmt.Sprintf(i))
	}

	resourceVariableRead(ctx, d, m)

	return diags
}

func resourceVariableRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	variableName := d.Id()

	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.ReadVariable(&c.CompanyID, variableName)
	}, nil)

	if err != nil {
		return diag.FromErr(err)
	}
	resp, ok := sdkRes.(map[string]dv.Variable)
	if !ok {
		err = fmt.Errorf("Unable to cast variable type to response from Davinci API for variable with name: %s", variableName)
		return diag.FromErr(err)
	}

	if len(resp) != 1 {
		return diag.FromErr(fmt.Errorf("Received error while attempting to retrieve variable"))
	}
	for _, res := range resp {
		if err := d.Set("type", res.Type); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("mutable", res.Mutable); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("description", res.DisplayName); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("value", res.Value); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("min", res.Min); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("max", res.Max); err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func resourceVariableUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}
	if d.HasChanges("name", "context") {
		return diag.Errorf(`Updates to name and context are not allowed`)
	}

	if d.HasChanges("value", "description", "mutable", "type") {
		variablePayload := getVariableAttributes(d)

		sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
			return c.UpdateVariable(&c.CompanyID, &variablePayload)
		}, nil)

		if err != nil {
			return diag.FromErr(err)
		}
		resp, ok := sdkRes.(map[string]dv.Variable)
		if !ok || len(resp) != 1 {
			err = fmt.Errorf("Unable to parse update response from Davinci API for variable with name: %s", variablePayload.Name)
			return diag.FromErr(err)
		}

	}

	return resourceVariableRead(ctx, d, m)
}

func resourceVariableDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}
	variableName := d.Id()

	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.DeleteVariable(&c.CompanyID, variableName)
	}, nil)

	if err != nil {
		return diag.FromErr(err)
	}
	resp, ok := sdkRes.(*dv.Message)
	if !ok || resp.Message == "" {
		err = fmt.Errorf("Unable to parse delete response from Davinci API for variable with name: %s", variableName)
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func getVariableAttributes(d *schema.ResourceData) dv.VariablePayload {
	variablePayload := dv.VariablePayload{
		Name:    d.Get("name").(string),
		Context: d.Get("context").(string),
		Type:    d.Get("type").(string),
	}
	if mutable, ok := d.GetOk("mutable"); ok {
		variablePayload.Mutable = mutable.(bool)
	}
	if description, ok := d.GetOk("description"); ok {
		variablePayload.Description = description.(string)
	}
	if value, ok := d.GetOk("value"); ok {
		variablePayload.Value = value.(string)
	}
	if min, ok := d.GetOk("min"); ok {
		variablePayload.Min = min.(int)
	}
	if max, ok := d.GetOk("max"); ok {
		variablePayload.Max = max.(int)
	}
	return variablePayload
}
