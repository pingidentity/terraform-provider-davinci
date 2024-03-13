package davinci_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
)

func TestAccDataSourceApplication_DefinedApplication_ByID_Clean(t *testing.T) {
	testAccDataSourceApplication_DefinedApplication_ByID(t, false)
}

func TestAccDataSourceApplication_DefinedApplication_ByID_WithBootstrap(t *testing.T) {
	testAccDataSourceApplication_DefinedApplication_ByID(t, true)
}

func testAccDataSourceApplication_DefinedApplication_ByID(t *testing.T, withBootstrapConfig bool) {

	resourceBase := "davinci_application"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)
	dataSourceFullName := fmt.Sprintf("data.%s", resourceFullName)

	name := resourceName

	hcl := testAccDataSourceApplication_DefinedApplication_ByID_Hcl(resourceName, name, withBootstrapConfig)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		//CheckDestroy:       acctest.CheckResourceDestroy([]string{"davinci_application"}),
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair(dataSourceFullName, "id", resourceFullName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "environment_id", resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "name", resourceFullName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "api_key_enabled", resourceFullName, "api_key_enabled"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "api_keys", resourceFullName, "api_keys"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "metadata", resourceFullName, "metadata"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "user_pools", resourceFullName, "user_pools"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "user_portal", resourceFullName, "user_portal"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "oauth", resourceFullName, "oauth"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "saml", resourceFullName, "saml"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "policy", resourceFullName, "policy"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "created_date", resourceFullName, "created_date"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "customer_id", resourceFullName, "customer_id"),
				),
			},
		},
	})
}

func TestAccDataSourceApplication_BootstrapApplication_ByID_WithBootstrap(t *testing.T) {

	withBootstrapConfig := true

	resourceBase := "davinci_application"
	resourceName := acctest.ResourceNameGen()
	dataSourceFullName := fmt.Sprintf("data.%s.%s", resourceBase, resourceName)

	name := resourceName

	hcl := testAccDataSourceApplication_BootstrapApplication_ByID_Hcl(resourceName, name, withBootstrapConfig)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		//CheckDestroy:       acctest.CheckResourceDestroy([]string{"davinci_application"}),
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(dataSourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
					resource.TestMatchResourceAttr(dataSourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
					resource.TestCheckResourceAttr(dataSourceFullName, "name", "PingOne SSO Connection"),
					resource.TestCheckResourceAttr(dataSourceFullName, "api_key_enabled", "true"),
					resource.TestCheckResourceAttr(dataSourceFullName, "api_keys.%", "2"),
					resource.TestMatchResourceAttr(dataSourceFullName, "api_keys.prod", regexp.MustCompile(`^[a-z0-9]+$`)),
					resource.TestMatchResourceAttr(dataSourceFullName, "api_keys.test", regexp.MustCompile(`^[a-z0-9]+$`)),
					resource.TestCheckResourceAttr(dataSourceFullName, "metadata.rp_name", "PingOne SSO Connection"),
					resource.TestCheckResourceAttr(dataSourceFullName, "user_pools.%", "2"),
					resource.TestCheckResourceAttr(dataSourceFullName, "user_pools.connection_id", "defaultUserPool"),
					resource.TestCheckResourceAttr(dataSourceFullName, "user_pools.connector_id", "skUserPool"),
					resource.TestCheckResourceAttr(dataSourceFullName, "user_portal.#", "0"),
					resource.TestCheckResourceAttr(dataSourceFullName, "oauth.#", "1"),
					resource.TestCheckResourceAttr(dataSourceFullName, "saml.#", "1"),
					resource.TestCheckResourceAttr(dataSourceFullName, "policy.#", "1"),
					resource.TestMatchResourceAttr(dataSourceFullName, "created_date", verify.EpochDateRegexpFullString),
					resource.TestMatchResourceAttr(dataSourceFullName, "customer_id", verify.P1DVResourceIDRegexpFullString),
				),
			},
		},
	})
}

func TestAccDataSourceApplication_NotFound(t *testing.T) {
	t.Parallel()

	resourceName := acctest.ResourceNameGen()

	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		//CheckDestroy:       acctest.CheckResourceDestroy([]string{"davinci_application"}),
		Steps: []resource.TestStep{
			{
				Config:      testAccDataSourceApplication_NotFoundByID(resourceName),
				ExpectError: regexp.MustCompile("App not found"),
			},
		},
	})
}

func testAccDataSourceApplication_DefinedApplication_ByID_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
  environment_id  = pingone_environment.%[2]s.id
  name            = "%[3]s"
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

data "davinci_application" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  id             = davinci_application.%[2]s.id
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccDataSourceApplication_BootstrapApplication_ByID_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

data "davinci_applications" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
}

data "davinci_application" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  application_id = data.davinci_applications.%[2]s.applications.* [index(data.davinci_applications.%[2]s.applications[*].name, "PingOne SSO Connection")].id
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccDataSourceApplication_NotFoundByID(resourceName string) string {
	return fmt.Sprintf(`


	%[1]s

data "davinci_application" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  application_id = "9c052a8a-14be-44e4-8f07-2662569994ce" // dummy ID that conforms to UUID v4
}`, acctest.PingoneEnvironmentSsoHcl(resourceName, false), resourceName)
}
