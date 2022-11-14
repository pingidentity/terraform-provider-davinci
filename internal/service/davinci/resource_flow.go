package davinci

import (
	"context"
	"encoding/json"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
				Required:    true,
				Description: "Deploy Flow after import.",
			},
			"flow_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Computed Flow ID after import.",
			},
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Environment to import flow into.",
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

	flowJson := d.Get("flow_json").(string)

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
	rString, err := json.Marshal(&res.Flow)

	if err := d.Set("flow_json", string(rString)); err != nil {
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
		_, err := c.UpdateFlowWithJson(&c.CompanyID, &flowJson, flowId)
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
	drift := true

	desired := dv.Flow{}
	actual := dv.Flow{}
	err := json.Unmarshal([]byte(new), &desired)
	if err != nil {
		panic(err)
	}
	json.Unmarshal([]byte(old), &actual)
	if err != nil {
		panic(err)
	}

	if !reflect.DeepEqual(desired.GraphData, actual.GraphData) {
		drift = false
	}

	if desired.Name != actual.Name {
		drift = false
	}

	if desired.FlowStatus != actual.FlowStatus {
		drift = false
	}

	return drift
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
