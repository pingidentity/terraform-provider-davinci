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

func ResourceApplicationFlowPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceApplicationFlowPolicyCreate,
		ReadContext:   resourceApplicationFlowPolicyRead,
		UpdateContext: resourceApplicationFlowPolicyUpdate,
		DeleteContext: resourceApplicationFlowPolicyDelete,
		Schema: map[string]*schema.Schema{
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PingOne environment id",
			},
			"application_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Id of the application this policy is associated with",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Policy Name",
			},
			"policy_flow": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Set of weighted flows that this application will use",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"flow_id": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Identifier of the flow that this policy will use.",
						},
						"version_id": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Version of the flow that this policy will use. Use -1 for latest",
						},
						"weight": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "If multiple flows are specified, the weight determines the probability of the flow being used. This must add up to 100",
						},
						"success_nodes": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of node ids used by analytics for tracking user interaction.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"status": {
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "enabled",
				Description:      "If Policy should be enabled. Valid values are: enabled, disabled",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"enabled", "disabled"}, false)),
			},
			"created_date": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Creation epoch of policy.",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: resourceApplicaionFlowPolicyImport,
		},
	}
}

func resourceApplicationFlowPolicyCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	appPolicy, err := expandAppPolicy(d)
	if err != nil {
		return diag.FromErr(err)
	}

	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.CreateFlowPolicy(&c.CompanyID, d.Get("application_id").(string), *appPolicy)
	}, nil)

	if err != nil {
		return diag.FromErr(err)
	}

	res, ok := sdkRes.(*dv.App)
	if !ok || res.Name == "" {
		err = fmt.Errorf("Unable to parse created policy response from Davinci API")
		return diag.FromErr(err)
	}
	var resAppPolicy *dv.Policy
	for _, v := range res.Policies {
		if v.Name == appPolicy.Name {
			resAppPolicy = &v
			break
		}
	}
	if resAppPolicy.PolicyID == "" {
		err = fmt.Errorf("Unable to find created policy in response from Davinci API")
		return diag.FromErr(err)
	}

	d.SetId(resAppPolicy.PolicyID)

	resourceApplicationFlowPolicyRead(ctx, d, meta)

	return diags
}

func resourceApplicationFlowPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	appId := d.Get("application_id").(string)
	policyId := d.Id()

	skdRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.ReadApplication(&c.CompanyID, appId)
	}, nil)
	fmt.Printf("err: %+v\n", err)
	if err != nil {
		ep, errErr := c.ParseDvHttpError(err)
		if errErr != nil {
			return diag.FromErr(errErr)
		}
		if ep.Status == 404 && strings.Contains(ep.Body, "App not found") {
			d.SetId("")
			// diags = append(diags, diag.Diagnostic{})
			return diags
		}
		return diag.FromErr(err)
	}

	resp, ok := skdRes.(*dv.App)
	if !ok {
		err = fmt.Errorf("failed to cast App type to response on Application with id: %s", appId)
		return diag.FromErr(err)
	}

	flatResp, err := flattenAppPolicy(resp, policyId)
	// fmt.Printf("flatResp: %+v\n", flatResp)
	if err != nil {
		if strings.Contains(err.Error(), "Unable to find policy with id") {
			d.SetId("")
			return diags
		}
		return diag.FromErr(err)
	}
	for i, v := range flatResp {
		if err = d.Set(i, v); err != nil {
			return diag.FromErr(err)
		}
	}
	d.SetId(policyId)
	return diags
}

func resourceApplicationFlowPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	appId := d.Get("application_id").(string)
	appPolicy, err := expandAppPolicy(d)
	if err != nil {
		return diag.FromErr(err)
	}
	appPolicy.PolicyID = d.Id()

	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.UpdateFlowPolicy(&c.CompanyID, appId, *appPolicy)
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	res, ok := sdkRes.(*dv.App)
	if !ok || res.Name == "" {
		err = fmt.Errorf("failed to cast update policy response to Application on id: %s", appId)
		return diag.FromErr(err)
	}

	return resourceApplicationFlowPolicyRead(ctx, d, meta)
}

func resourceApplicationFlowPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	appId := d.Get("application_id").(string)
	policyId := d.Id()

	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.DeleteFlowPolicy(&c.CompanyID, appId, policyId)
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	res, ok := sdkRes.(*dv.Message)
	if !ok || res.Message != "" {
		err = fmt.Errorf("failed to delete application policy with id: %s", appId)
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func expandAppPolicy(d *schema.ResourceData) (*dv.Policy, error) {
	policy := dv.Policy{
		Name:   d.Get("name").(string),
		Status: d.Get("status").(string),
	}

	if v, ok := d.GetOk("policy_flow"); ok {
		var policyFlows []dv.PolicyFlow
		for _, vv := range v.(*schema.Set).List() {
			policyFlowMap := vv.(map[string]interface{})
			thisPolicyFlow := dv.PolicyFlow{
				FlowID:    policyFlowMap["flow_id"].(string),
				VersionID: policyFlowMap["version_id"].(int),
				Weight:    policyFlowMap["weight"].(int),
			}
			policyFlows = append(policyFlows, thisPolicyFlow)
		}
		policy.PolicyFlows = policyFlows
	}

	return &policy, nil
}

func flattenAppPolicy(app *dv.App, policyId string) (map[string]interface{}, error) {
	var policy dv.Policy
	for _, v := range app.Policies {
		if v.PolicyID == policyId {
			policy = v
			break
		}
	}
	if policy.PolicyID == "" {
		return nil, fmt.Errorf("Unable to find policy with id: %s", policyId)
	}
	a := map[string]interface{}{
		"environment_id": app.CompanyID,
		"application_id": app.AppID,
		"name":           policy.Name,
		"status":         policy.Status,
		"created_date":   policy.CreatedDate,
	}
	polFlows := []interface{}{}
	for _, w := range policy.PolicyFlows {
		thisPolFlow := map[string]interface{}{
			"flow_id":    w.FlowID,
			"weight":     w.Weight,
			"version_id": w.VersionID,
		}
		polFlows = append(polFlows, thisPolFlow)
	}

	a["policy_flow"] = polFlows

	//Return
	return a, nil
}

func resourceApplicaionFlowPolicyImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	idComponents := []utils.ImportComponent{
		{
			Label:  "environment_id",
			Regexp: regexp.MustCompile(`[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}`),
		},
		{
			Label:  "davinci_application_id",
			Regexp: regexp.MustCompile(`[a-f0-9]{32}`),
		},
		{
			Label:     "davinci_application_flow_policy_id",
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
	if err = d.Set("application_id", attributes["davinci_application_id"]); err != nil {
		return nil, err
	}
	d.SetId(attributes["davinci_application_flow_policy_id"])

	resourceApplicationFlowPolicyRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
}
