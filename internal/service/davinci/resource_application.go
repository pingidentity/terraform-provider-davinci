package davinci

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pingidentity/terraform-provider-davinci/internal/framework"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
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
				Description: "The ID of the PingOne environment to create the DaVinci application. Must be a valid PingOne resource ID. This field is immutable and will trigger a replace plan if changed.",

				ValidateDiagFunc: validation.ToDiagFunc(verify.ValidP1ResourceID),
				ForceNew:         true,
			},
			"customer_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An ID that represents the customer tenant.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The application name.",
			},
			"created_date": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Resource creation date as epoch.",
			},
			"api_key_enabled": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				Description: "A boolean that specifies whether the API key is enabled for the application.",
			},
			"api_keys": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "A map of strings that represents the application's API Key.",
				Elem: &schema.Schema{
					Type:      schema.TypeString,
					Sensitive: true,
				},
			},
			"metadata": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Application Metadata.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"user_pools": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Application User Pools.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"user_portal": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.  A single object that describes user portal settings.",
				Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"up_title": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
						},
						"add_auth_method_title": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
						},
						"remove_auth_method_title": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
						},
						"cred_page_title": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
						},
						"cred_page_subtitle": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
						},
						"name_auth_method_title": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
						},
						"name_confirm_btn_text": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
						},
						"update_message": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
						},
						"update_body_message": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
						},
						"remove_message": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
						},
						"remove_body_message": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
						},
						"remove_confirm_btn_text": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
						},
						"remove_cancel_btn_text": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
						},
						"flow_timeout_seconds": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
						},
						"show_user_info": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
						},
						"show_mfa_button": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
						},
						"show_variables": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
						},
						"show_logout_button": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
						},
					},
				},
			},
			"oauth": {
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "A single list item specifying OIDC/OAuth 2.0 configuration.",
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"enabled": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     true,
							Description: "A boolean that specifies whether OIDC/OAuth 2.0 settings are enabled for the application.",
						},
						"values": {
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "A single list item specifying OIDC/OAuth 2.0 configuration values.",
							MaxItems:    1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     true,
										Description: "A boolean that enables/disables the OAuth 2.0 configuration for the application.",
									},
									"client_secret": {
										Type:        schema.TypeString,
										Computed:    true,
										Sensitive:   true,
										Description: "The client secret for the OIDC application.",
									},
									"enforce_signed_request_openid": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "A boolean that specifies whether to enforce receiving signed requests.",
									},
									"sp_jwks_url": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A string that specifies a service provider (SP) JWKS URL.",
									},
									"sp_jwks_openid": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A string that specifies service provider (SP) JWKS keys to verify the authorization request signature.",
									},
									"redirect_uris": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Redirect URLs for the application.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"logout_uris": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Logout URLs for the application.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"allowed_scopes": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Allowed scopes for the application. Available scopes are `openid`, `profile`, `flow_analytics`.",
										Elem: &schema.Schema{
											Type:             schema.TypeString,
											ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"openid", "profile", "flow_analytics"}, false)),
										},
									},
									"allowed_grants": {
										Type:        schema.TypeSet,
										Optional:    true,
										Description: "Allowed grants for the application. Available grants are `authorizationCode`, `clientCredentials`, `implicit`.",
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
				Type:        schema.TypeList,
				Optional:    true,
				Computed:    true,
				Description: "**Deprecation notice**: SAML configuration is now deprecated in the service and will be removed in the next major release.  A single list item that specifies SAML configuration.",
				Deprecated:  "SAML configuration is now deprecated in the service and will be removed in the next major release.",
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
							Type:        schema.TypeList,
							Optional:    true,
							MaxItems:    1,
							Description: "SAML configuration",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Set to true if using saml block. This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"redirect_uri": {
										Type:        schema.TypeString,
										Optional:    true,
										Sensitive:   true,
										Description: "The redirect URI for the SAML application. This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"enforce_signed_request": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Field: 'Enforce Receiving Signed Requests' in UI. This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"audience": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Field: 'Audience' in UI. This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"sp_cert": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "This is now deprecated in the service and will be removed from the provider in the next major release.",
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
				Computed:    true,
				Description: "**Deprecation Notice** The `policy` block should be removed from application and managed as a `davinci_application_flow_policy` resource. Review the migration guidance here: https://registry.terraform.io/providers/pingidentity/davinci/latest/docs/guides/migrate-application-flow-policy. Flow Policy Configuration",
				Deprecated:  "The `policy` block should be removed from application and managed as a `davinci_application_flow_policy` resource. Review the migration guidance here: https://registry.terraform.io/providers/pingidentity/davinci/latest/docs/guides/migrate-application-flow-policy. This attribute will be removed in the next major release.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
						"name": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Policy friendly name",
						},
						"status": {
							Type:             schema.TypeString,
							Optional:         true,
							Default:          "enabled",
							Description:      "Policy status. Valid values are: enabled, disabled",
							ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice([]string{"enabled", "disabled"}, false)),
						},
						"policy_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Generated identifier of a created policy.",
						},
						"created_date": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Creation epoch of policy.",
						},
					},
				},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: resourceApplicationImport,
		},
	}
}

func resourceApplicationCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	app, err := expandApp(d)
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
			return c.CreateInitializedApplicationWithResponse(environmentID, app)
		},
	)

	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	res, ok := sdkRes.(*dv.App)
	if !ok || res.Name == "" {
		err = fmt.Errorf("Unable to parse created app response from Davinci API")
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId(*res.AppID)

	resourceApplicationRead(ctx, d, meta)

	return diags
}

func resourceApplicationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	appId := d.Id()

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

	flatResp, err := flattenApp(resp)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	for i, v := range flatResp {
		if err = d.Set(i, v); err != nil {
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}
	}
	d.SetId(*resp.AppID)
	return diags
}

func resourceApplicationUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := meta.(*dv.APIClient)

	appId := d.Id()

	environmentID := d.Get("environment_id").(string)

	// Policy CRUD
	// if policies have changes, compare changes. if new policy is added, create it. if old policy is removed, delete it. if policy is updated, update it.
	if d.HasChange("policy") {
		oldPols, newPols := d.GetChange("policy")
		oldPolsList := expandFlowPolicies(oldPols)
		newPolsList := expandFlowPolicies(newPols)
		oldPolicyMap := make(map[string]dv.Policy)
		for _, v := range oldPolsList {
			oldPolicyMap[*v.PolicyID] = v
		}
		for _, newPol := range newPolsList {
			oldPol, exists := oldPolicyMap[*newPol.PolicyID]
			if exists {
				//update policy
				sdkRes, err := sdk.DoRetryable(
					ctx,
					c,
					environmentID,
					func() (interface{}, *http.Response, error) {
						return c.UpdateFlowPolicyWithResponse(environmentID, appId, newPol)
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
				delete(oldPolicyMap, *oldPol.PolicyID)
			} else {
				// create policy
				sdkRes, err := sdk.DoRetryable(
					ctx,
					c,
					environmentID,
					func() (interface{}, *http.Response, error) {
						return c.CreateFlowPolicyWithResponse(environmentID, appId, newPol)
					},
				)
				if err != nil {
					diags = append(diags, diag.FromErr(err)...)
					return diags
				}
				res, ok := sdkRes.(*dv.App)
				if !ok || res.Name == "" {
					err = fmt.Errorf("failed to cast create policy response to Application on id: %s", appId)
					diags = append(diags, diag.FromErr(err)...)
					return diags
				}
				delete(oldPolicyMap, *oldPol.PolicyID)
			}
		}
		//delete old policies that are not in new policies
		for _, oldPol := range oldPolicyMap {
			_, err := sdk.DoRetryable(
				ctx,
				c,
				environmentID,
				func() (interface{}, *http.Response, error) {
					return c.DeleteFlowPolicyWithResponse(environmentID, appId, *oldPol.PolicyID)
				},
			)
			if err != nil {
				diags = append(diags, diag.FromErr(err)...)
				return diags
			}
		}
	}

	if d.HasChangesExcept("policy") {
		app, err := expandApp(d)
		if err != nil {
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}
		appId := d.Id()
		app.AppID = &appId

		sdkRes, err := sdk.DoRetryable(
			ctx,
			c,
			environmentID,
			func() (interface{}, *http.Response, error) {
				return c.UpdateApplicationWithResponse(environmentID, app)
			},
		)
		if err != nil {
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}
		res, ok := sdkRes.(*dv.App)
		if !ok || res.Name == "" {
			err = fmt.Errorf("failed to cast update application response to Application on id: %v", *app.AppID)
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}
	}

	return resourceApplicationRead(ctx, d, meta)
}

func resourceApplicationDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	appId := d.Id()

	environmentID := d.Get("environment_id").(string)

	sdkRes, err := sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.DeleteApplicationWithResponse(environmentID, appId)
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
	if !ok || res.Message == nil || *res.Message == "" {
		err = fmt.Errorf("Error when deleting application with id %s, no response message found.", appId)
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId("")

	return diags
}

func resourceApplicationImport(ctx context.Context, d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {

	idComponents := []framework.ImportComponent{
		{
			Label:  "environment_id",
			Regexp: verify.P1ResourceIDRegexp,
		},
		{
			Label:     "davinci_application_id",
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

	d.SetId(attributes["davinci_application_id"])

	resourceApplicationRead(ctx, d, meta)

	return []*schema.ResourceData{d}, nil
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
		oal := oa.([]interface{})
		oaUpdate := &dv.Oauth{}

		//Set OAuthEnabled
		oam := oal[0].(map[string]interface{})
		if oalmEnabled, ok := oam["enabled"].(bool); ok {
			oaUpdate.Enabled = oalmEnabled
		}
		// Set OAuth Values
		oamValues := oam["values"].([]interface{})
		if len(oamValues) == 1 {
			oamValuesMap := oamValues[0].(map[string]interface{})
			oaUpdate.Values = &dv.OauthValues{
				Enabled: oamValuesMap["enabled"].(bool),
			}

			if v, ok := oamValuesMap["client_secret"].(string); ok {
				oaUpdate.Values.ClientSecret = &v
			}

			if v, ok := oamValuesMap["enforce_signed_request_openid"].(bool); ok {
				oaUpdate.Values.EnforceSignedRequestOpenid = &v
			}

			if v, ok := oamValuesMap["sp_jwks_url"].(string); ok {
				oaUpdate.Values.SpjwksUrl = &v
			}

			if v, ok := oamValuesMap["sp_jwks_openid"].(string); ok {
				oaUpdate.Values.SpJwksOpenid = &v
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
		sl := saml.([]interface{})
		if len(sl) == 1 {
			svUpdate := &dv.Saml{}
			sm := sl[0].(map[string]interface{})
			samlValues := sm["values"].([]interface{})
			if len(samlValues) == 1 {
				svMap := samlValues[0].(map[string]interface{})
				svUpdate.Values = &dv.SamlValues{
					Enabled: svMap["enabled"].(bool),
				}

				if v, ok := svMap["enforce_signed_request"].(bool); ok {
					svUpdate.Values.EnforceSignedRequest = &v
				}

				if v, ok := svMap["redirect_uri"].(string); ok {
					svUpdate.Values.RedirectURI = &v
				}

				if v, ok := svMap["audience"].(string); ok {
					svUpdate.Values.Audience = &v
				}

				if v, ok := svMap["sp_cert"].(string); ok {
					svUpdate.Values.SpCert = &v
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
		upValues := uv.([]interface{})
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
	policyList := fp.(*schema.Set).List()
	fvUpdate := []dv.Policy{}
	if len(policyList) > 0 {
		for _, policy := range policyList {
			flMap := policy.(map[string]interface{})
			thisFvUpdate := dv.Policy{}

			if v, ok := flMap["name"].(string); ok {
				thisFvUpdate.Name = &v
			}

			if v, ok := flMap["status"].(string); ok {
				thisFvUpdate.Status = &v
			}

			if v, ok := flMap["policy_id"].(string); ok {
				thisFvUpdate.PolicyID = &v
			}

			thisPolicyFlows := flMap["policy_flow"].(*schema.Set).List()
			for _, w := range thisPolicyFlows {
				flPMap := w.(map[string]interface{})
				thisFvPUpdate := dv.PolicyFlow{
					FlowID:    flPMap["flow_id"].(string),
					VersionID: flPMap["version_id"].(int),
				}

				if v, ok := flPMap["weight"].(int); ok {
					thisFvPUpdate.Weight = &v
				}

				thisFvUpdate.PolicyFlows = append(thisFvUpdate.PolicyFlows, thisFvPUpdate)
			}

			fvUpdate = append(fvUpdate, thisFvUpdate)
		}
	}
	return fvUpdate
}
