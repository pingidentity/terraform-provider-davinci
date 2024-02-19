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

func TestAccResourceVariable_RemovalDrift(t *testing.T) {

	resourceBase := "davinci_variable"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	hcl := testAccResourceVariable_CompanyContext_Full_Hcl(resourceName, resourceName, true)

	var variableID, environmentID string

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
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		CheckDestroy:      davinci.Variable_CheckDestroy(),
		Steps: []resource.TestStep{
			// Configure
			{
				Config: hcl,
				Check:  davinci.Variable_GetIDs(resourceFullName, &environmentID, &variableID),
			},
			// Replan after removal preconfig
			{
				PreConfig: func() {
					davinci.Variable_RemovalDrift_PreConfig(t, environmentID, variableID)
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
			},
			// Test removal of the environment
			{
				Config:   hcl,
				Check:    davinci.Variable_GetIDs(resourceFullName, &environmentID, &variableID),
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

func TestAccResourceVariable_Full_CompanyContext_Clean(t *testing.T) {
	testAccResourceVariable_Full_CompanyContext(t, false)
}
func TestAccResourceVariable_Full_CompanyContext_WithBootstrap(t *testing.T) {
	testAccResourceVariable_Full_CompanyContext(t, true)
}

func testAccResourceVariable_Full_CompanyContext(t *testing.T, withBootstrapConfig bool) {

	resourceBase := "davinci_variable"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	fullStep := resource.TestStep{
		Config: testAccResourceVariable_CompanyContext_Full_Hcl(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", regexp.MustCompile(`^[a-zA-Z0-9]+##SK##company$`)),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", name),
			resource.TestCheckResourceAttr(resourceFullName, "context", "company"),
			resource.TestCheckResourceAttr(resourceFullName, "description", fmt.Sprintf(("desc-%s"), name)),
			resource.TestCheckResourceAttr(resourceFullName, "value", "7"),
			resource.TestCheckResourceAttr(resourceFullName, "type", "number"),
			resource.TestCheckResourceAttr(resourceFullName, "min", "5"),
			resource.TestCheckResourceAttr(resourceFullName, "max", "10"),
			resource.TestCheckResourceAttr(resourceFullName, "mutable", "false"),
		),
	}

	minimalStep := resource.TestStep{
		Config: testAccResourceVariable_CompanyContext_Minimal_Hcl(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", regexp.MustCompile(`^[a-zA-Z0-9]+##SK##company$`)),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", name),
			resource.TestCheckResourceAttr(resourceFullName, "context", "company"),
			resource.TestCheckResourceAttr(resourceFullName, "description", ""),
			resource.TestCheckResourceAttr(resourceFullName, "value", ""),
			resource.TestCheckResourceAttr(resourceFullName, "type", "number"),
			resource.TestCheckResourceAttr(resourceFullName, "min", "0"),
			resource.TestCheckResourceAttr(resourceFullName, "max", "2000"),
			resource.TestCheckResourceAttr(resourceFullName, "mutable", "true"),
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
		CheckDestroy:      davinci.Variable_CheckDestroy(),
		Steps: []resource.TestStep{
			// Create from scratch
			fullStep,
			{
				Config:  testAccResourceVariable_CompanyContext_Full_Hcl(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// Create from scratch
			minimalStep,
			{
				Config:  testAccResourceVariable_CompanyContext_Minimal_Hcl(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// Test updates
			fullStep,
			{ //minimalStep,
				Config: testAccResourceVariable_CompanyContext_Full_Hcl(resourceName, name, withBootstrapConfig),
				SkipFunc: func() (bool, error) {
					return true, nil // skip due to https://github.com/pingidentity/terraform-provider-davinci/issues/253
				},
			},
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
				ImportStateVerifyIgnore: []string{
					"context", // This shouldn't be ignored, can be solved on migration to the plugin framework
				},
			},
		},
	})
}

func TestAccResourceVariable_Full_FlowInstanceContext_Clean(t *testing.T) {
	testAccResourceVariable_Full_FlowInstanceContext(t, false)
}
func TestAccResourceVariable_Full_FlowInstanceContext_WithBootstrap(t *testing.T) {
	testAccResourceVariable_Full_FlowInstanceContext(t, true)
}

func testAccResourceVariable_Full_FlowInstanceContext(t *testing.T, withBootstrapConfig bool) {

	resourceBase := "davinci_variable"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	fullStep := resource.TestStep{
		Config: testAccResourceVariable_FlowInstanceContext_Full_Hcl(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", regexp.MustCompile(`^[a-zA-Z0-9]+##SK##flowInstance$`)),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", name),
			resource.TestCheckResourceAttr(resourceFullName, "context", "flowInstance"),
			resource.TestCheckResourceAttr(resourceFullName, "description", fmt.Sprintf(("desc-%s"), name)),
			resource.TestCheckResourceAttr(resourceFullName, "value", "7"),
			resource.TestCheckResourceAttr(resourceFullName, "type", "number"),
			resource.TestCheckResourceAttr(resourceFullName, "min", "5"),
			resource.TestCheckResourceAttr(resourceFullName, "max", "10"),
			resource.TestCheckResourceAttr(resourceFullName, "mutable", "false"),
		),
	}

	minimalStep := resource.TestStep{
		Config: testAccResourceVariable_FlowInstanceContext_Minimal_Hcl(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", regexp.MustCompile(`^[a-zA-Z0-9]+##SK##flowInstance$`)),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", name),
			resource.TestCheckResourceAttr(resourceFullName, "context", "flowInstance"),
			resource.TestCheckResourceAttr(resourceFullName, "description", ""),
			resource.TestCheckResourceAttr(resourceFullName, "value", ""),
			resource.TestCheckResourceAttr(resourceFullName, "type", "number"),
			resource.TestCheckResourceAttr(resourceFullName, "min", "0"),
			resource.TestCheckResourceAttr(resourceFullName, "max", "2000"),
			resource.TestCheckResourceAttr(resourceFullName, "mutable", "true"),
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
		CheckDestroy:      davinci.Variable_CheckDestroy(),
		Steps: []resource.TestStep{
			// Create from scratch
			fullStep,
			{
				Config:  testAccResourceVariable_FlowInstanceContext_Full_Hcl(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// Create from scratch
			minimalStep,
			{
				Config:  testAccResourceVariable_FlowInstanceContext_Minimal_Hcl(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// Test updates
			fullStep,
			{ //minimalStep,
				Config: testAccResourceVariable_FlowInstanceContext_Full_Hcl(resourceName, name, withBootstrapConfig),
				SkipFunc: func() (bool, error) {
					return true, nil // skip due to https://github.com/pingidentity/terraform-provider-davinci/issues/253
				},
			},
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
				ImportStateVerifyIgnore: []string{
					"context", // This shouldn't be ignored, can be solved on migration to the plugin framework
				},
			},
		},
	})
}

func TestAccResourceVariable_Full_UserContext_Clean(t *testing.T) {
	testAccResourceVariable_Full_UserContext(t, false)
}
func TestAccResourceVariable_Full_UserContext_WithBootstrap(t *testing.T) {
	testAccResourceVariable_Full_UserContext(t, true)
}

func testAccResourceVariable_Full_UserContext(t *testing.T, withBootstrapConfig bool) {

	resourceBase := "davinci_variable"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	fullStep := resource.TestStep{
		Config: testAccResourceVariable_UserContext_Full_Hcl(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", regexp.MustCompile(`^[a-zA-Z0-9]+##SK##user$`)),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", name),
			resource.TestCheckResourceAttr(resourceFullName, "context", "user"),
			resource.TestCheckResourceAttr(resourceFullName, "description", fmt.Sprintf(("desc-%s"), name)),
			resource.TestCheckResourceAttr(resourceFullName, "value", "7"),
			resource.TestCheckResourceAttr(resourceFullName, "type", "number"),
			resource.TestCheckResourceAttr(resourceFullName, "min", "5"),
			resource.TestCheckResourceAttr(resourceFullName, "max", "10"),
			resource.TestCheckResourceAttr(resourceFullName, "mutable", "false"),
		),
	}

	minimalStep := resource.TestStep{
		Config: testAccResourceVariable_UserContext_Minimal_Hcl(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", regexp.MustCompile(`^[a-zA-Z0-9]+##SK##user$`)),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", name),
			resource.TestCheckResourceAttr(resourceFullName, "context", "user"),
			resource.TestCheckResourceAttr(resourceFullName, "description", ""),
			resource.TestCheckResourceAttr(resourceFullName, "value", ""),
			resource.TestCheckResourceAttr(resourceFullName, "type", "number"),
			resource.TestCheckResourceAttr(resourceFullName, "min", "0"),
			resource.TestCheckResourceAttr(resourceFullName, "max", "2000"),
			resource.TestCheckResourceAttr(resourceFullName, "mutable", "true"),
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
		CheckDestroy:      davinci.Variable_CheckDestroy(),
		Steps: []resource.TestStep{
			// Create from scratch
			fullStep,
			{
				Config:  testAccResourceVariable_UserContext_Full_Hcl(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// Create from scratch
			minimalStep,
			{
				Config:  testAccResourceVariable_UserContext_Minimal_Hcl(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// Test updates
			fullStep,
			{ //minimalStep,
				Config: testAccResourceVariable_UserContext_Full_Hcl(resourceName, name, withBootstrapConfig),
				SkipFunc: func() (bool, error) {
					return true, nil // skip due to https://github.com/pingidentity/terraform-provider-davinci/issues/253
				},
			},
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
				ImportStateVerifyIgnore: []string{
					"context", // This shouldn't be ignored, can be solved on migration to the plugin framework
				},
			},
		},
	})
}

func TestAccResourceVariable_ChangeDataType(t *testing.T) {

	resourceBase := "davinci_variable"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	numberStep := resource.TestStep{
		Config: testAccResourceVariable_Minimal_Hcl(resourceName, name, "company", false),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "type", "number"),
		),
	}

	stringStep := resource.TestStep{
		Config: testAccResourceVariable_Minimal_String_Hcl(resourceName, name, "company", false),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "type", "string"),
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
		CheckDestroy:      davinci.Variable_CheckDestroy(),
		Steps: []resource.TestStep{
			numberStep,
			stringStep,
			numberStep,
		},
	})
}

func TestAccResourceVariable_BadParameters(t *testing.T) {

	resourceBase := "davinci_variable"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	hcl := testAccResourceVariable_UserContext_Minimal_Hcl(resourceName, resourceName, false)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		CheckDestroy:      davinci.Variable_CheckDestroy(),
		Steps: []resource.TestStep{
			// Configure
			{
				Config: hcl,
			},
			// Errors
			{
				ResourceName: resourceFullName,
				ImportState:  true,
				ExpectError:  regexp.MustCompile(`Invalid import ID specified \(".*"\).  The ID should be in the format "environment_id/davinci_variable_id" and must match regex: .*`),
			},
			{
				ResourceName:  resourceFullName,
				ImportStateId: "/",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`Invalid import ID specified \(".*"\).  The ID should be in the format "environment_id/davinci_variable_id" and must match regex: .*`),
			},
			{
				ResourceName:  resourceFullName,
				ImportStateId: "badformat/badformat",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`Invalid import ID specified \(".*"\).  The ID should be in the format "environment_id/davinci_variable_id" and must match regex: .*`),
			},
		},
	})
}

