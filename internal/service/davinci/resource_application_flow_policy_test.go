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

func TestAccResourceApplicationFlowPolicy_RemovalDrift(t *testing.T) {

	resourceBase := "davinci_application_flow_policy"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	hcl, err := testAccResourceApplicationFlowPolicy_Full_HCL(resourceName, resourceName, false)
	if err != nil {
		t.Fatalf("Failed to generate HCL: %v", err)
	}

	var applicationFlowPolicyID, applicationID, environmentID string

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
		CheckDestroy:             davinci.ApplicationFlowPolicy_CheckDestroy(),
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
				// Check:    davinci.Application_GetIDs(resourceFullName, &environmentID, &applicationID),
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

func TestAccResourceApplicationFlowPolicy_Full(t *testing.T) {

	withBootstrapConfig := false

	resourceBase := "davinci_application_flow_policy"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	fullStepHcl, err := testAccResourceApplicationFlowPolicy_Full_HCL(resourceName, fmt.Sprintf("%s-1", name), withBootstrapConfig)
	if err != nil {
		t.Fatalf("Failed to generate HCL: %v", err)
	}

	fullStep := resource.TestStep{
		Config: fullStepHcl,
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "application_id", verify.P1DVResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", fmt.Sprintf("%s-1", name)),
			resource.TestCheckResourceAttr(resourceFullName, "status", "disabled"),
			resource.TestCheckResourceAttr(resourceFullName, "policy_flow.#", "3"),
			// https://github.com/pingidentity/terraform-provider-davinci/issues/257
			//resource.TestMatchResourceAttr(resourceFullName, "created_date", verify.EpochDateRegexpFullString),
		),
	}

	minimalStepHcl, err := testAccResourceApplicationFlowPolicy_Minimal_HCL(resourceName, fmt.Sprintf("%s-2", name), withBootstrapConfig)
	if err != nil {
		t.Fatalf("Failed to generate HCL: %v", err)
	}

	minimalStep := resource.TestStep{
		Config: minimalStepHcl,
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "application_id", verify.P1DVResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", fmt.Sprintf("%s-2", name)),
			resource.TestCheckResourceAttr(resourceFullName, "status", "enabled"),
			resource.TestCheckResourceAttr(resourceFullName, "policy_flow.#", "1"),
			// https://github.com/pingidentity/terraform-provider-davinci/issues/257
			//resource.TestMatchResourceAttr(resourceFullName, "created_date", verify.EpochDateRegexpFullString),
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
		CheckDestroy:             davinci.ApplicationFlowPolicy_CheckDestroy(),
		Steps: []resource.TestStep{
			// Create from scratch
			fullStep,
			{
				Config:  fullStepHcl,
				Destroy: true,
			},
			// Create from scratch
			minimalStep,
			{
				Config:  minimalStepHcl,
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
				ImportStateVerifyIgnore: []string{
					"created_date", // https://github.com/pingidentity/terraform-provider-davinci/issues/257
				},
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

	fullStepHcl, err := testAccResourceApplicationFlowPolicy_WithPolicyFlow_Full_HCL(resourceName, name, withBootstrapConfig)
	if err != nil {
		t.Fatalf("Failed to generate HCL: %v", err)
	}

	fullStep := resource.TestStep{
		Config: fullStepHcl,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "policy_flow.#", "3"),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "policy_flow.*", map[string]*regexp.Regexp{
				"flow_id":           verify.P1DVResourceIDRegexpFullString,
				"weight":            regexp.MustCompile(`^35$`),
				"version_id":        regexp.MustCompile(`^-1$`),
				"success_nodes.0":   regexp.MustCompile(`^node-1$`),
				"success_nodes.1":   regexp.MustCompile(`^node-2$`),
				"allowed_ip_list.0": regexp.MustCompile(`^10.1.2.3/23$`),
				"allowed_ip_list.1": regexp.MustCompile(`^10.1.2.4/23$`),
			}),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "policy_flow.*", map[string]*regexp.Regexp{
				"flow_id":           verify.P1DVResourceIDRegexpFullString,
				"weight":            regexp.MustCompile(`^45$`),
				"version_id":        regexp.MustCompile(`^-1$`),
				"success_nodes.0":   regexp.MustCompile(`^node-1$`),
				"success_nodes.1":   regexp.MustCompile(`^node-2$`),
				"allowed_ip_list.0": regexp.MustCompile(`^10.1.2.5/23$`),
				"allowed_ip_list.1": regexp.MustCompile(`^10.1.2.6/23$`),
			}),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "policy_flow.*", map[string]*regexp.Regexp{
				"flow_id":           verify.P1DVResourceIDRegexpFullString,
				"weight":            regexp.MustCompile(`^20$`),
				"version_id":        regexp.MustCompile(`^-1$`),
				"success_nodes.0":   regexp.MustCompile(`^node-1$`),
				"success_nodes.1":   regexp.MustCompile(`^node-2$`),
				"allowed_ip_list.0": regexp.MustCompile(`^10.1.2.3/23$`),
			}),
		),
	}

	minimalStep1Hcl, err := testAccResourceApplicationFlowPolicy_WithPolicyFlow_Minimal1_HCL(resourceName, name, withBootstrapConfig)
	if err != nil {
		t.Fatalf("Failed to generate HCL: %v", err)
	}

	minimalStep1 := resource.TestStep{
		Config: minimalStep1Hcl,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "policy_flow.#", "1"),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "policy_flow.*", map[string]*regexp.Regexp{
				"flow_id":         verify.P1DVResourceIDRegexpFullString,
				"version_id":      regexp.MustCompile(`^-1$`),
				"success_nodes.0": regexp.MustCompile(`^node-3$`),
			}),
		),
	}

	minimalStep2Hcl, err := testAccResourceApplicationFlowPolicy_WithPolicyFlow_Minimal2_HCL(resourceName, name, withBootstrapConfig)
	if err != nil {
		t.Fatalf("Failed to generate HCL: %v", err)
	}

	minimalStep2 := resource.TestStep{
		Config: minimalStep2Hcl,
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
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.ApplicationFlowPolicy_CheckDestroy(),
		Steps: []resource.TestStep{
			// Create from scratch
			fullStep,
			{
				Config:  fullStepHcl,
				Destroy: true,
			},
			// Create from scratch
			minimalStep1,
			{
				Config:  minimalStep1Hcl,
				Destroy: true,
			},
			// Create from scratch
			minimalStep2,
			{
				Config:  minimalStep2Hcl,
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
				ImportStateVerifyIgnore: []string{
					"created_date", // https://github.com/pingidentity/terraform-provider-davinci/issues/257
				},
			},
		},
	})
}

