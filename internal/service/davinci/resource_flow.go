package davinci

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/pingidentity/terraform-provider-davinci/internal/framework"
	stringvalidatorinternal "github.com/pingidentity/terraform-provider-davinci/internal/framework/stringvalidator"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

// Types
type FlowResource serviceClientType

type FlowResourceModel struct {
	Id               types.String `tfsdk:"id"`
	EnvironmentId    types.String `tfsdk:"environment_id"`
	FlowJSON         types.String `tfsdk:"flow_json"`
	FlowJSONResponse types.String `tfsdk:"flow_json_response"`
	Deploy           types.Bool   `tfsdk:"deploy"`
	Name             types.String `tfsdk:"name"`
	Description      types.String `tfsdk:"description"`
	ConnectionLinks  types.Set    `tfsdk:"connection_link"`
	SubFlowLinks     types.Set    `tfsdk:"subflow_link"`
	FlowVariables    types.Set    `tfsdk:"flow_variables"`
}

type FlowConnectionLinkResourceModel struct {
	Id                        types.String `tfsdk:"id"`
	ReplaceImportConnectionId types.String `tfsdk:"replace_import_connection_id"`
	Name                      types.String `tfsdk:"name"`
}

type FlowSubflowLinkResourceModel struct {
	Id                        types.String `tfsdk:"id"`
	ReplaceImportConnectionId types.String `tfsdk:"replace_import_connection_id"`
	Name                      types.String `tfsdk:"name"`
}

type FlowVariablesResourceModel struct {
	Id          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	FlowId      types.String `tfsdk:"flow_id"`
	Context     types.String `tfsdk:"context"`
	Type        types.String `tfsdk:"type"`
	Mutable     types.Bool   `tfsdk:"mutable"`
	Min         types.Int64  `tfsdk:"min"`
	Max         types.Int64  `tfsdk:"max"`
}

// Framework interfaces
var (
	_ resource.Resource                = &FlowResource{}
	_ resource.ResourceWithConfigure   = &FlowResource{}
	_ resource.ResourceWithModifyPlan  = &FlowResource{}
	_ resource.ResourceWithImportState = &FlowResource{}
)

// New Object
func NewFlowResource() resource.Resource {
	return &FlowResource{}
}

// Metadata
func (r *FlowResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_flow"
}

