package davinci

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/boolvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/pingidentity/terraform-provider-davinci/internal/framework"
	stringvalidatorinternal "github.com/pingidentity/terraform-provider-davinci/internal/framework/stringvalidator"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	"github.com/pingidentity/terraform-provider-davinci/internal/utils"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
	"github.com/samir-gandhi/davinci-client-go/davinci"
)

// Types
type VariableResource serviceClientType

type VariableResourceModel struct {
	Id            types.String `tfsdk:"id"`
	EnvironmentId types.String `tfsdk:"environment_id"`
	FlowId        types.String `tfsdk:"flow_id"`
	Name          types.String `tfsdk:"name"`
	Context       types.String `tfsdk:"context"`
	Description   types.String `tfsdk:"description"`
	Type          types.String `tfsdk:"type"`
	Mutable       types.Bool   `tfsdk:"mutable"`
	Value         types.String `tfsdk:"value"`
	EmptyValue    types.Bool   `tfsdk:"empty_value"`
	ValueService  types.String `tfsdk:"value_service"`
	Min           types.Int64  `tfsdk:"min"`
	Max           types.Int64  `tfsdk:"max"`
}

// Framework interfaces
var (
	_ resource.Resource                = &VariableResource{}
	_ resource.ResourceWithConfigure   = &VariableResource{}
	_ resource.ResourceWithModifyPlan  = &VariableResource{}
	_ resource.ResourceWithImportState = &VariableResource{}
)

const (
	contextCompany      = "company"
	contextFlowInstance = "flowInstance"
	contextUser         = "user"
	contextFlow         = "flow"
)

// New Object
func NewVariableResource() resource.Resource {
	return &VariableResource{}
}

// Metadata
func (r *VariableResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_variable"
}

