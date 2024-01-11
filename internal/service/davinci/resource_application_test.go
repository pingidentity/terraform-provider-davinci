package davinci_test

import (
	"context"
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest/service/base"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest/service/davinci"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
)

func TestAccResourceApplication_RemovalDrift(t *testing.T) {

	resourceBase := "davinci_application"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	hcl := testAccResourceApplication_Full_HCL(resourceName, resourceName, true)

	var applicationID, environmentID string

	ctx := context.Background()

	p1Client, err := acctest.PingOneTestClient(ctx)
	if err != nil {
		t.Fatalf("Failed to get API client: %v", err)
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		CheckDestroy:      davinci.Application_CheckDestroy(),
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
				Config: hcl,
				Check:  davinci.Application_GetIDs(resourceFullName, &environmentID, &applicationID),
				//SkipFunc: func() (bool, error) { return true, nil },
			},
			{
				PreConfig: func() {
					base.Environment_RemovalDrift_PreConfig(ctx, p1Client.API.ManagementAPIClient, t, environmentID)
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
				//SkipFunc:           func() (bool, error) { return true, nil },
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
			// resource.TestCheckResourceAttr(resourceFullName, "oauth.0.enabled", "false"),
			// resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.#", "1"),
			// resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.enabled", "false"),
			// resource.TestMatchResourceAttr(resourceFullName, "oauth.0.values.0.client_secret", regexp.MustCompile(`^[a-z0-9]+$`)),
			// resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.enforce_signed_request_openid", "1"),
			// resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.sp_jwks_url", "https://www.ping-eng.com/testjwks"),
			// resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.sp_jwks_openid", "1"),
			// resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.redirect_uris.*", "https://auth.ping-eng.com/env-id/rp/callback/oidc"),
			// resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.redirect_uris.*", "https://auth.ping-eng.com/env-id/rp/callback/1"),
			// resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.redirect_uris.*", "https://auth.ping-eng.com/env-id/rp/callback/openid_connect"),
			// resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.logout_uris.*", "https://auth.ping-eng.com/env-id/logout"),
			// resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.logout_uris.*", "https://auth.ping-eng.com/env-id/logout1"),
			// resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_scopes.*", "openid"),
			// resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_scopes.*", "profile"),
			// resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_grants", "authorizationCode"),
			// resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_grants", "implicit"),
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
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		CheckDestroy:      davinci.Application_CheckDestroy(),
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
		Config: testAccResourceApplication_WithOAuth_Full_HCL(resourceName, fmt.Sprintf("%s-1", name), withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "oauth.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.enabled", "false"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.enabled", "false"),
			resource.TestMatchResourceAttr(resourceFullName, "oauth.0.values.0.client_secret", regexp.MustCompile(`^[a-z0-9]+$`)),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.enforce_signed_request_openid", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.sp_jwks_url", "https://www.ping-eng.com/testjwks"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.sp_jwks_openid", "1"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.redirect_uris.*", "https://auth.ping-eng.com/env-id/rp/callback/oidc"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.redirect_uris.*", "https://auth.ping-eng.com/env-id/rp/callback/1"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.redirect_uris.*", "https://auth.ping-eng.com/env-id/rp/callback/openid_connect"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.logout_uris.*", "https://auth.ping-eng.com/env-id/logout"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.logout_uris.*", "https://auth.ping-eng.com/env-id/logout1"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_scopes.*", "openid"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_scopes.*", "profile"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_grants", "authorizationCode"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_grants", "implicit"),
		),
	}

	minimalStep := resource.TestStep{
		Config: testAccResourceApplication_WithOAuth_Minimal_HCL(resourceName, fmt.Sprintf("%s-2", name), withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "oauth.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.enabled", "false"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.enabled", "false"),
			resource.TestMatchResourceAttr(resourceFullName, "oauth.0.values.0.client_secret", regexp.MustCompile(`^[a-z0-9]+$`)),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.enforce_signed_request_openid", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.sp_jwks_url", "https://www.ping-eng.com/testjwks"),
			resource.TestCheckResourceAttr(resourceFullName, "oauth.0.values.0.sp_jwks_openid", "1"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.redirect_uris.*", "https://auth.ping-eng.com/env-id/rp/callback/oidc"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.redirect_uris.*", "https://auth.ping-eng.com/env-id/rp/callback/1"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.redirect_uris.*", "https://auth.ping-eng.com/env-id/rp/callback/openid_connect"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.logout_uris.*", "https://auth.ping-eng.com/env-id/logout"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.logout_uris.*", "https://auth.ping-eng.com/env-id/logout1"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_scopes.*", "openid"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_scopes.*", "profile"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_grants", "authorizationCode"),
			resource.TestCheckTypeSetElemAttr(resourceFullName, "oauth.0.values.0.allowed_grants", "implicit"),
		),
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		CheckDestroy:      davinci.Application_CheckDestroy(),
		Steps: []resource.TestStep{
			// Create from scratch
			fullStep,
			{
				Config:  testAccResourceApplication_WithOAuth_Full_HCL(resourceName, fmt.Sprintf("%s-1", name), withBootstrapConfig),
				Destroy: true,
			},
			// Create from scratch
			minimalStep,
			{
				Config:  testAccResourceApplication_WithOAuth_Minimal_HCL(resourceName, fmt.Sprintf("%s-2", name), withBootstrapConfig),
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

func TestAccResourceApplication_WithFlowPolicy_Full(t *testing.T) {

	withBootstrapConfig := false

	resourceBase := "davinci_application"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	fullStep := resource.TestStep{
		Config: testAccResourceApplication_WithFlowPolicy_Full_HCL(resourceName, fmt.Sprintf("%s-1", name), withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", fmt.Sprintf("%s-1", name)),
			resource.TestCheckResourceAttr(resourceFullName, "connector_id", "pingOneMfaConnector"),
			resource.TestCheckResourceAttr(resourceFullName, "property.#", "5"),
			resource.TestMatchResourceAttr(resourceFullName, "created_date", verify.EpochDateRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "customer_id", verify.P1DVResourceIDRegexpFullString),
		),
	}

	minimalStep := resource.TestStep{
		Config: testAccResourceApplication_WithFlowPolicy_Minimal_HCL(resourceName, fmt.Sprintf("%s-2", name), withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", fmt.Sprintf("%s-2", name)),
			resource.TestCheckResourceAttr(resourceFullName, "connector_id", "annotationConnector"),
			resource.TestCheckResourceAttr(resourceFullName, "property.#", "0"),
			resource.TestMatchResourceAttr(resourceFullName, "created_date", verify.EpochDateRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "customer_id", verify.P1DVResourceIDRegexpFullString),
		),
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		CheckDestroy:      davinci.Application_CheckDestroy(),
		Steps: []resource.TestStep{
			// Create from scratch
			fullStep,
			{
				Config:  testAccResourceApplication_WithFlowPolicy_Full_HCL(resourceName, fmt.Sprintf("%s-1", name), withBootstrapConfig),
				Destroy: true,
			},
			// Create from scratch
			minimalStep,
			{
				Config:  testAccResourceApplication_WithFlowPolicy_Minimal_HCL(resourceName, fmt.Sprintf("%s-2", name), withBootstrapConfig),
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

func TestAccResourceApplication_P1SessionFlowPolicy(t *testing.T) {
	resourceAppBase := "davinci_application"
	resourceName := acctest.ResourceNameGen()
	resourceAppFullName := fmt.Sprintf("%s.%s", resourceAppBase, resourceName)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		CheckDestroy:      davinci.Application_CheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApplication_P1SessionFlowPolicy_Hcl(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceAppFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceAppFullName, "name"),
					resource.TestCheckResourceAttrSet(resourceAppFullName, "policy.0.policy_id"),
					resource.TestCheckNoResourceAttr(resourceAppFullName, "policy.1.policy_id"),
				),
			},
			{
				Config: testAccResourceApplication_P1SessionFlowPolicyUpdate_Hcl(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceAppFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceAppFullName, "name"),
					resource.TestCheckResourceAttrSet(resourceAppFullName, "policy.0.policy_id"),
					resource.TestCheckNoResourceAttr(resourceAppFullName, "policy.1.policy_id"),
				),
			},
			// Test importing the resource
			{
				ResourceName: resourceAppFullName,
				ImportStateIdFunc: func() resource.ImportStateIdFunc {
					return func(s *terraform.State) (string, error) {
						rs, ok := s.RootModule().Resources[resourceAppFullName]
						if !ok {
							return "", fmt.Errorf("Resource Not found: %s", resourceAppFullName)
						}

						return fmt.Sprintf("%s/%s", rs.Primary.Attributes["environment_id"], rs.Primary.ID), nil
					}
				}(),
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{
					// "context", // This shouldn't be ignored, can be solved on migration to the plugin framework
				},
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
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		CheckDestroy:      davinci.Application_CheckDestroy(),
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
	environment_id = pingone_environment.%[2]s.id
  name           = "%[3]s"
  api_key_enabled = false

  oauth {
    enabled = false
    values {
		enabled                       = false
      allowed_grants                = ["implicit", "authorizationCode"]
      allowed_scopes                = ["openid", "profile"]
      enforce_signed_request_openid = false
      redirect_uris                 = [
		"https://auth.ping-eng.com/env-id/rp/callback/openid_connect",
		"https://auth.ping-eng.com/env-id/rp/callback/oidc",
		"https://auth.ping-eng.com/env-id/rp/callback/1",
		]
      logout_uris                 = [
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

func testAccResourceApplication_WithFlowPolicy_Full_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
	environment_id = pingone_environment.%[2]s.id
  name           = "%[3]s"
  
  oauth {
    enabled = true
    values {
      allowed_grants                = ["authorizationCode"]
      allowed_scopes                = ["openid", "profile"]
      enabled                       = true
      enforce_signed_request_openid = false
      redirect_uris                 = [
		"https://auth.ping-eng.com/env-id/rp/callback/openid_connect",
		]
    }
  }

  policy {
    name = "simpleflow-1"
	status = "enabled"

    policy_flow {
      flow_id    = davinci_flow.%[3]s-1.id
      version_id = -1
      weight     = 25
    }

	policy_flow {
		flow_id    = davinci_flow.%[3]s-2.id
		version_id = -1
		weight     = 45
	  }

	policy_flow {
		flow_id    = davinci_flow.%[3]s-3.id
		version_id = -1
		weight     = 30
	  }
  }

  policy {
    name = "simpleflow-2"
    policy_flow {
      flow_id    = davinci_flow.%[3]s-1.id
      version_id = -1
      weight     = 100
    }
    status = "enabled"
  }

  policy {
    name = "simpleflow-3"
    policy_flow {
      flow_id    = davinci_flow.%[3]s-2.id
      version_id = -1
      weight     = 100
    }
    status = "disabled"
  }
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceApplication_WithFlowPolicy_Minimal_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
	environment_id = pingone_environment.%[2]s.id
  name           = "%[3]s"

  oauth {
    enabled = true
    values {
      allowed_grants                = ["authorizationCode"]
      allowed_scopes                = ["openid", "profile"]
      enabled                       = true
      enforce_signed_request_openid = false
      redirect_uris                 = [
		"https://auth.ping-eng.com/env-id/rp/callback/openid_connect",
		]
    }
  }

  policy {
    name = "simpleflow-1"
	status = "enabled"

    policy_flow {
      flow_id    = davinci_flow.%[3]s-1.id
      version_id = -1
      weight     = 100
    }
  }
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceApplication_P1SessionFlowPolicy_Hcl(resourceName string) (hcl string) {
	flows := acctest.FlowsForTests(resourceName)

	baseHcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{flows.PingOneSessionMainFlow.Hcl, flows.PingOneSessionSubFlow.Hcl})
	hcl = fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s-flow" {
  name           = "Flow"
  connector_id   = "flowConnector"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
}

resource "davinci_application" "%[2]s" {
  name           = "TF ACC Test-%[2]s"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  oauth {
    enabled = true
    values {
      allowed_grants                = ["authorizationCode"]
      allowed_scopes                = ["openid", "profile"]
      enabled                       = true
      enforce_signed_request_openid = false
      redirect_uris                 = ["https://auth.pingone.com/env-id/rp/callback/openid_connect"]
    }
  }
  saml {
    values {
      enabled                = false
      enforce_signed_request = false
    }
  }
  policy {
    name = "simpleflow"
    policy_flow {
      flow_id    = resource.davinci_flow.%[3]s.id
      version_id = -1
      weight     = 100
    }
    status = "enabled"
  }
}
`, baseHcl, resourceName, flows.PingOneSessionMainFlow.Name)
	return hcl
}

func testAccResourceApplication_P1SessionFlowPolicyUpdate_Hcl(resourceName string) (hcl string) {
	flows := acctest.FlowsForTests(resourceName)

	baseHcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{flows.PingOneSessionMainFlowUpdate.Hcl, flows.PingOneSessionSubFlow.Hcl})
	hcl = fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s-flow" {
  name           = "Flow"
  connector_id   = "flowConnector"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
}

resource "davinci_application" "%[2]s" {
  name           = "TF ACC Test-%[2]s"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  oauth {
    enabled = true
    values {
      allowed_grants                = ["authorizationCode"]
      allowed_scopes                = ["openid", "profile"]
      enabled                       = true
      enforce_signed_request_openid = false
      redirect_uris                 = ["https://auth.pingone.com/env-id/rp/callback/openid_connect"]
    }
  }
  saml {
    values {
      enabled                = false
      enforce_signed_request = false
    }
  }
  policy {
    name = "simpleflow"
    policy_flow {
      flow_id    = resource.davinci_flow.%[3]s.id
      version_id = -1
      weight     = 100
    }
    status = "enabled"
  }
}


`, baseHcl, resourceName, flows.PingOneSessionMainFlowUpdate.Name)
	return hcl
}
