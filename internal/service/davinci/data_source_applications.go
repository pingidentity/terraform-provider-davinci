package davinci

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

func DataSourceApplications() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplicationsRead,
		Schema: map[string]*schema.Schema{
			"environment_id": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The ID of the PingOne environment to retrieve applications from. Must be a valid PingOne resource ID.",
				ValidateDiagFunc: validation.ToDiagFunc(verify.ValidP1ResourceID),
			},
			"applications": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "A set of applications retrieved from the environment.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID for the application.",
						},
						"environment_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the PingOne environment that contains the application.",
						},
						"customer_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "An ID that represents the customer tenant.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The application name",
						},
						"created_date": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Resource creation date as epoch.",
						},
						"api_key_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
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
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.  A single object that describes user portal settings.",
							Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"up_title": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
										Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"add_auth_method_title": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
										Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"remove_auth_method_title": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
										Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"cred_page_title": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
										Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"cred_page_subtitle": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
										Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"name_auth_method_title": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
										Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"name_confirm_btn_text": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
										Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"update_message": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
										Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"update_body_message": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
										Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"remove_message": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
										Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"remove_body_message": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
										Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"remove_confirm_btn_text": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
										Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"remove_cancel_btn_text": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
										Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"flow_timeout_seconds": {
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
										Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"show_user_info": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
										Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"show_mfa_button": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
										Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"show_variables": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
										Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
									"show_logout_button": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "**Deprecation notice** This is now deprecated in the service and will be removed from the provider in the next major release.",
										Deprecated:  "This is now deprecated in the service and will be removed from the provider in the next major release.",
									},
								},
							},
						},
						"oauth": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "A single set item specifying OIDC/OAuth 2.0 configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "A boolean that specifies whether OIDC/OAuth 2.0 settings are enabled for the application.",
									},
									"values": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "A single list item specifying OIDC/OAuth 2.0 configuration values.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Computed:    true,
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
													Computed:    true,
													Description: "A boolean that specifies whether to enforce receiving signed requests.",
												},
												"sp_jwks_url": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "A string that specifies a service provider (SP) JWKS URL.",
												},
												"sp_jwks_openid": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "A string that specifies service provider (SP) JWKS keys to verify the authorization request signature.",
												},
												"redirect_uris": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Redirect URLs for the application.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"logout_uris": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Logout URLs for the application.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"allowed_scopes": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Allowed scopes for the application. Available scopes are `openid`, `profile`, `flow_analytics`.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"allowed_grants": {
													Type:        schema.TypeSet,
													Computed:    true,
													Description: "Allowed grants for the application. Available grants are `authorizationCode`, `clientCredentials`, `implicit`.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"saml": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "**Deprecation notice**: SAML configuration is now deprecated in the service and will be removed in the next major release.  A single list item that specifies SAML configuration.",
							Deprecated:  "SAML configuration is now deprecated in the service and will be removed in the next major release.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"values": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "SAML configuration. This is deprecated in the UI and will be removed in a future release.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Set to true if using saml block. This is deprecated in the UI and will be removed in a future release.",
												},
												"redirect_uri": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The redirect URI for the SAML application. This is deprecated in the UI and will be removed in a future release.",
												},
												"enforce_signed_request": {
													Type:        schema.TypeBool,
													Computed:    true,
													Description: "Field: 'Enforce Receiving Signed Requests' in UI. This is deprecated in the UI and will be removed in a future release.",
												},
												"audience": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Field: 'Audience' in UI. This is deprecated in the UI and will be removed in a future release.",
												},
												"sp_cert": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "This is deprecated in the UI and will be removed in a future release.",
												},
											},
										},
									},
								},
							},
						},
						"policy": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "A set of Flow Policies assigned to the application.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"policy_flow": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Set of weighted flows that this application will use",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"flow_id": {
													Type:        schema.TypeString,
													Computed:    true,
													Description: "Identifier of the flow that this policy will use.",
												},
												"version_id": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Version of the flow that this policy will use. Use -1 for latest",
												},
												"weight": {
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "If multiple flows are specified, the weight determines the probability of the flow being used. This must add up to 100",
												},
												"success_nodes": {
													Type:        schema.TypeList,
													Computed:    true,
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
										Computed:    true,
										Description: "Policy friendly name",
									},
									"status": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Policy status. Valid values are: enabled, disabled",
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
				},
			},
		},
		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(20 * time.Minute),
		},
	}
}

func dataSourceApplicationsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	environmentID := d.Get("environment_id").(string)

	res, err := readAllApplications(ctx, c, environmentID, d.Timeout(schema.TimeoutRead))
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	apps := make([]map[string]interface{}, 0)
	for _, thisApp := range res {
		thisApp := thisApp // G601 (CWE-118)
		app, err := flattenApp(&thisApp)
		if err != nil {
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}
		app["id"] = thisApp.AppID
		apps = append(apps, app)
	}

	if err := d.Set("applications", apps); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId(fmt.Sprintf("id-%s-applications", c.CompanyID))
	return diags
}

