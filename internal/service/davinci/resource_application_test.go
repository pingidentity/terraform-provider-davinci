package davinci_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest/service/davinci"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
)

func TestAccResourceApplication_RemovalDrift(t *testing.T) {

	resourceBase := "davinci_application"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	hcl := testAccResourceApplication_Full_HCL(resourceName, resourceName, true)

	var applicationID, environmentID string

	// ctx := context.Background()

	// p1Client, err := acctest.PingOneTestClient(ctx)
	// if err != nil {
	// 	t.Fatalf("Failed to get API client: %v", err)
	// }

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Application_CheckDestroy(),
		Steps: []resource.TestStep{
			// Configure
			{
				Config: hcl,
				Check:  davinci.Application_GetIDs(resourceFullName, &environmentID, &applicationID),
			},
			// Replan after removal preconfig
			{
				PreConfig: func() {
					davinci.Application_RemovalDrift_PreConfig(t, environmentID, applicationID)
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
			},
			// Test removal of the environment
			{
				Config:   hcl,
				Check:    davinci.Application_GetIDs(resourceFullName, &environmentID, &applicationID),
				SkipFunc: func() (bool, error) { return true, nil },
			},
			{
				// PreConfig: func() {
				// 	base.Environment_RemovalDrift_PreConfig(ctx, p1Client.API.ManagementAPIClient, t, environmentID)
				// },
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
				SkipFunc:           func() (bool, error) { return true, nil },
			},
		},
	})
}

func TestAccResourceApplication_Full_Clean(t *testing.T) {
	testAccResourceApplication_Full(t, false)
}

func TestAccResourceApplication_Full_WithBootstrap(t *testing.T) {
	testAccResourceApplication_Full(t, true)
}

