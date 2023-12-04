package davinci

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	"github.com/pingidentity/terraform-provider-davinci/internal/utils"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

func ResourceFlow() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceFlowCreate,
		ReadContext:   resourceFlowRead,
		UpdateContext: resourceFlowUpdate,
		DeleteContext: resourceFlowDelete,
		Schema: map[string]*schema.Schema{
			"flow_json": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "DaVinci Flow in raw json format.",
				DiffSuppressFunc: computeFlowDrift,
				Sensitive:        true,
			},
			"deploy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Deploy Flow after import. Flows must be deployed to be used.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Computed Flow Name after import. Matches 'name' in flow_json",
			},
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PingOne Environment to import flow into.",
			},
			"connection_link": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "Connections this flow depends on. flow_json connectionId will be updated to id matching name .",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Connection ID that will be used when flow is imported.",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Connection Name to match when updating flow_json connectionId.",
						},
					},
				},
			},
			"subflow_link": {
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Description: "Child flows of this resource. Required to keep mapping if flow_json contains subflows. flow_json subflowId will be updated to id matching name. Note, subflow will automatically point to latest version (-1).",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Subflow Flow ID that will be used when flow is imported.",
						},
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Subflow Name to match when updating flow_json subflowId.",
						},
						//TODO implement subflow version
						// "subflow_version": {
						// 	Type:        schema.TypeString,
						// 	Optional: true,
						// 	Computed: true,
						// 	Description: "Subflow Version to use",
						// },
					},
				},
			},
			"flow_variables": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Returned list of Flow Context variables. These are Variable resources that are created and managed by the Flow resource via flow_json",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "DaVinci internal name of variable",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Friendly Name of Variable in UI",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of Variable in UI",
						},
						"flow_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Should match id of this davinci_flow",
						},
						"context": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Should always return 'flow'",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Underlying type of variable",
						},
						"mutable": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "If true, the variable can be modified by the flow. If false, the variable is read-only and cannot be modified by the flow.",
						},
						"min": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "",
						},
						"max": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "",
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: resourceFlowImport,
		},
	}
}

func resourceFlowCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	validateFlowDeps(d, &diags)

	var flowJson string
	if fj, ok := d.GetOk("flow_json"); ok {
		// check that all connections and subflows dependencies are met
		flowJson = fj.(string)
		//Update subflows if needed
		subsJson, err := mapSubFlows(d, flowJson)
		if err != nil {
			return diag.FromErr(err)
		}
		//Update connections if needed
		connsJson, err := mapFlowConnections(d, *subsJson)
		if err != nil {
			return diag.FromErr(err)
		}
		flowJson = *connsJson
	} else {
		return diag.FromErr(fmt.Errorf("Error: flow_json not found"))
	}

	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.CreateFlowWithJson(&c.CompanyID, &flowJson)
	}, nil)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error Importing Flow",
			Detail:   fmt.Sprintf(`This may indicate the flow contains unconfigured nodes. Additionally the flow may have been imported as an unmanaged resource and may require manual intervention. API Error: %v`, err),
		})
		return diags
	}
	res, ok := sdkRes.(*dv.Flow)
	if !ok || res.Name == "" {
		err = fmt.Errorf("Unable to parse create response from Davinci API on flow")
		return diag.FromErr(err)
	}

	err = deployIfNeeded(ctx, c, d, res.FlowID)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Error deploying flow",
			Detail:   fmt.Sprintf(`This may indicate flow '%s' contains unconfigured nodes.`, res.Name),
		})
		_, deleteErr := sdk.DoRetryable(ctx, func() (interface{}, error) {
			return c.DeleteFlow(&c.CompanyID, res.FlowID)
		}, nil)
		if deleteErr != nil {
			return diag.FromErr(deleteErr)
		}
		return diags
	}

	d.SetId(res.FlowID)

	resourceFlowRead(ctx, d, meta)

	return diags
}

func resourceFlowRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	flowId := d.Id()

	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.ReadFlowVersion(&c.CompanyID, flowId, nil)
	}, nil)
	if err != nil {
		httpErr, _ := dv.ParseDvHttpError(err)
		if strings.Contains(httpErr.Body, "Error retrieving flow version") {
			d.SetId("")
			return diags
		}
		return diag.FromErr(err)
	}
	res, ok := sdkRes.(*dv.FlowInfo)
	if !ok {
		err = fmt.Errorf("Unable to cast FlowInfo type to response from Davinci API on flow id: %v", flowId)
		return diag.FromErr(err)
	}

	if res.Flow.FlowID == "" {
		d.SetId("")
		return diags
	}

	d.SetId(res.Flow.FlowID)

	if err := d.Set("name", res.Flow.Name); err != nil {
		return diag.FromErr(err)
	}
	rString, err := json.Marshal(&res.Flow)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("flow_json", string(rString)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("environment_id", string(res.Flow.CompanyID)); err != nil {
		return diag.FromErr(err)
	}

	// Set subflows, this is important for terraform import
	subflows, err := flattenSubflows(res.Flow)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("subflow_link", subflows); err != nil {
		return diag.FromErr(err)
	}

	// Set connections, this is important for terraform import
	connections, err := flattenConnections(res.Flow)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("connection_link", connections); err != nil {
		return diag.FromErr(err)
	}

	// Set flow_variables, this is important for terraform import)
	flowVariables, err := flattenFlowVariables(res.Flow)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("flow_variables", flowVariables); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func flattenSubflows(flow dv.Flow) ([]interface{}, error) {
	var subflows []interface{}
	for _, v := range flow.GraphData.Elements.Nodes {
		if v.Data.ConnectorID == "flowConnector" && (v.Data.CapabilityName == "startSubFlow" || v.Data.CapabilityName == "startUiSubFlow") {
			sfProp, err := expandSubFlowProps(v.Data.Properties)
			if err != nil {
				return nil, err
			}
			subflowMap := map[string]interface{}{
				"id":   sfProp.SubFlowID.Value.Value,
				"name": sfProp.SubFlowID.Value.Label,
			}
			subflows = append(subflows, subflowMap)
		}
	}
	return subflows, nil
}

func flattenConnections(flow dv.Flow) ([]interface{}, error) {
	var connections []interface{}
	for _, node := range flow.GraphData.Elements.Nodes {
		if node.Data.ConnectionID != "" && node.Data.Name != "" {
			connectionMap := map[string]interface{}{
				"id":   node.Data.ConnectionID,
				"name": node.Data.Name,
			}
			connections = append(connections, connectionMap)
		}
	}
	return connections, nil
}

func flattenFlowVariables(flow dv.Flow) ([]interface{}, error) {
	flowJson, err := json.Marshal(flow)
	if err != nil {
		return nil, err
	}
	flowVars, err := getFlowVariables(string(flowJson))
	if err != nil {
		return nil, err
	}

	var flowVariables []interface{}
	for _, flowVar := range flowVars {
		varStateSimpleName := strings.Split(flowVar.Name, "##SK##")
		if varStateSimpleName[0] == "" || len(varStateSimpleName) == 1 {
			return nil, fmt.Errorf("Unable to parse variable name: %s for state", flowVar.Name)
		}
		flowVariableMap := map[string]interface{}{
			"id":          flowVar.Name,
			"name":        varStateSimpleName[0],
			"description": flowVar.Fields.DisplayName,
			"flow_id":     flowVar.FlowID,
			"context":     flowVar.Context,
			"type":        flowVar.Type,
			"mutable":     flowVar.Fields.Mutable,
			"min":         flowVar.Fields.Min,
			"max":         flowVar.Fields.Max,
		}
		flowVariables = append(flowVariables, flowVariableMap)
	}
	return flowVariables, nil
}

func resourceFlowUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	flowId := d.Id()

	// If changes are detected:
	// 1. Map subflows
	// 2. Map connections
	// 3. Update flow via API
	// 4. Update flow variables via variables api
	// 5. Deploy flow if needed
	if d.HasChanges("flow_json", "connection_link", "subflow_link") {
		flowJson := d.Get("flow_json").(string)
		subsJson, err := mapSubFlows(d, flowJson)
		if err != nil {
			return diag.FromErr(err)
		}
		connsJson, err := mapFlowConnections(d, *subsJson)
		if err != nil {
			return diag.FromErr(err)
		}
		flowJson = *connsJson

		flowVars, err := getFlowVariables(flowJson)
		if err != nil {
			return diag.FromErr(err)
		}

		sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
			return c.UpdateFlowWithJson(&c.CompanyID, &flowJson, flowId)
		}, nil)
		if err != nil {
			return diag.FromErr(err)
		}

		// Update Flow Variables
		// this will ONLY be flow variables because getFlowVariables removes all other variables contexts
		// If variable exists in state, it should be updated. Else it should be created.
		flowVariablesState, variableStateOk := d.GetOk("flow_variables")
		for _, v := range flowVars {
			variablePayload := dv.VariablePayload{
				// Name:        v.Name,
				Description: v.Fields.DisplayName,
				FlowId:      flowId,
				Context:     v.Context,
				Type:        v.Fields.Type,
				Mutable:     v.Fields.Mutable,
				Min:         v.Fields.Min,
				Max:         v.Fields.Max,
			}
			// Variable Payload is identified by the simple name of the variable because it is suffixed with an unknown unique id.
			varSimpleName := strings.Split(v.Name, "##SK##")
			existsInState := false
			if varSimpleName[0] == "" || len(varSimpleName) == 1 {
				return diag.FromErr(fmt.Errorf("Unable to parse variable name: %s from flow_json", v.Name))
			}
			if variableStateOk {
				for _, stateVar := range flowVariablesState.([]interface{}) {
					stateVarMap := stateVar.(map[string]interface{})
					if stateVarMap["name"].(string) == varSimpleName[0] {
						variablePayload.Name = stateVarMap["name"].(string)
						existsInState = true
						// Update SHOULD be safe because the variable should exist if the flow exists.
						_, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
							return c.UpdateVariable(&c.CompanyID, &variablePayload)
						}, nil)
						if err != nil {
							return diag.FromErr(err)
						}
						break
					}
				}
			}
			if !existsInState {
				variablePayload.Name = varSimpleName[0]

				_, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
					return c.CreateVariable(&c.CompanyID, &variablePayload)
				}, nil)
				if err != nil {
					// In rare scenarios, the variable may exist in the environment but not in state, if so it should update instead.
					httpErr, _ := dv.ParseDvHttpError(err)
					if httpErr != nil && strings.Contains(httpErr.Body, "Record already exists") {
						_, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
							return c.UpdateVariable(&c.CompanyID, &variablePayload)
						}, nil)
						if err != nil {
							return diag.FromErr(err)
						}
					} else {
						return diag.FromErr(err)
					}
				}
			}
		}

		res, ok := sdkRes.(*dv.Flow)
		if !ok || res.Name == "" {
			err = fmt.Errorf("Unable to parse update response from Davinci API on flow id: %v", flowId)
			return diag.FromErr(err)
		}
	}

	if err := deployIfNeeded(ctx, c, d, flowId); err != nil {
		return diag.FromErr(err)
	}

	return resourceFlowRead(ctx, d, meta)
}

func resourceFlowDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	flowId := d.Id()

	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.DeleteFlow(&c.CompanyID, flowId)
	}, nil)

	if err != nil {
		return diag.FromErr(err)
	}
	res, ok := sdkRes.(*dv.Message)
	if !ok || res.Message == "" {
		err = fmt.Errorf("Unable to parse delete response from Davinci API on flow id: %v", flowId)
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func resourceFlowImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	idComponents := []utils.ImportComponent{
		{
			Label:  "environment_id",
			Regexp: regexp.MustCompile(`[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`),
		},
		{
			Label:     "davinci_flow_id",
			Regexp:    regexp.MustCompile(`[a-f0-9]{32}`),
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

	d.SetId(attributes["davinci_flow_id"])

	resourceFlowRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}

func computeFlowDrift(k, old, new string, d *schema.ResourceData) bool {
	var err error
	// exit quickly if new resource
	if old == "" && new != "" {
		return false
	}

	// Apply subflow dependencies if needed.
	if _, ok := d.GetOk("subflow_link"); ok {
		if new != "" {
			newFlowJson, err := mapSubFlows(d, new)
			if err != nil {
				//lintignore:R009
				panic(err)
			}
			new = *newFlowJson
		}
		if old != "" {
			oldFlowJson, err := mapSubFlows(d, old)
			if err != nil {
				//lintignore:R009
				panic(err)
			}
			old = *oldFlowJson
		}
	}

	// Apply connection dependencies if needed.
	if _, ok := d.GetOk("connection_link"); ok {
		if new != "" {
			newFlowJson, err := mapFlowConnections(d, new)
			if err != nil {
				//lintignore:R009
				panic(err)
			}
			new = *newFlowJson
		}
		if old != "" {
			oldFlowJson, err := mapFlowConnections(d, old)
			if err != nil {
				//lintignore:R009
				panic(err)
			}
			old = *oldFlowJson
		}
	}

	//Prepare current and desired inputs for drift detection
	current := dv.Flow{}
	desired := dv.Flow{}
	if old != "" {
		currentFi := dv.FlowImport{}
		currentF := dv.Flow{}
		err = json.Unmarshal([]byte(old), &currentFi)
		if err != nil {
			//lintignore:R009
			panic(err)
		}
		// convert to type Flow if needed
		if currentFi.FlowInfo.Name != "" {
			current = currentFi.FlowInfo
		}
		err = json.Unmarshal([]byte(old), &currentF)
		if err != nil {
			//lintignore:R009
			panic(err)
		}
		if currentF.GraphData.Elements.Nodes != nil {
			current = currentF
		}
	}
	if new != "" {
		desiredFi := dv.FlowImport{}
		desiredF := dv.Flow{}
		err = json.Unmarshal([]byte(new), &desiredFi)
		if err != nil {
			//lintignore:R009
			panic(err)
		}
		// convert to type Flow if needed
		if desiredFi.FlowInfo.Name != "" {
			desired = desiredFi.FlowInfo
		}
		err = json.Unmarshal([]byte(new), &desiredF)
		if err != nil {
			//lintignore:R009
			panic(err)
		}
		if desiredF.GraphData.Elements.Nodes != nil {
			desired = desiredF
		}
	}

	// Check for Settings, inputSchema, outputSchema, name, FlowStatus, trigger drift

	//ignore logLevel
	var cSettings, dSettings map[string]interface{}

	if _, ok := current.Settings.(map[string]interface{}); ok {
		cSettings = current.Settings.(map[string]interface{})
		delete(cSettings, "logLevel")
	}
	if _, ok := desired.Settings.(map[string]interface{}); ok {
		dSettings = desired.Settings.(map[string]interface{})
		delete(dSettings, "logLevel")
	}
	if !reflect.DeepEqual(cSettings, dSettings) {
		return false
	}

	//account for deletion of inputSchema
	if current.InputSchema != nil && desired.InputSchema == nil {
		desired.InputSchema = make([]interface{}, 0)
	}
	if !reflect.DeepEqual(current.InputSchema, desired.InputSchema) {
		return false
	}

	//OutputSchema Diff
	if !reflect.DeepEqual(current.OutputSchemaCompiled, desired.OutputSchemaCompiled) {
		return false
	}

	if !reflect.DeepEqual(current.Trigger, desired.Trigger) {
		return false
	}

	if current.Name != desired.Name {
		return false
	}
	if current.FlowStatus != desired.FlowStatus {
		return false
	}

	// Variables Diff
	dVar := []dv.FlowVariable{}
	for _, v := range desired.Variables {
		if v.Context == "flow" {
			dVar = append(dVar, v)
		}
	}
	desired.Variables = dVar
	cVar := []dv.FlowVariable{}
	for _, v := range current.Variables {
		if v.Context == "flow" {
			cVar = append(cVar, v)
		}
	}
	current.Variables = cVar
	sort.SliceStable(current.Variables, func(i, j int) bool {
		return current.Variables[i].Name < current.Variables[j].Name
	})
	sort.SliceStable(desired.Variables, func(i, j int) bool {
		return desired.Variables[i].Name < desired.Variables[j].Name
	})
	for i, currentV := range current.Variables {
		//Check relevant fields only
		desiredV := desired.Variables[i]
		if currentV.Fields.DisplayName != desiredV.Fields.DisplayName ||
			currentV.Type != desiredV.Type ||
			currentV.Visibility != desiredV.Visibility ||
			currentV.Fields.Mutable != desiredV.Fields.Mutable ||
			currentV.Fields.Max != desiredV.Fields.Max ||
			currentV.Fields.Min != desiredV.Fields.Min {
			return false
		}
	}

	// Overall GraphData Diff
	if !reflect.DeepEqual(current.GraphData, desired.GraphData) {
		cGraph := current.GraphData
		cGraph.Elements.Nodes = nil
		dGraph := desired.GraphData
		dGraph.Elements.Nodes = nil
		// GraphData Diff without Nodes
		if !reflect.DeepEqual(cGraph, dGraph) {
			return false
		}
		currentNodes, err := json.Marshal(current.GraphData.Elements.Nodes)
		if err != nil {
			//lintignore:R009
			panic(err)
		}
		desiredNodes, err := json.Marshal(desired.GraphData.Elements.Nodes)
		if err != nil {
			//lintignore:R009
			panic(err)
		}

		// Nodes Diff - This is mainly to account for json null vs go nil
		if string(currentNodes) != string(desiredNodes) {
			return false
		}
	}

	return true
}

func deployIfNeeded(ctx context.Context, c *dv.APIClient, d *schema.ResourceData, flowId string) error {
	isDeploy := d.Get("deploy").(bool)
	if isDeploy {
		_, err := c.DeployFlow(&c.CompanyID, flowId)
		if err != nil {
			return fmt.Errorf("Possible misconfigured flow: %v", err)
		}
	}
	return nil
}

func mapSubFlows(d *schema.ResourceData, flowJson string) (*string, error) {
	if sf, ok := d.GetOk("subflow_link"); ok {
		fjMap, err := dv.ParseFlowImportJson(&flowJson)
		if err != nil {
			return nil, err
		}
		sfList := sf.(*schema.Set).List()
		for i, v := range fjMap.FlowInfo.GraphData.Elements.Nodes {
			sfProp := &dv.SubFlowProperties{}
			// Only two types of subflow capabilities use the subflowId and subflowVersionId properties
			if v.Data.ConnectorID == "flowConnector" && (v.Data.CapabilityName == "startSubFlow" || v.Data.CapabilityName == "startUiSubFlow") {
				sfProp, err = expandSubFlowProps(v.Data.Properties)
				if err != nil {
					return nil, err
				}
				for _, sfMap := range sfList {
					sfValues := sfMap.(map[string]interface{})
					if sfValues["name"].(string) == sfProp.SubFlowID.Value.Label {
						sfProp.SubFlowID.Value.Value = sfValues["id"].(string)
						//TODO implement subflow version
						// sfProp.SubFlowVersionID.Value = sfValues["subflow_version"].(string)
					}
				}
				err = mapstructure.Decode(sfProp, &v.Data.Properties)
				if err != nil {
					return nil, err
				}
				fjMap.FlowInfo.GraphData.Elements.Nodes[i] = v
			}
		}
		fjByte, err := json.Marshal(fjMap)
		if err != nil {
			return nil, err
		}
		flowJson = string(fjByte)
	}
	return &flowJson, nil
}

func expandSubFlowProps(subflowProps map[string]interface{}) (*dv.SubFlowProperties, error) {

	sfp, ok := subflowProps["subFlowId"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Flow Validation Error: subFlowId key not found in subflow properties")
	}
	sfpVal, ok := sfp["value"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Flow Validation Error: subFlowId value not found in subflow properties")
	}
	sfId := dv.SubFlowID{
		Value: dv.SubFlowValue{
			Value: sfpVal["value"].(string),
			Label: sfpVal["label"].(string),
		},
	}
	subflowVersionId, ok := subflowProps["subFlowVersionId"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("Flow Validation Error: subFlowVersionId not found in subflow properties")
	}
	var sfvidString string
	if subflowVersionId["value"] == nil {
		return nil, fmt.Errorf("Flow Validation Error: subFlowVersionId.value not found in subflow properties")
	}
	switch subflowVersionId["value"].(type) {
	case int:
		sfvidString = strconv.Itoa(subflowVersionId["value"].(int))
	case float64:
		sfvidString = strconv.FormatFloat(subflowVersionId["value"].(float64), 'f', -1, 64)
	case string:
		sfvidString = subflowVersionId["value"].(string)
	default:
		return nil, fmt.Errorf("Flow Validation Error: subflow versionId is not a string or int")
	}

	sfv := dv.SubFlowVersionID{
		Value: sfvidString,
	}
	if sfId.Value.Value == "" || sfv.Value == "" {
		return nil, fmt.Errorf("Flow Validation Error: subflow value or versionId is empty")
	}
	subflow := dv.SubFlowProperties{
		SubFlowID:        sfId,
		SubFlowVersionID: sfv,
	}
	return &subflow, nil
}

func mapFlowConnections(d *schema.ResourceData, flowJson string) (*string, error) {
	if conns, ok := d.GetOk("connection_link"); ok {
		fjMap, err := dv.ParseFlowImportJson(&flowJson)
		if err != nil {
			return nil, err
		}
		connList := conns.(*schema.Set).List()
		for i, v := range fjMap.FlowInfo.GraphData.Elements.Nodes {
			for _, connMap := range connList {
				connValues := connMap.(map[string]interface{})
				if connValues["name"].(string) == v.Data.Name {
					v.Data.ConnectionID = connValues["id"].(string)
				}
			}
			fjMap.FlowInfo.GraphData.Elements.Nodes[i] = v
		}
		fjByte, err := json.Marshal(fjMap)
		if err != nil {
			return nil, err
		}
		flowJson = string(fjByte)
	}
	return &flowJson, nil
}

// validate that each connection and subflow in flow_json has a corresponding dependency
func validateFlowDeps(d *schema.ResourceData, diags *diag.Diagnostics) {
	if flowJson, ok := d.Get("flow_json").(string); ok {
		fjMap, err := dv.ParseFlowImportJson(&flowJson)
		if err != nil {
			*diags = append(*diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Error parsing flow_json",
				Detail:   err.Error(),
			})
			return
		}
		//prepare de-duped map of connections
		flowConns := map[string]struct {
			ConnectionID   string
			ConnectionName string
		}{}
		for _, v := range fjMap.FlowInfo.GraphData.Elements.Nodes {
			if v.Data.Name == "" || v.Data.ConnectionID == "" {
				continue
			}
			flowConns[v.Data.ConnectionID] = struct {
				ConnectionID   string
				ConnectionName string
			}{
				ConnectionID:   v.Data.ConnectionID,
				ConnectionName: v.Data.Name,
			}
			// validate subflows
			if v.Data.ConnectorID == "flowConnector" && (v.Data.CapabilityName == "startSubFlow" || v.Data.CapabilityName == "startUiSubFlow") {
				foundSubflow := false
				sfProp, err := expandSubFlowProps(v.Data.Properties)
				if err != nil {
					*diags = append(*diags, diag.Diagnostic{
						Severity: diag.Error,
						Summary:  "Error Validating flow_json",
						Detail:   err.Error(),
					})
					return
				}
				if subflows, ok := d.GetOk("subflow_link"); ok {
					sfList := subflows.(*schema.Set).List()
					for _, sfMap := range sfList {
						sfValues := sfMap.(map[string]interface{})
						if sfValues["name"].(string) == sfProp.SubFlowID.Value.Label {
							foundSubflow = true
						}
					}
				}
				if !foundSubflow {
					*diags = append(*diags, diag.Diagnostic{
						Severity: diag.Warning,
						Summary:  "Unmapped Subflow Dependency",
						Detail:   fmt.Sprintf("Flow '%s' contains subflow named '%s' which is not defined in dependent subflows. \nThis may lead to incorrect subflow mapping", fjMap.Name, sfProp.SubFlowID.Value.Label),
					})
				}
			}
		}
		// validate connections
		for _, v := range flowConns {
			foundConnection := false
			if conns, ok := d.GetOk("connection_link"); ok {
				connList := conns.(*schema.Set).List()
				for _, connMap := range connList {
					connValues := connMap.(map[string]interface{})
					if connValues["name"].(string) == v.ConnectionName {
						foundConnection = true
					}
				}
			}
			if !foundConnection {
				*diags = append(*diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  "Unmapped Connection Dependency",
					Detail:   fmt.Sprintf("Flow '%s' contains connection '%s' which is not defined in dependent connections. \nThis may lead to autogeneration of unmanaged connections or incorrect connection mapping", fjMap.Name, v.ConnectionName),
				})
			}
		}
	}
}

func getFlowVariables(flowJson string) ([]dv.FlowVariable, error) {
	flowOutput, err := dv.MakeFlowPayload(&flowJson, "Flow")
	if err == nil {
		var flow dv.Flow
		if err = json.Unmarshal([]byte(*flowOutput), &flow); err != nil {
			return nil, err
		}
		return flow.Variables, nil
	}
	flowOutput, err = dv.MakeFlowPayload(&flowJson, "FlowImport")
	if err == nil {
		var flow dv.FlowImport
		if err = json.Unmarshal([]byte(*flowOutput), &flow); err != nil {
			return nil, err
		}
		return flow.FlowInfo.Variables, nil
	}
	return nil, fmt.Errorf("Error: Unable to abstract flow variables from flow_json")
}