func testAccResourceApplicationFlowPolicy_Full_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string, err error) {
	flowResources, err := flowResources(resourceName, name, 3)
	if err != nil {
		return "", err
	}

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

resource "davinci_application_flow_policy" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  application_id = davinci_application.%[2]s.id

  name   = "%[3]s"
  status = "disabled"

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

%[4]s
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name, flowResources), nil
}

func testAccResourceApplicationFlowPolicy_Minimal_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string, err error) {
	flowResources, err := flowResources(resourceName, name, 3)
	if err != nil {
		return "", err
	}

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

resource "davinci_application_flow_policy" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  application_id = davinci_application.%[2]s.id

  name = "%[3]s"

  policy_flow {
    flow_id    = davinci_flow.%[2]s-1.id
    version_id = -1
    weight     = 100
  }
}

%[4]s
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name, flowResources), nil
}

func testAccResourceApplicationFlowPolicy_WithPolicyFlow_Full_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string, err error) {
	flowResources, err := flowResources(resourceName, name, 3)
	if err != nil {
		return "", err
	}

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

resource "davinci_application_flow_policy" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  application_id = davinci_application.%[2]s.id

  name   = "%[3]s"
  status = "disabled"

  policy_flow {
    flow_id         = davinci_flow.%[2]s-1.id
    version_id      = -1
    weight          = 35
    success_nodes   = ["node-1", "node-2"]
    allowed_ip_list = ["10.1.2.3/23", "10.1.2.4/23"]
  }

  policy_flow {
    flow_id         = davinci_flow.%[2]s-2.id
    version_id      = -1
    weight          = 45
    success_nodes   = ["node-1", "node-2"]
    allowed_ip_list = ["10.1.2.6/23", "10.1.2.5/23"]
  }

  policy_flow {
    flow_id         = davinci_flow.%[2]s-3.id
    version_id      = -1
    weight          = 20
    success_nodes   = ["node-1", "node-2"]
    allowed_ip_list = ["10.1.2.3/23"]
  }
}

%[4]s
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name, flowResources), nil
}

func testAccResourceApplicationFlowPolicy_WithPolicyFlow_Minimal1_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string, err error) {
	flowResources, err := flowResources(resourceName, name, 3)
	if err != nil {
		return "", err
	}

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

resource "davinci_application_flow_policy" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  application_id = davinci_application.%[2]s.id

  name = "%[3]s"

  policy_flow {
    flow_id       = davinci_flow.%[2]s-1.id
    version_id    = -1
    success_nodes = ["node-3"]
  }
}

