package davinci

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/mitchellh/mapstructure"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
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
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "DaVinci generated identifier after import.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Computed Flow Name after import. Will match 'name' in flow_json",
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
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceFlowCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)
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
		return diag.FromErr(err)
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

	resourceFlowRead(ctx, d, m)

	return diags
}

func resourceFlowRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	flowId := d.Id()

	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.ReadFlow(&c.CompanyID, flowId)
	}, nil)

	if err != nil {
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
	if err := d.Set("id", res.Flow.FlowID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", res.Flow.Name); err != nil {
		return diag.FromErr(err)
	}
	rString, err := json.Marshal(&res.Flow)

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

func resourceFlowUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	flowId := d.Get("id").(string)

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

		sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
			return c.UpdateFlowWithJson(&c.CompanyID, &flowJson, flowId)
		}, nil)

		if err != nil {
			return diag.FromErr(err)
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

	return resourceFlowRead(ctx, d, m)
}

func resourceFlowDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)
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
				panic(err)
			}
			new = *newFlowJson
		}
		if old != "" {
			oldFlowJson, err := mapSubFlows(d, old)
			if err != nil {
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
				panic(err)
			}
			new = *newFlowJson
		}
		if old != "" {
			oldFlowJson, err := mapFlowConnections(d, old)
			if err != nil {
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
			panic(err)
		}
		// convert to type Flow if needed
		if currentFi.FlowInfo.Name != "" {
			current = currentFi.FlowInfo
		}
		err = json.Unmarshal([]byte(old), &currentF)
		if err != nil {
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
			panic(err)
		}
		// convert to type Flow if needed
		if desiredFi.FlowInfo.Name != "" {
			desired = desiredFi.FlowInfo
		}
		err = json.Unmarshal([]byte(new), &desiredF)
		if err != nil {
			panic(err)
		}
		if desiredF.GraphData.Elements.Nodes != nil {
			desired = desiredF
		}
	}

	if current.Name != desired.Name {
		return false
	}

	if current.FlowStatus != desired.FlowStatus {
		return false
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
			panic(err)
		}
		desiredNodes, err := json.Marshal(desired.GraphData.Elements.Nodes)
		if err != nil {
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
			return err
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

	sfp := subflowProps["subFlowId"].(map[string]interface{})
	sfpVal := sfp["value"].(map[string]interface{})
	sfId := dv.SubFlowID{
		Value: dv.SubFlowValue{
			Value: sfpVal["value"].(string),
			Label: sfpVal["label"].(string),
		},
	}
	subflowVersionId := subflowProps["subFlowVersionId"].(map[string]interface{})
	sfv := dv.SubFlowVersionID{
		Value: subflowVersionId["value"].(string),
	}
	if sfId.Value.Value == "" || sfv.Value == "" {
		return nil, fmt.Errorf("Error: subflow value or versionId is empty")
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
				sfProp, _ := expandSubFlowProps(v.Data.Properties)
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