func flattenApp(app *dv.App) (map[string]interface{}, error) {
	a := map[string]interface{}{
		"environment_id":  app.CompanyID,
		"customer_id":     app.CustomerID,
		"name":            app.Name,
		"created_date":    app.CreatedDate,
		"api_key_enabled": app.APIKeyEnabled,

		"api_keys": map[string]interface{}{
			"prod": app.APIKeys.Prod,
			"test": app.APIKeys.Test,
		},
		"metadata": map[string]interface{}{
			"rp_name": app.Metadata.RpName,
		},
		//User pools seems to always only have one array item
		"user_pools": map[string]interface{}{
			"connection_id": app.UserPools[0].ConnectionID,
			"connector_id":  app.UserPools[0].ConnectorID,
		},
	}

	if app.Saml != nil {
		if app.Saml.Values != nil {
			smap := []map[string]interface{}{{"values": []map[string]interface{}{}}}
			smap[0] = map[string]interface{}{
				"values": []map[string]interface{}{{
					"enabled":                app.Saml.Values.Enabled,
					"redirect_uri":           app.Saml.Values.RedirectURI,
					"enforce_signed_request": app.Saml.Values.EnforceSignedRequest,
					"audience":               app.Saml.Values.Audience,
					"sp_cert":                app.Saml.Values.SpCert,
				}},
			}
			a["saml"] = smap
		}
	}

	if app.Oauth != nil {

		amap := []map[string]interface{}{{"enabled": app.Oauth.Enabled, "values": []map[string]interface{}{}}}
		if app.Oauth.Values != nil {
			amap[0] = map[string]interface{}{"enabled": app.Oauth.Enabled, "values": []map[string]interface{}{{
				"enabled":                       app.Oauth.Values.Enabled,
				"client_secret":                 app.Oauth.Values.ClientSecret,
				"enforce_signed_request_openid": app.Oauth.Values.EnforceSignedRequestOpenid,
				"sp_jwks_url":                   app.Oauth.Values.SpjwksUrl,
				"sp_jwks_openid":                app.Oauth.Values.SpJwksOpenid,
				"redirect_uris":                 app.Oauth.Values.RedirectUris,
				"logout_uris":                   app.Oauth.Values.LogoutUris,
				"allowed_scopes":                app.Oauth.Values.AllowedScopes,
				"allowed_grants":                app.Oauth.Values.AllowedGrants,
			}}}
		}

		a["oauth"] = amap
	}

	//User Portal
	if app.UserPortal != nil {
		if app.UserPortal.Values != nil {
			a["user_portal"] = []map[string]interface{}{{
				"up_title":                 app.UserPortal.Values.UpTitle,
				"add_auth_method_title":    app.UserPortal.Values.AddAuthMethodTitle,
				"remove_auth_method_title": app.UserPortal.Values.RemoveAuthMethodTitle,
				"cred_page_title":          app.UserPortal.Values.CredentialPageTitle,
				"cred_page_subtitle":       app.UserPortal.Values.CredentialPageSubTitle,
				"name_auth_method_title":   app.UserPortal.Values.NameAuthMethodTitle,
				"name_confirm_btn_text":    app.UserPortal.Values.NameConfirmButtonText,
				"update_message":           app.UserPortal.Values.UpdateMessage,
				"update_body_message":      app.UserPortal.Values.UpdateBodyMessage,
				"remove_message":           app.UserPortal.Values.RemoveMessage,
				"remove_body_message":      app.UserPortal.Values.RemoveBodyMessage,
				"remove_confirm_btn_text":  app.UserPortal.Values.RemoveConfirmButtonText,
				"remove_cancel_btn_text":   app.UserPortal.Values.RemoveCancelButtonText,
				"flow_timeout_seconds":     app.UserPortal.Values.FlowTimeoutInSeconds,
				"show_user_info":           app.UserPortal.Values.ShowUserInfo,
				"show_mfa_button":          app.UserPortal.Values.ShowMfaButton,
				"show_variables":           app.UserPortal.Values.ShowVariables,
				"show_logout_button":       app.UserPortal.Values.ShowLogoutButton,
			}}
		}
	}

	//Policies
	pols := []interface{}{}

	for _, v := range app.Policies {
		polFlows := []interface{}{}
		for _, w := range v.PolicyFlows {
			thisPolFlow := map[string]interface{}{
				"flow_id":    w.FlowID,
				"weight":     w.Weight,
				"version_id": w.VersionID,
			}
			polFlows = append(polFlows, thisPolFlow)
		}

		pols = append(pols, map[string]interface{}{
			"policy_flow":  polFlows,
			"name":         v.Name,
			"status":       v.Status,
			"policy_id":    v.PolicyID,
			"created_date": v.CreatedDate,
		})
	}
	a["policy"] = pols

	//Return
	return a, nil
}

func readAllApplications(ctx context.Context, c *dv.APIClient, environmentID string, timeout time.Duration) ([]dv.App, error) {

	applications := -1

	stateConf := &retry.StateChangeConf{
		Pending: []string{
			"false",
		},
		Target: []string{
			"true",
			"err",
		},
		Refresh: func() (interface{}, string, error) {

			// Run the API call
			sdkRes, err := sdk.DoRetryable(
				ctx,
				c,
				environmentID,
				func() (interface{}, *http.Response, error) {
					return c.ReadApplicationsWithResponse(&environmentID, nil)
				},
			)

			if err != nil {
				return nil, "err", err
			}

			res, ok := sdkRes.([]dv.App)
			if !ok {
				err = fmt.Errorf("Unable to parse applications response from Davinci API")
				return nil, "err", err
			}

			// If the number of applications has changed since last time, we need to keep waiting
			if len(res) != applications {
				applications = len(res)
				return res, "false", nil
			}

			return res, "true", nil
		},
		Timeout:                   timeout - time.Minute,
		Delay:                     20 * time.Second,
		MinTimeout:                2 * time.Second,
		ContinuousTargetOccurence: 5, // we want five consecutive successful reads of the same number of applications
	}
	sdkRes, err := stateConf.WaitForStateContext(ctx)
	if err != nil {
		return nil, err
	}

	res, ok := sdkRes.([]dv.App)
	if !ok {
		err = fmt.Errorf("Unable to parse applications response from Davinci API")
		return nil, err
	}

	return res, nil

}
