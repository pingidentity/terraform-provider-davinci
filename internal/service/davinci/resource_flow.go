package davinci

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	// "github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
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
		// Flatten subflow dependencies if needed.
		expandedFlowJson, err := expandSubFlow(d, flowJson)
		if err != nil {
			return diag.FromErr(err)
		}
		flowJson = *expandedFlowJson
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
		expandedFlowJson, err := expandSubFlow(d, flowJson)
		if err != nil {
			return diag.FromErr(err)
		}
		flowJson = *expandedFlowJson
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
	desired := dv.Flow{}
	current := dv.Flow{}
	if _, ok := d.GetOk("subflows"); ok {
		if new != "" {
			newFlowJson, err := expandSubFlow(d, new)
			if err != nil {
				panic(err)
			}
			new = *newFlowJson
		}
		if old != "" {
			oldFlowJson, err := expandSubFlow(d, old)
			if err != nil {
				panic(err)
			}
			old = *oldFlowJson
		}
	}

	json.Unmarshal([]byte(old), &current)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(new), &desired)
	if err != nil {
		panic(err)
	}

	if current.Name != desired.Name {
		return false
	}

	if current.FlowStatus != desired.FlowStatus {
		return false
	}

	if !reflect.DeepEqual(current.GraphData, desired.GraphData) {
		return false
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

func expandSubFlow(d *schema.ResourceData, flowJson string) (*string, error) {
	if sf, ok := d.GetOk("subflows"); ok {
		fjMap, err := dv.ParseFlowJson(&flowJson)
		if err != nil {
			return nil, err
		}
		sfList := sf.(*schema.Set).List()
		for i, v := range fjMap.FlowInfo.GraphData.Elements.Nodes {
			sfProp := &dv.SubFlowProperties{}
			if v.Data.ConnectorID == "flowConnector" {
				sfProp, err = expandSubFlowProps(v.Data.Properties)
				for _, sfMap := range sfList {
					sfValues := sfMap.(map[string]interface{})
					if sfValues["subflow_name"] == sfProp.SubFlowID.Value.Label {
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
