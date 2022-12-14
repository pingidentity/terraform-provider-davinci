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
			},
			"deploy": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Deploy Flow after import.",
			},
			"flow_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Computed Flow ID after import.",
			},
			"flow_name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Computed Flow Name after import.",
			},
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Environment to import flow into.",
			},
			"connections": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Connections this flow depends on. flow_json will be updated with provided connection IDs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"connection_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Connection ID",
						},
						"connection_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Connection Name",
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
			"subflows": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Child flows of this resource. Required if flow_json contains subflows.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"subflow_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Subflow Flow ID",
						},
						"subflow_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Subflow Name",
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
	}
}

func resourceFlowCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	var flowJson string
	if fj, ok := d.GetOk("flow_json"); ok {
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

	res, err := c.CreateFlowWithJson(&c.CompanyID, &flowJson)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := deployIfNeeded(ctx, c, d, res.FlowID); err != nil {
		return diag.FromErr(err)
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

	res, err := c.ReadFlow(&c.CompanyID, flowId)
	if err != nil {
		return diag.FromErr(err)
	}
	if res.Flow.FlowID == "" {
		d.SetId("")
		return diags
	}
	if err := d.Set("flow_id", res.Flow.FlowID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("flow_name", res.Flow.Name); err != nil {
		return diag.FromErr(err)
	}
	rString, err := json.Marshal(&res.Flow)

	if err := d.Set("flow_json", string(rString)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("environment_id", string(res.Flow.CompanyID)); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceFlowUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	flowId := d.Get("flow_id").(string)

	if d.HasChanges("flow_json") {
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
		_, err = c.UpdateFlowWithJson(&c.CompanyID, &flowJson, flowId)
		if err != nil {
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

	_, err = c.DeleteFlow(&c.CompanyID, flowId)
	if err != nil {
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
	if _, ok := d.GetOk("subflows"); ok {
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
	if _, ok := d.GetOk("connections"); ok {
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

func mapSubFlows(d *schema.ResourceData, flowJson string) (*string, error) {
	if sf, ok := d.GetOk("subflows"); ok {
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
					if sfValues["subflow_name"].(string) == sfProp.SubFlowID.Value.Label {
						sfProp.SubFlowID.Value.Value = sfValues["subflow_id"].(string)
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

func mapFlowConnections(d *schema.ResourceData, flowJson string) (*string, error) {
	if conns, ok := d.GetOk("connections"); ok {
		fjMap, err := dv.ParseFlowImportJson(&flowJson)
		if err != nil {
			return nil, err
		}
		connList := conns.(*schema.Set).List()
		for i, v := range fjMap.FlowInfo.GraphData.Elements.Nodes {
			for _, connMap := range connList {
				connValues := connMap.(map[string]interface{})
				if connValues["connection_name"].(string) == v.Data.Name {
					v.Data.ConnectionID = connValues["connection_id"].(string)
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
