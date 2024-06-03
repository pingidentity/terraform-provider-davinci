package davinci

import (
	"context"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pingidentity/terraform-provider-davinci/internal/framework"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
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
				Description: "The name of the variable.  This field is immutable and will trigger a replace plan if changed.",
				ForceNew:    true,
			},
			"context": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"company", "flowInstance", "user"}, false),
				Description:  "The variable context.  Must be one of: `company`, `flowInstance`, `user`.   This field is immutable and will trigger a replace plan if changed.",
				ForceNew:     true,
			},
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the PingOne environment to create the DaVinci connection. Must be a valid PingOne resource ID. This field is immutable and will trigger a replace plan if changed.",

				ValidateDiagFunc: validation.ToDiagFunc(verify.ValidP1ResourceID),
				ForceNew:         true,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A string that specifies the description of the variable.",
			},
			"type": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringInSlice([]string{"string", "number", "boolean", "object"}, false),
				Description:  "The variable's data type.  Must be one of `string`, `number`, `boolean`, `object`.",
			},
			"mutable": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "A boolean that specifies whether the variable is mutable.  If `true`, the variable can be modified by the flow. If `false`, the variable is read-only and cannot be modified by the flow.",
			},
			"value": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Variable value as string, type will be inferred from the value specified in the `type` parameter.",
				Sensitive:   true,
			},
			"min": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The minimum value of the variable, if the `type` parameter is set as `number`.",
				Default:     0,
			},
			"max": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "The maximum value of the variable, if the `type` parameter is set as `number`.",
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

	variablePayload := getVariableAttributes(d)

	environmentID := d.Get("environment_id").(string)

	sdkRes, err := sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.CreateVariableWithResponse(environmentID, &variablePayload)
		},
	)

	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	res, ok := sdkRes.(map[string]dv.Variable)
	if !ok {
		err = fmt.Errorf("Unable to parse response from Davinci API for variable")
		diags = append(diags, diag.FromErr(err)...)
		return diags
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

	variableName := d.Id()

	environmentID := d.Get("environment_id").(string)

	params := dv.Params{
		Page:  "0",
		Limit: "100",
	}

	sdkRes, err := sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.ReadVariablesWithResponse(environmentID, &params)
		},
	)

	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	resp, ok := sdkRes.(map[string]dv.Variable)
	if !ok {
		err = fmt.Errorf("Unable to cast variable type to response from Davinci API for variable with name: %s", variableName)
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	//variable not found
	if len(resp) == 0 {
		d.SetId("")
		return diags
	}

	found := false

	for key, res := range resp {
		if strings.Contains(key, variableName) {
			found = true

			s := strings.Split(variableName, "##SK##")
			name := s[0]
			context := s[1]
			if err := d.Set("name", name); err != nil {
				diags = append(diags, diag.FromErr(err)...)
				return diags
			}
			if err := d.Set("context", context); err != nil {
				diags = append(diags, diag.FromErr(err)...)
				return diags
			}
			if err := d.Set("environment_id", res.CompanyID); err != nil {
				diags = append(diags, diag.FromErr(err)...)
				return diags
			}
			if err := d.Set("type", res.Type); err != nil {
				diags = append(diags, diag.FromErr(err)...)
				return diags
			}
			if err := d.Set("mutable", res.Mutable); err != nil {
				diags = append(diags, diag.FromErr(err)...)
				return diags
			}
			if err := d.Set("description", res.DisplayName); err != nil {
				diags = append(diags, diag.FromErr(err)...)
				return diags
			}
			if err := d.Set("value", res.Value); err != nil {
				diags = append(diags, diag.FromErr(err)...)
				return diags
			}
			if err := d.Set("min", res.Min); err != nil {
				diags = append(diags, diag.FromErr(err)...)
				return diags
			}
			if err := d.Set("max", res.Max); err != nil {
				diags = append(diags, diag.FromErr(err)...)
				return diags
			}
		}
	}

	if !found {
		d.SetId("")
	}

	return diags
}

func resourceVariableUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := meta.(*dv.APIClient)

	environmentID := d.Get("environment_id").(string)

	if d.HasChanges("name", "context") {
		return diag.Errorf(`Updates to name and context are not allowed`)
	}

	variablePayload := getVariableAttributes(d)

	sdkRes, err := sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.UpdateVariableWithResponse(environmentID, &variablePayload)
		},
	)

	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	resp, ok := sdkRes.(map[string]dv.Variable)
	if !ok || len(resp) != 1 {
		err = fmt.Errorf("Unable to parse update response from Davinci API for variable with name: %v", *variablePayload.Name)
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	return resourceVariableRead(ctx, d, meta)
}

func resourceVariableDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	variableName := d.Id()

	environmentID := d.Get("environment_id").(string)

	sdkRes, err := sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.DeleteVariableWithResponse(environmentID, variableName)
		},
	)

	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	resp, ok := sdkRes.(*dv.Message)
	if !ok || resp.Message == nil || *resp.Message == "" {
		err = fmt.Errorf("Unable to parse delete response from Davinci API for variable with name: %s", variableName)
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId("")

	return diags
}

func resourceVariableImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	idComponents := []framework.ImportComponent{
		{
			Label:  "environment_id",
			Regexp: verify.P1ResourceIDRegexp,
		},
		{
			Label:     "davinci_variable_id",
			Regexp:    regexp.MustCompile(`.*`),
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

	d.SetId(attributes["davinci_variable_id"])

	resourceVariableRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func getVariableAttributes(d *schema.ResourceData) dv.VariablePayload {
	variablePayload := dv.VariablePayload{
		Context: d.Get("context").(string),
		Type:    d.Get("type").(string),
	}

	if v, ok := d.Get("name").(string); ok {
		variablePayload.Name = &v
	}

	if flowId, ok := d.GetOk("flow_id"); ok {
		if v, ok := flowId.(string); ok {
			variablePayload.FlowId = &v
		}
	}
	if mutable, ok := d.GetOk("mutable"); ok {
		if v, ok := mutable.(bool); ok {
			variablePayload.Mutable = &v
		}
	}
	if description, ok := d.GetOk("description"); ok {
		if v, ok := description.(string); ok {
			variablePayload.Description = &v
		}
	}
	if value, ok := d.GetOk("value"); ok {
		if v, ok := value.(string); ok {
			variablePayload.Value = &v
		}
	}
	if min, ok := d.GetOk("min"); ok {
		if v, ok := min.(int); ok {
			variablePayload.Min = &v
		}
	}
	if max, ok := d.GetOk("max"); ok {
		if v, ok := max.(int); ok {
			variablePayload.Max = &v
		}
	}
	return variablePayload
}

// Framework
type VariableResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	FlowId      types.String `tfsdk:"flow_id"`
	Context     types.String `tfsdk:"context"`
	Type        types.String `tfsdk:"type"`
	Value       types.String `tfsdk:"value"`
	Mutable     types.Bool   `tfsdk:"mutable"`
	Min         types.Int64  `tfsdk:"min"`
	Max         types.Int64  `tfsdk:"max"`
}

func (p *VariableResourceModel) expand() *dv.VariablePayload {

	data := dv.VariablePayload{
		Context: p.Context.ValueString(),
		Type:    p.Type.ValueString(),
	}

	if !p.Name.IsNull() {
		data.Name = p.Name.ValueStringPointer()
	}

	if !p.Description.IsNull() {
		data.Description = p.Description.ValueStringPointer()
	}

	if !p.FlowId.IsNull() {
		data.FlowId = p.FlowId.ValueStringPointer()
	}

	if !p.Value.IsNull() {
		data.Value = p.Value.ValueStringPointer()
	}

	if !p.Mutable.IsNull() {
		data.Mutable = p.Mutable.ValueBoolPointer()
	}

	if !p.Min.IsNull() {
		min := int(p.Min.ValueInt64())
		data.Min = &min
	}

	if !p.Max.IsNull() {
		max := int(p.Max.ValueInt64())
		data.Max = &max
	}

	return &data
}