%[4]s
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name, flowResources), nil
}

func testAccResourceApplicationFlowPolicy_WithPolicyFlow_Minimal2_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string, err error) {
	flowResources, err := flowResources(resourceName, name, 3)
	if err != nil {
		return "", err
	}

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

resource "davinci_application_flow_policy" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  application_id = davinci_application.%[2]s.id

  name = "%[3]s"

  policy_flow {
    flow_id    = davinci_flow.%[2]s-1.id
    version_id = -1
    weight     = 100
  }
}

%[4]s
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name, flowResources), nil
}

func TestAccResourceApplicationFlowPolicy_WithTrigger_Full(t *testing.T) {

	withBootstrapConfig := false

	resourceBase := "davinci_application_flow_policy"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	fullStepHcl, err := testAccResourceApplicationFlowPolicy_WithTrigger_Full_HCL(resourceName, name, withBootstrapConfig)
	if err != nil {
		t.Fatalf("Failed to generate HCL: %v", err)
	}

	fullStep := resource.TestStep{
		Config: fullStepHcl,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "trigger.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "trigger.0.type", "AUTHENTICATION"),
			resource.TestCheckResourceAttr(resourceFullName, "trigger.0.configuration.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "trigger.0.configuration.0.mfa.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "trigger.0.configuration.0.mfa.0.enabled", "true"),
			resource.TestCheckResourceAttr(resourceFullName, "trigger.0.configuration.0.mfa.0.time", "30"),
			resource.TestCheckResourceAttr(resourceFullName, "trigger.0.configuration.0.mfa.0.time_format", "min"),
			resource.TestCheckResourceAttr(resourceFullName, "trigger.0.configuration.0.pwd.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "trigger.0.configuration.0.pwd.0.enabled", "false"),
			resource.TestCheckResourceAttr(resourceFullName, "trigger.0.configuration.0.pwd.0.time", "60"),
			resource.TestCheckResourceAttr(resourceFullName, "trigger.0.configuration.0.pwd.0.time_format", "min"),
		),
	}

	minimalStepHcl, err := testAccResourceApplicationFlowPolicy_WithTrigger_Minimal_HCL(resourceName, name, withBootstrapConfig)
	if err != nil {
		t.Fatalf("Failed to generate HCL: %v", err)
	}

	minimalStep := resource.TestStep{
		Config: minimalStepHcl,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "trigger.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "trigger.0.type", "AUTHENTICATION"),
			resource.TestCheckResourceAttr(resourceFullName, "trigger.0.configuration.#", "0"),
		),
	}

	noTriggerStepHcl, err := testAccResourceApplicationFlowPolicy_WithP1Flow_NoTrigger_HCL(resourceName, name, withBootstrapConfig)
	if err != nil {
		t.Fatalf("Failed to generate HCL: %v", err)
	}

	noTriggerStep := resource.TestStep{
		Config: noTriggerStepHcl,
		Check: resource.ComposeTestCheckFunc(
			// Trigger is retained by the DaVinci API for PingOne flow policies even when omitted
			// from HCL. With trigger as Optional+Computed, no plan diff occurs — this step
			// validates no breaking state change when upgrading from HCL without a trigger block.
			resource.TestCheckResourceAttr(resourceFullName, "trigger.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "trigger.0.type", "AUTHENTICATION"),
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
		CheckDestroy:             davinci.ApplicationFlowPolicy_CheckDestroy(),
		Steps: []resource.TestStep{
			// Create from scratch with full trigger
			fullStep,
			{
				Config:  fullStepHcl,
				Destroy: true,
			},
			// Create from scratch with minimal trigger
			minimalStep,
			{
				Config:  minimalStepHcl,
				Destroy: true,
			},
			// Test updates
			fullStep,
			minimalStep,
			fullStep,
			// Apply config without trigger block - validates no breaking state change on upgrade
			noTriggerStep,
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
				ImportStateVerifyIgnore: []string{
					"created_date", // https://github.com/pingidentity/terraform-provider-davinci/issues/257
				},
			},
		},
	})
}

func testAccResourceApplicationFlowPolicy_WithTrigger_Full_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string, err error) {
	p1FlowResources, err := p1FlowResources(resourceName, name)
	if err != nil {
		return "", err
	}

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

resource "davinci_application_flow_policy" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  application_id = davinci_application.%[2]s.id

  name   = "%[3]s"
  status = "enabled"

  policy_flow {
    flow_id    = davinci_flow.%[2]s-p1.id
    version_id = -1
    weight     = 100
  }

  trigger {
    configuration {
      mfa {
        enabled     = true
        time        = 30
        time_format = "min"
      }
      pwd {
        enabled     = false
        time        = 60
        time_format = "min"
      }
    }
  }
}

