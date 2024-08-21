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
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
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
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Variable_CheckDestroy(),
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

	var variableID, environmentID string

	fullStep := resource.TestStep{
		Config: testAccResourceVariable_CompanyContext_Full_Hcl(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", regexp.MustCompile(`^[a-zA-Z0-9]+##SK##company$`)),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", name),
			resource.TestCheckResourceAttr(resourceFullName, "context", "company"),
			resource.TestCheckResourceAttr(resourceFullName, "description", fmt.Sprintf(("desc-%s"), name)),
			resource.TestCheckResourceAttr(resourceFullName, "value", "7"),
			resource.TestCheckNoResourceAttr(resourceFullName, "empty_value"),
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
			resource.TestCheckNoResourceAttr(resourceFullName, "description"),
			resource.TestCheckNoResourceAttr(resourceFullName, "value"),
			resource.TestCheckNoResourceAttr(resourceFullName, "empty_value"),
			resource.TestCheckResourceAttr(resourceFullName, "type", "number"),
			resource.TestCheckResourceAttr(resourceFullName, "min", "0"),
			resource.TestCheckResourceAttr(resourceFullName, "max", "2000"),
			resource.TestCheckResourceAttr(resourceFullName, "mutable", "true"),
		),
	}

	secretDynamicStep := resource.TestStep{
		Config: testAccResourceVariable_CompanyContext_SecretDynamic_Hcl(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			davinci.Variable_GetIDs(resourceFullName, &environmentID, &variableID),
			resource.TestMatchResourceAttr(resourceFullName, "id", regexp.MustCompile(`^[a-zA-Z0-9]+##SK##company$`)),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", name),
			resource.TestCheckResourceAttr(resourceFullName, "context", "company"),
			resource.TestCheckNoResourceAttr(resourceFullName, "description"),
			resource.TestCheckNoResourceAttr(resourceFullName, "value"),
			resource.TestCheckNoResourceAttr(resourceFullName, "empty_value"),
			resource.TestCheckResourceAttr(resourceFullName, "type", "secret"),
			resource.TestCheckResourceAttr(resourceFullName, "min", "0"),
			resource.TestCheckResourceAttr(resourceFullName, "max", "2000"),
			resource.TestCheckResourceAttr(resourceFullName, "mutable", "true"),
		),
	}

	secretStaticStep := resource.TestStep{
		Config: testAccResourceVariable_CompanyContext_SecretStatic_Hcl(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			davinci.Variable_GetIDs(resourceFullName, &environmentID, &variableID),
			resource.TestMatchResourceAttr(resourceFullName, "id", regexp.MustCompile(`^[a-zA-Z0-9]+##SK##company$`)),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", name),
			resource.TestCheckResourceAttr(resourceFullName, "context", "company"),
			resource.TestCheckNoResourceAttr(resourceFullName, "description"),
			resource.TestCheckResourceAttr(resourceFullName, "value", "mysecret"),
			resource.TestCheckNoResourceAttr(resourceFullName, "empty_value"),
			resource.TestCheckResourceAttr(resourceFullName, "type", "secret"),
			resource.TestCheckResourceAttr(resourceFullName, "min", "0"),
			resource.TestCheckResourceAttr(resourceFullName, "max", "2000"),
			resource.TestCheckResourceAttr(resourceFullName, "mutable", "true"),
		),
	}

	variableName := name
	mutable := true

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Variable_CheckDestroy(),
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
			{
				Config:  testAccResourceVariable_CompanyContext_Full_Hcl(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// Test secret
			secretDynamicStep,
			{
				Config: testAccResourceVariable_CompanyContext_SecretDynamic_Hcl(resourceName, name, withBootstrapConfig),
				PreConfig: func() {
					davinci.Variable_RandomVariableValue_PreConfig(t, environmentID, nil, &dv.VariablePayload{
						Name:    &variableName,
						Context: "company",
						Type:    "secret",
						Mutable: &mutable,
					})
				},
				ExpectNonEmptyPlan: false,
				PlanOnly:           true,
			},
			secretStaticStep,
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
					"value",   // This shouldn't be ignored, can be solved on migration to the plugin framework
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
			resource.TestCheckNoResourceAttr(resourceFullName, "empty_value"),
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
			resource.TestCheckNoResourceAttr(resourceFullName, "description"),
			resource.TestCheckNoResourceAttr(resourceFullName, "value"),
			resource.TestCheckNoResourceAttr(resourceFullName, "empty_value"),
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
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Variable_CheckDestroy(),
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
					"value",   // This shouldn't be ignored, can be solved on migration to the plugin framework
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
			resource.TestCheckNoResourceAttr(resourceFullName, "empty_value"),
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
			resource.TestCheckNoResourceAttr(resourceFullName, "description"),
			resource.TestCheckNoResourceAttr(resourceFullName, "value"),
			resource.TestCheckNoResourceAttr(resourceFullName, "empty_value"),
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
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Variable_CheckDestroy(),
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
					"value",   // This shouldn't be ignored, can be solved on migration to the plugin framework
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
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Variable_CheckDestroy(),
		Steps: []resource.TestStep{
			numberStep,
			stringStep,
			numberStep,
		},
	})
}

