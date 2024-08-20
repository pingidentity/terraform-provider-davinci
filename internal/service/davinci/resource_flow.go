package davinci

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-framework-validators/setvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/pingidentity/terraform-provider-davinci/internal/framework"
	"github.com/pingidentity/terraform-provider-davinci/internal/framework/customtypes/davinciexporttype"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
	"github.com/samir-gandhi/davinci-client-go/davinci"
)

// Types
type FlowResource serviceClientType

type FlowResourceModel struct {
	Id                    types.String                  `tfsdk:"id"`
	EnvironmentId         types.String                  `tfsdk:"environment_id"`
	FlowJSON              davinciexporttype.ParsedValue `tfsdk:"flow_json"`
	FlowConfigurationJSON davinciexporttype.ParsedValue `tfsdk:"flow_configuration_json"`
	FlowExportJSON        davinciexporttype.ParsedValue `tfsdk:"flow_export_json"`
	Deploy                types.Bool                    `tfsdk:"deploy"`
	Name                  types.String                  `tfsdk:"name"`
	Description           types.String                  `tfsdk:"description"`
	ConnectionLinks       types.Set                     `tfsdk:"connection_link"`
	SubFlowLinks          types.Set                     `tfsdk:"subflow_link"`
	FlowVariables         types.Set                     `tfsdk:"flow_variables"`
}

type FlowVariableResourceModel struct {
	Id      types.String `tfsdk:"id"`
	FlowId  types.String `tfsdk:"flow_id"`
	Name    types.String `tfsdk:"name"`
	Context types.String `tfsdk:"context"`
	Type    types.String `tfsdk:"type"`
}

type FlowConnectionLinkResourceModel struct {
	Id                        types.String `tfsdk:"id"`
	ReplaceImportConnectionId types.String `tfsdk:"replace_import_connection_id"`
	Name                      types.String `tfsdk:"name"`
}

type FlowSubflowLinkResourceModel struct {
	Id                     types.String `tfsdk:"id"`
	ReplaceImportSubflowId types.String `tfsdk:"replace_import_subflow_id"`
	Name                   types.String `tfsdk:"name"`
}

var (
	flowVariablesTFObjectTypes = map[string]attr.Type{
		"id":      types.StringType,
		"name":    types.StringType,
		"flow_id": types.StringType,
		"context": types.StringType,
		"type":    types.StringType,
	}
)

var (
	flowJsonCmpOptsConfiguration = davinci.ExportCmpOpts{
		IgnoreConfig:              false,
		IgnoreDesignerCues:        false,
		IgnoreEnvironmentMetadata: true,
		IgnoreUnmappedProperties:  false,
		IgnoreVersionMetadata:     true,
		IgnoreFlowMetadata:        false,
		IgnoreFlowVariables:       false,
		NodeOpts: &davinci.ExportNodeCmpOpts{
			VariablesConnector: &davinci.ExportNodeVariablesCmpOpts{
				ExpectVariableValues: true, // The input field needs validation
			},
		},
	}

	flowConfigurationJsonCmpOptsConfiguration = davinci.ExportCmpOpts{
		IgnoreConfig:              false,
		IgnoreDesignerCues:        false,
		IgnoreEnvironmentMetadata: true,
		IgnoreUnmappedProperties:  false,
		IgnoreVersionMetadata:     true,
		IgnoreFlowMetadata:        true,
		IgnoreFlowVariables:       true,
		NodeOpts: &davinci.ExportNodeCmpOpts{
			VariablesConnector: &davinci.ExportNodeVariablesCmpOpts{
				ExpectVariableValues: false, // We don't need to validate this
			},
		},
	}

	flowExportJsonCmpOptsConfiguration = davinci.ExportCmpOpts{
		IgnoreConfig:              false,
		IgnoreDesignerCues:        false,
		IgnoreEnvironmentMetadata: false,
		IgnoreUnmappedProperties:  true, // because we don't do calculation with this object
		IgnoreVersionMetadata:     false,
		IgnoreFlowMetadata:        false,
		IgnoreFlowVariables:       true, // because this is handled by another resource
		NodeOpts: &davinci.ExportNodeCmpOpts{
			VariablesConnector: &davinci.ExportNodeVariablesCmpOpts{
				ExpectVariableValues: false, // We don't need to validate this
			},
		},
	}
)