%[4]s
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name, p1FlowResources), nil
}

func testAccResourceApplicationFlowPolicy_WithTrigger_Minimal_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string, err error) {
	p1FlowResources, err := p1FlowResources(resourceName, name)
	if err != nil {
		return "", err
	}

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

resource "davinci_application_flow_policy" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  application_id = davinci_application.%[2]s.id

  name   = "%[3]s"
  status = "enabled"

  policy_flow {
    flow_id    = davinci_flow.%[2]s-p1.id
    version_id = -1
    weight     = 100
  }

  trigger {
  }
}

%[4]s
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name, p1FlowResources), nil
}

func testAccResourceApplicationFlowPolicy_WithP1Flow_NoTrigger_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string, err error) {
	p1FlowResources, err := p1FlowResources(resourceName, name)
	if err != nil {
		return "", err
	}

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

resource "davinci_application_flow_policy" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  application_id = davinci_application.%[2]s.id

  name   = "%[3]s"
  status = "enabled"

  policy_flow {
    flow_id    = davinci_flow.%[2]s-p1.id
    version_id = -1
    weight     = 100
  }
}

%[4]s
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name, p1FlowResources), nil
}

func p1FlowResources(resourceName, name string) (hcl string, err error) {
	mainFlowJson, err := acctest.ReadFlowJsonFile("flows/p1sessionmainflow.json")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`
resource "davinci_connection" "%[1]s-p1auth" {
  environment_id = pingone_environment.%[1]s.id
  connector_id   = "pingOneAuthenticationConnector"
  name           = "%[2]s-p1auth"
}

resource "davinci_connection" "%[1]s-flow" {
  environment_id = pingone_environment.%[1]s.id
  connector_id   = "flowConnector"
  name           = "%[2]s-flow"
}

resource "davinci_connection" "%[1]s-node" {
  environment_id = pingone_environment.%[1]s.id
  connector_id   = "nodeConnector"
  name           = "%[2]s-node"
}

resource "davinci_connection" "%[1]s-annotation" {
  environment_id = pingone_environment.%[1]s.id
  connector_id   = "annotationConnector"
  name           = "%[2]s-annotation"
}

resource "davinci_flow" "%[1]s-p1" {
  environment_id = pingone_environment.%[1]s.id

  flow_json = <<EOT
%[3]s
EOT

  connection_link {
    id                           = davinci_connection.%[1]s-p1auth.id
    name                         = davinci_connection.%[1]s-p1auth.name
    replace_import_connection_id = "054d0c6ac38a15f82497469675634cab"
  }

  connection_link {
    id                           = davinci_connection.%[1]s-flow.id
    name                         = davinci_connection.%[1]s-flow.name
    replace_import_connection_id = "4f45de91338525b684c30eb2faccd568"
  }

  connection_link {
    id                           = davinci_connection.%[1]s-node.id
    name                         = davinci_connection.%[1]s-node.name
    replace_import_connection_id = "3566e86a35c26e575396dcfb89a3dcc0"
  }

  connection_link {
    id                           = davinci_connection.%[1]s-annotation.id
    name                         = davinci_connection.%[1]s-annotation.name
    replace_import_connection_id = "921bfae85c38ed45045e07be703d86b8"
  }
}
`, resourceName, name, mainFlowJson), nil
}

func flowResources(resourceName, name string, count int) (hcl string, err error) {

	mainFlowJson, err := acctest.ReadFlowJsonFile("flows/simple.json")
	if err != nil {
		return "", err
	}

	hcl += fmt.Sprintf(`
resource "davinci_connection" "%[1]s" {
  environment_id = pingone_environment.%[1]s.id
  connector_id   = "httpConnector"
  name           = "%[2]s"
}
	  `, resourceName, name)

	for i := 1; i <= count; i++ {
		hcl += fmt.Sprintf(`
resource "davinci_flow" "%[1]s-%[2]d" {
  environment_id = pingone_environment.%[1]s.id

  flow_json = <<EOT
%[3]s
EOT

  connection_link {
    id                           = davinci_connection.%[1]s.id
    name                         = davinci_connection.%[1]s.name
    replace_import_connection_id = "867ed4363b2bc21c860085ad2baa817d"
  }
}
`, resourceName, i, mainFlowJson)
	}

	return hcl, nil
}
