package davinci

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

func ResourceApplication() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceApplicationCreate,
		ReadContext:   resourceApplicationRead,
		UpdateContext: resourceApplicationUpdate,
		DeleteContext: resourceApplicationDelete,
		Schema: map[string]*schema.Schema{
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PingOne environment id",
			},
			"customer_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Internal DaVinci id. Should not be set by user.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Application name",
			},
			"created_date": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Creation date as epoch.",
			},
			"api_key_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "Enabled by default in UI",
			},
			"api_keys": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Appplication Api Key. Returned value for prod field is most commonly used.",
				Elem: &schema.Schema{
					Type:      schema.TypeString,
					Sensitive: true,
				},
			},
			"metadata": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Appplication Metadata",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"user_pools": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Appplication User Pools. Not implemented",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"user_portal": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "User Profile in UI",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"up_title": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"add_auth_method_title": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"remove_auth_method_title": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"cred_page_title": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"cred_page_subtitle": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name_auth_method_title": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"name_confirm_btn_text": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"update_message": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"update_body_message": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"remove_message": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"remove_body_message": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"remove_confirm_btn_text": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"remove_cancel_btn_text": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"flow_timeout_seconds": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"show_user_info": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"show_mfa_button": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"show_variables": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
						"show_logout_button": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
						},
					},
				},
			},
			"oauth": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "OIDC configuration",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},
						"values": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "OIDC configuration",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Optional: true,
										Default:  true,
									},
									"client_secret": {
										Type:      schema.TypeString,
										Computed:  true,
										Sensitive: true,
									},
									"enforce_signed_request_openid": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"sp_jwks_url": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"sp_jwks_openid": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"redirect_uris": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"logout_uris": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"allowed_scopes": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Schema{
											Type:             schema.TypeString,
											ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"openid", "profile", "flow_analytics"}, false)),
										},
									},
									"allowed_grants": {
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Schema{
											Type:             schema.TypeString,
											ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"authorizationCode", "clientCredentials", "implicit"}, false)),
										},
									},
								},
							},
						},
					},
				},
			},
			"saml": {
				Type: schema.TypeSet,
				//TODO - Remove the need for this
				// requiring this to accound for returned nil values.
				Required:    true,
				Description: "SAML configuration",
				MaxItems:    1,
				// DefaultFunc: func() (interface{}, error) {
				// 	smap := []map[string]interface{}{{
				// 		"values": []map[string]interface{}{{
				// 			"enabled":                true,
				// 			"enforce_signed_request": false,
				// 		}},
				// 	}}
				// 	return smap, nil
				// },
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"values": {
							Type:        schema.TypeSet,
							Optional:    true,
							MaxItems:    1,
							Description: "SAML configuration",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"redirect_uri": {
										Type:      schema.TypeString,
										Optional:  true,
										Sensitive: true,
									},
									"enforce_signed_request": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"audience": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"sp_cert": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},
			"policy": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "Flow Policy Config",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"policy_flow": {
							Type:        schema.TypeSet,
							Optional:    true,
							Description: "Weighted flows that this Application will use",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"flow_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"version_id": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"weight": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"success_nodes": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"status": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "enabled",
						},
						"policy_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"created_date": {
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceApplicationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	app, err := expandApp(d)
	if err != nil {
		return diag.FromErr(err)
	}

	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.CreateInitializedApplication(&c.CompanyID, app)
	}, nil)

	if err != nil {
		return diag.FromErr(err)
	}

	res, ok := sdkRes.(*dv.App)
	if !ok || res.Name == "" {
		err = fmt.Errorf("Unable to parse created app response from Davinci API")
		return diag.FromErr(err)
	}

	d.SetId(res.AppID)

	resourceApplicationRead(ctx, d, m)

	return diags
}

func resourceApplicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	appId := d.Id()

	skdRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.ReadApplication(&c.CompanyID, appId)
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	resp, ok := skdRes.(*dv.App)
	if !ok {
		err = fmt.Errorf("failed to cast App type to response on Application with id: %s", appId)
		return diag.FromErr(err)
	}

	if err != nil {
		ep, err := c.ParseDvHttpError(err)
		if ep.Status == 404 && strings.Contains(ep.Body, "App not found") {
			d.SetId("")
			// diags = append(diags, diag.Diagnostic{})
			return diags
		}
		return diag.FromErr(err)
	}
	flatResp, err := flattenApp(resp)
	for i, v := range flatResp {

		d.Set(i, v)
	}
	d.SetId(resp.AppID)
	return diags
}

func resourceApplicationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	appId := d.Id()

	// if a new policy is added it must be created
	if d.HasChanges("policy") {
		_, new := d.GetChange("policy")
		pols := expandFlowPolicies(new)
		for _, v := range pols {
			if v.PolicyID == "" {
				sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
					return c.CreateFlowPolicy(&c.CompanyID, appId, v)
				}, nil)
				if err != nil {
					return diag.FromErr(err)
				}
				res, ok := sdkRes.(*dv.App)
				if !ok || res.Name == "" {
					err = fmt.Errorf("failed to cast create policy response to Application on id: %s", appId)
					return diag.FromErr(err)
				}
			} else {
				sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
					return c.UpdateFlowPolicy(&c.CompanyID, appId, v)
				}, nil)
				if err != nil {
					return diag.FromErr(err)
				}
				res, ok := sdkRes.(*dv.App)
				if !ok || res.Name == "" {
					err = fmt.Errorf("failed to cast update policy response to Application on id: %s", appId)
					return diag.FromErr(err)
				}
			}
		}
	}

	if d.HasChanges("name", "api_key_enabled", "user_portal", "oauth", "saml") {
		app, err := expandApp(d)
		if err != nil {
			return diag.FromErr(err)
		}
		app.AppID = d.Id()

		sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
			return c.UpdateApplication(&c.CompanyID, app)
		}, nil)
		if err != nil {
			return diag.FromErr(err)
		}
		res, ok := sdkRes.(*dv.App)
		if !ok || res.Name == "" {
			err = fmt.Errorf("failed to cast update application response to Application on id: %s", app.AppID)
			return diag.FromErr(err)
		}
	}

	return resourceApplicationRead(ctx, d, m)
}

func resourceApplicationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	appId := d.Id()

	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.DeleteApplication(&c.CompanyID, appId)
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}
	res, ok := sdkRes.(*dv.Message)
	if !ok || res.Message == "" {
		err = fmt.Errorf("failed to delete update application response to Application on id: %s", appId)
		return diag.FromErr(err)
	}

	d.SetId("")

	return diags
}