// func testAccResourceVariable_Values_UserContext(t *testing.T) {
// 	testAccResourceVariable_Values(t, "user")
// }

func TestAccResourceVariable_Values_CompanyContext(t *testing.T) {
	testAccResourceVariable_Values(t, "company", "testVariable")
}

func TestAccResourceVariable_Values_FlowInstanceContext(t *testing.T) {
	testAccResourceVariable_Values(t, "flowInstance", "flowInstanceVariable1")
}

func TestAccResourceVariable_Values_FlowContext(t *testing.T) {
	testAccResourceVariable_Values(t, "flow", "fdgdfgfdg")
}

func testAccResourceVariable_Values(t *testing.T, context, variableName string) {

	resourceBase := "davinci_variable"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s-%s", resourceBase, resourceName, context)
	flowResourceFullName := fmt.Sprintf("davinci_flow.%s", resourceName)

	name := resourceName

	withBootstrapConfig := false

	var variableID, flowID, environmentID string

	dynamicValueStep := resource.TestStep{
		Config: testAccResourceVariable_DynamicValue_Hcl(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			davinci.Variable_GetIDs(resourceFullName, &environmentID, &variableID),
			davinci.Flow_GetIDs(flowResourceFullName, &environmentID, &flowID),
			resource.TestCheckNoResourceAttr(resourceFullName, "value"),
			resource.TestCheckNoResourceAttr(resourceFullName, "empty_value"),
			resource.TestCheckNoResourceAttr(resourceFullName, "value_service"),
		),
	}

	emptyValueStep := resource.TestStep{
		Config: testAccResourceVariable_EmptyValue_Hcl(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			davinci.Variable_GetIDs(resourceFullName, &environmentID, &variableID),
			davinci.Flow_GetIDs(flowResourceFullName, &environmentID, &flowID),
			resource.TestCheckNoResourceAttr(resourceFullName, "value"),
			resource.TestCheckResourceAttr(resourceFullName, "empty_value", "true"),
			resource.TestCheckNoResourceAttr(resourceFullName, "value_service"),
		),
	}

	staticValueStep1 := resource.TestStep{
		Config: testAccResourceVariable_StaticValue_Hcl(resourceName, name, withBootstrapConfig, "myvar"),
		Check: resource.ComposeTestCheckFunc(
			davinci.Variable_GetIDs(resourceFullName, &environmentID, &variableID),
			davinci.Flow_GetIDs(flowResourceFullName, &environmentID, &flowID),
			resource.TestCheckResourceAttr(resourceFullName, "value", "myvar"),
			resource.TestCheckNoResourceAttr(resourceFullName, "empty_value"),
			resource.TestCheckResourceAttr(resourceFullName, "value_service", "myvar"),
		),
	}

	staticValueStep2 := resource.TestStep{
		Config: testAccResourceVariable_StaticValue_Hcl(resourceName, name, withBootstrapConfig, "myvar2"),
		Check: resource.ComposeTestCheckFunc(
			davinci.Variable_GetIDs(resourceFullName, &environmentID, &variableID),
			davinci.Flow_GetIDs(flowResourceFullName, &environmentID, &flowID),
			resource.TestCheckResourceAttr(resourceFullName, "value", "myvar2"),
			resource.TestCheckNoResourceAttr(resourceFullName, "empty_value"),
			resource.TestCheckResourceAttr(resourceFullName, "value_service", "myvar2"),
		),
	}

	mutable := true

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Variable_CheckDestroy(),
		Steps: []resource.TestStep{
			dynamicValueStep,
			{
				Config: testAccResourceVariable_DynamicValue_Hcl(resourceName, name, withBootstrapConfig),
				PreConfig: func() {
					davinci.Variable_RandomVariableValue_PreConfig(t, environmentID, &flowID, &dv.VariablePayload{
						Name:    &variableName,
						Context: context,
						Type:    "string",
						Mutable: &mutable,
					})
				},
				ExpectNonEmptyPlan: false,
				PlanOnly:           true,
			},
			emptyValueStep,
			{
				Config: testAccResourceVariable_EmptyValue_Hcl(resourceName, name, withBootstrapConfig),
				PreConfig: func() {
					davinci.Variable_RandomVariableValue_PreConfig(t, environmentID, &flowID, &dv.VariablePayload{
						Name:    &variableName,
						Context: context,
						Type:    "string",
						Mutable: &mutable,
					})
				},
				ExpectNonEmptyPlan: true,
				PlanOnly:           true,
			},
			staticValueStep1,
			{
				Config: testAccResourceVariable_StaticValue_Hcl(resourceName, name, withBootstrapConfig, "myvar"),
				PreConfig: func() {
					davinci.Variable_RandomVariableValue_PreConfig(t, environmentID, &flowID, &dv.VariablePayload{
						Name:    &variableName,
						Context: context,
						Type:    "string",
						Mutable: &mutable,
					})
				},
				ExpectNonEmptyPlan: true,
				PlanOnly:           true,
			},
			staticValueStep2,
		},
	})
}

