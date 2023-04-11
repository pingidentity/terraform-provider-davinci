package davinci

import (
	"context"
	"fmt"
	// "log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

func DataSourceApplications() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplicationsRead,
		Schema: map[string]*schema.Schema{
			"applications": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "application_id for this application.",
						},
						"environment_id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "PingOne environment id",
						},
						"customer_id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Application name",
						},
						"created_date": {
							Type:     schema.TypeInt,
							Computed: true,
						},
						"api_key_enabled": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Enabled by default in UI",
						},
						"api_keys": {
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Appplication Api Key",
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
							Type:     schema.TypeMap,
							Computed: true,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"user_portal": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "User Profile in UI",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"up_title": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"add_auth_method_title": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"remove_auth_method_title": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"cred_page_title": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"cred_page_subtitle": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name_auth_method_title": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"name_confirm_btn_text": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"update_message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"update_body_message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"remove_message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"remove_body_message": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"remove_confirm_btn_text": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"remove_cancel_btn_text": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"flow_timeout_seconds": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"show_user_info": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"show_mfa_button": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"show_variables": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"show_logout_button": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"oauth": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "OIDC configuration",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"values": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "OIDC configuration",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"client_secret": {
													Type:      schema.TypeString,
													Computed:  true,
													Sensitive: true,
												},
												"enforce_signed_request_openid": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"sp_jwks_url": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"sp_jwks_openid": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"redirect_uris": {
													Type:     schema.TypeSet,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"logout_uris": {
													Type:     schema.TypeSet,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"allowed_scopes": {
													Type:     schema.TypeSet,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"allowed_grants": {
													Type:     schema.TypeSet,
													Computed: true,
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
							Description: "SAML configuration",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"values": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "SAML configuration",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"enabled": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"redirect_uri": {
													Type:      schema.TypeString,
													Computed:  true,
													Sensitive: true,
												},
												"enforce_signed_request": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"audience": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"sp_cert": {
													Type:     schema.TypeString,
													Computed: true,
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
							Description: "Flow Policy Config",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"policy_flow": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Weighted flows that this Application will use",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"flow_id": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"version_id": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"weight": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"success_nodes": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
									"name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"status": {
										Type:     schema.TypeString,
										Computed: true,
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
				},
			},
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PingOne environment id",
			},
		},
	}
}

func dataSourceApplicationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.ReadApplications(&c.CompanyID, nil)
	}, nil)

	if err != nil {
		return diag.FromErr(err)
	}

	resp, ok := sdkRes.([]dv.App)
	if !ok {
		err = fmt.Errorf("Unable to cast apps response from Davinci API")
		return diag.FromErr(err)
	}

	apps := make([]interface{}, len(resp))
	for i, thisApp := range resp {
		app, err := flattenApp(&thisApp)
		if err != nil {
			return diag.FromErr(err)
		}
		app["id"] = thisApp.AppID
		apps[i] = app
	}

	if err := d.Set("applications", apps); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
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