func expandApp(d *schema.ResourceData) (*dv.AppUpdate, error) {
	// Set Top layer.
	a := dv.AppUpdate{
		Name:          d.Get("name").(string),
		APIKeyEnabled: d.Get("api_key_enabled").(bool),
	}

	// Set OAuth
	oa, ok := d.GetOk("oauth")
	if ok {
		oal := oa.(*schema.Set).List()
		oaUpdate := &dv.Oauth{}

		//Set OAuthEnabled
		oam := oal[0].(map[string]interface{})
		if oalmEnabled, ok := oam["enabled"].(bool); ok {
			oaUpdate.Enabled = oalmEnabled
		}
		// Set OAuth Values
		oamValues := oam["values"].(*schema.Set).List()
		if len(oamValues) == 1 {
			oamValuesMap := oamValues[0].(map[string]interface{})
			oaUpdate.Values = &dv.OauthValues{
				Enabled:                    oamValuesMap["enabled"].(bool),
				ClientSecret:               oamValuesMap["client_secret"].(string),
				EnforceSignedRequestOpenid: oamValuesMap["enforce_signed_request_openid"].(bool),
				SpjwksUrl:                  oamValuesMap["sp_jwks_url"].(string),
				SpJwksOpenid:               oamValuesMap["sp_jwks_openid"].(string),
			}
			slist := expandStringList(oamValuesMap["redirect_uris"].(*schema.Set).List())
			oaUpdate.Values.RedirectUris = slist
			slist = expandStringList(oamValuesMap["logout_uris"].(*schema.Set).List())
			oaUpdate.Values.LogoutUris = slist
			slist = expandStringList(oamValuesMap["allowed_scopes"].(*schema.Set).List())
			oaUpdate.Values.AllowedScopes = slist
			slist = expandStringList(oamValuesMap["allowed_grants"].(*schema.Set).List())
			oaUpdate.Values.AllowedGrants = slist

		}
		if len(oamValues) > 1 {
			return nil, fmt.Errorf("Only one set for OAuth Values allowed")
		}
		a.Oauth = oaUpdate
	}

	//Set Saml
	saml, ok := d.GetOk("saml")
	if ok {
		sl := saml.(*schema.Set).List()
		if len(sl) == 1 {
			svUpdate := &dv.Saml{}
			sm := sl[0].(map[string]interface{})
			samlValues := sm["values"].(*schema.Set).List()
			if len(samlValues) == 1 {
				svMap := samlValues[0].(map[string]interface{})
				svUpdate.Values = &dv.SamlValues{
					Enabled:              svMap["enabled"].(bool),
					EnforceSignedRequest: svMap["enforce_signed_request"].(bool),
					RedirectURI:          svMap["redirect_uri"].(string),
					Audience:             svMap["audience"].(string),
					SpCert:               svMap["sp_cert"].(string),
				}
				a.Saml = svUpdate
			}
			if len(samlValues) > 1 {
				return nil, fmt.Errorf("Only one set for Saml Values allowed")
			}
		}
		if len(sl) > 1 {
			return nil, fmt.Errorf("Only one set for Saml Values allowed")
		}
	}

	//Set User Portal
	uv, ok := d.GetOk("user_portal")
	if ok {
		upValues := uv.(*schema.Set).List()
		if len(upValues) == 1 {
			up := &dv.UserPortal{}
			up.Values = upValues[0].(*dv.UserPortalValues)
			a.UserPortal = up
		}
		if len(upValues) > 1 {
			return nil, fmt.Errorf("Only one set for User Portal Values allowed")
		}
	}

	//Set Flow Policies
	fp, ok := d.GetOk("policy")
	if ok {
		fvUpdate := expandFlowPolicies(fp)
		if len(fvUpdate) > 0 {
			a.Policies = fvUpdate
		}
	}
	return &a, nil
}

func expandStringList(configured []interface{}) []string {
	vs := make([]string, 0, len(configured))
	for _, v := range configured {
		val, ok := v.(string)
		if ok && val != "" {
			vs = append(vs, v.(string))
		}
	}
	return vs
}

func expandFlowPolicies(fp interface{}) []dv.Policy {
	fl := fp.(*schema.Set).List()
	fvUpdate := []dv.Policy{}
	if len(fl) > 0 {
		for _, v := range fl {
			flMap := v.(map[string]interface{})
			thisFvUpdate := dv.Policy{
				Name:     flMap["name"].(string),
				Status:   flMap["status"].(string),
				PolicyID: flMap["policy_id"].(string),
			}
			thisPolicyFlows := flMap["policy_flow"].(*schema.Set).List()
			for _, w := range thisPolicyFlows {
				flPMap := w.(map[string]interface{})
				thisFvPUpdate := dv.PolicyFlow{
					FlowID:    flPMap["flow_id"].(string),
					VersionID: flPMap["version_id"].(int),
					Weight:    flPMap["weight"].(int),
				}
				thisFvUpdate.PolicyFlows = append(thisFvUpdate.PolicyFlows, thisFvPUpdate)
			}

			fvUpdate = append(fvUpdate, thisFvUpdate)
		}
	}
	return fvUpdate
}