// Framework interfaces
var (
	_ resource.Resource                   = &FlowResource{}
	_ resource.ResourceWithConfigure      = &FlowResource{}
	_ resource.ResourceWithValidateConfig = &FlowResource{}
	_ resource.ResourceWithModifyPlan     = &FlowResource{}
	_ resource.ResourceWithImportState    = &FlowResource{}
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
	const minimumConnectionMappings = 1

	nameDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"A string that identifies the flow name after import.  If the field is left blank, a flow name will be derived by the service from the name in the import JSON (the `flow_json` parameter).",
	)

	descriptionDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"A string that specifies a description of the flow.  If the field is left undefined, the description from the flow export will be used.  If this field is left undefined and the flow export does not contain a description, the service will define a description on import.",
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
		"List of Flow variables that will be updated in the DaVinci instance. These are variable resources that are created and managed by the Flow resource, where variables are embedded in the `flow_json` DaVinci export file.",
	)

	flowVariablesContextDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"The variable context.  Should always return `flow`.",
	)

	flowVariablesTypeDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"The variable's data type.  Expected to be one of `string`, `number`, `boolean`, `object`.",
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

				Validators: []validator.String{
					stringvalidator.LengthAtLeast(attrMinLength),
				},
			},

			"flow_json": schema.StringAttribute{
				Description: framework.SchemaAttributeDescriptionFromMarkdown("The DaVinci Flow to import, in raw JSON format. Should be a JSON file of a single flow (without subflows) that has been exported from a source DaVinci environment.  Must be a valid JSON string.").Description,
				Required:    true,
				Sensitive:   true,

				CustomType: davinciexporttype.ParsedType{
					ExportCmpOpts: flowJsonCmpOptsConfiguration,
				},
			},

			"flow_configuration_json": schema.StringAttribute{
				Description: framework.SchemaAttributeDescriptionFromMarkdown("The parsed configuration of the DaVinci Flow import JSON.  Drift is calculated based on this attribute.").Description,
				Computed:    true,
				Sensitive:   true,

				CustomType: davinciexporttype.ParsedType{
					ExportCmpOpts: flowConfigurationJsonCmpOptsConfiguration,
				},
			},

			"flow_export_json": schema.StringAttribute{
				Description: framework.SchemaAttributeDescriptionFromMarkdown("The DaVinci Flow export in raw JSON format following successful import, including target environment metadata.").Description,
				Computed:    true,
				Sensitive:   true,

				CustomType: davinciexporttype.ParsedType{
					ExportCmpOpts: flowExportJsonCmpOptsConfiguration,
				},
			},

			"deploy": schema.BoolAttribute{
				Description:         deployDescription.Description,
				MarkdownDescription: deployDescription.MarkdownDescription,
				DeprecationMessage:  "This attribute is deprecated and will be removed in a future release.  Flows are automatically deployed on import.",
				Optional:            true,
				Computed:            true,

				Default: booldefault.StaticBool(true),
			},

			"flow_variables": schema.SetNestedAttribute{
				Description:         flowVariablesDescription.Description,
				MarkdownDescription: flowVariablesDescription.MarkdownDescription,
				Computed:            true,

				NestedObject: schema.NestedAttributeObject{

					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Description: framework.SchemaAttributeDescriptionFromMarkdown("The DaVinci internal ID of the variable.").Description,
							Computed:    true,
						},

						"name": schema.StringAttribute{
							Description: framework.SchemaAttributeDescriptionFromMarkdown("The user friendly name of the variable in the UI.").Description,
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
					},
				},
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

				Validators: []validator.Set{
					setvalidator.SizeAtLeast(minimumConnectionMappings),
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
		},
	}
}

func (p *FlowResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {

	// Destruction plan
	if req.Plan.Raw.IsNull() {
		return
	}

	var plan, config FlowResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read Terraform config data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	unknownFlowConfigPlan := true
	var flowConfigObject davinci.FlowConfiguration

	// Compute the Flow Configuration (the drift of the import file is calculated based on this attribute)
	if !plan.FlowJSON.IsUnknown() {
		var flowObject davinci.Flow
		err := davinci.Unmarshal([]byte(plan.FlowJSON.ValueString()), &flowObject, davinci.ExportCmpOpts{})
		if err != nil {
			resp.Diagnostics.AddAttributeError(
				path.Root("flow_json"),
				"Error parsing `flow_json`",
				fmt.Sprintf("Error parsing `flow_json` into flow configuration object when applying plan modifications: %s", err),
			)
			return
		}

		if config.Name.IsNull() {
			resp.Plan.SetAttribute(ctx, path.Root("name"), flowObject.Name)
		}

		if v := flowObject.Description; config.Description.IsNull() && v != nil && *v != "" {
			resp.Plan.SetAttribute(ctx, path.Root("description"), *v)
		}

		flowConfigObject = flowObject.FlowConfiguration

		var d diag.Diagnostics

		unknownFlowConfigPlan, d = modifyPlanForConnectionSubflowLinkMappings(ctx, &flowConfigObject, plan.ConnectionLinks, plan.SubFlowLinks)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return
		}

		if !unknownFlowConfigPlan && !req.State.Raw.IsNull() {

			var state FlowResourceModel

			// Read Terraform state data into the model
			resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
			if resp.Diagnostics.HasError() {
				return
			}

			// Flow configuration
			if !state.FlowConfigurationJSON.IsNull() {
				// Compute the Flow Configuration (the drift of the import file is calculated based on this attribute)
				var flowConfigStateObject davinci.FlowConfiguration
				err = davinci.Unmarshal([]byte(state.FlowConfigurationJSON.ValueString()), &flowConfigStateObject, davinci.ExportCmpOpts{})
				if err != nil {
					resp.Diagnostics.AddAttributeError(
						path.Root("flow_json"),
						"Error parsing `flow_configuration_json` in state",
						fmt.Sprintf("Error parsing `flow_configuration_json` in state into flow configuration object: %s", err),
					)
					return
				}

				resp.Diagnostics.Append(modifyPlanForMergedProperties(&flowConfigObject.FlowUpdateConfiguration, flowConfigStateObject.FlowUpdateConfiguration)...)
				if resp.Diagnostics.HasError() {
					return
				}
			}

			// Flow variables
			flowVariablesPlan := make([]davinci.FlowVariable, 0)
			for _, flowVariable := range flowObject.FlowMetadata.Variables {
				if v := flowVariable.FlowID; v != nil {
					flowVariableIDOld := *flowVariable.FlowID
					flowVariable.FlowID = state.Id.ValueStringPointer()
					flowVariable.Name = strings.Replace(flowVariable.Name, flowVariableIDOld, state.Id.ValueString(), -1)
				}

				flowVariablesPlan = append(flowVariablesPlan, flowVariable)
			}

			var d diag.Diagnostics
			flowVariables, d := flowVariablesToTF(flowVariablesPlan)
			resp.Diagnostics.Append(d...)
			resp.Plan.SetAttribute(ctx, path.Root("flow_variables"), flowVariables)
		}
	}

	var flowConfigurationJSON basetypes.StringValue

	if unknownFlowConfigPlan {

		flowConfigurationJSON = types.StringUnknown()

		resp.Plan.SetAttribute(ctx, path.Root("flow_variables"), types.SetUnknown(types.ObjectType{AttrTypes: flowVariablesTFObjectTypes}))

	} else {

		// Flow configuration
		jsonFlowConfigBytes, err := davinci.Marshal(flowConfigObject, flowConfigurationJsonCmpOptsConfiguration)
		if err != nil {
			resp.Diagnostics.AddError(
				"Error converting the flow object to JSON",
				fmt.Sprintf("Error converting the flow object (from the API response) to JSON in plan modification.  This is a bug in the provider, please report this to the provider maintainers. Error: %s", err),
			)
			return
		}

		flowConfigurationJSON = framework.StringToTF(string(jsonFlowConfigBytes[:]))
	}

	resp.Plan.SetAttribute(ctx, path.Root("flow_configuration_json"), flowConfigurationJSON)

	if !req.State.Raw.IsNull() {
		var state FlowResourceModel

		// Read Terraform plan data into the model
		resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
		if resp.Diagnostics.HasError() {
			return
		}

		// Read Terraform state data into the model
		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
		if resp.Diagnostics.HasError() {
			return
		}

		if !plan.FlowJSON.Equal(state.FlowJSON) || !plan.FlowConfigurationJSON.Equal(state.FlowConfigurationJSON) || !plan.FlowVariables.Equal(state.FlowVariables) {
			resp.Plan.SetAttribute(ctx, path.Root("flow_export_json"), types.StringUnknown())
		}
	}

}

