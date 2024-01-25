package davinci_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
)

func TestAccDataSourceApplications_AllApplications_Clean(t *testing.T) {
	testAccDataSourceApplications_AllApplications(t, false)
}

func TestAccDataSourceApplications_AllApplications_WithBootstrap(t *testing.T) {
	testAccDataSourceApplications_AllApplications(t, true)
}

func testAccDataSourceApplications_AllApplications(t *testing.T, withBootstrapConfig bool) {

	resourceBase := "davinci_applications"
	resourceName := acctest.ResourceNameGen()
	dataSourceFullName := fmt.Sprintf("data.%s.%s", resourceBase, resourceName)

	name := resourceName

	hcl := testAccDataSourceApplications_AllApplications_Hcl(resourceName, name, withBootstrapConfig)

	var applicationsCountRegex *regexp.Regexp
	if withBootstrapConfig {
		applicationsCountRegex = regexp.MustCompile(`^[456789]$`) // up to 9
	} else {
		applicationsCountRegex = regexp.MustCompile(`^3$`)
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		//CheckDestroy:       acctest.CheckResourceDestroy([]string{"davinci_application"}),
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(dataSourceFullName, "id", regexp.MustCompile(fmt.Sprintf(`^id-%s-applications$`, verify.P1ResourceIDRegexp.String()))),
					resource.TestMatchResourceAttr(dataSourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
					resource.TestMatchResourceAttr(dataSourceFullName, "applications.#", applicationsCountRegex),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceFullName, "applications.*", map[string]*regexp.Regexp{
						"id":                                verify.P1DVResourceIDRegexpFullString,
						"environment_id":                    verify.P1ResourceIDRegexpFullString,
						"customer_id":                       verify.P1DVResourceIDRegexpFullString,
						"name":                              regexp.MustCompile(fmt.Sprintf(`^%s-1$`, name)),
						"created_date":                      verify.EpochDateRegexpFullString,
						"api_key_enabled":                   regexp.MustCompile(`^false$`),
						"api_keys.%":                        regexp.MustCompile(`^2$`),
						"api_keys.prod":                     regexp.MustCompile(`^[a-z0-9]+$`),
						"api_keys.test":                     regexp.MustCompile(`^[a-z0-9]+$`),
						"metadata.rp_name":                  regexp.MustCompile(fmt.Sprintf(`^%s-1$`, name)),
						"user_pools.%":                      regexp.MustCompile(`^2$`),
						"user_pools.connection_id":          regexp.MustCompile(`^defaultUserPool$`),
						"user_pools.connector_id":           regexp.MustCompile(`^skUserPool$`),
						"user_portal.#":                     regexp.MustCompile(`^0$`),
						"oauth.#":                           regexp.MustCompile(`^1$`),
						"oauth.0.enabled":                   regexp.MustCompile(`^false$`),
						"oauth.0.values.#":                  regexp.MustCompile(`^1$`),
						"oauth.0.values.0.enabled":          regexp.MustCompile(`^false$`),
						"oauth.0.values.0.allowed_grants.#": regexp.MustCompile(`^2$`),
						"oauth.0.values.0.allowed_grants.0": regexp.MustCompile(`^implicit|authorizationCode$`),
						"oauth.0.values.0.allowed_grants.1": regexp.MustCompile(`^implicit|authorizationCode$`),
						"oauth.0.values.0.allowed_scopes.#": regexp.MustCompile(`^2$`),
						"oauth.0.values.0.allowed_scopes.0": regexp.MustCompile(`^openid|profile$`),
						"oauth.0.values.0.allowed_scopes.1": regexp.MustCompile(`^openid|profile$`),
						"oauth.0.values.0.client_secret":    regexp.MustCompile(`^[a-z0-9]+$`),
						"oauth.0.values.0.enforce_signed_request_openid": regexp.MustCompile(`^false$`),
						"oauth.0.values.0.redirect_uris.#":               regexp.MustCompile(`^3$`),
						"oauth.0.values.0.redirect_uris.0":               regexp.MustCompile(`^https:\/\/auth\.ping-eng\.com\/env-id\/rp\/callback\/openid_connect|https:\/\/auth\.ping-eng\.com\/env-id\/rp\/callback\/oidc|https:\/\/auth\.ping-eng\.com\/env-id\/rp\/callback\/1$`),
						"oauth.0.values.0.redirect_uris.1":               regexp.MustCompile(`^https:\/\/auth\.ping-eng\.com\/env-id\/rp\/callback\/openid_connect|https:\/\/auth\.ping-eng\.com\/env-id\/rp\/callback\/oidc|https:\/\/auth\.ping-eng\.com\/env-id\/rp\/callback\/1$`),
						"oauth.0.values.0.redirect_uris.2":               regexp.MustCompile(`^https:\/\/auth\.ping-eng\.com\/env-id\/rp\/callback\/openid_connect|https:\/\/auth\.ping-eng\.com\/env-id\/rp\/callback\/oidc|https:\/\/auth\.ping-eng\.com\/env-id\/rp\/callback\/1$`),
						"oauth.0.values.0.logout_uris.#":                 regexp.MustCompile(`^2$`),
						"oauth.0.values.0.logout_uris.0":                 regexp.MustCompile(`^https:\/\/auth\.ping-eng\.com\/env-id\/logout1|https:\/\/auth\.ping-eng\.com\/env-id\/logout$`),
						"oauth.0.values.0.logout_uris.1":                 regexp.MustCompile(`^https:\/\/auth\.ping-eng\.com\/env-id\/logout1|https:\/\/auth\.ping-eng\.com\/env-id\/logout$`),
						"oauth.0.values.0.sp_jwks_openid":                regexp.MustCompile(`^$`),
						"oauth.0.values.0.sp_jwks_url":                   regexp.MustCompile(`^https:\/\/www\.ping-eng\.com\/testjwks$`),
						"saml.#":                                         regexp.MustCompile(`^1$`),
						"policy.#":                                       regexp.MustCompile(`^0$`),
					}),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceFullName, "applications.*", map[string]*regexp.Regexp{
						"id":                                verify.P1DVResourceIDRegexpFullString,
						"environment_id":                    verify.P1ResourceIDRegexpFullString,
						"customer_id":                       verify.P1DVResourceIDRegexpFullString,
						"name":                              regexp.MustCompile(fmt.Sprintf(`^%s-2$`, name)),
						"created_date":                      verify.EpochDateRegexpFullString,
						"api_key_enabled":                   regexp.MustCompile(`^true$`),
						"api_keys.%":                        regexp.MustCompile(`^2$`),
						"api_keys.prod":                     regexp.MustCompile(`^[a-z0-9]+$`),
						"api_keys.test":                     regexp.MustCompile(`^[a-z0-9]+$`),
						"metadata.rp_name":                  regexp.MustCompile(fmt.Sprintf(`^%s-2$`, name)),
						"user_pools.%":                      regexp.MustCompile(`^2$`),
						"user_pools.connection_id":          regexp.MustCompile(`^defaultUserPool$`),
						"user_pools.connector_id":           regexp.MustCompile(`^skUserPool$`),
						"user_portal.#":                     regexp.MustCompile(`^0$`),
						"oauth.#":                           regexp.MustCompile(`^1$`),
						"oauth.0.enabled":                   regexp.MustCompile(`^true$`),
						"oauth.0.values.#":                  regexp.MustCompile(`^1$`),
						"oauth.0.values.0.enabled":          regexp.MustCompile(`^true$`),
						"oauth.0.values.0.allowed_grants.#": regexp.MustCompile(`^1$`),
						"oauth.0.values.0.allowed_grants.0": regexp.MustCompile(`^authorizationCode$`),
						"oauth.0.values.0.allowed_scopes.#": regexp.MustCompile(`^2$`),
						"oauth.0.values.0.allowed_scopes.0": regexp.MustCompile(`^openid|profile$`),
						"oauth.0.values.0.allowed_scopes.1": regexp.MustCompile(`^openid|profile$`),
						"oauth.0.values.0.client_secret":    regexp.MustCompile(`^[a-z0-9]+$`),
						"oauth.0.values.0.enforce_signed_request_openid": regexp.MustCompile(`^false$`),
						"oauth.0.values.0.redirect_uris.#":               regexp.MustCompile(`^0$`),
						"oauth.0.values.0.logout_uris.#":                 regexp.MustCompile(`^0$`),
						"oauth.0.values.0.sp_jwks_openid":                regexp.MustCompile(`^$`),
						"oauth.0.values.0.sp_jwks_url":                   regexp.MustCompile(`^$`),
						"saml.#":                                         regexp.MustCompile(`^1$`),
						"policy.#":                                       regexp.MustCompile(`^0$`),
					}),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceFullName, "applications.*", map[string]*regexp.Regexp{
						"id":                                verify.P1DVResourceIDRegexpFullString,
						"environment_id":                    verify.P1ResourceIDRegexpFullString,
						"customer_id":                       verify.P1DVResourceIDRegexpFullString,
						"name":                              regexp.MustCompile(fmt.Sprintf(`^%s-3$`, name)),
						"created_date":                      verify.EpochDateRegexpFullString,
						"api_key_enabled":                   regexp.MustCompile(`^true$`),
						"api_keys.%":                        regexp.MustCompile(`^2$`),
						"api_keys.prod":                     regexp.MustCompile(`^[a-z0-9]+$`),
						"api_keys.test":                     regexp.MustCompile(`^[a-z0-9]+$`),
						"metadata.rp_name":                  regexp.MustCompile(fmt.Sprintf(`^%s-3$`, name)),
						"user_pools.%":                      regexp.MustCompile(`^2$`),
						"user_pools.connection_id":          regexp.MustCompile(`^defaultUserPool$`),
						"user_pools.connector_id":           regexp.MustCompile(`^skUserPool$`),
						"user_portal.#":                     regexp.MustCompile(`^0$`),
						"oauth.#":                           regexp.MustCompile(`^1$`),
						"oauth.0.enabled":                   regexp.MustCompile(`^true$`),
						"oauth.0.values.#":                  regexp.MustCompile(`^1$`),
						"oauth.0.values.0.enabled":          regexp.MustCompile(`^true$`),
						"oauth.0.values.0.allowed_grants.#": regexp.MustCompile(`^1$`),
						"oauth.0.values.0.allowed_grants.0": regexp.MustCompile(`^authorizationCode$`),
						"oauth.0.values.0.allowed_scopes.#": regexp.MustCompile(`^2$`),
						"oauth.0.values.0.allowed_scopes.0": regexp.MustCompile(`^openid|profile$`),
						"oauth.0.values.0.allowed_scopes.1": regexp.MustCompile(`^openid|profile$`),
						"oauth.0.values.0.client_secret":    regexp.MustCompile(`^[a-z0-9]+$`),
						"oauth.0.values.0.enforce_signed_request_openid": regexp.MustCompile(`^false$`),
						"oauth.0.values.0.redirect_uris.#":               regexp.MustCompile(`^1$`),
						"oauth.0.values.0.redirect_uris.0":               regexp.MustCompile(`^https:\/\/auth\.ping-eng\.com\/env-id\/rp\/callback\/openid_connect$`),
						"oauth.0.values.0.logout_uris.#":                 regexp.MustCompile(`^0$`),
						"oauth.0.values.0.sp_jwks_openid":                regexp.MustCompile(`^$`),
						"oauth.0.values.0.sp_jwks_url":                   regexp.MustCompile(`^$`),
						"saml.#":                                         regexp.MustCompile(`^1$`),
						"policy.#":                                       regexp.MustCompile(`^0$`),
					}),
				),
			},
		},
	})
}