func TestAccResourceVariable_UnknownValue(t *testing.T) {

	resourceBase := "davinci_variable"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	withBootstrapConfig := false

	step1 := resource.TestStep{
		Config: testAccResourceVariable_UnknownValue1_Hcl(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "value", "testVariable##SK##company"),
			resource.TestCheckNoResourceAttr(resourceFullName, "empty_value"),
			resource.TestCheckResourceAttr(resourceFullName, "value_service", "testVariable##SK##company"),
		),
	}

	step2 := resource.TestStep{
		Config: testAccResourceVariable_UnknownValue2_Hcl(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "value", "testVariable3##SK##company"),
			resource.TestCheckNoResourceAttr(resourceFullName, "empty_value"),
			resource.TestCheckResourceAttr(resourceFullName, "value_service", "testVariable3##SK##company"),
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
		CheckDestroy:             davinci.Variable_CheckDestroy(),
		Steps: []resource.TestStep{
			step1,
			step2,
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
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Variable_CheckDestroy(),
		Steps: []resource.TestStep{
			// Configure
			{
				Config: hcl,
			},
			// Errors
			{
				ResourceName: resourceFullName,
				ImportState:  true,
				ExpectError:  regexp.MustCompile(`Unexpected Import Identifier`),
			},
			{
				ResourceName:  resourceFullName,
				ImportStateId: "/",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`Unexpected Import Identifier`),
			},
			{
				ResourceName:  resourceFullName,
				ImportStateId: "badformat/badformat",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`Unexpected Import Identifier`),
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

func testAccResourceVariable_CompanyContext_SecretDynamic_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	context := "company"
	return fmt.Sprintf(`
%[1]s

resource "davinci_variable" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  name           = "%[3]s"
  context        = "%[4]s"
  type           = "secret"
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name, context)
}

func testAccResourceVariable_CompanyContext_SecretStatic_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	context := "company"
	return fmt.Sprintf(`
%[1]s

resource "davinci_variable" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  name           = "%[3]s"
  context        = "%[4]s"
  type           = "secret"
  value          = "mysecret"
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name, context)
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

func testAccResourceVariable_DynamicValue_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	mainFlowJson, err := acctest.ReadFlowJsonFile("flows/full-basic-vars.json")
	if err != nil {
		return ""
	}

	prevariablesHCL := fmt.Sprintf(`
resource "davinci_variable" "%[1]s-company" {
  environment_id = pingone_environment.%[1]s.id

  name        = "testVariable"
  context     = "company"
  description = "testVariable description"
  type        = "string"
}

resource "davinci_variable" "%[1]s-flowInstance" {
  environment_id = pingone_environment.%[1]s.id

  name        = "flowInstanceVariable1"
  context     = "flowInstance"
  description = "flowInstanceVariable1 description"
  type        = "string"
}`, resourceName)

	prevariablesDependsHCL := fmt.Sprintf(`depends_on = [davinci_variable.%[1]s-flowInstance, davinci_variable.%[1]s-company]`, resourceName)

	postvariablesHCL := fmt.Sprintf(`
resource "davinci_variable" "%[1]s-flow" {
  environment_id = pingone_environment.%[1]s.id
  flow_id        = davinci_flow.%[1]s.id

  name        = "fdgdfgfdg"
  context     = "flow"
  description = "fdgdfgfdg description"
  type        = "string"
}`, resourceName)

	hcl, _ = testAccResourceFlow_Variable(resourceName, name, withBootstrapConfig, mainFlowJson, prevariablesHCL, prevariablesDependsHCL, postvariablesHCL)

	return hcl
}

func testAccResourceVariable_StaticValue_Hcl(resourceName, name string, withBootstrapConfig bool, variableValue string) (hcl string) {
	mainFlowJson, err := acctest.ReadFlowJsonFile("flows/full-basic-vars.json")
	if err != nil {
		return ""
	}

	prevariablesHCL := fmt.Sprintf(`
resource "davinci_variable" "%[1]s-company" {
  environment_id = pingone_environment.%[1]s.id

  name        = "testVariable"
  context     = "company"
  description = "testVariable description"
  type        = "string"
  value       = "%[2]s"
}

resource "davinci_variable" "%[1]s-flowInstance" {
  environment_id = pingone_environment.%[1]s.id

  name        = "flowInstanceVariable1"
  context     = "flowInstance"
  description = "flowInstanceVariable1 description"
  type        = "string"
  value       = "%[2]s"
}`, resourceName, variableValue)

	prevariablesDependsHCL := fmt.Sprintf(`depends_on = [davinci_variable.%[1]s-flowInstance, davinci_variable.%[1]s-company]`, resourceName)

	postvariablesHCL := fmt.Sprintf(`
resource "davinci_variable" "%[1]s-flow" {
  environment_id = pingone_environment.%[1]s.id
  flow_id        = davinci_flow.%[1]s.id

  name        = "fdgdfgfdg"
  context     = "flow"
  description = "fdgdfgfdg description"
  type        = "string"
  value       = "%[2]s"
}`, resourceName, variableValue)

	hcl, _ = testAccResourceFlow_Variable(resourceName, name, withBootstrapConfig, mainFlowJson, prevariablesHCL, prevariablesDependsHCL, postvariablesHCL)

	return hcl
}

func testAccResourceVariable_UnknownValue1_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	hcl = fmt.Sprintf(`
resource "davinci_variable" "%[1]s" {
  environment_id = pingone_environment.%[1]s.id

  name        = "testVariable2"
  context     = "company"
  description = "testVariable2 description"
  type        = "string"
  value       = davinci_variable.%[1]s-company.id
}

%s`, resourceName, testAccResourceVariable_StaticValue_Hcl(resourceName, name, withBootstrapConfig, "myVal123"))

	return
}