func testAccResourceVariable_Full_Hcl(resourceName, name, context string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_variable" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  name           = "%[3]s"
  context        = "%[4]s"
  description    = "desc-%[3]s"
  value          = "7"
  type           = "number"
  min            = "5"
  max            = "10"
  mutable        = false
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name, context)
}

func testAccResourceVariable_Minimal_Hcl(resourceName, name, context string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_variable" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  name           = "%[3]s"
  context        = "%[4]s"
  type           = "number"
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name, context)
}

func testAccResourceVariable_Minimal_String_Hcl(resourceName, name, context string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_variable" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  name           = "%[3]s"
  context        = "%[4]s"
  type           = "string"
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name, context)
}

func testAccResourceVariable_CompanyContext_Full_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return testAccResourceVariable_Full_Hcl(resourceName, name, "company", withBootstrapConfig)
}

func testAccResourceVariable_CompanyContext_Minimal_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return testAccResourceVariable_Minimal_Hcl(resourceName, name, "company", withBootstrapConfig)
}

func testAccResourceVariable_FlowInstanceContext_Full_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return testAccResourceVariable_Full_Hcl(resourceName, name, "flowInstance", withBootstrapConfig)
}

func testAccResourceVariable_FlowInstanceContext_Minimal_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return testAccResourceVariable_Minimal_Hcl(resourceName, name, "flowInstance", withBootstrapConfig)
}

func testAccResourceVariable_UserContext_Full_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return testAccResourceVariable_Full_Hcl(resourceName, name, "user", withBootstrapConfig)
}

func testAccResourceVariable_UserContext_Minimal_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return testAccResourceVariable_Minimal_Hcl(resourceName, name, "user", withBootstrapConfig)
}