// Schema.
func (r *FlowResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {

	const attrMinLength = 1

	nameDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"A string that identifies the flow name after import.  If the field is left blank, a flow name will be derived by the service from the name in the import JSON (the `flow_json` parameter).",
	)

	descriptionDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"A string that specifies a description of the flow.  If the field is left blank, a description value will be derived by the service.",
	)

	deployDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"**Deprecation notice:** This attribute is deprecated and will be removed in a future release.  Flows are automatically deployed on import. A boolean that specifies whether to deploy the flow after import.",
	).DefaultValue(true)

	connectionLinkDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"Mappings to connections that this flow depends on.  Connections should be managed (with the `davinci_connection` resource) or retrieved (with the `davinci_connection` data source) to provide the mappings needed for this configuration block.",
	)

	connectionLinkNameDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"The connector name.  If `replace_import_connection_id` is also specified, this value is used when the flow is imported.  If `replace_import_connection_id` is not specified, the name must match that of the connector in the import file, so the connector ID in the `id` parameter can be updated.",
	)

	connectionLinkReplaceImportConnectionIdDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"Connection ID of the connector in the import to replace with the connector described in `id` and `name` parameters.  This can be found in the source system in the \"Connectors\" menu, but is also at the following path in the JSON file: `[enabledGraphData|graphData].elements.nodes.data.connectionId`.",
	)

	subflowLinkDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"Child flows of this resource, where the `flow_json` contains reference to subflows.  If the `flow_json` contains subflows, this one `subflow_link` block is required per contained subflow.",
	)

	subflowLinkNameDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"The subflow name.  If `replace_import_subflow_id` is also specified, this value is used when the flow is imported.  If `replace_import_subflow_id` is not specified, the name must match that of the connector in the import file, so the connector ID in the `id` parameter can be updated.",
	)

	subflowLinkReplaceImportConnectionIdDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"Subflow ID of the subflow in the import to replace with the subflow described in `id` and `name` parameters.  This can be found in the source system in the \"Connectors\" menu, but is also at the following path in the JSON file: `[enabledGraphData|graphData].elements.nodes.data.connectionId`.",
	)

	flowVariablesDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"Returned list of Flow Context variables. These are variable resources that are created and managed by the Flow resource via `flow_json`.",
	)

	flowVariablesContextDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"The variable context.  Should always return `flow`.",
	)

	flowVariablesTypeDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"The variable's data type.  Expected to be one of `string`, `number`, `boolean`, `object`.",
	)

	flowVariablesMutableDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"A boolean that specifies whether the variable is mutable.  If `true`, the variable can be modified by the flow. If `false`, the variable is read-only and cannot be modified by the flow.",
	)

	flowVariablesMinDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"The minimum value of the variable, if the `type` parameter is set as `number`.",
	)

	flowVariablesMaxDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"The maximum value of the variable, if the `type` parameter is set as `number`.",
	)

	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		Description: "Resource to import and manage a DaVinci flow in an environment.  Connection and Subflow references in the JSON export can be overridden with ones managed by Terraform, see the examples and schema below for details.",

		Attributes: map[string]schema.Attribute{
			"id": framework.Attr_ID(),

			"environment_id": framework.Attr_LinkID(
				framework.SchemaAttributeDescriptionFromMarkdown("The ID of the PingOne environment to import the DaVinci flow to."),
			),

			"name": schema.StringAttribute{
				Description:         nameDescription.Description,
				MarkdownDescription: nameDescription.MarkdownDescription,
				Optional:            true,
				Computed:            true,

				Validators: []validator.String{
					stringvalidator.LengthAtLeast(attrMinLength),
				},
			},

			"description": schema.StringAttribute{
				Description:         descriptionDescription.Description,
				MarkdownDescription: descriptionDescription.MarkdownDescription,
				Optional:            true,
				Computed:            true,
			},

			"flow_json": schema.StringAttribute{
				Description: framework.SchemaAttributeDescriptionFromMarkdown("The DaVinci Flow export in raw json format.  Must be a valid JSON string.").Description,
				Required:    true,
				Sensitive:   true,

				Validators: []validator.String{
					stringvalidatorinternal.IsParseableJSON(),
				},
			},

			"flow_json_response": schema.StringAttribute{
				Description: framework.SchemaAttributeDescriptionFromMarkdown("The DaVinci Flow export in raw json format following successful import, including target environment metadata.").Description,
				Computed:    true,
				Sensitive:   true,
			},

			"deploy": schema.BoolAttribute{
				Description:         deployDescription.Description,
				MarkdownDescription: deployDescription.MarkdownDescription,
				DeprecationMessage:  "This attribute is deprecated and will be removed in a future release.  Flows are automatically deployed on import.",
				Optional:            true,
				Computed:            true,

				Default: booldefault.StaticBool(true),
			},
		},

		Blocks: map[string]schema.Block{

			"connection_link": schema.SetNestedBlock{
				Description:         connectionLinkDescription.Description,
				MarkdownDescription: connectionLinkDescription.MarkdownDescription,

				NestedObject: schema.NestedBlockObject{

					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: framework.SchemaAttributeDescriptionFromMarkdown("A string that specifies the connector ID that will be applied when flow is imported.").Description,
							Required:    true,

							Validators: []validator.String{
								verify.P1DVResourceIDValidator(),
							},
						},

						"name": schema.StringAttribute{
							Description:         connectionLinkNameDescription.Description,
							MarkdownDescription: connectionLinkNameDescription.MarkdownDescription,
							Required:            true,
						},

						"replace_import_connection_id": schema.StringAttribute{
							Description:         connectionLinkReplaceImportConnectionIdDescription.Description,
							MarkdownDescription: connectionLinkReplaceImportConnectionIdDescription.MarkdownDescription,
							Optional:            true,
						},
					},
				},
			},

			"subflow_link": schema.SetNestedBlock{
				Description:         subflowLinkDescription.Description,
				MarkdownDescription: subflowLinkDescription.MarkdownDescription,

				NestedObject: schema.NestedBlockObject{

					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: framework.SchemaAttributeDescriptionFromMarkdown("A string that specifies the subflow ID that will be applied when flow is imported.").Description,
							Required:    true,

							Validators: []validator.String{
								verify.P1DVResourceIDValidator(),
							},
						},

						"name": schema.StringAttribute{
							Description:         subflowLinkNameDescription.Description,
							MarkdownDescription: subflowLinkNameDescription.MarkdownDescription,
							Required:            true,
						},

						"replace_import_subflow_id": schema.StringAttribute{
							Description:         subflowLinkReplaceImportConnectionIdDescription.Description,
							MarkdownDescription: subflowLinkReplaceImportConnectionIdDescription.MarkdownDescription,
							Optional:            true,
						},
					},
				},
			},

			"flow_variables": schema.SetNestedBlock{
				Description:         flowVariablesDescription.Description,
				MarkdownDescription: flowVariablesDescription.MarkdownDescription,

				NestedObject: schema.NestedBlockObject{

					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: framework.SchemaAttributeDescriptionFromMarkdown("The DaVinci internal ID of the variable.").Description,
							Computed:    true,
						},

						"name": schema.StringAttribute{
							Description: framework.SchemaAttributeDescriptionFromMarkdown("The user friendly name of the variable in the UI.").Description,
							Computed:    true,
						},

						"description": schema.StringAttribute{
							Description: framework.SchemaAttributeDescriptionFromMarkdown("A string that specifies the description of the variable.").Description,
							Computed:    true,
						},

						"flow_id": schema.StringAttribute{
							Description: framework.SchemaAttributeDescriptionFromMarkdown("The flow ID that the variable belongs to, which should match the ID of this resource.").Description,
							Computed:    true,
						},

						"context": schema.StringAttribute{
							Description:         flowVariablesContextDescription.Description,
							MarkdownDescription: flowVariablesContextDescription.MarkdownDescription,
							Computed:            true,
						},

						"type": schema.StringAttribute{
							Description:         flowVariablesTypeDescription.Description,
							MarkdownDescription: flowVariablesTypeDescription.MarkdownDescription,
							Computed:            true,
						},

						"mutable": schema.BoolAttribute{
							Description:         flowVariablesMutableDescription.Description,
							MarkdownDescription: flowVariablesMutableDescription.MarkdownDescription,
							Computed:            true,
						},

						"min": schema.Int64Attribute{
							Description:         flowVariablesMinDescription.Description,
							MarkdownDescription: flowVariablesMinDescription.MarkdownDescription,
							Computed:            true,
						},

						"max": schema.Int64Attribute{
							Description:         flowVariablesMaxDescription.Description,
							MarkdownDescription: flowVariablesMaxDescription.MarkdownDescription,
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (p *FlowResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// TODO
}

func (r *FlowResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *FlowResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan, state FlowResourceModel

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
	daVinciImport, d := plan.expand()
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	environmentID := plan.EnvironmentId.ValueString()

	sdkRes, err := sdk.DoRetryable(
		ctx,
		r.Client,
		environmentID,
		func() (any, *http.Response, error) {
			return r.Client.CreateFlowWithResponse(environmentID, daVinciImport)
		},
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error importing flow",
			fmt.Sprintf("Error creating flow: %s", err),
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	response, ok := sdkRes.(*dv.Flow)
	if !ok || response.Name == "" {
		resp.Diagnostics.AddError(
			"Unexpected response",
			fmt.Sprintf("Unable to parse create response from Davinci API on flow"),
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	// Create the state to save
	state = plan

	// Save updated data into Terraform state
	resp.Diagnostics.Append(state.toState(response)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *FlowResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data *FlowResourceModel

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
	flowID := data.Id.ValueString()

	// Run the API call
	sdkRes, err := sdk.DoRetryable(
		ctx,
		r.Client,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return r.Client.ReadFlowVersionWithResponse(environmentID, flowID, nil)
		},
	)

	if err != nil {
		//httpErr, _ := dv.ParseDvHttpError(err)
		// if strings.Contains(httpErr.Body, "Error retrieving flow version") {
		// 	d.SetId("")
		// 	return diags
		// }
		resp.Diagnostics.AddError(
			"Error reading flow",
			fmt.Sprintf("Error reading flow: %s", err),
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	response, ok := sdkRes.(*dv.FlowInfo)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected response",
			fmt.Sprintf("Unable to parse read response from Davinci API on flow id: %v", flowID),
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	// Remove from state if resource is not found
	if response == nil || response.Flow.FlowID == "" {
		resp.State.RemoveResource(ctx)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(data.toState(&response.Flow)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *FlowResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FlowResourceModel

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

	daVinciImport, d := plan.expand()
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	environmentID := plan.EnvironmentId.ValueString()
	flowID := plan.Id.ValueString()

	sdkRes, err := sdk.DoRetryable(
		ctx,
		r.Client,
		environmentID,
		func() (any, *http.Response, error) {
			return r.Client.UpdateFlowWithResponse(environmentID, flowID, daVinciImport)
		},
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error importing flow",
			fmt.Sprintf("Error creating flow: %s", err),
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	response, ok := sdkRes.(*dv.Flow)
	if !ok || response.Name == "" {
		resp.Diagnostics.AddError(
			"Unexpected response",
			fmt.Sprintf("Unable to parse create response from Davinci API on flow"),
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	// Create the state to save
	state = plan

	// Save updated data into Terraform state
	resp.Diagnostics.Append(state.toState(response)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *FlowResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data FlowResourceModel

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
	flowID := data.Id.ValueString()

	sdkRes, err := sdk.DoRetryable(
		ctx,
		r.Client,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return r.Client.DeleteFlowWithResponse(environmentID, flowID)
		},
	)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error deleting flow",
			fmt.Sprintf("Error deleting flow: %s", err),
		)
	}
	res, ok := sdkRes.(*dv.Message)
	if !ok || res.Message == "" {
		resp.Diagnostics.AddWarning(
			"Unexpected response",
			fmt.Sprintf("Unable to parse delete response from Davinci API on flow id: %v", flowID),
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

}

func (r *FlowResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {

	idComponents := []framework.ImportComponent{
		{
			Label:  "environment_id",
			Regexp: verify.P1ResourceIDRegexp,
		},
		{
			Label:  "davinci_flow_id",
			Regexp: verify.P1DVResourceIDRegexp,
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

func (p *FlowResourceModel) expand() (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := dv.FlowImport{
		Name:        p.Name.ValueString(),
		Description: p.Description.ValueString(),
	}

	err := json.Unmarshal([]byte(p.FlowJSON.ValueString()), &data.FlowInfo)
	if err != nil {
		diags.AddError(
			"Error parsing flow_json",
			fmt.Sprintf("Error parsing flow_json: %s", err),
		)
		return nil, diags
	}

	data.FlowNameMapping = map[string]string{
		data.FlowInfo.FlowID: p.Name.ValueString(),
	}

	return data, diags
}

func (p *FlowResourceModel) toState(apiObject *dv.Flow) diag.Diagnostics {
	var diags diag.Diagnostics

	if apiObject == nil {
		diags.AddError(
			"Data object missing",
			"Cannot convert the data object to state as the data object is nil.  Please report this to the provider maintainers.",
		)

		return diags
	}

	p.Id = framework.StringToTF(apiObject.FlowID)
	p.EnvironmentId = framework.StringToTF(apiObject.CompanyID)

	jsonBytes, err := json.Marshal(apiObject)
	if err != nil {
		diags.AddError(
			"Error converting the flow object to JSON",
			fmt.Sprintf("Error converting the flow object (from the API response) to JSON.  This is a bug in the provider, please report this to the provider maintainers. Error: %s", err),
		)
		return diags
	}

	p.FlowJSONResponse = framework.StringToTF(string(jsonBytes[:]))

	if apiObject.DeployedDate != nil && *apiObject.DeployedDate > 0 {
		p.Deploy = types.BoolValue(true)
	} else {
		p.Deploy = types.BoolValue(false)
	}

	p.Name = framework.StringToTF(apiObject.Name)

	if apiObject.Description != nil {
		p.Description = framework.StringToTF(*apiObject.Description)
	}

	// TODO

	return diags
}
