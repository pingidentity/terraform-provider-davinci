package davinci

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pingidentity/terraform-provider-davinci/internal/framework"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

func ResourceApplicationFlowPolicy() *schema.Resource {
	return &schema.Resource{
		DeprecationMessage: "This resource is deprecated and will be removed in a future release. Use the `pingone_davinci_application_flow_policy` resource in the PingOne Terraform provider instead (https://registry.terraform.io/providers/pingidentity/pingone/latest/docs/resources/davinci_application_flow_policy). For more information, see https://github.com/pingidentity/terraform-provider-davinci/issues/601",
		CreateContext:      resourceApplicationFlowPolicyCreate,
		ReadContext:        resourceApplicationFlowPolicyRead,
		UpdateContext:      resourceApplicationFlowPolicyUpdate,
		DeleteContext:      resourceApplicationFlowPolicyDelete,
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
			"trigger": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "A block that specifies the trigger configuration for PingOne flow policies.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A string that specifies the trigger type. Set by the DaVinci API.",
						},
						"configuration": {
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "A block that specifies the trigger configuration details.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"mfa": {
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
										Description: "A block that specifies the MFA trigger configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "A boolean that specifies whether MFA trigger is enabled.",
												},
												"time": {
													Type:        schema.TypeFloat,
													Optional:    true,
													Description: "A number that specifies the MFA trigger time.",
												},
												"time_format": {
													Type:             schema.TypeString,
													Optional:         true,
													Description:      "A string that specifies the MFA trigger time format. Valid values are `min`, `hour`, `day`.",
													ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"min", "hour", "day"}, false)),
												},
											},
										},
									},
									"pwd": {
										Type:        schema.TypeList,
										Optional:    true,
										MaxItems:    1,
										Description: "A block that specifies the password trigger configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Optional:    true,
													Description: "A boolean that specifies whether password trigger is enabled.",
												},
												"time": {
													Type:        schema.TypeFloat,
													Optional:    true,
													Description: "A number that specifies the password trigger time.",
												},
												"time_format": {
													Type:             schema.TypeString,
													Optional:         true,
													Description:      "A string that specifies the password trigger time format. Valid values are `min`, `hour`, `day`.",
													ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"min", "hour", "day"}, false)),
												},
											},
										},
									},
								},
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
			return c.CreateFlowPolicyWithResponse(environmentID, d.Get("application_id").(string), *appPolicy)
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
		if v.Name != nil && appPolicy.Name != nil && *v.Name == *appPolicy.Name {
			v := v // G601 (CWE-118)
			resAppPolicy = &v
			break
		}
	}
	if resAppPolicy == nil || resAppPolicy.PolicyID == nil || *resAppPolicy.PolicyID == "" {
		err = fmt.Errorf("Unable to find created policy in response from Davinci API")
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId(*resAppPolicy.PolicyID)

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
			return c.ReadApplicationWithResponse(environmentID, appId)
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
	policyId := d.Id()
	appPolicy.PolicyID = &policyId

	environmentID := d.Get("environment_id").(string)

	sdkRes, err := sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.UpdateFlowPolicyWithResponse(environmentID, appId, *appPolicy)
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
			return c.DeleteFlowPolicyWithResponse(environmentID, appId, policyId)
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
	if !ok || (res != nil && res.Message != nil && *res.Message != "") {
		err = fmt.Errorf("failed to delete application policy with id: %s", appId)
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId("")

	return diags
}

func expandAppPolicy(d *schema.ResourceData) (*dv.Policy, error) {
	policyName := d.Get("name").(string)
	policyStatus := d.Get("status").(string)
	policy := dv.Policy{
		Name:   &policyName,
		Status: &policyStatus,
	}

	if v, ok := d.GetOk("policy_flow"); ok {
		var policyFlows []dv.PolicyFlow
		for _, vv := range v.(*schema.Set).List() {
			policyFlowMap := vv.(map[string]interface{})
			thisPolicyFlow := dv.PolicyFlow{
				FlowID:    policyFlowMap["flow_id"].(string),
				VersionID: policyFlowMap["version_id"].(int),
			}

			if v, ok := policyFlowMap["weight"].(int); ok {
				thisPolicyFlow.Weight = &v
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

	if v, ok := d.GetOk("trigger"); ok {
		triggerList := v.([]interface{})
		if len(triggerList) > 0 && triggerList[0] != nil {
			triggerMap := triggerList[0].(map[string]interface{})
			policyTrigger := &dv.PolicyTrigger{}

			// NOTE: policyTrigger.Type is intentionally NOT set here.
			// The DaVinci API rejects "type" as an additional property on POST/PUT.
			// The API sets the type itself and returns it on GET; flattenAppPolicy reads it from state.

			if configList, ok := triggerMap["configuration"].([]interface{}); ok && len(configList) > 0 && configList[0] != nil {
				configMap := configList[0].(map[string]interface{})
				triggerConfig := &dv.TriggerConfiguration{}

				if mfaList, ok := configMap["mfa"].([]interface{}); ok && len(mfaList) > 0 && mfaList[0] != nil {
					mfaMap := mfaList[0].(map[string]interface{})
					mfaConfig := &dv.TriggerConfigurationMFA{}

					if enabled, ok := mfaMap["enabled"].(bool); ok {
						mfaConfig.Enabled = &enabled
					}
					if timeVal, ok := mfaMap["time"].(float64); ok {
						v := float32(timeVal)
						mfaConfig.Time = &v
					}
					if timeFormat, ok := mfaMap["time_format"].(string); ok && timeFormat != "" {
						mfaConfig.TimeFormat = &timeFormat
					}

					triggerConfig.MFA = mfaConfig
				}

				if pwdList, ok := configMap["pwd"].([]interface{}); ok && len(pwdList) > 0 && pwdList[0] != nil {
					pwdMap := pwdList[0].(map[string]interface{})
					pwdConfig := &dv.TriggerConfigurationPassword{}

					if enabled, ok := pwdMap["enabled"].(bool); ok {
						pwdConfig.Enabled = &enabled
					}
					if timeVal, ok := pwdMap["time"].(float64); ok {
						v := float32(timeVal)
						pwdConfig.Time = &v
					}
					if timeFormat, ok := pwdMap["time_format"].(string); ok && timeFormat != "" {
						pwdConfig.TimeFormat = &timeFormat
					}

					triggerConfig.PWD = pwdConfig
				}

				policyTrigger.Configuration = triggerConfig
			}

			policy.Trigger = policyTrigger
		}
	}

	return &policy, nil
}

func flattenAppPolicy(app *dv.App, policyId string) (map[string]interface{}, error) {
	var policy dv.Policy
	for _, v := range app.Policies {
		if *v.PolicyID == policyId {
			policy = v
			break
		}
	}
	if policy.PolicyID == nil || *policy.PolicyID == "" {
		return nil, fmt.Errorf("Unable to find policy with id: %s", policyId)
	}
	a := map[string]interface{}{}

	if app.CompanyID != nil {
		a["environment_id"] = app.CompanyID
	}

	if app.AppID != nil {
		a["application_id"] = app.AppID
	}

	if policy.Name != nil {
		a["name"] = *policy.Name
	}

	if policy.Status != nil {
		a["status"] = *policy.Status
	}

	if policy.CreatedDate != nil {
		a["created_date"] = policy.CreatedDate.UnixMilli()
	}

	polFlows := []interface{}{}
	for _, w := range policy.PolicyFlows {
		thisPolFlow := map[string]interface{}{
			"flow_id":         w.FlowID,
			"version_id":      w.VersionID,
			"success_nodes":   w.SuccessNodes,
			"allowed_ip_list": w.IP,
		}

		if w.Weight != nil {
			thisPolFlow["weight"] = *w.Weight
		}

		polFlows = append(polFlows, thisPolFlow)
	}

	a["policy_flow"] = polFlows

	if policy.Trigger != nil {
		triggerMap := map[string]interface{}{}

		if policy.Trigger.Type != nil {
			triggerMap["type"] = *policy.Trigger.Type
		}

		if policy.Trigger.Configuration != nil {
			configMap := map[string]interface{}{}

			if policy.Trigger.Configuration.MFA != nil {
				mfaMap := map[string]interface{}{}
				mfa := policy.Trigger.Configuration.MFA

				if mfa.Enabled != nil {
					mfaMap["enabled"] = *mfa.Enabled
				}
				if mfa.Time != nil {
					mfaMap["time"] = float64(*mfa.Time)
				}
				if mfa.TimeFormat != nil {
					mfaMap["time_format"] = *mfa.TimeFormat
				}

				configMap["mfa"] = []interface{}{mfaMap}
			}

			if policy.Trigger.Configuration.PWD != nil {
				pwdMap := map[string]interface{}{}
				pwd := policy.Trigger.Configuration.PWD

				if pwd.Enabled != nil {
					pwdMap["enabled"] = *pwd.Enabled
				}
				if pwd.Time != nil {
					pwdMap["time"] = float64(*pwd.Time)
				}
				if pwd.TimeFormat != nil {
					pwdMap["time_format"] = *pwd.TimeFormat
				}

				configMap["pwd"] = []interface{}{pwdMap}
			}

			triggerMap["configuration"] = []interface{}{configMap}
		}

		a["trigger"] = []interface{}{triggerMap}
	}

	//Return
	return a, nil
}

func resourceApplicaionFlowPolicyImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	idComponents := []framework.ImportComponent{
		{
			Label:  "environment_id",
			Regexp: verify.P1ResourceIDRegexp,
		},
		{
			Label:  "davinci_application_id",
			Regexp: verify.P1DVResourceIDRegexp,
		},
		{
			Label:     "davinci_application_flow_policy_id",
			Regexp:    verify.P1DVResourceIDRegexp,
			PrimaryID: true,
		},
	}
	attributes, err := framework.ParseImportID(d.Id(), idComponents...)
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
