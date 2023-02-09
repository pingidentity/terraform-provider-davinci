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
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the application to retrieve.",
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
	}
}

func dataSourceApplicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}

	appId := d.Get("application_id").(string)
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
	for i, v := range flatResp {
		d.Set(i, v)
	}

	d.SetId(resp.AppID)
	return diags
}