// Schema.
func (r *VariableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	const attrMinLength = 1

	flowIdDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"A string that specifies the ID of the flow to which the variable is assigned.  This field is required when the `context` field is set to `flow`.",
	).AppendMarkdownString("Must be a valid PingOne resource ID.").RequiresReplace()

	nameDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"A string that specifies the name of the variable.",
	).RequiresReplace()

	contexts := []string{contextCompany, contextFlowInstance, contextUser, contextFlow}
	contextDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"A string that specifies the context of the variable.",
	).AllowedValues(utils.StringSliceToAnySlice(contexts)...).RequiresReplace()

	descriptionDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"A string that specifies the description of the variable.",
	)

	varTypes := []string{"string", "number", "boolean", "object", "secret"}
	typeDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"A string that specifies the variable's data type.",
	).AllowedValues(utils.StringSliceToAnySlice(varTypes)...)

	mutableDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"A boolean that specifies whether the variable is mutable.  If `true`, the variable can be modified by the flow. If `false`, the variable is read-only and cannot be modified by the flow.",
	).DefaultValue(true)

	valueDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"A string that specifies the default value of the variable, the type will be inferred from the value specified in the `type` parameter.  If left blank or omitted, the resource will not track the variable's value in state.  If the variable value should be tracked in state as an empty string, use the `empty_value` parameter.  Note that if the `type` is `secret`, the provider will not be able to remediate the value's configuration drift in the DaVinci service.",
	).ConflictsWith([]string{"empty_value"})

	valueServiceDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"A string that specifies the value of the variable in the service, the type will be inferred from the value specified in the `type` parameter.",
	)

	EmptyValueDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"A boolean that specifies whether the variable's `value` must be kept as an empty string.",
	).ConflictsWith([]string{"value"})

	minDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"An integer that specifies the minimum value of the variable, if the `type` parameter is set as `number`.",
	).DefaultValue(0)

	maxDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"An integer that specifies the maximum value of the variable, if the `type` parameter is set as `number`.",
	).DefaultValue(2000)

	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		Description: "Resource to import and manage a DaVinci variable in an environment.  Connection and Subvariable references in the JSON export can be overridden with ones managed by Terraform, see the examples and schema below for details.",

		Attributes: map[string]schema.Attribute{
			"id": framework.Attr_ID(),

			"environment_id": framework.Attr_LinkID(
				framework.SchemaAttributeDescriptionFromMarkdown("The ID of the PingOne environment to manage the DaVinci variable in."),
			),

			"flow_id": schema.StringAttribute{
				Description:         flowIdDescription.Description,
				MarkdownDescription: flowIdDescription.MarkdownDescription,
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
				Validators: []validator.String{
					verify.P1DVResourceIDValidator(),
					stringvalidatorinternal.IsRequiredIfMatchesPathValue(
						types.StringValue(contextFlow),
						path.MatchRelative().AtParent().AtName("context"),
					),
					stringvalidatorinternal.ConflictsIfMatchesPathValue(
						types.StringValue(contextCompany),
						path.MatchRelative().AtParent().AtName("context"),
					),
					stringvalidatorinternal.ConflictsIfMatchesPathValue(
						types.StringValue(contextFlowInstance),
						path.MatchRelative().AtParent().AtName("context"),
					),
					stringvalidatorinternal.ConflictsIfMatchesPathValue(
						types.StringValue(contextUser),
						path.MatchRelative().AtParent().AtName("context"),
					),
				},
			},

			"name": schema.StringAttribute{
				Description:         nameDescription.Description,
				MarkdownDescription: nameDescription.MarkdownDescription,
				Required:            true,

				Validators: []validator.String{
					stringvalidator.LengthAtLeast(attrMinLength),
				},

				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},

			"context": schema.StringAttribute{
				Description:         contextDescription.Description,
				MarkdownDescription: contextDescription.MarkdownDescription,
				Required:            true,

				Validators: []validator.String{
					stringvalidator.OneOf(contexts...),
				},

				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},

			"description": schema.StringAttribute{
				Description:         descriptionDescription.Description,
				MarkdownDescription: descriptionDescription.MarkdownDescription,
				Optional:            true,
			},

			"type": schema.StringAttribute{
				Description:         typeDescription.Description,
				MarkdownDescription: typeDescription.MarkdownDescription,
				Required:            true,

				Validators: []validator.String{
					stringvalidator.OneOf(varTypes...),
				},
			},

			"mutable": schema.BoolAttribute{
				Description:         mutableDescription.Description,
				MarkdownDescription: mutableDescription.MarkdownDescription,
				Optional:            true,
				Computed:            true,

				Default: booldefault.StaticBool(true),
			},

			"value": schema.StringAttribute{
				Description:         valueDescription.Description,
				MarkdownDescription: valueDescription.MarkdownDescription,
				Optional:            true,
				Sensitive:           true,

				Validators: []validator.String{
					stringvalidator.LengthAtLeast(attrMinLength),
					stringvalidator.ConflictsWith(
						path.MatchRelative().AtParent().AtName("value"),
						path.MatchRelative().AtParent().AtName("empty_value"),
					),
				},
			},

			"empty_value": schema.BoolAttribute{
				Description:         EmptyValueDescription.Description,
				MarkdownDescription: EmptyValueDescription.MarkdownDescription,
				Optional:            true,

				Validators: []validator.Bool{
					boolvalidator.ConflictsWith(
						path.MatchRelative().AtParent().AtName("value"),
						path.MatchRelative().AtParent().AtName("empty_value"),
					),
				},
			},

			"value_service": schema.StringAttribute{
				Description:         valueServiceDescription.Description,
				MarkdownDescription: valueServiceDescription.MarkdownDescription,
				Computed:            true,
				Sensitive:           true,
			},

			"min": schema.Int64Attribute{
				Description:         minDescription.Description,
				MarkdownDescription: minDescription.MarkdownDescription,
				Optional:            true,
				Computed:            true,

				Default: int64default.StaticInt64(0),
			},

			"max": schema.Int64Attribute{
				Description:         maxDescription.Description,
				MarkdownDescription: maxDescription.MarkdownDescription,
				Optional:            true,
				Computed:            true,

				Default: int64default.StaticInt64(2000),
			},
		},
	}
}