func testAccResourceApplication_Full(t *testing.T, withBootstrapConfig bool) {

	resourceBase := "davinci_application"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	fullStep := resource.TestStep{
		Config: testAccResourceApplication_Full_HCL(resourceName, fmt.Sprintf("%s-1", name), withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", fmt.Sprintf("%s-1", name)),
			resource.TestCheckResourceAttr(resourceFullName, "api_key_enabled", "false"),
			resource.TestCheckResourceAttr(resourceFullName, "api_keys.%", "2"),
			resource.TestMatchResourceAttr(resourceFullName, "api_keys.prod", regexp.MustCompile(`^[a-z0-9]+$`)),
			resource.TestMatchResourceAttr(resourceFullName, "api_keys.test", regexp.MustCompile(`^[a-z0-9]+$`)),
			resource.TestCheckResourceAttr(resourceFullName, "metadata.rp_name", fmt.Sprintf("%s-1", name)),
			resource.TestCheckResourceAttr(resourceFullName, "user_pools.%", "2"),
			resource.TestCheckResourceAttr(resourceFullName, "user_pools.connection_id", "defaultUserPool"),
			resource.TestCheckResourceAttr(resourceFullName, "user_pools.connector_id", "skUserPool"),
			resource.TestCheckResourceAttr(resourceFullName, "user_portal.#", "0"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "saml.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "policy.#", "0"),
			resource.TestMatchResourceAttr(resourceFullName, "created_date", verify.EpochDateRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "customer_id", verify.P1DVResourceIDRegexpFullString),
		),
	}

	minimalStep := resource.TestStep{
		Config: testAccResourceApplication_Minimal_HCL(resourceName, fmt.Sprintf("%s-2", name), withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", fmt.Sprintf("%s-2", name)),
			resource.TestCheckResourceAttr(resourceFullName, "api_key_enabled", "true"),
			resource.TestCheckResourceAttr(resourceFullName, "api_keys.%", "2"),
			resource.TestMatchResourceAttr(resourceFullName, "api_keys.prod", regexp.MustCompile(`^[a-z0-9]+$`)),
			resource.TestMatchResourceAttr(resourceFullName, "api_keys.test", regexp.MustCompile(`^[a-z0-9]+$`)),
			//resource.TestCheckResourceAttr(resourceFullName, "metadata.rp_name", fmt.Sprintf("%s-2", name)),
			resource.TestCheckResourceAttr(resourceFullName, "user_pools.%", "2"),
			resource.TestCheckResourceAttr(resourceFullName, "user_pools.connection_id", "defaultUserPool"),
			resource.TestCheckResourceAttr(resourceFullName, "user_pools.connector_id", "skUserPool"),
			resource.TestCheckResourceAttr(resourceFullName, "user_portal.#", "0"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "saml.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "policy.#", "0"),
			resource.TestMatchResourceAttr(resourceFullName, "created_date", verify.EpochDateRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "customer_id", verify.P1DVResourceIDRegexpFullString),
		),
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Application_CheckDestroy(),
		Steps: []resource.TestStep{
			// Create from scratch
			fullStep,
			{
				Config:  testAccResourceApplication_Full_HCL(resourceName, fmt.Sprintf("%s-1", name), withBootstrapConfig),
				Destroy: true,
			},
			// Create from scratch
			minimalStep,
			{
				Config:  testAccResourceApplication_Minimal_HCL(resourceName, fmt.Sprintf("%s-2", name), withBootstrapConfig),
				Destroy: true,
			},
			// Test updates
			fullStep,
			minimalStep,
			fullStep,
			// Test importing the resource
			{
				ResourceName: resourceFullName,
				ImportStateIdFunc: func() resource.ImportStateIdFunc {
					return func(s *terraform.State) (string, error) {
						rs, ok := s.RootModule().Resources[resourceFullName]
						if !ok {
							return "", fmt.Errorf("Resource Not found: %s", resourceFullName)
						}

						return fmt.Sprintf("%s/%s", rs.Primary.Attributes["environment_id"], rs.Primary.ID), nil
					}
				}(),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceApplication_WithOAuth_Full(t *testing.T) {

	withBootstrapConfig := false

	resourceBase := "davinci_application"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	fullStep := resource.TestStep{
		Config: testAccResourceApplication_WithOAuth_Full_HCL(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "oauth.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.enabled", "false"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.enabled", "false"),
			resource.TestMatchResourceAttr(resourceFullName, "oauth.0.values.0.client_secret", regexp.MustCompile(`^[a-z0-9]+$`)),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.enforce_signed_request_openid", "false"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.sp_jwks_url", "https://www.ping-eng.com/testjwks"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.sp_jwks_openid", ""),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.redirect_uris.#", "3"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.redirect_uris.*", "https://auth.ping-eng.com/env-id/rp/callback/oidc"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.redirect_uris.*", "https://auth.ping-eng.com/env-id/rp/callback/1"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.redirect_uris.*", "https://auth.ping-eng.com/env-id/rp/callback/openid_connect"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.logout_uris.#", "2"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.logout_uris.*", "https://auth.ping-eng.com/env-id/logout"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.logout_uris.*", "https://auth.ping-eng.com/env-id/logout1"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.allowed_scopes.#", "2"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_scopes.*", "openid"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_scopes.*", "profile"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.allowed_grants.#", "2"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_grants.*", "authorizationCode"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_grants.*", "implicit"),
		),
	}

	minimalStep1 := resource.TestStep{
		Config: testAccResourceApplication_WithOAuth_Minimal1_HCL(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "oauth.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.enabled", "true"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.enabled", "true"),
			resource.TestMatchResourceAttr(resourceFullName, "oauth.0.values.0.client_secret", regexp.MustCompile(`^[a-z0-9]+$`)),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.enforce_signed_request_openid", "false"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.sp_jwks_url", ""),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.sp_jwks_openid", ""),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.redirect_uris.#", "0"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.logout_uris.#", "0"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.allowed_scopes.#", "2"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_scopes.*", "openid"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_scopes.*", "profile"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.allowed_grants.#", "1"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_grants.*", "authorizationCode"),
		),
	}

	minimalStep2 := resource.TestStep{
		Config: testAccResourceApplication_WithOAuth_Minimal2_HCL(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "oauth.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.enabled", "true"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.enabled", "true"),
			resource.TestMatchResourceAttr(resourceFullName, "oauth.0.values.0.client_secret", regexp.MustCompile(`^[a-z0-9]+$`)),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.enforce_signed_request_openid", "false"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.sp_jwks_url", ""),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.sp_jwks_openid", ""),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.redirect_uris.#", "0"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.logout_uris.#", "0"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.allowed_scopes.#", "0"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.allowed_grants.#", "0"),
		),
	}

	minimalStep3 := resource.TestStep{
		Config: testAccResourceApplication_WithOAuth_Minimal3_HCL(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "oauth.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.enabled", "true"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.enabled", "true"),
			resource.TestMatchResourceAttr(resourceFullName, "oauth.0.values.0.client_secret", regexp.MustCompile(`^[a-z0-9]+$`)),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.enforce_signed_request_openid", "false"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.sp_jwks_url", ""),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.sp_jwks_openid", ""),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.redirect_uris.#", "1"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.redirect_uris.*", "https://auth.ping-eng.com/env-id/rp/callback/openid_connect"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.logout_uris.#", "0"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.allowed_scopes.#", "2"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_scopes.*", "openid"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_scopes.*", "profile"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.allowed_grants.#", "1"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_grants.*", "authorizationCode"),
		),
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Application_CheckDestroy(),
		Steps: []resource.TestStep{
			// Create from scratch
			fullStep,
			{
				Config:  testAccResourceApplication_WithOAuth_Full_HCL(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// Create from scratch
			minimalStep1,
			{
				Config:  testAccResourceApplication_WithOAuth_Minimal1_HCL(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// Create from scratch
			minimalStep2,
			{
				Config:  testAccResourceApplication_WithOAuth_Minimal2_HCL(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// Create from scratch
			minimalStep3,
			{
				Config:  testAccResourceApplication_WithOAuth_Minimal3_HCL(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// Test updates
			fullStep,
			// Skipping steps because of SDKv2 bug for computed blocks
			// minimalStep1,
			// fullStep,
			// minimalStep2,
			// fullStep,
			minimalStep3,
			fullStep,
			// Test importing the resource
			{
				ResourceName: resourceFullName,
				ImportStateIdFunc: func() resource.ImportStateIdFunc {
					return func(s *terraform.State) (string, error) {
						rs, ok := s.RootModule().Resources[resourceFullName]
						if !ok {
							return "", fmt.Errorf("Resource Not found: %s", resourceFullName)
						}

						return fmt.Sprintf("%s/%s", rs.Primary.Attributes["environment_id"], rs.Primary.ID), nil
					}
				}(),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceApplication_BadParameters(t *testing.T) {

	resourceBase := "davinci_application"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	hcl := testAccResourceApplication_Full_HCL(resourceName, resourceName, false)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Application_CheckDestroy(),
		Steps: []resource.TestStep{
			// Configure
			{
				Config: hcl,
			},
			// Errors
			{
				ResourceName: resourceFullName,
				ImportState:  true,
				ExpectError:  regexp.MustCompile(`Invalid import ID specified \(".*"\).  The ID should be in the format "environment_id/davinci_application_id" and must match regex: .*`),
			},
			{
				ResourceName:  resourceFullName,
				ImportStateId: "/",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`Invalid import ID specified \(".*"\).  The ID should be in the format "environment_id/davinci_application_id" and must match regex: .*`),
			},
			{
				ResourceName:  resourceFullName,
				ImportStateId: "badformat/badformat",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`Invalid import ID specified \(".*"\).  The ID should be in the format "environment_id/davinci_application_id" and must match regex: .*`),
			},
		},
	})
}

func testAccResourceApplication_Full_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
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
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceApplication_Minimal_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  name           = "%[3]s"
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceApplication_WithOAuth_Full_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
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
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceApplication_WithOAuth_Minimal1_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  name           = "%[3]s"

  oauth {}
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceApplication_WithOAuth_Minimal2_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  name           = "%[3]s"

  oauth {
    values {}
  }
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceApplication_WithOAuth_Minimal3_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  name           = "%[3]s"

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
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}
