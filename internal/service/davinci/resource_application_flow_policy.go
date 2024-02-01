package davinci

import (
	"context"
	"fmt"
	"net/http"
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
				Description: "The ID of the PingOne environment to manage the flow policy in. Must be a valid PingOne resource ID. This field is immutable and will trigger a replace plan if changed.",
				ForceNew:    true,
			},
			"application_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the DaVinci application to manage the flow policy for. Must be a valid DaVinci resource ID. This field is immutable and will trigger a replace plan if changed.",
				ForceNew:    true,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A string that specifies the name of the policy.",
			},
			"policy_flow": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "Set of weighted flows that this application will use.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"flow_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "Identifier of the flow that this policy will use.",
						},
						"version_id": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Version of the flow that this policy will use. Use `-1` for the latest version.",
						},
						"weight": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "If multiple flows are specified, the weight determines the probability of the flow being used. The weights across all policy flows must add up to `100`.",
						},
						"success_nodes": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "A list of node ids used by analytics for tracking user interaction.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"allowed_ip_list": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "A list of IP CIDR entries that are allowed use of the application policy flow.",
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
				Description:      "A boolan that specifies whether the policy should be enabled. Valid values are: `enabled`, `disabled`.",
				ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"enabled", "disabled"}, false)),
			},
			"created_date": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Resource creation date as epoch.",
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

	appPolicy, err := expandAppPolicy(d)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	environmentID := d.Get("environment_id").(string)

	sdkRes, err := sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.CreateFlowPolicyWithResponse(&environmentID, d.Get("application_id").(string), *appPolicy)
		},
	)

	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	res, ok := sdkRes.(*dv.App)
	if !ok || res.Name == "" {
		err = fmt.Errorf("Unable to parse created policy response from Davinci API")
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	var resAppPolicy *dv.Policy
	for _, v := range res.Policies {
		if v.Name == appPolicy.Name {
			v := v // G601 (CWE-118)
			resAppPolicy = &v
			break
		}
	}
	if resAppPolicy.PolicyID == "" {
		err = fmt.Errorf("Unable to find created policy in response from Davinci API")
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId(resAppPolicy.PolicyID)

	resourceApplicationFlowPolicyRead(ctx, d, meta)

	return diags
}

func resourceApplicationFlowPolicyRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	appId := d.Get("application_id").(string)
	policyId := d.Id()

	environmentID := d.Get("environment_id").(string)

	skdRes, err := sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.ReadApplicationWithResponse(&environmentID, appId)
		},
	)
	if err != nil {
		if dvError, ok := err.(dv.ErrorResponse); ok {
			if dvError.HttpResponseCode == http.StatusNotFound || dvError.Code == dv.DV_ERROR_CODE_APPLICATION_NOT_FOUND {
				d.SetId("")
				return diags
			}
		}
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	resp, ok := skdRes.(*dv.App)
	if !ok {
		err = fmt.Errorf("failed to cast App type to response on Application with id: %s", appId)
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	flatResp, err := flattenAppPolicy(resp, policyId)
	if err != nil {
		if strings.Contains(err.Error(), "Unable to find policy with id") {
			d.SetId("")
			return diags
		}
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	for i, v := range flatResp {
		if err = d.Set(i, v); err != nil {
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}
	}
	d.SetId(policyId)
	return diags
}

func resourceApplicationFlowPolicyUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := meta.(*dv.APIClient)

	appId := d.Get("application_id").(string)
	appPolicy, err := expandAppPolicy(d)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	appPolicy.PolicyID = d.Id()

	environmentID := d.Get("environment_id").(string)

	sdkRes, err := sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.UpdateFlowPolicyWithResponse(&environmentID, appId, *appPolicy)
		},
	)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	res, ok := sdkRes.(*dv.App)
	if !ok || res.Name == "" {
		err = fmt.Errorf("failed to cast update policy response to Application on id: %s", appId)
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	return resourceApplicationFlowPolicyRead(ctx, d, meta)
}

func resourceApplicationFlowPolicyDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	appId := d.Get("application_id").(string)
	policyId := d.Id()

	environmentID := d.Get("environment_id").(string)

	sdkRes, err := sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.DeleteFlowPolicyWithResponse(&environmentID, appId, policyId)
		},
	)
	if err != nil {
		if dvError, ok := err.(dv.ErrorResponse); ok {
			if dvError.HttpResponseCode == http.StatusNotFound || dvError.Code == dv.DV_ERROR_CODE_APPLICATION_NOT_FOUND {
				return diags
			}
		}
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	res, ok := sdkRes.(*dv.Message)
	if !ok || res.Message != "" {
		err = fmt.Errorf("failed to delete application policy with id: %s", appId)
		diags = append(diags, diag.FromErr(err)...)
		return diags
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

			successNodes := make([]string, 0)
			for _, successNode := range policyFlowMap["success_nodes"].(*schema.Set).List() {
				successNodes = append(successNodes, successNode.(string))
			}

			ips := make([]string, 0)
			for _, ip := range policyFlowMap["allowed_ip_list"].(*schema.Set).List() {
				ips = append(ips, ip.(string))
			}

			thisPolicyFlow.SuccessNodes = successNodes
			thisPolicyFlow.IP = ips

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
			"flow_id":         w.FlowID,
			"weight":          w.Weight,
			"version_id":      w.VersionID,
			"success_nodes":   w.SuccessNodes,
			"allowed_ip_list": w.IP,
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