func testAccDataSourceApplications_AllApplications_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s-1" {
	environment_id  = pingone_environment.%[2]s.id
	name            = "%[3]s-1"
	api_key_enabled = false
  
	oauth {
	  enabled = false
	  values {
		enabled                       = false
		allowed_grants                = ["implicit", "authorizationCode"]
		allowed_scopes                = ["openid", "profile"]
		enforce_signed_request_openid = false
		redirect_uris = [
		  "https://auth.ping-eng.com/env-id/rp/callback/openid_connect",
		  "https://auth.ping-eng.com/env-id/rp/callback/oidc",
		  "https://auth.ping-eng.com/env-id/rp/callback/1",
		]
		logout_uris = [
		  "https://auth.ping-eng.com/env-id/logout1",
		  "https://auth.ping-eng.com/env-id/logout"
		]
		sp_jwks_url = "https://www.ping-eng.com/testjwks"
	  }
	}
  }

  resource "davinci_application" "%[2]s-2" {
	environment_id = pingone_environment.%[2]s.id
	name           = "%[3]s-2"
  }

  resource "davinci_application" "%[2]s-3" {
	environment_id = pingone_environment.%[2]s.id
	name           = "%[3]s-3"
  
	oauth {
	  values {
		allowed_grants = ["authorizationCode"]
		allowed_scopes = ["openid", "profile"]
		redirect_uris = [
		  "https://auth.ping-eng.com/env-id/rp/callback/openid_connect",
		]
	  }
	}
  }

data "davinci_applications" "%[2]s" {
  environment_id = resource.pingone_environment.%[2]s.id

  depends_on = [
    davinci_application.%[2]s-1,
    davinci_application.%[2]s-2,
    davinci_application.%[2]s-3,
  ]
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}
