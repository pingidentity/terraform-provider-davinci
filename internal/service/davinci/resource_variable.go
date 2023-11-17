package davinci

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	"github.com/pingidentity/terraform-provider-davinci/internal/utils"
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
				ForceNew:    true,
			},
			"context": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"company", "flowInstance", "user"}, false),
				Description:  "Must be one of: company, flowInstance, user",
				ForceNew:     true,
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
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "If true, the variable can be modified by the flow. If false, the variable is read-only and cannot be modified by the flow.",
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Variable value as string, type will be inferred",
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
			StateContext: resourceVariableImport,
		},
	}
}

func resourceVariableCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
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
		d.SetId(fmt.Sprint(i))
	}

	resourceVariableRead(ctx, d, meta)

	return diags
}

func resourceVariableRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
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

	//variable not found
	if len(resp) != 1 {
		d.SetId("")
		return diags
	}

	for _, res := range resp {
		s := strings.Split(variableName, "##SK##")
		name := s[0]
		context := s[1]
		if err := d.Set("name", name); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("context", context); err != nil {
			return diag.FromErr(err)
		}
		if err := d.Set("environment_id", res.CompanyID); err != nil {
			return diag.FromErr(err)
		}
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

func resourceVariableUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)

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

	return resourceVariableRead(ctx, d, meta)
}

func resourceVariableDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
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

func resourceVariableImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	idComponents := []utils.ImportComponent{
		{
			Label:  "environment_id",
			Regexp: regexp.MustCompile(`[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`),
		},
		{
			Label:     "davinci_variable_id",
			Regexp:    regexp.MustCompile(`.*`),
			PrimaryID: true,
		},
	}

	attributes, err := utils.ParseImportID(d.Id(), idComponents...)
	if err != nil {
		return nil, err
	}

	if err = d.Set("environment_id", attributes["environment_id"]); err != nil {
		return nil, err
	}

	d.SetId(attributes["davinci_variable_id"])

	resourceVariableRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func getVariableAttributes(d *schema.ResourceData) dv.VariablePayload {
	variablePayload := dv.VariablePayload{
		Name:    d.Get("name").(string),
		Context: d.Get("context").(string),
		Type:    d.Get("type").(string),
	}
	if flowId, ok := d.GetOk("flow_id"); ok {
		variablePayload.FlowId = flowId.(string)
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