func (p *VariableResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {

	// Destruction plan
	if req.Plan.Raw.IsNull() {
		return
	}

	var varValuePlan basetypes.StringValue
	resp.Diagnostics.Append(resp.Plan.GetAttribute(ctx, path.Root("value"), &varValuePlan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	var varEmptyValuePlan basetypes.BoolValue
	resp.Diagnostics.Append(resp.Plan.GetAttribute(ctx, path.Root("empty_value"), &varEmptyValuePlan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	if !varEmptyValuePlan.IsNull() && !varEmptyValuePlan.IsUnknown() {
		resp.Plan.SetAttribute(ctx, path.Root("value_service"), types.StringNull())
	} else {

		if !varValuePlan.IsNull() && !varValuePlan.IsUnknown() {
			resp.Plan.SetAttribute(ctx, path.Root("value_service"), varValuePlan)
		}
	}

	if varEmptyValuePlan.IsUnknown() || varValuePlan.IsUnknown() {
		resp.Plan.SetAttribute(ctx, path.Root("value_service"), types.StringUnknown())
	}
}

func (r *VariableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	resourceConfig, ok := req.ProviderData.(framework.ResourceType)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected the provider client, got: %T. Please report this issue to the provider maintainers.", req.ProviderData),
		)

		return
	}

	r.Client = resourceConfig.Client
	if r.Client == nil {
		resp.Diagnostics.AddError(
			"Client not initialised",
			"Expected the DaVinci client, got nil.  Please report this issue to the provider maintainers.",
		)
		return
	}
}

func (r *VariableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, state VariableResourceModel

	if r.Client == nil {
		resp.Diagnostics.AddError(
			"Client not initialized",
			"Expected the DaVinci client, got nil.  Please report this issue to the provider maintainers.")
		return
	}

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build the model for the API
	variablePayload := plan.expand()

	environmentID := plan.EnvironmentId.ValueString()

	var response interface{}
	var err error

	response, err = sdk.DoRetryable(
		ctx,
		r.Client,
		environmentID,
		func() (any, *http.Response, error) {
			return r.Client.CreateVariableWithResponse(environmentID, variablePayload)
		},
	)
	if err != nil {
		if strings.Contains(err.Error(), "Record already exists") {
			if plan.Context.ValueString() == contextFlow {
				// we override "flow" variables
				response, err = sdk.DoRetryable(
					ctx,
					r.Client,
					environmentID,
					func() (any, *http.Response, error) {
						return r.Client.UpdateVariableWithResponse(environmentID, variablePayload)
					},
				)
				if err != nil {
					resp.Diagnostics.AddError(
						"Error overriding flow variable",
						fmt.Sprintf("Error overriding flow variable: %s. %s", *variablePayload.Name, err),
					)
					return
				}
			}
		} else {
			resp.Diagnostics.AddError(
				"Error creating variable",
				fmt.Sprintf("Error creating variable: %s. %s", *variablePayload.Name, err),
			)
			return
		}
	}

	res, ok := response.(map[string]davinci.Variable)
	if !ok {
		resp.Diagnostics.AddError(
			"Error creating variable",
			"Unable to parse response from Davinci API for variable",
		)
		return
	}

	if len(res) != 1 {
		resp.Diagnostics.AddError(
			"Error creating variable",
			"Received error while attempting to create variable",
		)
		return
	}

	// Create the state to save
	state = plan

	// Save updated data into Terraform state
	resp.Diagnostics.Append(state.toState(res)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *VariableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *VariableResourceModel

	if r.Client == nil {
		resp.Diagnostics.AddError(
			"Client not initialized",
			"Expected the DaVinci client, got nil.  Please report this issue to the provider maintainers.")
		return
	}

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	environmentID := data.EnvironmentId.ValueString()
	variableID := data.Id.ValueString()

	response, err := sdk.DoRetryable(
		ctx,
		r.Client,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return r.Client.ReadVariableWithResponse(environmentID, variableID)
		},
	)
	if err != nil {
		// if err starts with "Variable not found" then return
		if strings.Contains(err.Error(), "Variable not found") {
			resp.State.RemoveResource(ctx)
			return
		}

		resp.Diagnostics.AddError(
			"Error reading variable",
			fmt.Sprintf("Error reading variable: %s", err),
		)

		return
	}
	res, ok := response.(map[string]davinci.Variable)
	if !ok {
		resp.Diagnostics.AddError(
			"Error creating variable",
			"Unable to parse response from Davinci API for variable",
		)
		return
	}

	if len(res) != 1 {
		resp.Diagnostics.AddError(
			"Error reading variable",
			"Received error while attempting to read variable, more than one variable returned",
		)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(data.toState(res)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *VariableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state VariableResourceModel

	if r.Client == nil {
		resp.Diagnostics.AddError(
			"Client not initialized",
			"Expected the DaVinci client, got nil.  Please report this issue to the provider maintainers.")
		return
	}

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build the model for the API
	variablePayload := plan.expand()

	environmentID := plan.EnvironmentId.ValueString()

	response, err := sdk.DoRetryable(
		ctx,
		r.Client,
		environmentID,
		func() (any, *http.Response, error) {
			return r.Client.UpdateVariableWithResponse(environmentID, variablePayload)
		},
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error updating variable",
			fmt.Sprintf("Error updating variable: %s. %s", *variablePayload.Name, err),
		)
		return
	}

	res, ok := response.(map[string]davinci.Variable)
	if !ok {
		resp.Diagnostics.AddError(
			"Error updating variable",
			"Unable to parse response from Davinci API for variable",
		)
		return
	}

	if len(res) != 1 {
		resp.Diagnostics.AddError(
			"Error updating variable",
			"Received error while attempting to updating variable",
		)
		return
	}

	// Create the state to save
	state = plan

	// Save updated data into Terraform state
	resp.Diagnostics.Append(state.toState(res)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *VariableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data VariableResourceModel

	if r.Client == nil {
		resp.Diagnostics.AddError(
			"Client not initialized",
			"Expected the DaVinci client, got nil.  Please report this issue to the provider maintainers.")
		return
	}

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	environmentID := data.EnvironmentId.ValueString()
	variableID := data.Id.ValueString()

	sdkRes, err := sdk.DoRetryable(
		ctx,
		r.Client,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return r.Client.DeleteVariableWithResponse(environmentID, variableID)
		},
	)

	if err != nil {
		if strings.Contains(err.Error(), "Variable not found") {
			return
		}

		resp.Diagnostics.AddError(
			"Error deleting variable",
			fmt.Sprintf("Error deleting variable: %s", err),
		)
	}
	res, ok := sdkRes.(*davinci.Message)
	if !ok || res.Message == nil || *res.Message == "" {
		resp.Diagnostics.AddWarning(
			"Unexpected response",
			fmt.Sprintf("Unable to parse delete response from Davinci API on variable id: %v", variableID),
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *VariableResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

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

	attributes, err := framework.ParseImportID(req.ID, idComponents...)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unexpected Import Identifier",
			err.Error(),
		)
		return
	}

	for _, idComponent := range idComponents {
		pathKey := idComponent.Label

		if idComponent.PrimaryID {
			pathKey = "id"
		}

		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root(pathKey), attributes[idComponent.Label])...)
	}
}

func (p *VariableResourceModel) expand() *davinci.VariablePayload {

	data := davinci.VariablePayload{
		Context: p.Context.ValueString(),
		Type:    p.Type.ValueString(),
	}

	if !p.Name.IsNull() && !p.Name.IsUnknown() {
		data.Name = p.Name.ValueStringPointer()
	}

	if !p.Description.IsNull() && !p.Description.IsUnknown() {
		data.Description = p.Description.ValueStringPointer()
	}

	if !p.Value.IsNull() && !p.Value.IsUnknown() {
		data.Value = p.Value.ValueStringPointer()
	}

	if !p.EmptyValue.IsNull() && !p.EmptyValue.IsUnknown() {
		v := ""
		data.Value = &v
	}

	if !p.Min.IsNull() && !p.Min.IsUnknown() {
		v := int(p.Min.ValueInt64())
		data.Min = &v
	}

	if !p.Max.IsNull() && !p.Max.IsUnknown() {
		v := int(p.Max.ValueInt64())
		data.Max = &v
	}

	if !p.Mutable.IsNull() && !p.Mutable.IsUnknown() {
		data.Mutable = p.Mutable.ValueBoolPointer()
	}

	if !p.FlowId.IsNull() && !p.FlowId.IsUnknown() {
		data.FlowId = p.FlowId.ValueStringPointer()
	}

	return &data
}

func (p *VariableResourceModel) toState(apiObject map[string]davinci.Variable) diag.Diagnostics {
	var diags diag.Diagnostics

	if len(apiObject) == 0 {
		diags.AddError(
			"Data object missing",
			"Cannot convert the data object to state as the data object is nil.  Please report this to the provider maintainers.",
		)

		return diags
	}

	if len(apiObject) > 1 {
		diags.AddError(
			"Too many data objects",
			"Cannot convert the data object to state as there are multiple returned.  Please report this to the provider maintainers.",
		)

		return diags
	}

	var variableKey string
	var variableObject davinci.Variable
	for i, v := range apiObject {
		variableKey = i
		variableObject = v
	}

	s := strings.Split(variableKey, "##SK##")
	variableName := s[0]
	variableContext := s[1]

	p.Id = framework.StringToTF(variableKey)
	//p.EnvironmentId = framework.StringToTF(apiObject.CompanyID)
	//p.FlowId = framework.StringToTF(variableObject.FlowID)
	p.Name = framework.StringToTF(variableName)
	p.Context = framework.StringToTF(variableContext)

	if v := variableObject.DisplayName; v != nil {
		p.Description = framework.StringToTF(*v)
	} else {
		p.Description = types.StringNull()
	}

	if v := variableObject.Type; v != nil {
		p.Type = framework.StringToTF(*v)
	} else {
		p.Type = types.StringNull()
	}

	if v := variableObject.Mutable; v != nil {
		p.Mutable = framework.BoolToTF(*v)
	} else {
		p.Mutable = types.BoolValue(false) // comes back null if false
	}

	// if v := variableObject.Value; v != nil {
	// 	p.Value = framework.StringToTF(*v)
	// } else {
	// 	p.Value = types.StringNull()
	// }

	if v := variableObject.Value; v != nil {

		value := ""
		// switch type of v
		switch v := v.(type) {
		case string:
			value = v
			value = regexp.MustCompile(`^\*+$`).ReplaceAllString(value, p.Value.ValueString())
		case bool:
			value = fmt.Sprintf("%v", v)
		case int:
			value = fmt.Sprintf("%d", v)
		default:
			bytes, err := json.Marshal(v)
			if err != nil {
				diags.AddError("Error marshalling variable value", err.Error())
				return diags
			}
			value = string(bytes)
		}

		p.ValueService = framework.StringToTF(value)
	} else {
		p.ValueService = types.StringNull()
	}

	if v := variableObject.Min; v != nil {
		safeInt, err := utils.SafeIntToInt32(*v)
		if err != nil {
			diags.AddError("Error converting min value", err.Error())
		} else {
			p.Min = framework.Int32ToTF(safeInt)
		}
	} else {
		p.Min = types.Int64Null()
	}

	if v := variableObject.Max; v != nil {
		safeInt, err := utils.SafeIntToInt32(*v)
		if err != nil {
			diags.AddError("Error converting max value", err.Error())
		} else {
			p.Max = framework.Int32ToTF(safeInt)
		}
	} else {
		p.Max = types.Int64Null()
	}

	return diags
}