func testAccResourceVariable_UnknownValue2_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	hcl = fmt.Sprintf(`
resource "davinci_variable" "%[1]s" {
  environment_id = pingone_environment.%[1]s.id

  name        = "testVariable2"
  context     = "company"
  description = "testVariable2 description"
  type        = "string"
  value       = davinci_variable.%[1]s-test.id
}

resource "davinci_variable" "%[1]s-test" {
  environment_id = pingone_environment.%[1]s.id

  name        = "testVariable3"
  context     = "company"
  description = "testVariable3 description"
  type        = "string"
}

%s`, resourceName, testAccResourceVariable_StaticValue_Hcl(resourceName, name, withBootstrapConfig, "myVal123"))

	return
}

func testAccResourceVariable_EmptyValue_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	mainFlowJson, err := acctest.ReadFlowJsonFile("flows/full-basic-vars.json")
	if err != nil {
		return ""
	}

	prevariablesHCL := fmt.Sprintf(`
resource "davinci_variable" "%[1]s-company" {
  environment_id = pingone_environment.%[1]s.id

  name        = "testVariable"
  context     = "company"
  description = "testVariable description"
  type        = "string"
  empty_value = true
}

resource "davinci_variable" "%[1]s-flowInstance" {
  environment_id = pingone_environment.%[1]s.id

  name        = "flowInstanceVariable1"
  context     = "flowInstance"
  description = "flowInstanceVariable1 description"
  type        = "string"
  empty_value = true
}`, resourceName)

	prevariablesDependsHCL := fmt.Sprintf(`depends_on = [davinci_variable.%[1]s-flowInstance, davinci_variable.%[1]s-company]`, resourceName)

	postvariablesHCL := fmt.Sprintf(`
resource "davinci_variable" "%[1]s-flow" {
  environment_id = pingone_environment.%[1]s.id
  flow_id        = davinci_flow.%[1]s.id

  name        = "fdgdfgfdg"
  context     = "flow"
  description = "fdgdfgfdg description"
  type        = "string"
  empty_value = true
}`, resourceName)

	hcl, _ = testAccResourceFlow_Variable(resourceName, name, withBootstrapConfig, mainFlowJson, prevariablesHCL, prevariablesDependsHCL, postvariablesHCL)

	return hcl
}
