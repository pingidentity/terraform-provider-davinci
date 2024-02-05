package davinci

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
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
				Description:  "The ID of the application to retrieve.",
				ExactlyOneOf: []string{"id", "application_id"},
			},
			"id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "The ID of the application.  Use of this parameter to fetch the application data is deprecated, use the `application_id` parameter instead.",
				ExactlyOneOf: []string{"id", "application_id"},
			},
			"environment_id": {
				Type:             schema.TypeString,
				Required:         true,
				Description:      "The ID of the PingOne environment to create the DaVinci application. Must be a valid PingOne resource ID.",
				ValidateDiagFunc: validation.ToDiagFunc(verify.ValidP1ResourceID),
			},
			"customer_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An ID that represents the customer tenant.",
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The application name.",
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
				Description: "A single list item specifying OIDC/OAuth 2.0 configuration.",
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
										Description: "The client secret for the OIDC application.",
										Sensitive:   true,
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

	appIdInt, ok := d.GetOk("id")
	if ok {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Use of `id` to select an application is deprecated, please use `application_id` instead",
			Detail:   "The use of `id` is deprecated and will be made a computed attribute in a future release, please use `application_id` instead",
		})
	}
	if !ok {
		appIdInt, ok = d.GetOk("application_id")
		if !ok {
			diags = append(diags, diag.FromErr(fmt.Errorf("application_id must be set"))...)
			return diags
		}
	}
	appId := appIdInt.(string)

	environmentID := d.Get("environment_id").(string)

	sdkRes, err := sdk.DoRetryable(
		ctx,
		c,
		environmentID,
		func() (interface{}, *http.Response, error) {
			return c.ReadApplicationWithResponse(environmentID, appId)
		},
	)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	resp, ok := sdkRes.(*dv.App)
	if !ok {
		err = fmt.Errorf("failed to cast response to Application for id: %s", appId)
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	flatResp, err := flattenApp(resp)
	if err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}
	for i, v := range flatResp {
		err := d.Set(i, v)
		if err != nil {
			diags = append(diags, diag.FromErr(err)...)
			return diags
		}
	}

	if err = d.Set("application_id", resp.AppID); err != nil {
		diags = append(diags, diag.FromErr(err)...)
		return diags
	}

	d.SetId(resp.AppID)

	return diags
}