func (p *FlowResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config FlowResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(validateConnectionSubflowLinkMappings(ctx, config.FlowJSON, config.ConnectionLinks, config.SubFlowLinks, true)...)

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

	resp.Diagnostics.Append(plan.validate(ctx)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Build the model for the API
	daVinciImport, d := plan.expand(ctx)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	environmentID := plan.EnvironmentId.ValueString()

	createResponse, err := sdk.DoRetryable(
		ctx,
		r.Client,
		environmentID,
		func() (any, *http.Response, error) {
			return r.Client.CreateFlowWithResponse(environmentID, *daVinciImport)
		},
	)
	if err != nil {

		additionalText := ""

		if dvError, ok := err.(davinci.ErrorResponse); ok {

			additionalText = "This error may indicate the flow is not fully configured."

			if dvError.Code == davinci.DV_ERROR_CODE_ERROR_CREATING_CONNECTOR {
				additionalText = "This error may indicate that a node in the flow is not fully configured."
			}
		}

		resp.Diagnostics.AddError(
			"Error importing flow",
			fmt.Sprintf("Error creating flow: %s. %s", err, additionalText),
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	createFlow, ok := createResponse.(*davinci.Flow)
	if !ok || createFlow.Name == "" {
		resp.Diagnostics.AddError(
			"Unexpected response",
			"Unable to parse create response from Davinci API on flow",
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	// do an update for the settings
	flowUpdate := davinci.FlowUpdate{
		FlowUpdateConfiguration: daVinciImport.FlowInfo.FlowUpdateConfiguration,
		CurrentVersion:          createFlow.CurrentVersion,
		Name:                    &createFlow.Name,
		Description:             createFlow.Description,
	}

	resp.Diagnostics.Append(r.updateFlow(ctx, environmentID, createFlow.FlowID, &flowUpdate, true)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Do an export for state
	// Run the API call
	sdkRes, err := sdk.DoRetryable(
		ctx,
		r.Client,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return r.Client.ReadFlowVersionOptionalVariableWithResponse(environmentID, createFlow.FlowID, nil, false)
		},
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading flow",
			fmt.Sprintf("Error reading flow: %s", err),
		)
		return
	}

	response, ok := sdkRes.(*davinci.FlowInfo)
	if !ok || response.Flow.Name == "" {
		resp.Diagnostics.AddError(
			"Unexpected response",
			"Unable to parse create export response from Davinci API on flow",
		)
		return
	}

	// Create the state to save
	state = plan

	// Save updated data into Terraform state
	resp.Diagnostics.Append(state.toState(&response.Flow)...)
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
			return r.Client.ReadFlowVersionOptionalVariableWithResponse(environmentID, flowID, nil, false)
		},
	)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading flow",
			fmt.Sprintf("Error reading flow: %s", err),
		)
	}
	if resp.Diagnostics.HasError() {
		return
	}

	response, ok := sdkRes.(*davinci.FlowInfo)
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

	// Read Terraform state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	environmentID := plan.EnvironmentId.ValueString()
	flowID := plan.Id.ValueString()

	if !plan.FlowConfigurationJSON.Equal(state.FlowConfigurationJSON) {
		daVinciUpdate, d := plan.expandUpdate(state)
		resp.Diagnostics.Append(d...)
		if resp.Diagnostics.HasError() {
			return
		}

		resp.Diagnostics.Append(r.updateFlow(ctx, environmentID, flowID, daVinciUpdate, true)...)
		if resp.Diagnostics.HasError() {
			return
		}
	} else if !plan.Description.Equal(state.Description) || !plan.Name.Equal(state.Name) {
		daVinciUpdate := &davinci.FlowUpdate{
			Name:        plan.Name.ValueStringPointer(),
			Description: plan.Description.ValueStringPointer(),
		}

		resp.Diagnostics.Append(r.updateFlow(ctx, environmentID, flowID, daVinciUpdate, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	// Variables
	if !plan.FlowVariables.Equal(state.FlowVariables) {
		var flowVarPlan, flowVarState []FlowVariableResourceModel
		resp.Diagnostics.Append(plan.FlowVariables.ElementsAs(ctx, &flowVarPlan, false)...)
		resp.Diagnostics.Append(state.FlowVariables.ElementsAs(ctx, &flowVarState, false)...)
		// If there are errors, keep going as it's a read to state left

		if !resp.Diagnostics.HasError() {
			for _, flowVar := range flowVarPlan {
				var flowVarStateFound bool
				for _, flowVarState := range flowVarState {
					if flowVar.Id.Equal(flowVarState.Id) {
						flowVarStateFound = true

						// test for modifications
						if !cmp.Equal(flowVar, flowVarState) {
							flowVariable := flowVar.expand()

							_, err := sdk.DoRetryable(
								ctx,
								r.Client,
								environmentID,
								func() (interface{}, *http.Response, error) {
									return r.Client.UpdateVariableWithResponse(environmentID, flowVariable)
								},
							)
							if err != nil {
								resp.Diagnostics.AddError(
									fmt.Sprintf("Error adding %s variable", flowVariable.Context),
									fmt.Sprintf("Error adding %s variable %s as part of flow update: %s", flowVariable.Context, *flowVariable.Name, err),
								)
							}
						}
					}
				}

				if !flowVarStateFound {
					// add the new variable
					flowVariable := flowVar.expand()

					_, err := sdk.DoRetryable(
						ctx,
						r.Client,
						environmentID,
						func() (interface{}, *http.Response, error) {
							return r.Client.CreateVariableWithResponse(environmentID, flowVariable)
						},
					)
					if err != nil {
						// if error starts with "Record already exists" then ignore it
						if !strings.HasPrefix(err.Error(), "Record already exists") {
							resp.Diagnostics.AddError(
								fmt.Sprintf("Error adding %s variable", flowVariable.Context),
								fmt.Sprintf("Error adding %s variable %s as part of flow update: %s", flowVariable.Context, *flowVariable.Name, err),
							)
						}
					}
				}
			}

			for _, flowVar := range flowVarState {
				var flowVarPlanFound bool
				for _, flowVarPlan := range flowVarPlan {
					if flowVar.Id.Equal(flowVarPlan.Id) {
						flowVarPlanFound = true
						break
					}
				}

				if !flowVarPlanFound && flowVar.Context.Equal(types.StringValue(contextFlow)) {
					// remove the variable
					_, err := sdk.DoRetryable(
						ctx,
						r.Client,
						environmentID,
						func() (interface{}, *http.Response, error) {
							return r.Client.DeleteVariableWithResponse(environmentID, flowVar.Id.ValueString())
						},
					)
					if err != nil {
						if !strings.HasPrefix(err.Error(), "Error deleting record") {
							resp.Diagnostics.AddError(
								fmt.Sprintf("Error removing %s variable", flowVar.Context.ValueString()),
								fmt.Sprintf("Error removing %s variable %s as part of flow update: %s", flowVar.Context.ValueString(), flowVar.Name.ValueString(), err),
							)
						}
					}
				}
			}
		}
	}

	// Do an export for state
	// Run the API call
	sdkRes, err := sdk.DoRetryable(
		ctx,
		r.Client,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return r.Client.ReadFlowVersionOptionalVariableWithResponse(environmentID, flowID, nil, false)
		},
	)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error reading flow",
			fmt.Sprintf("Error reading flow: %s", err),
		)
		return
	}

	response, ok := sdkRes.(*davinci.FlowInfo)
	if !ok || response.Flow.Name == "" {
		resp.Diagnostics.AddError(
			"Unexpected response",
			"Unable to parse update export response from Davinci API on flow",
		)
		return
	}

	// Create the state to save
	state = plan

	// Save updated data into Terraform state
	resp.Diagnostics.Append(state.toState(&response.Flow)...)
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *FlowResource) updateFlow(ctx context.Context, environmentID, flowID string, daVinciUpdate *davinci.FlowUpdate, deploy bool) diag.Diagnostics {
	var diags diag.Diagnostics

	_, err := sdk.DoRetryable(
		ctx,
		r.Client,
		environmentID,
		func() (any, *http.Response, error) {
			return r.Client.UpdateFlowWithResponse(environmentID, flowID, *daVinciUpdate)
		},
	)
	if err != nil {
		diags.AddError(
			"Error importing flow",
			fmt.Sprintf("Error updating flow: %s", err),
		)
	}
	if diags.HasError() {
		return diags
	}

	if deploy {
		_, err = sdk.DoRetryable(
			ctx,
			r.Client,
			environmentID,
			func() (any, *http.Response, error) {
				return r.Client.DeployFlowWithResponse(environmentID, flowID)
			},
		)
		if err != nil {
			diags.AddError(
				"Error deploying flow",
				fmt.Sprintf("Error deploying flow, this might indicate a misconfiguration of the flow, or an unmapped node connection: %s", err),
			)
		}
	}

	return diags
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
	res, ok := sdkRes.(*davinci.Message)
	if !ok || res.Message == nil || *res.Message == "" {
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
			Label:     "davinci_flow_id",
			Regexp:    verify.P1DVResourceIDRegexp,
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

func (p *FlowResourceModel) validate(ctx context.Context) (diags diag.Diagnostics) {

	diags.Append(validateConnectionSubflowLinkMappings(ctx, p.FlowJSON, p.ConnectionLinks, p.SubFlowLinks, false)...)

	return diags
}

func (p *FlowResourceModel) expand(ctx context.Context) (*davinci.FlowImport, diag.Diagnostics) {
	var diags diag.Diagnostics

	data := davinci.FlowImport{
		Name:        p.Name.ValueStringPointer(),
		Description: p.Description.ValueStringPointer(),
	}

	err := davinci.Unmarshal([]byte(p.FlowJSON.ValueString()), &data.FlowInfo, davinci.ExportCmpOpts{})
	if err != nil {
		diags.AddError(
			"Error parsing `flow_json`",
			fmt.Sprintf("Error parsing `flow_json`: %s", err),
		)
		return nil, diags
	}

	if data.Name == nil {
		data.Name = &data.FlowInfo.Name
	}

	if data.Description == nil && data.FlowInfo.Description != nil {
		data.Description = data.FlowInfo.Description
	}

	data.FlowNameMapping = map[string]string{
		data.FlowInfo.FlowID: p.Name.ValueString(),
	}

	var flowConfigObject davinci.FlowConfiguration
	if !p.FlowConfigurationJSON.IsUnknown() && !p.FlowConfigurationJSON.IsNull() {

		err = davinci.Unmarshal([]byte(p.FlowConfigurationJSON.ValueString()), &flowConfigObject, davinci.ExportCmpOpts{})
		if err != nil {
			diags.AddError(
				"Error parsing `flow_configuration_json`",
				fmt.Sprintf("Error parsing `flow_configuration_json` into flow configuration object: %s", err),
			)
			return nil, diags
		}

		data.FlowInfo.FlowConfiguration = flowConfigObject

	} else {

		// Connection and subflow links
		err = davinci.Unmarshal([]byte(p.FlowJSON.ValueString()), &flowConfigObject, davinci.ExportCmpOpts{})
		if err != nil {
			diags.AddError(
				"Error parsing `flow_json`",
				fmt.Sprintf("Error parsing `flow_json` into flow configuration object: %s", err),
			)
			return nil, diags
		}

		unknownFlowConfigPlan, d := modifyPlanForConnectionSubflowLinkMappings(ctx, &flowConfigObject, p.ConnectionLinks, p.SubFlowLinks)
		diags.Append(d...)
		if diags.HasError() {
			return nil, diags
		}

		if unknownFlowConfigPlan {

			diags.AddError(
				"Unknown Flow Import",
				"The `flow_configuration_json` parameter is unknown.  Cannot complete the plan calculation.",
			)
			return nil, diags

		}

		jsonFlowConfigBytes, err := davinci.Marshal(flowConfigObject, flowConfigurationJsonCmpOptsConfiguration)
		if err != nil {
			diags.AddError(
				"Error converting the flow object to JSON",
				fmt.Sprintf("Error converting the flow object (from the API response) to JSON in API model parsing.  This is a bug in the provider, please report this to the provider maintainers. Error: %s", err),
			)
			return nil, diags
		}

		err = davinci.Unmarshal(jsonFlowConfigBytes, &data.FlowInfo.FlowConfiguration, flowConfigurationJsonCmpOptsConfiguration)
		if err != nil {
			diags.AddError(
				"Error parsing `flow_configuration_json`",
				fmt.Sprintf("Error parsing `flow_configuration_json`: %s", err),
			)
			return nil, diags
		}
	}

	return &data, diags
}

func (p *FlowResourceModel) expandUpdate(state FlowResourceModel) (*davinci.FlowUpdate, diag.Diagnostics) {
	var diags diag.Diagnostics

	var data, stateData davinci.FlowUpdate

	err := davinci.Unmarshal([]byte(p.FlowConfigurationJSON.ValueString()), &data, davinci.ExportCmpOpts{})
	if err != nil {
		diags.AddError(
			"Error parsing `flow_json`",
			fmt.Sprintf("Error parsing `flow_json`: %s", err),
		)
		return nil, diags
	}

	if !p.Name.IsNull() {
		data.Name = p.Name.ValueStringPointer()
	}

	if !p.Description.IsNull() {
		data.Description = p.Description.ValueStringPointer()
	}

	err = davinci.Unmarshal([]byte(state.FlowConfigurationJSON.ValueString()), &stateData, davinci.ExportCmpOpts{})
	if err != nil {
		diags.AddError(
			"Error parsing `flow_configuration_json`",
			fmt.Sprintf("Error parsing `flow_configuration_json`: %s", err),
		)
		return nil, diags
	}

	diags.Append(modifyPlanForMergedProperties(&data.FlowUpdateConfiguration, stateData.FlowUpdateConfiguration)...)
	if diags.HasError() {
		return nil, diags
	}

	return &data, diags
}

func (p *FlowVariableResourceModel) expand() *davinci.VariablePayload {

	mutableValue := true

	data := davinci.VariablePayload{
		Context: p.Context.ValueString(),
		Type:    p.Type.ValueString(),
		Mutable: &mutableValue,
	}

	if !p.Name.IsNull() && !p.Name.IsUnknown() {
		data.Name = p.Name.ValueStringPointer()
	}

	if !p.FlowId.IsNull() && !p.FlowId.IsUnknown() {
		data.FlowId = p.FlowId.ValueStringPointer()
	}

	return &data
}

func (p *FlowResourceModel) toState(apiObject *davinci.Flow) diag.Diagnostics {
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

	jsonConfigurationBytes, err := davinci.Marshal(apiObject.FlowConfiguration, flowConfigurationJsonCmpOptsConfiguration)
	if err != nil {
		diags.AddError(
			"Error converting the flow object configuration to JSON",
			fmt.Sprintf("Error converting the flow object configuration (from the API response) to JSON.  This is a bug in the provider, please report this to the provider maintainers. Error: %s", err),
		)
		return diags
	}

	p.FlowConfigurationJSON = framework.DaVinciExportTypeToTF(string(jsonConfigurationBytes[:]), flowConfigurationJsonCmpOptsConfiguration)

	jsonBytes, err := json.Marshal(apiObject)
	if err != nil {
		diags.AddError(
			"Error converting the flow object to JSON",
			fmt.Sprintf("Error converting the flow object (from the API response) to JSON in recording state.  This is a bug in the provider, please report this to the provider maintainers. Error: %s", err),
		)
		return diags
	}

	p.FlowExportJSON = framework.DaVinciExportTypeToTF(string(jsonBytes[:]), flowExportJsonCmpOptsConfiguration)

	if p.FlowJSON.IsNull() {
		p.FlowJSON = framework.DaVinciExportTypeToTF(string(jsonBytes[:]), flowJsonCmpOptsConfiguration)
	}

	if apiObject.DeployedDate != nil {
		p.Deploy = types.BoolValue(true)
	} else {
		p.Deploy = types.BoolValue(false)
	}

	p.Name = framework.StringToTF(apiObject.Name)

	if v := apiObject.Description; v != nil {
		p.Description = framework.StringToTF(*v)
	}

	var d diag.Diagnostics
	p.FlowVariables, d = flowVariablesToTF(apiObject.Variables)
	diags.Append(d...)

	return diags
}

func flowVariablesToTF(apiObject []davinci.FlowVariable) (types.Set, diag.Diagnostics) {
	var diags diag.Diagnostics

	tfObjType := types.ObjectType{AttrTypes: flowVariablesTFObjectTypes}

	if apiObject == nil {
		return types.SetNull(tfObjType), diags
	}

	variablesSet := make([]attr.Value, 0)

	for _, variable := range apiObject {

		varStateSimpleName := strings.Split(variable.Name, "##SK##")
		if varStateSimpleName[0] == "" || len(varStateSimpleName) == 1 {

			diags.AddError(
				"Error parsing variable name",
				fmt.Sprintf("Unable to parse variable name: %s for state", variable.Name),
			)

			return types.SetNull(tfObjType), diags
		}

		attributesMap := map[string]attr.Value{
			"id":      framework.StringToTF(variable.Name),
			"name":    framework.StringToTF(varStateSimpleName[0]),
			"flow_id": framework.StringOkToTF(variable.FlowID, true),
			"context": framework.StringOkToTF(variable.Context, true),
			"type":    framework.StringOkToTF(variable.Fields.Type, true),
		}

		flattenedObj, d := types.ObjectValue(flowVariablesTFObjectTypes, attributesMap)
		diags.Append(d...)

		variablesSet = append(variablesSet, flattenedObj)

	}

	returnVar, d := types.SetValue(tfObjType, variablesSet)
	diags.Append(d...)

	return returnVar, diags
}

type flowConnectionNodeModel struct {
	connectionID        string
	connectorID         string
	id                  string
	name                string
	nodeType            string
	subFlowIDValueLabel string
	subFlowIDValueValue string
	hasSubflowRef       bool
}

func parseFlowNodeProperties(node davinci.Node) flowConnectionNodeModel {

	returnVar := flowConnectionNodeModel{
		hasSubflowRef: false,
	}

	if nd := node.Data; nd != nil {

		if v := nd.ConnectionID; v != nil {
			returnVar.connectionID = *v
		}

		if v := nd.ConnectorID; v != nil {
			returnVar.connectorID = *v
		}

		if v := nd.ID; v != nil {
			returnVar.id = *v
		}

		if v := nd.Name; v != nil {
			returnVar.name = *v
		}

		if v := nd.NodeType; v != nil {
			returnVar.nodeType = *v
		}

		if ndp := nd.Properties; ndp != nil {
			if ndps := ndp.SubFlowID; ndps != nil {
				if ndpsv := ndps.Value; ndpsv != nil {

					returnVar.hasSubflowRef = true

					if v := ndpsv.Label; v != nil {
						returnVar.subFlowIDValueLabel = *v
					}

					if v := ndpsv.Value; v != nil {
						returnVar.subFlowIDValueValue = *v
					}

				}
			}
		}
	}

	return returnVar
}

// Validate if there are connections in the flow that should have a connection mapping, and flow connector instances that should have a subflow mapping
func validateConnectionSubflowLinkMappings(ctx context.Context, flowJSON davinciexporttype.ParsedValue, connectionLinks basetypes.SetValue, subFlowLinks basetypes.SetValue, allowUnknownValues bool) (diags diag.Diagnostics) {

	if !flowJSON.IsUnknown() && !connectionLinks.IsUnknown() && !subFlowLinks.IsUnknown() {

		var flowConfigObject davinci.FlowConfiguration
		err := davinci.Unmarshal([]byte(flowJSON.ValueString()), &flowConfigObject, davinci.ExportCmpOpts{})
		if err != nil {
			diags.AddError(
				"Error parsing `flow_json`",
				fmt.Sprintf("Error parsing `flow_json` into flow configuration object: %s", err),
			)
			return
		}

		var connectionLinksPlan []FlowConnectionLinkResourceModel
		diags.Append(connectionLinks.ElementsAs(ctx, &connectionLinksPlan, false)...)

		var subflowLinksPlan []FlowSubflowLinkResourceModel
		diags.Append(subFlowLinks.ElementsAs(ctx, &subflowLinksPlan, false)...)

		if flowConfigObject.GraphData.Elements != nil && flowConfigObject.GraphData.Elements.Nodes != nil && len(flowConfigObject.GraphData.Elements.Nodes) > 0 {

			for _, node := range flowConfigObject.GraphData.Elements.Nodes {

				nodeObject := parseFlowNodeProperties(node)

				if nodeObject.nodeType == "CONNECTION" || nodeObject.connectorID != "" {

					connectionLinkFound := false

					// Validate the connection Link mapping
					for _, connectionLinkPlan := range connectionLinksPlan {

						if !allowUnknownValues && (connectionLinkPlan.ReplaceImportConnectionId.IsUnknown() || connectionLinkPlan.Name.IsUnknown() || connectionLinkPlan.Id.IsUnknown()) {
							diags.AddAttributeError(
								path.Root("connection_link"),
								"Unknown Connection Links",
								"One of `connection_link.replace_connection_id`, `connection_link.name`, `connection_link.id` is unknown.  Cannot validate the connection link mappings.",
							)

							return diags
						}

						if !connectionLinkPlan.ReplaceImportConnectionId.IsNull() && connectionLinkPlan.ReplaceImportConnectionId.ValueString() == nodeObject.connectionID {
							connectionLinkFound = true
						}

						if connectionLinkPlan.Name.ValueString() == nodeObject.name {
							connectionLinkFound = true
						}

						if connectionLinkPlan.ReplaceImportConnectionId.IsUnknown() || connectionLinkPlan.Name.IsUnknown() {
							connectionLinkFound = true // defer this validation to the plan step
						}

					}

					if !connectionLinkFound {
						diags.AddAttributeWarning(
							path.Root("connection_link"),
							"Unmapped node connection",
							fmt.Sprintf("The flow JSON to import (provided in the `flow_json` parameter) contains a node connection that does not have a `connection_link` mapping.  This behaviour is deprecated - going forward all connections in a flow must have a `connection_link` block parameter defined.\n\n"+
								"Consider using the `davinci_connection` resource (to create a Terraform managed connection) with the `davinci_flow.connection_link` parameter (to map the Terraform managed connection with the connection in the flow).\n"+
								"For more information and guidance on how to correctly configure flow connections, visit https://github.com/pingidentity/terraform-provider-davinci/issues/272\n\n"+
								"Connection ID: %s\nConnector ID: %s\nConnection Name: %s\nNode Type: %s\nNode ID: %s", nodeObject.connectionID, nodeObject.connectorID, nodeObject.name, nodeObject.nodeType, nodeObject.id),
						)
					}

					// Validate the subflow link mapping if necessary
					if nodeObject.connectorID == "flowConnector" && nodeObject.hasSubflowRef {
						subflowLinkFound := false

						for _, subflowLinkPlan := range subflowLinksPlan {

							if !allowUnknownValues && (subflowLinkPlan.ReplaceImportSubflowId.IsUnknown() || subflowLinkPlan.Name.IsUnknown() || subflowLinkPlan.Id.IsUnknown()) {
								diags.AddAttributeError(
									path.Root("subflow_link"),
									"Unknown SubFlow Links",
									"One of `subflow_link.replace_subflow_id`, `subflow_link.name`, `subflow_link.id` is unknown.  Cannot validate the flow connector subflow link mappings.",
								)

								return diags
							}

							if !subflowLinkPlan.ReplaceImportSubflowId.IsNull() && subflowLinkPlan.ReplaceImportSubflowId.ValueString() == nodeObject.subFlowIDValueValue {
								subflowLinkFound = true
							}

							if subflowLinkPlan.Name.ValueString() == nodeObject.subFlowIDValueLabel {
								subflowLinkFound = true
							}

							if subflowLinkPlan.ReplaceImportSubflowId.IsUnknown() && subflowLinkPlan.Name.IsUnknown() {
								subflowLinkFound = true // defer this validation to the plan step
							}
						}

						if !subflowLinkFound {
							diags.AddAttributeWarning(
								path.Root("subflow_link"),
								"Unmapped flow connector subflow",
								fmt.Sprintf("The flow JSON to import (provided in the `flow_json` parameter) contains a subflow referenced in a flow connector that does not have a `subflow_link` mapping.  This behaviour is deprecated - going forward all subflows defined in a flow must have a `subflow_link` block parameter defined.\n\n"+
									"Consider using the `davinci_flow` resource (to create the subflow) with the `davinci_flow.subflow_link` parameter (to map the Terraform managed subflow with the main flow).\n"+
									"For more information and guidance on how to correctly configure subflow connections, visit https://github.com/pingidentity/terraform-provider-davinci/issues/273\n\n"+
									"Connection ID: %s\nConnector ID: %s\nConnection Name: %s\nNode Type: %s\nNode ID: %s\nSubflow Name: %s\nSubflow ID: %s", nodeObject.connectionID, nodeObject.connectorID, nodeObject.name, nodeObject.nodeType, nodeObject.id, nodeObject.subFlowIDValueLabel, nodeObject.subFlowIDValueValue),
							)
						}
					}
				}
			}
		}
	}

	if !allowUnknownValues && flowJSON.IsUnknown() {
		diags.AddAttributeError(
			path.Root("flow_json"),
			"Unknown Flow Import",
			"The `flow_json` parameter is unknown.  Cannot validate the connection link mappings.",
		)
	}

	if !allowUnknownValues && connectionLinks.IsUnknown() {
		diags.AddAttributeError(
			path.Root("connection_link"),
			"Unknown Connection Links",
			"The `connection_link` parameter is unknown.  Cannot validate the connection link mappings.",
		)
	}

	if !allowUnknownValues && subFlowLinks.IsUnknown() {
		diags.AddAttributeError(
			path.Root("subflow_link"),
			"Unknown Subflow Links",
			"The `subflow_link` parameter is unknown.  Cannot validate the flow connector subflow link mappings.",
		)
	}

	return diags
}

// Modify the plan for connector and subflow re-mapping
func modifyPlanForConnectionSubflowLinkMappings(ctx context.Context, flowConfigObject *davinci.FlowConfiguration, connectionLinks basetypes.SetValue, subflowLinks basetypes.SetValue) (unknownFlowConfigPlan bool, diags diag.Diagnostics) {

	// Update connectors if we know the config of connection links
	unknownFlowConfigPlan = connectionLinks.IsUnknown() || subflowLinks.IsUnknown()

	if !unknownFlowConfigPlan && flowConfigObject.GraphData.Elements != nil && flowConfigObject.GraphData.Elements.Nodes != nil && len(flowConfigObject.GraphData.Elements.Nodes) > 0 {

		var connectionLinksPlan []FlowConnectionLinkResourceModel
		diags.Append(connectionLinks.ElementsAs(ctx, &connectionLinksPlan, false)...)

		var subflowLinksPlan []FlowSubflowLinkResourceModel
		diags.Append(subflowLinks.ElementsAs(ctx, &subflowLinksPlan, false)...)

		newNodes := make([]davinci.Node, 0)
		for _, node := range flowConfigObject.GraphData.Elements.Nodes {

			nodeObject := parseFlowNodeProperties(node)

			if !unknownFlowConfigPlan && (nodeObject.nodeType == "CONNECTION" || nodeObject.connectorID != "") {

				// Find the connection_link reference
				for _, connectionLinkPlan := range connectionLinksPlan {

					if connectionLinkPlan.ReplaceImportConnectionId.IsUnknown() || connectionLinkPlan.Name.IsUnknown() || connectionLinkPlan.Id.IsUnknown() {
						unknownFlowConfigPlan = true
						break
					}

					if node.Data.ConnectionID != nil && !connectionLinkPlan.ReplaceImportConnectionId.IsNull() && connectionLinkPlan.ReplaceImportConnectionId.ValueString() == *node.Data.ConnectionID {
						// If the replace import connection ID is known and not null, we can replace the connection ID in the flowConfigObject

						// replace the ID and name in the flowConfigObject
						connectionID := connectionLinkPlan.Id.ValueString()
						connectionName := connectionLinkPlan.Name.ValueString()
						node.Data.ConnectionID = &connectionID
						node.Data.Name = &connectionName

					} else if connectionLinkPlan.Name.ValueString() == nodeObject.name {
						// If we're here, the replace import connection ID is known to be null, so we do name matching

						// replace the ID in the flowConfigObject
						connectionID := connectionLinkPlan.Id.ValueString()
						node.Data.ConnectionID = &connectionID
					}
				}

				if !unknownFlowConfigPlan && nodeObject.connectorID == "flowConnector" && nodeObject.hasSubflowRef {

					// Find the subflow_link reference
					for _, subflowLinkPlan := range subflowLinksPlan {

						if subflowLinkPlan.ReplaceImportSubflowId.IsUnknown() || subflowLinkPlan.Name.IsUnknown() || subflowLinkPlan.Id.IsUnknown() {
							unknownFlowConfigPlan = true
							break
						}

						// If the replace import connection ID is known and not null, we can replace the connection ID in the flowConfigObject
						if !subflowLinkPlan.ReplaceImportSubflowId.IsNull() && subflowLinkPlan.ReplaceImportSubflowId.ValueString() == nodeObject.subFlowIDValueValue {

							// replace the ID and label
							subflowPlanValue := subflowLinkPlan.Id.ValueString()
							subflowPlanLabel := subflowLinkPlan.Name.ValueString()
							node.Data.Properties.SubFlowID.Value.Value = &subflowPlanValue
							node.Data.Properties.SubFlowID.Value.Label = &subflowPlanLabel

						} else if subflowLinkPlan.Name.ValueString() == nodeObject.subFlowIDValueLabel {
							// If we're here, the replace import connection ID is known to be null, so we do name matching

							// replace the ID in the flowConfigObject
							subflowPlanValue := subflowLinkPlan.Id.ValueString()
							node.Data.Properties.SubFlowID.Value.Value = &subflowPlanValue
						}
					}
				}
			}

			newNodes = append(newNodes, node)
		}

		flowConfigObject.GraphData.Elements.Nodes = newNodes
	}

	return unknownFlowConfigPlan, diags
}

func modifyPlanForMergedProperties(flowConfigObject *davinci.FlowUpdateConfiguration, stateFlowConfigObject davinci.FlowUpdateConfiguration) (diags diag.Diagnostics) {

	// Must merge settings back into flow import if it's missing
	if stateFlowConfigObject.Settings != nil && flowConfigObject.Settings == nil {
		flowConfigObject.Settings = stateFlowConfigObject.Settings

		diags.AddAttributeWarning(
			path.Root("flow_configuration_json"),
			"Implicitly merged settings",
			"The `flow_json` import does not contain flow settings, but the existing state does.  To maintain consistency with the DaVinci service, the settings from the Terraform state have been merged into the import.\n\nIf this is not the desired behaviour, please update the DaVinci flow in the source environment to include the desired flow settings and re-export.",
		)
	}

	return diags

}
