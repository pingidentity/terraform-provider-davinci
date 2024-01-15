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

func TestAccResourceApplicationFlowPolicy_RemovalDrift(t *testing.T) {

	resourceBase := "davinci_application_flow_policy"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	hcl := testAccResourceApplicationFlowPolicy_Full_HCL(resourceName, resourceName, true)

	var applicationFlowPolicyID, applicationID, environmentID string

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
		CheckDestroy:      davinci.ApplicationFlowPolicy_CheckDestroy(),
		Steps: []resource.TestStep{
			// Configure
			{
				Config: hcl,
				Check:  davinci.ApplicationFlowPolicy_GetIDs(resourceFullName, &environmentID, &applicationID, &applicationFlowPolicyID),
			},
			// Replan after removal preconfig
			{
				PreConfig: func() {
					davinci.ApplicationFlowPolicy_RemovalDrift_PreConfig(t, environmentID, applicationID, applicationFlowPolicyID)
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
			},
			// Configure
			{
				Config: hcl,
				Check:  davinci.ApplicationFlowPolicy_GetIDs(resourceFullName, &environmentID, &applicationID, &applicationFlowPolicyID),
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

func TestAccResourceApplicationFlowPolicy_Full(t *testing.T) {

	withBootstrapConfig := false

	resourceBase := "davinci_application_flow_policy"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	fullStep := resource.TestStep{
		Config: testAccResourceApplicationFlowPolicy_Full_HCL(resourceName, fmt.Sprintf("%s-1", name), withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "application_id", verify.P1DVResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", fmt.Sprintf("%s-1", name)),
			resource.TestCheckResourceAttr(resourceFullName, "status", "disabled"),
			resource.TestCheckResourceAttr(resourceFullName, "policy_flow.#", "3"),
			resource.TestMatchResourceAttr(resourceFullName, "created_date", verify.EpochDateRegexpFullString),
		),
	}

	minimalStep := resource.TestStep{
		Config: testAccResourceApplicationFlowPolicy_Minimal_HCL(resourceName, fmt.Sprintf("%s-2", name), withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "application_id", verify.P1DVResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", fmt.Sprintf("%s-1", name)),
			resource.TestCheckResourceAttr(resourceFullName, "status", "enabled"),
			resource.TestCheckResourceAttr(resourceFullName, "policy_flow.#", "0"),
			resource.TestMatchResourceAttr(resourceFullName, "created_date", verify.EpochDateRegexpFullString),
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
		CheckDestroy:      davinci.ApplicationFlowPolicy_CheckDestroy(),
		Steps: []resource.TestStep{
			// Create from scratch
			fullStep,
			{
				Config:  testAccResourceApplicationFlowPolicy_Full_HCL(resourceName, fmt.Sprintf("%s-1", name), withBootstrapConfig),
				Destroy: true,
			},
			// Create from scratch
			minimalStep,
			{
				Config:  testAccResourceApplicationFlowPolicy_Minimal_HCL(resourceName, fmt.Sprintf("%s-2", name), withBootstrapConfig),
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

						return fmt.Sprintf("%s/%s/%s", rs.Primary.Attributes["environment_id"], rs.Primary.Attributes["application_id"], rs.Primary.ID), nil
					}
				}(),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccResourceApplicationFlowPolicy_WithPolicyFlow_Full(t *testing.T) {

	withBootstrapConfig := false

	resourceBase := "davinci_application_flow_policy"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	fullStep := resource.TestStep{
		Config: testAccResourceApplicationFlowPolicy_WithPolicyFlow_Full_HCL(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "policy_flow.#", "3"),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "policy_flow.*", map[string]*regexp.Regexp{
				"flow_id":    verify.P1DVResourceIDRegexpFullString,
				"weight":     regexp.MustCompile(`^35$`),
				"version_id": regexp.MustCompile(`^-1$`),
			}),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "policy_flow.*", map[string]*regexp.Regexp{
				"flow_id":    verify.P1DVResourceIDRegexpFullString,
				"weight":     regexp.MustCompile(`^45$`),
				"version_id": regexp.MustCompile(`^-1$`),
			}),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "policy_flow.*", map[string]*regexp.Regexp{
				"flow_id":    verify.P1DVResourceIDRegexpFullString,
				"weight":     regexp.MustCompile(`^20$`),
				"version_id": regexp.MustCompile(`^-1$`),
			}),
		),
	}

	minimalStep1 := resource.TestStep{
		Config: testAccResourceApplicationFlowPolicy_WithPolicyFlow_Minimal1_HCL(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "policy_flow.#", "1"),
		),
	}

	minimalStep2 := resource.TestStep{
		Config: testAccResourceApplicationFlowPolicy_WithPolicyFlow_Minimal2_HCL(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "policy_flow.#", "1"),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "policy_flow.*", map[string]*regexp.Regexp{
				"flow_id":    verify.P1DVResourceIDRegexpFullString,
				"weight":     regexp.MustCompile(`^100$`),
				"version_id": regexp.MustCompile(`^-1$`),
			}),
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
		CheckDestroy:      davinci.ApplicationFlowPolicy_CheckDestroy(),
		Steps: []resource.TestStep{
			// Create from scratch
			fullStep,
			{
				Config:  testAccResourceApplicationFlowPolicy_WithPolicyFlow_Full_HCL(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// Create from scratch
			minimalStep1,
			{
				Config:  testAccResourceApplicationFlowPolicy_WithPolicyFlow_Minimal1_HCL(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// Create from scratch
			minimalStep2,
			{
				Config:  testAccResourceApplicationFlowPolicy_WithPolicyFlow_Minimal2_HCL(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// Test updates
			fullStep,
			minimalStep1,
			fullStep,
			minimalStep2,
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

						return fmt.Sprintf("%s/%s/%s", rs.Primary.Attributes["environment_id"], rs.Primary.Attributes["application_id"], rs.Primary.ID), nil
					}
				}(),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccResourceApplicationFlowPolicy_Full_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
  environment_id  = pingone_environment.%[2]s.id
  name            = "%[3]s"

  oauth {
    values {
      allowed_grants                = ["authorizationCode"]
      allowed_scopes                = ["openid", "profile"]
      redirect_uris = [
        "https://auth.ping-eng.com/env-id/rp/callback/openid_connect",
      ]
    }
  }
}

resource "davinci_application_flow_policy" "%[2]s" {
	environment_id = pingone_environment.%[2]s.id
	application_id = davinci_application.%[2]s.id

	name           = "%[3]s"
	status         = "disabled"

	policy_flow {
	  flow_id    = davinci_flow.%[2]s-1.id
	  version_id = -1
	  weight     = 35
	}

	policy_flow {
		flow_id    = davinci_flow.%[2]s-2.id
		version_id = -1
		weight     = 45
	  }

	  policy_flow {
		flow_id    = davinci_flow.%[2]s-3.id
		version_id = -1
		weight     = 20
	  }
  }
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceApplicationFlowPolicy_Minimal_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
	environment_id  = pingone_environment.%[2]s.id
	name            = "%[3]s"
  
	oauth {
	  values {
		allowed_grants                = ["authorizationCode"]
		allowed_scopes                = ["openid", "profile"]
		redirect_uris = [
		  "https://auth.ping-eng.com/env-id/rp/callback/openid_connect",
		]
	  }
	}
  }

  resource "davinci_application_flow_policy" "%[2]s" {
	environment_id = pingone_environment.%[2]s.id
	application_id = davinci_application.%[2]s.id

	name           = "%[3]s"
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceApplicationFlowPolicy_WithPolicyFlow_Full_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
  environment_id  = pingone_environment.%[2]s.id
  name            = "%[3]s"

  oauth {
    values {
      allowed_grants                = ["authorizationCode"]
      allowed_scopes                = ["openid", "profile"]
      redirect_uris = [
        "https://auth.ping-eng.com/env-id/rp/callback/openid_connect",
      ]
    }
  }
}

resource "davinci_application_flow_policy" "%[2]s" {
	environment_id = pingone_environment.%[2]s.id
	application_id = davinci_application.%[2]s.id

	name           = "%[3]s"
	status         = "disabled"

	policy_flow {
	  flow_id    = davinci_flow.%[2]s-1.id
	  version_id = -1
	  weight     = 35
	}

	policy_flow {
		flow_id    = davinci_flow.%[2]s-2.id
		version_id = -1
		weight     = 45
	  }

	  policy_flow {
		flow_id    = davinci_flow.%[2]s-3.id
		version_id = -1
		weight     = 20
	  }
  }
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceApplicationFlowPolicy_WithPolicyFlow_Minimal1_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
	environment_id  = pingone_environment.%[2]s.id
	name            = "%[3]s"
  
	oauth {
	  values {
		allowed_grants                = ["authorizationCode"]
		allowed_scopes                = ["openid", "profile"]
		redirect_uris = [
		  "https://auth.ping-eng.com/env-id/rp/callback/openid_connect",
		]
	  }
	}
  }

  resource "davinci_application_flow_policy" "%[2]s" {
	environment_id = pingone_environment.%[2]s.id
	application_id = davinci_application.%[2]s.id

	name           = "%[3]s"

	policy_flow {}
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceApplicationFlowPolicy_WithPolicyFlow_Minimal2_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
	environment_id  = pingone_environment.%[2]s.id
	name            = "%[3]s"
  
	oauth {
	  values {
		allowed_grants                = ["authorizationCode"]
		allowed_scopes                = ["openid", "profile"]
		redirect_uris = [
		  "https://auth.ping-eng.com/env-id/rp/callback/openid_connect",
		]
	  }
	}
  }

  resource "davinci_application_flow_policy" "%[2]s" {
	environment_id = pingone_environment.%[2]s.id
	application_id = davinci_application.%[2]s.id

	name           = "%[3]s"

	policy_flow {
		flow_id    = davinci_flow.%[2]s-1.id
		version_id = -1
		weight     = 100
	  }
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}
