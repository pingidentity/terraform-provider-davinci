package davinci

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

func DataSourceApplication() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplicationRead,
		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "DEPRECATED: Use field 'id'. id of the application to retrieve.",
				ExactlyOneOf: []string{"id", "application_id"},
			},
			"id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "id of the application to retrieve.",
				ExactlyOneOf: []string{"id", "application_id"},
			},
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PingOne environment id",
			},
			"customer_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Internal used field",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Application name",
			},
			"created_date": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Resource creation date",
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
				Type:        schema.TypeMap,
				Computed:    true,
				Description: "Internal read only field.",
				Elem: &schema.Schema{
					Type:        schema.TypeString,
					Description: "Internal read only field.",
				},
			},
			"user_portal": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "This is deprecated in the UI and will be removed in a future release.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"up_title": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "This is deprecated in the UI and will be removed in a future release.",
						},
						"add_auth_method_title": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "This is deprecated in the UI and will be removed in a future release.",
						},
						"remove_auth_method_title": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "This is deprecated in the UI and will be removed in a future release.",
						},
						"cred_page_title": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "This is deprecated in the UI and will be removed in a future release.",
						},
						"cred_page_subtitle": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "This is deprecated in the UI and will be removed in a future release.",
						},
						"name_auth_method_title": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "This is deprecated in the UI and will be removed in a future release.",
						},
						"name_confirm_btn_text": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "This is deprecated in the UI and will be removed in a future release.",
						},
						"update_message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "This is deprecated in the UI and will be removed in a future release.",
						},
						"update_body_message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "This is deprecated in the UI and will be removed in a future release.",
						},
						"remove_message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "This is deprecated in the UI and will be removed in a future release.",
						},
						"remove_body_message": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "This is deprecated in the UI and will be removed in a future release.",
						},
						"remove_confirm_btn_text": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "This is deprecated in the UI and will be removed in a future release.",
						},
						"remove_cancel_btn_text": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "This is deprecated in the UI and will be removed in a future release.",
						},
						"flow_timeout_seconds": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "This is deprecated in the UI and will be removed in a future release.",
						},
						"show_user_info": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "This is deprecated in the UI and will be removed in a future release.",
						},
						"show_mfa_button": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "This is deprecated in the UI and will be removed in a future release.",
						},
						"show_variables": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "This is deprecated in the UI and will be removed in a future release.",
						},
						"show_logout_button": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "This is deprecated in the UI and will be removed in a future release.",
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
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "This will automatically be set to true if the oauth block is used.",
						},
						"values": {
							Type:        schema.TypeSet,
							Computed:    true,
							Description: "OIDC configuration",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Set to true if using oauth block",
									},
									"client_secret": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The client secret for the OIDC application.",
										Sensitive:   true,
									},
									"enforce_signed_request_openid": {
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Field: 'Enforce Receiving Signed Requests' in UI.",
									},
									"sp_jwks_url": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Field: 'Service Provider (SP) JWKS URL' in UI.",
									},
									"sp_jwks_openid": {
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Field: 'Service Provider (SP) JWKS Keys to Verify Authorization Request Signature' in UI. ",
									},
									"redirect_uris": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Redirect URLs for the application",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"logout_uris": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Logout URLs for the application",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"allowed_scopes": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Allowed scopes for the application. Available scopes are: openid, profile, flow_analytics.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"allowed_grants": {
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "Allowed grants for the application. Available grants are: authorizationCode, clientCredentials, implicit. ",
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
										Type:        schema.TypeBool,
										Computed:    true,
										Description: "Set to true if using saml block. This is deprecated in the UI and will be removed in a future release.",
									},
									"redirect_uri": {
										Type:        schema.TypeString,
										Computed:    true,
										Sensitive:   true,
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
				Description: "Flow Policy Configuration",
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
										Description: "List of node ids to be used in analytics.",
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
							Description: "Returned identifier of a created policy.",
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
	}
}

func dataSourceApplicationRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	// TODO: remove this once application_id is removed
	appIdInt, ok := d.GetOk("application_id")
	if ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "application_id is deprecated, please use id instead",
			Detail:   "application_id is deprecated and will be removed in a future release, please use id instead",
		})
	}
	if !ok {
		appIdInt, ok = d.GetOk("id")
		if !ok {
			return diag.FromErr(fmt.Errorf("id must be set"))
		}
	}
	appId := appIdInt.(string)

	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.ReadApplication(&c.CompanyID, appId)
	}, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	resp, ok := sdkRes.(*dv.App)
	if !ok {
		err = fmt.Errorf("failed to cast response to Application for id: %s", appId)
		return diag.FromErr(err)
	}

	flatResp, err := flattenApp(resp)
	if err != nil {
		return diag.FromErr(err)
	}
	for i, v := range flatResp {
		err := d.Set(i, v)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// TODO: remove this once application_id is removed
	if err = d.Set("application_id", resp.AppID); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(resp.AppID)

	return diags
}
