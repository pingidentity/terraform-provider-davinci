package davinci_test

import (
	"encoding/json"
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

func TestAccResourceFlow_RemovalDrift(t *testing.T) {

	resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	minimalStepHcl, _, err := testAccResourceFlow_Minimal_WithMappingIDs_HCL(resourceName, name, false)
	if err != nil {
		t.Fatalf("Failed to get HCL: %v", err)
	}

	var flowID, environmentID string

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
		CheckDestroy:             davinci.Flow_CheckDestroy(),
		Steps: []resource.TestStep{
			// Configure
			{
				Config: minimalStepHcl,
				Check:  davinci.Flow_GetIDs(resourceFullName, &environmentID, &flowID),
			},
			// Replan after removal preconfig
			{
				PreConfig: func() {
					davinci.Flow_RemovalDrift_PreConfig(t, environmentID, flowID)
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
			},
			// Test removal of the environment
			{
				Config:   minimalStepHcl,
				Check:    davinci.Flow_GetIDs(resourceFullName, &environmentID, &flowID),
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

func TestAccResourceFlow_Basic_Clean(t *testing.T) {
	testAccResourceFlow_Basic(t, false)
}

func TestAccResourceFlow_Basic_WithBootstrap(t *testing.T) {
	testAccResourceFlow_Basic(t, true)
}

func testAccResourceFlow_Basic(t *testing.T, withBootstrapConfig bool) {

	resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	fullStepHcl, fullStepJson, err := testAccResourceFlow_Full_WithMappingIDs_HCL(resourceName, name, withBootstrapConfig)
	if err != nil {
		t.Fatalf("Failed to get HCL: %v", err)
	}

	fullStep := resource.TestStep{
		Config: fullStepHcl,
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", "my awesome flow"),
			resource.TestCheckResourceAttr(resourceFullName, "description", "my awesome flow description"),
			resource.TestCheckResourceAttr(resourceFullName, "flow_json", fmt.Sprintf("%s\n", fullStepJson)),
			resource.TestCheckResourceAttrSet(resourceFullName, "flow_configuration_json"),
			resource.TestCheckResourceAttrSet(resourceFullName, "flow_export_json"),
			resource.TestCheckResourceAttr(resourceFullName, "connection_link.#", "5"),
			resource.TestCheckResourceAttr(resourceFullName, "deploy", "true"),
			resource.TestCheckResourceAttr(resourceFullName, "subflow_link.#", "2"),
			resource.TestCheckResourceAttr(resourceFullName, "flow_variables.#", "2"),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "flow_variables.*", map[string]*regexp.Regexp{
				"context": regexp.MustCompile(`^flow$`),
				"flow_id": verify.P1DVResourceIDRegexpFullString,
				"id":      regexp.MustCompile(fmt.Sprintf(`^fdgdfgfdg##SK##flow##SK##%s$`, verify.P1DVResourceIDRegexp.String())),
				"max":     regexp.MustCompile(`^2000$`),
				"min":     regexp.MustCompile(`^0$`),
				"mutable": regexp.MustCompile(`^true$`),
				"name":    regexp.MustCompile(`^fdgdfgfdg$`),
				"type":    regexp.MustCompile(`^property$`),
			}),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "flow_variables.*", map[string]*regexp.Regexp{
				"context": regexp.MustCompile(`^flow$`),
				"flow_id": verify.P1DVResourceIDRegexpFullString,
				"id":      regexp.MustCompile(fmt.Sprintf(`^test123##SK##flow##SK##%s$`, verify.P1DVResourceIDRegexp.String())),
				"max":     regexp.MustCompile(`^20$`),
				"min":     regexp.MustCompile(`^4$`),
				"mutable": regexp.MustCompile(`^true$`),
				"name":    regexp.MustCompile(`^test123$`),
				"type":    regexp.MustCompile(`^property$`),
			}),
		),
	}

	minimalStepHcl, minimalStepJson, err := testAccResourceFlow_Minimal_WithMappingIDs_HCL(resourceName, name, withBootstrapConfig)
	if err != nil {
		t.Fatalf("Failed to get HCL: %v", err)
	}

	minimalStep := resource.TestStep{
		Config: minimalStepHcl,
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", "simple"),
			resource.TestCheckResourceAttr(resourceFullName, "description", "This is a fallback description"),
			resource.TestCheckResourceAttr(resourceFullName, "flow_json", fmt.Sprintf("%s\n", minimalStepJson)),
			resource.TestCheckResourceAttr(resourceFullName, "connection_link.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "deploy", "true"),
			resource.TestCheckResourceAttr(resourceFullName, "subflow_link.#", "0"),
			resource.TestCheckResourceAttr(resourceFullName, "flow_variables.#", "0"),
		),
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)

			if withBootstrapConfig {
				t.Skipf("Skipping test with bootstrap config: https://github.com/pingidentity/terraform-provider-davinci/issues/266")
			}
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Flow_CheckDestroy(),
		Steps: []resource.TestStep{
			// Create full from scratch
			fullStep,
			{
				Config:  fullStepHcl,
				Destroy: true,
			},
			// Create minimal from scratch
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

						return fmt.Sprintf("%s/%s", rs.Primary.Attributes["environment_id"], rs.Primary.ID), nil
					}
				}(),
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"connection_link",
					"subflow_link",
					"flow_json",
				},
			},
		},
	})
}

func TestAccResourceFlow_ConnectionSubflowLinks_WithMappingIDs_Clean(t *testing.T) {
	testAccResourceFlow_ConnectionSubflowLinks_WithMappingIDs(t, false)
}

func TestAccResourceFlow_ConnectionSubflowLinks_WithMappingIDs_WithBootstrap(t *testing.T) {
	testAccResourceFlow_ConnectionSubflowLinks_WithMappingIDs(t, true)
}

func testAccResourceFlow_ConnectionSubflowLinks_WithMappingIDs(t *testing.T, withBootstrapConfig bool) {

	resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	fullStepHcl, fullStepJson, err := testAccResourceFlow_Full_WithMappingIDs_HCL(resourceName, name, withBootstrapConfig)
	if err != nil {
		t.Fatalf("Failed to get HCL: %v", err)
	}

	fullStep := resource.TestStep{
		Config: fullStepHcl,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "flow_json", fmt.Sprintf("%s\n", fullStepJson)),
			resource.TestCheckResourceAttrSet(resourceFullName, "flow_configuration_json"),
			resource.TestCheckResourceAttrSet(resourceFullName, "flow_export_json"),
			resource.TestCheckResourceAttr(resourceFullName, "connection_link.#", "5"),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "connection_link.*", map[string]*regexp.Regexp{
				"id":                           verify.P1DVResourceIDRegexpFullString,
				"replace_import_connection_id": verify.P1DVResourceIDRegexpFullString,
				"name":                         regexp.MustCompile(fmt.Sprintf(`^%s-variables$`, name)),
			}),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "connection_link.*", map[string]*regexp.Regexp{
				"id":                           verify.P1DVResourceIDRegexpFullString,
				"replace_import_connection_id": verify.P1DVResourceIDRegexpFullString,
				"name":                         regexp.MustCompile(fmt.Sprintf(`^%s-http$`, name)),
			}),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "connection_link.*", map[string]*regexp.Regexp{
				"id":                           verify.P1DVResourceIDRegexpFullString,
				"replace_import_connection_id": verify.P1DVResourceIDRegexpFullString,
				"name":                         regexp.MustCompile(fmt.Sprintf(`^%s-functions$`, name)),
			}),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "connection_link.*", map[string]*regexp.Regexp{
				"id":                           verify.P1DVResourceIDRegexpFullString,
				"replace_import_connection_id": verify.P1DVResourceIDRegexpFullString,
				"name":                         regexp.MustCompile(fmt.Sprintf(`^%s-flow$`, name)),
			}),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "connection_link.*", map[string]*regexp.Regexp{
				"id":                           verify.P1DVResourceIDRegexpFullString,
				"replace_import_connection_id": verify.P1DVResourceIDRegexpFullString,
				"name":                         regexp.MustCompile(fmt.Sprintf(`^%s-error$`, name)),
			}),
			resource.TestCheckResourceAttr(resourceFullName, "subflow_link.#", "2"),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "subflow_link.*", map[string]*regexp.Regexp{
				"id":                        verify.P1DVResourceIDRegexpFullString,
				"replace_import_subflow_id": verify.P1DVResourceIDRegexpFullString,
				"name":                      regexp.MustCompile(fmt.Sprintf(`^%s-subflow-1$`, name)),
			}),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "subflow_link.*", map[string]*regexp.Regexp{
				"id":                        verify.P1DVResourceIDRegexpFullString,
				"replace_import_subflow_id": verify.P1DVResourceIDRegexpFullString,
				"name":                      regexp.MustCompile(fmt.Sprintf(`^%s-subflow-2$`, name)),
			}),
		),
	}

	minimalStepHcl, minimalStepJson, err := testAccResourceFlow_Minimal_WithMappingIDs_HCL(resourceName, name, withBootstrapConfig)
	if err != nil {
		t.Fatalf("Failed to get HCL: %v", err)
	}

	minimalStep := resource.TestStep{
		Config: minimalStepHcl,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "flow_json", fmt.Sprintf("%s\n", minimalStepJson)),
			resource.TestCheckResourceAttrSet(resourceFullName, "flow_configuration_json"),
			resource.TestCheckResourceAttrSet(resourceFullName, "flow_export_json"),
			resource.TestCheckResourceAttr(resourceFullName, "connection_link.#", "1"),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "connection_link.*", map[string]*regexp.Regexp{
				"id":                           verify.P1DVResourceIDRegexpFullString,
				"replace_import_connection_id": verify.P1DVResourceIDRegexpFullString,
				"name":                         regexp.MustCompile(fmt.Sprintf(`^%s-error$`, name)),
			}),
			resource.TestCheckResourceAttr(resourceFullName, "subflow_link.#", "0"),
		),
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)

			if withBootstrapConfig {
				t.Skipf("Skipping test with bootstrap config: https://github.com/pingidentity/terraform-provider-davinci/issues/266")
			}
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Flow_CheckDestroy(),
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

						return fmt.Sprintf("%s/%s", rs.Primary.Attributes["environment_id"], rs.Primary.ID), nil
					}
				}(),
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"connection_link",
					"subflow_link",
					"flow_json",
				},
			},
		},
	})
}

func TestAccResourceFlow_ConnectionSubflowLinks_WithoutMappingIDs_Clean(t *testing.T) {
	testAccResourceFlow_ConnectionSubflowLinks_WithoutMappingIDs(t, false)
}

func TestAccResourceFlow_ConnectionSubflowLinks_WithoutMappingIDs_WithBootstrap(t *testing.T) {
	testAccResourceFlow_ConnectionSubflowLinks_WithoutMappingIDs(t, true)
}

func testAccResourceFlow_ConnectionSubflowLinks_WithoutMappingIDs(t *testing.T, withBootstrapConfig bool) {

	resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	fullStepHcl, fullStepJson, err := testAccResourceFlow_Full_WithoutMappingIDs_HCL(resourceName, name, withBootstrapConfig)
	if err != nil {
		t.Fatalf("Failed to get HCL: %v", err)
	}

	fullStep := resource.TestStep{
		Config: fullStepHcl,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "flow_json", fmt.Sprintf("%s\n", fullStepJson)),
			resource.TestCheckResourceAttrSet(resourceFullName, "flow_configuration_json"),
			resource.TestCheckResourceAttrSet(resourceFullName, "flow_export_json"),
			resource.TestCheckResourceAttr(resourceFullName, "connection_link.#", "5"),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "connection_link.*", map[string]*regexp.Regexp{
				"id":   verify.P1DVResourceIDRegexpFullString,
				"name": regexp.MustCompile(`^Variables$`),
			}),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "connection_link.*", map[string]*regexp.Regexp{
				"id":   verify.P1DVResourceIDRegexpFullString,
				"name": regexp.MustCompile(`^Http$`),
			}),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "connection_link.*", map[string]*regexp.Regexp{
				"id":   verify.P1DVResourceIDRegexpFullString,
				"name": regexp.MustCompile(`^Functions$`),
			}),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "connection_link.*", map[string]*regexp.Regexp{
				"id":   verify.P1DVResourceIDRegexpFullString,
				"name": regexp.MustCompile(`^Flow Connector$`),
			}),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "connection_link.*", map[string]*regexp.Regexp{
				"id":   verify.P1DVResourceIDRegexpFullString,
				"name": regexp.MustCompile(`^Error Message$`),
			}),
			resource.TestCheckResourceAttr(resourceFullName, "subflow_link.#", "2"),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "subflow_link.*", map[string]*regexp.Regexp{
				"id":   verify.P1DVResourceIDRegexpFullString,
				"name": regexp.MustCompile(`^subflow 1$`),
			}),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "subflow_link.*", map[string]*regexp.Regexp{
				"id":   verify.P1DVResourceIDRegexpFullString,
				"name": regexp.MustCompile(`^subflow 2$`),
			}),
		),
	}

	minimalStepHcl, minimalStepJson, err := testAccResourceFlow_Minimal_WithoutMappingIDs_HCL(resourceName, name, withBootstrapConfig)
	if err != nil {
		t.Fatalf("Failed to get HCL: %v", err)
	}

	minimalStep := resource.TestStep{
		Config: minimalStepHcl,
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "flow_json", fmt.Sprintf("%s\n", minimalStepJson)),
			resource.TestCheckResourceAttrSet(resourceFullName, "flow_configuration_json"),
			resource.TestCheckResourceAttrSet(resourceFullName, "flow_export_json"),
			resource.TestCheckResourceAttr(resourceFullName, "connection_link.#", "1"),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "connection_link.*", map[string]*regexp.Regexp{
				"id":   verify.P1DVResourceIDRegexpFullString,
				"name": regexp.MustCompile(`^Error Message$`),
			}),
			resource.TestCheckResourceAttr(resourceFullName, "subflow_link.#", "0"),
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
		CheckDestroy:             davinci.Flow_CheckDestroy(),
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

						return fmt.Sprintf("%s/%s", rs.Primary.Attributes["environment_id"], rs.Primary.ID), nil
					}
				}(),
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"connection_link",
					"subflow_link",
					"flow_json",
				},
			},
		},
	})
}

func TestAccResourceFlow_ComputeDifferences_ModifySettings(t *testing.T) {

	// Baseline
	mainFlowJson, err := acctest.ReadFlowJsonFile("flows/full-minimal.json")
	if err != nil {
		t.Fatalf("Failed to get HCL: %v", err)
	}

	var flow dv.Flow
	if err := json.Unmarshal([]byte(mainFlowJson), &flow); err != nil {
		t.Fatalf("Failed to unmarshal flow: %v", err)
	}

	testAccResourceFlow_ComputeDifferences(t, computeDifferencesTest{
		BaselineFlow: flow,
		ModifiedFlow: func() dv.Flow {
			newFlow := flow
			newFlow.Settings = map[string]interface{}{
				"csp":                           "worker-src 'self' blob:; script-src 'self' https://cdn.jsdelivr.net https://code.jquery.com https://devsdk.singularkey.com http://cdnjs.cloudflare.com 'unsafe-inline' 'unsafe-eval';",
				"intermediateLoadingScreenCSS":  "",
				"intermediateLoadingScreenHTML": "",
				"flowHttpTimeoutInSeconds":      301,
			}

			return newFlow
		}(),
		ExpectNonEmptyPlan: true,
	})
}

func TestAccResourceFlow_ComputeDifferences_CompanyId(t *testing.T) {

	// Baseline
	mainFlowJson, err := acctest.ReadFlowJsonFile("flows/full-minimal.json")
	if err != nil {
		t.Fatalf("Failed to get HCL: %v", err)
	}

	var flow dv.Flow
	if err := json.Unmarshal([]byte(mainFlowJson), &flow); err != nil {
		t.Fatalf("Failed to unmarshal flow: %v", err)
	}

	testAccResourceFlow_ComputeDifferences(t, computeDifferencesTest{
		BaselineFlow: flow,
		ModifiedFlow: func() dv.Flow {
			newFlow := flow
			newFlow.CompanyID = "12345678-1234-1234-1234-123456789012" // dummy UUID

			return newFlow
		}(),
		ExpectNonEmptyPlan: true,
	})
}

func TestAccResourceFlow_ComputeDifferences_Version(t *testing.T) {

	// Baseline
	mainFlowJson, err := acctest.ReadFlowJsonFile("flows/full-minimal.json")
	if err != nil {
		t.Fatalf("Failed to get HCL: %v", err)
	}

	var flow dv.Flow
	if err := json.Unmarshal([]byte(mainFlowJson), &flow); err != nil {
		t.Fatalf("Failed to unmarshal flow: %v", err)
	}

	testAccResourceFlow_ComputeDifferences(t, computeDifferencesTest{
		BaselineFlow: flow,
		ModifiedFlow: func() dv.Flow {
			newFlow := flow
			newFlow.CurrentVersion = func() *int32 {
				v := int32(32) // dummy version
				return &v
			}()

			return newFlow
		}(),
		ExpectNonEmptyPlan: true,
	})
}

func TestAccResourceFlow_ComputeDifferences_Description(t *testing.T) {

	// Baseline
	mainFlowJson, err := acctest.ReadFlowJsonFile("flows/full-minimal.json")
	if err != nil {
		t.Fatalf("Failed to get HCL: %v", err)
	}

	var flow dv.Flow
	if err := json.Unmarshal([]byte(mainFlowJson), &flow); err != nil {
		t.Fatalf("Failed to unmarshal flow: %v", err)
	}

	testAccResourceFlow_ComputeDifferences(t, computeDifferencesTest{
		BaselineFlow: flow,
		ModifiedFlow: func() dv.Flow {
			newFlow := flow
			newFlow.Description = func() *string {
				v := "What do you call a dinosaur playing hide and seek?  Doyouthinkhesawus Rex"
				return &v
			}()

			return newFlow
		}(),
		ExpectNonEmptyPlan: true,
	})
}

func TestAccResourceFlow_ComputeDifferences_AdditionalProperties(t *testing.T) {

	// Baseline
	mainFlowJson, err := acctest.ReadFlowJsonFile("flows/full-minimal.json")
	if err != nil {
		t.Fatalf("Failed to get HCL: %v", err)
	}

	var flow dv.Flow
	if err := json.Unmarshal([]byte(mainFlowJson), &flow); err != nil {
		t.Fatalf("Failed to unmarshal flow: %v", err)
	}

	testAccResourceFlow_ComputeDifferences(t, computeDifferencesTest{
		BaselineFlow: flow,
		ModifiedFlow: func() dv.Flow {
			newFlow := flow
			newFlow.AdditionalProperties = map[string]interface{}{
				"foo": "bar",
			}

			return newFlow
		}(),
		ExpectNonEmptyPlan: true,
	})
}

func TestAccResourceFlow_ComputeDifferences_NewNode(t *testing.T) {

	// Baseline
	mainFlowJson, err := acctest.ReadFlowJsonFile("flows/full-minimal.json")
	if err != nil {
		t.Fatalf("Failed to get HCL: %v", err)
	}

	var flow dv.Flow
	if err := json.Unmarshal([]byte(mainFlowJson), &flow); err != nil {
		t.Fatalf("Failed to unmarshal flow: %v", err)
	}

	testAccResourceFlow_ComputeDifferences(t, computeDifferencesTest{
		BaselineFlow: flow,
		ModifiedFlow: func() dv.Flow {
			var newFlow dv.Flow
			err := acctest.DeepCloneFlow(&flow, &newFlow)
			if err != nil {
				t.Fatalf("Failed to clone flow: %v", err)
			}

			newFlow.GraphData.Elements.Nodes = append(newFlow.GraphData.Elements.Nodes, dv.Node{
				Data: func() *dv.NodeData {
					v := dv.NodeData{
						ID: func() *string {
							v := "1u2m5vzr50"
							return &v
						}(),
						NodeType: func() *string {
							v := "CONNECTION"
							return &v
						}(),
						ConnectionID: func() *string {
							v := "867ed4363b2bc21c860085ad2baa817d"
							return &v
						}(),
						ConnectorID: func() *string {
							v := "httpConnector"
							return &v
						}(),
						Name: func() *string {
							v := "Http"
							return &v
						}(),
						Label: func() *string {
							v := "Http"
							return &v
						}(),
						Status: func() *string {
							v := "configured"
							return &v
						}(),
						CapabilityName: func() *string {
							v := "customHtmlMessage"
							return &v
						}(),
						Type: func() *string {
							v := "trigger"
							return &v
						}(),
						Properties: func() *dv.Properties {
							v := dv.Properties{
								AdditionalProperties: map[string]interface{}{
									"message": map[string]interface{}{
										"value": "[\n  {\n    \"children\": [\n      {\n        \"text\": \"Hello, world?\"\n      }\n    ]\n  }\n]",
									},
								},
							}
							return &v
						}(),
					}
					return &v
				}(),
				Position: func() *dv.Position {
					v := dv.Position{
						X: func() *float64 {
							v := float64(277)
							return &v
						}(),
						Y: func() *float64 {
							v := float64(336)
							return &v
						}(),
					}
					return &v
				}(),
				Group:      "nodes",
				Removed:    false,
				Selected:   false,
				Selectable: true,
				Locked:     false,
				Grabbable:  true,
				Pannable:   false,
			})

			return newFlow
		}(),
		ExpectNonEmptyPlan: true,
	})
}

type computeDifferencesTest struct {
	BaselineFlow       dv.Flow
	ModifiedFlow       dv.Flow
	ExpectNonEmptyPlan bool
}

func testAccResourceFlow_ComputeDifferences(t *testing.T, cd computeDifferencesTest) {

	resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	baseLineHcl, err := testAccResourceFlow_ComputeDifferences_HCL(resourceName, cd.BaselineFlow)
	if err != nil {
		t.Fatalf("Failed to get baseline flow HCL: %v", err)
	}

	modifiedHcl, err := testAccResourceFlow_ComputeDifferences_HCL(resourceName, cd.ModifiedFlow)
	if err != nil {
		t.Fatalf("Failed to get modified flow HCL: %v", err)
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Flow_CheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: baseLineHcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
				),
			},
			{
				Config:             modifiedHcl,
				ExpectNonEmptyPlan: cd.ExpectNonEmptyPlan,
				PlanOnly:           true,
			},
			{
				Config: modifiedHcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
				),
			},
		},
	})
}

func TestAccResourceFlow_UnknownFlowString(t *testing.T) {

	resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	stepHcl, err := testAccResourceFlow_UnknownFlow_HCL(resourceName, name)
	if err != nil {
		t.Fatalf("Failed to get HCL: %v", err)
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Flow_CheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config: stepHcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
				),
			},
		},
	})
}

func TestAccResourceFlow_BrokenFlow(t *testing.T) {

	resourceName := acctest.ResourceNameGen()

	brokenStepHcl, err := testAccResourceFlow_Broken_HCL(resourceName)
	if err != nil {
		t.Fatalf("Failed to get HCL: %v", err)
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Flow_CheckDestroy(),
		Steps: []resource.TestStep{
			{
				Config:      brokenStepHcl,
				ExpectError: regexp.MustCompile(`Error importing flow`),
			},
		},
	})
}

func testAccResourceFlow_Common_WithMappingIDs_HCL(resourceName, name string) (hcl string, err error) {

	subFlow1Json, err := acctest.ReadFlowJsonFile("flows/full-basic-subflow-1.json")
	if err != nil {
		return "", err
	}
	subFlow2Json, err := acctest.ReadFlowJsonFile("flows/full-basic-subflow-2.json")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`
// Variables connector
resource "davinci_connection" "%[1]s-variables" {
  environment_id = pingone_environment.%[1]s.id
  connector_id   = "variablesConnector"
  name           = "%[2]s-variables"
}

// Http connector
resource "davinci_connection" "%[1]s-http" {
  environment_id = pingone_environment.%[1]s.id
  connector_id   = "httpConnector"
  name           = "%[2]s-http"
}

// Functions connector
resource "davinci_connection" "%[1]s-functions" {
  environment_id = pingone_environment.%[1]s.id
  connector_id   = "functionsConnector"
  name           = "%[2]s-functions"
}

// Flow conductor connector
resource "davinci_connection" "%[1]s-flow" {
  environment_id = pingone_environment.%[1]s.id
  connector_id   = "flowConnector"
  name           = "%[2]s-flow"
}

// Error connector
resource "davinci_connection" "%[1]s-error" {
  environment_id = pingone_environment.%[1]s.id
  connector_id   = "errorConnector"
  name           = "%[2]s-error"
}

resource "davinci_flow" "%[1]s-subflow-1" {
  environment_id = pingone_environment.%[1]s.id

  name = "%[2]s-subflow-1"

  flow_json = <<EOT
%[3]s
EOT

  // Http connector
  connection_link {
    id                           = davinci_connection.%[1]s-http.id
    name                         = davinci_connection.%[1]s-http.name
    replace_import_connection_id = "867ed4363b2bc21c860085ad2baa817d"
  }
}

resource "davinci_flow" "%[1]s-subflow-2" {
  environment_id = pingone_environment.%[1]s.id

  name = "%[2]s-subflow-2"

  flow_json = <<EOT
%[4]s
EOT

  // Http connector
  connection_link {
    id                           = davinci_connection.%[1]s-http.id
    name                         = davinci_connection.%[1]s-http.name
    replace_import_connection_id = "867ed4363b2bc21c860085ad2baa817d"
  }
}
`, resourceName, name, subFlow1Json, subFlow2Json), nil
}

func testAccResourceFlow_Full_WithMappingIDs_HCL(resourceName, name string, withBootstrapConfig bool) (hcl, mainFlowJson string, err error) {

	mainFlowJson, err = acctest.ReadFlowJsonFile("flows/full-basic.json")
	if err != nil {
		return "", "", err
	}

	commonHcl, err := testAccResourceFlow_Common_WithMappingIDs_HCL(resourceName, name)
	if err != nil {
		return "", "", err
	}

	return fmt.Sprintf(`
%[1]s

%[2]s

resource "davinci_flow" "%[3]s" {
  environment_id = pingone_environment.%[3]s.id

  name        = "my awesome flow"
  description = "my awesome flow description"

  flow_json = <<EOT
%[4]s
EOT

  // Variables connector
  connection_link {
    id                           = davinci_connection.%[3]s-variables.id
    name                         = davinci_connection.%[3]s-variables.name
    replace_import_connection_id = "06922a684039827499bdbdd97f49827b"
  }

  // Http connector
  connection_link {
    id                           = davinci_connection.%[3]s-http.id
    name                         = davinci_connection.%[3]s-http.name
    replace_import_connection_id = "867ed4363b2bc21c860085ad2baa817d"
  }

  // Functions connector
  connection_link {
    id                           = davinci_connection.%[3]s-functions.id
    name                         = davinci_connection.%[3]s-functions.name
    replace_import_connection_id = "de650ca45593b82c49064ead10b9fe17"
  }

  // Flow connector
  connection_link {
    id                           = davinci_connection.%[3]s-flow.id
    name                         = davinci_connection.%[3]s-flow.name
    replace_import_connection_id = "2581eb287bb1d9bd29ae9886d675f89f"
  }

  // Error connector
  connection_link {
    id                           = davinci_connection.%[3]s-error.id
    name                         = davinci_connection.%[3]s-error.name
    replace_import_connection_id = "53ab83a4a4ab919d9f2cb02d9e111ac8"
  }

  // Subflow 2
  subflow_link {
    id                        = davinci_flow.%[3]s-subflow-2.id
    name                      = davinci_flow.%[3]s-subflow-2.name
    replace_import_subflow_id = "07503fed5c02849dbbd5ee932da654b2"
  }

  // Subflow 1
  subflow_link {
    id                        = davinci_flow.%[3]s-subflow-1.id
    name                      = davinci_flow.%[3]s-subflow-1.name
    replace_import_subflow_id = "00f66e8926ced6ef5b83619fde4a314a"
  }
}`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), commonHcl, resourceName, mainFlowJson), mainFlowJson, nil
}

func testAccResourceFlow_Minimal_WithMappingIDs_HCL(resourceName, name string, withBootstrapConfig bool) (hcl, mainFlowJson string, err error) {

	mainFlowJson, err = acctest.ReadFlowJsonFile("flows/full-minimal.json")
	if err != nil {
		return "", "", err
	}

	commonHcl, err := testAccResourceFlow_Common_WithMappingIDs_HCL(resourceName, name)
	if err != nil {
		return "", "", err
	}

	return fmt.Sprintf(`
%[1]s

%[2]s

resource "davinci_flow" "%[3]s" {
  environment_id = pingone_environment.%[3]s.id

  flow_json = <<EOT
%[4]s
EOT

  // Error connector
  connection_link {
    id                           = davinci_connection.%[3]s-error.id
    name                         = davinci_connection.%[3]s-error.name
    replace_import_connection_id = "53ab83a4a4ab919d9f2cb02d9e111ac8"
  }
}`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), commonHcl, resourceName, mainFlowJson), mainFlowJson, nil
}

func testAccResourceFlow_Common_WithoutMappingIDs_HCL_Resources(resourceName, name string) (hcl string, err error) {

	subFlow1Json, err := acctest.ReadFlowJsonFile("flows/full-basic-subflow-1.json")
	if err != nil {
		return "", err
	}
	subFlow2Json, err := acctest.ReadFlowJsonFile("flows/full-basic-subflow-2.json")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`
// Variables connector
resource "davinci_connection" "%[1]s-variables" {
  environment_id = pingone_environment.%[1]s.id
  connector_id   = "variablesConnector"
  name           = "Variables"
}

// Http connector
resource "davinci_connection" "%[1]s-http" {
  environment_id = pingone_environment.%[1]s.id
  connector_id   = "httpConnector"
  name           = "Http"
}

// Functions connector
resource "davinci_connection" "%[1]s-functions" {
  environment_id = pingone_environment.%[1]s.id
  connector_id   = "functionsConnector"
  name           = "Functions"
}

// Flow conductor connector
resource "davinci_connection" "%[1]s-flow" {
  environment_id = pingone_environment.%[1]s.id
  connector_id   = "flowConnector"
  name           = "Flow Connector"
}

// Error connector
resource "davinci_connection" "%[1]s-error" {
  environment_id = pingone_environment.%[1]s.id
  connector_id   = "errorConnector"
  name           = "Error Message"
}

resource "davinci_flow" "%[1]s-subflow-1" {
  environment_id = pingone_environment.%[1]s.id

  name = "subflow 1"

  flow_json = <<EOT
%[3]s
EOT

  // Http connector
  connection_link {
    id   = davinci_connection.%[1]s-http.id
    name = davinci_connection.%[1]s-http.name
  }
}

resource "davinci_flow" "%[1]s-subflow-2" {
  environment_id = pingone_environment.%[1]s.id

  name = "subflow 2"

  flow_json = <<EOT
%[4]s
EOT

  // Http connector
  connection_link {
    id   = davinci_connection.%[1]s-http.id
    name = davinci_connection.%[1]s-http.name
  }
}

locals {
  davinci_connection_variables_id   = davinci_connection.%[1]s-variables.id
  davinci_connection_variables_name = davinci_connection.%[1]s-variables.name

  davinci_connection_http_id   = davinci_connection.%[1]s-http.id
  davinci_connection_http_name = davinci_connection.%[1]s-http.name

  davinci_connection_functions_id   = davinci_connection.%[1]s-functions.id
  davinci_connection_functions_name = davinci_connection.%[1]s-functions.name

  davinci_connection_flow_id   = davinci_connection.%[1]s-flow.id
  davinci_connection_flow_name = davinci_connection.%[1]s-flow.name

  davinci_connection_error_id   = davinci_connection.%[1]s-error.id
  davinci_connection_error_name = davinci_connection.%[1]s-error.name
}
`, resourceName, name, subFlow1Json, subFlow2Json), nil
}

func testAccResourceFlow_Common_WithoutMappingIDs_HCL_Datasources(resourceName, name string) (hcl string, err error) {

	subFlow1Json, err := acctest.ReadFlowJsonFile("flows/full-basic-subflow-1.json")
	if err != nil {
		return "", err
	}
	subFlow2Json, err := acctest.ReadFlowJsonFile("flows/full-basic-subflow-2.json")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`
// Variables connector
data "davinci_connection" "%[1]s-variables" {
  environment_id = pingone_environment.%[1]s.id
  name           = "Variables"
}

// Http connector
data "davinci_connection" "%[1]s-http" {
  environment_id = pingone_environment.%[1]s.id
  name           = "Http"
}

// Functions connector
data "davinci_connection" "%[1]s-functions" {
  environment_id = pingone_environment.%[1]s.id
  name           = "Functions"
}

// Flow conductor connector
data "davinci_connection" "%[1]s-flow" {
  environment_id = pingone_environment.%[1]s.id
  name           = "Flow Connector"
}

// Error connector
data "davinci_connection" "%[1]s-error" {
  environment_id = pingone_environment.%[1]s.id
  name           = "Error Message"
}

resource "davinci_flow" "%[1]s-subflow-1" {
  environment_id = pingone_environment.%[1]s.id

  name = "subflow 1"

  flow_json = <<EOT
%[3]s
EOT

  // Http connector
  connection_link {
    id   = data.davinci_connection.%[1]s-http.id
    name = data.davinci_connection.%[1]s-http.name
  }
}

resource "davinci_flow" "%[1]s-subflow-2" {
  environment_id = pingone_environment.%[1]s.id

  name = "subflow 2"

  flow_json = <<EOT
%[4]s
EOT

  // Http connector
  connection_link {
    id   = data.davinci_connection.%[1]s-http.id
    name = data.davinci_connection.%[1]s-http.name
  }
}

locals {
  davinci_connection_variables_id   = data.davinci_connection.%[1]s-variables.id
  davinci_connection_variables_name = data.davinci_connection.%[1]s-variables.name

  davinci_connection_http_id   = data.davinci_connection.%[1]s-http.id
  davinci_connection_http_name = data.davinci_connection.%[1]s-http.name

  davinci_connection_functions_id   = data.davinci_connection.%[1]s-functions.id
  davinci_connection_functions_name = data.davinci_connection.%[1]s-functions.name

  davinci_connection_flow_id   = data.davinci_connection.%[1]s-flow.id
  davinci_connection_flow_name = data.davinci_connection.%[1]s-flow.name

  davinci_connection_error_id   = data.davinci_connection.%[1]s-error.id
  davinci_connection_error_name = data.davinci_connection.%[1]s-error.name
}
`, resourceName, name, subFlow1Json, subFlow2Json), nil
}

func testAccResourceFlow_Full_WithoutMappingIDs_HCL(resourceName, name string, withBootstrapConfig bool) (hcl, mainFlowJson string, err error) {

	mainFlowJson, err = acctest.ReadFlowJsonFile("flows/full-basic.json")
	if err != nil {
		return "", "", err
	}

	var commonHcl string

	if !withBootstrapConfig {
		commonHcl, err = testAccResourceFlow_Common_WithoutMappingIDs_HCL_Resources(resourceName, name)
	} else {
		commonHcl, err = testAccResourceFlow_Common_WithoutMappingIDs_HCL_Datasources(resourceName, name)
	}
	if err != nil {
		return "", "", err
	}

	return fmt.Sprintf(`
%[1]s

%[2]s

resource "davinci_flow" "%[3]s" {
  environment_id = pingone_environment.%[3]s.id

  name        = "my awesome flow"
  description = "my awesome flow description"

  flow_json = <<EOT
%[4]s
EOT

  // Variables connector
  connection_link {
    id   = local.davinci_connection_variables_id
    name = local.davinci_connection_variables_name
  }

  // Http connector
  connection_link {
    id   = local.davinci_connection_http_id
    name = local.davinci_connection_http_name
  }

  // Functions connector
  connection_link {
    id   = local.davinci_connection_functions_id
    name = local.davinci_connection_functions_name
  }

  // Flow connector
  connection_link {
    id   = local.davinci_connection_flow_id
    name = local.davinci_connection_flow_name
  }

  // Error connector
  connection_link {
    id   = local.davinci_connection_error_id
    name = local.davinci_connection_error_name
  }

  // Subflow 2
  subflow_link {
    id   = davinci_flow.%[3]s-subflow-2.id
    name = davinci_flow.%[3]s-subflow-2.name
  }

  // Subflow 1
  subflow_link {
    id   = davinci_flow.%[3]s-subflow-1.id
    name = davinci_flow.%[3]s-subflow-1.name
  }
}`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), commonHcl, resourceName, mainFlowJson), mainFlowJson, nil
}

func testAccResourceFlow_Minimal_WithoutMappingIDs_HCL(resourceName, name string, withBootstrapConfig bool) (hcl, mainFlowJson string, err error) {

	mainFlowJson, err = acctest.ReadFlowJsonFile("flows/full-minimal.json")
	if err != nil {
		return "", "", err
	}

	var commonHcl string

	if !withBootstrapConfig {
		commonHcl, err = testAccResourceFlow_Common_WithoutMappingIDs_HCL_Resources(resourceName, name)
	} else {
		commonHcl, err = testAccResourceFlow_Common_WithoutMappingIDs_HCL_Datasources(resourceName, name)
	}
	if err != nil {
		return "", "", err
	}

	return fmt.Sprintf(`
%[1]s

%[2]s

resource "davinci_flow" "%[3]s" {
  environment_id = pingone_environment.%[3]s.id

  flow_json = <<EOT
%[4]s
EOT

  // Error connector
  connection_link {
    id   = local.davinci_connection_error_id
    name = local.davinci_connection_error_name
  }
}`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), commonHcl, resourceName, mainFlowJson), mainFlowJson, nil
}

func testAccResourceFlow_Broken_HCL(resourceName string) (hcl string, err error) {

	mainFlowJson, err := acctest.ReadFlowJsonFile("flows/broken-flow.json")
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`
%[1]s

// Error connector
resource "davinci_connection" "%[2]s-error" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "errorConnector"
  name           = "%[2]s-error"
}

resource "davinci_flow" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id

  flow_json = <<EOT
%[3]s
EOT

  // Error connector
  connection_link {
    id                           = davinci_connection.%[2]s-error.id
    name                         = davinci_connection.%[2]s-error.name
    replace_import_connection_id = "6d8f6f706c45fd459a86b3f092602544"
  }
}`, acctest.PingoneEnvironmentSsoHcl(resourceName, false), resourceName, mainFlowJson), nil
}

func testAccResourceFlow_ComputeDifferences_HCL(resourceName string, flow dv.Flow) (hcl string, err error) {

	flowJson, err := json.Marshal(flow)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`
%[1]s

// Error connector
resource "davinci_connection" "%[2]s-error" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "errorConnector"
  name           = "%[2]s-error"
}

resource "davinci_flow" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id

  flow_json = <<EOT
%[3]s
EOT

  // Error connector
  connection_link {
    id                           = davinci_connection.%[2]s-error.id
    name                         = davinci_connection.%[2]s-error.name
    replace_import_connection_id = "53ab83a4a4ab919d9f2cb02d9e111ac8"
  }
}`, acctest.PingoneEnvironmentSsoHcl(resourceName, false), resourceName, flowJson), nil
}

func testAccResourceFlow_UnknownFlow_HCL(resourceName, name string) (hcl string, err error) {

	commonHcl, err := testAccResourceFlow_Common_WithMappingIDs_HCL(resourceName, name)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(`
%[1]s

%[2]s

resource "davinci_flow" "%[3]s" {
  environment_id = pingone_environment.%[3]s.id

  flow_json = <<EOT
  {
	"companyId": "942b4724-d83d-418c-966c-ed7d352a985c",
	"authTokenExpireIds": [],
	"connectorIds": [
	  "errorConnector",
	  "variablesConnector",
	  "httpConnector"
	],
	"createdDate": 1717075715105,
	"currentVersion": 4,
	"customerId": "db5f4450b2bd8a56ce076dec0c358a9a",
	"deployedDate": 1717075718651,
	"description": "Imported on Wed Jan 31 2024 13:29:13 GMT+0000 (Coordinated Universal Time)",
	"flowStatus": "enabled",
	"isOutputSchemaSaved": false,
	"name": "simple1",
	"publishedVersion": 4,
	"settings": {
	  "csp": "worker-src 'self' blob:; script-src 'self' https://cdn.jsdelivr.net https://code.jquery.com https://devsdk.singularkey.com http://cdnjs.cloudflare.com 'unsafe-inline' 'unsafe-eval';",
	  "intermediateLoadingScreenCSS": "",
	  "intermediateLoadingScreenHTML": "",
	  "flowHttpTimeoutInSeconds": 300,
	  "logLevel": 2,
	  "useCustomCSS": true
	},
	"timeouts": "null",
	"updatedDate": 1717075718670,
	"flowId": "0b6f609ba8442a8703aa46ece557ef73",
	"versionId": 4,
	"graphData": {
	  "elements": {
		"nodes": [
		  {
			"data": {
			  "id": "2pzouq7el7",
			  "nodeType": "CONNECTION",
			  "connectionId": "fa497c1ceaea43c0886d8d360874a53d",
			  "connectorId": "errorConnector",
			  "name": "abcd123-error",
			  "label": "Error Message",
			  "status": "configured",
			  "capabilityName": "customErrorMessage",
			  "type": "action",
			  "properties": {
				"errorMessage": {
				  "value": "[\n  {\n    \"children\": [\n      {\n        \"text\": \"This is an error - ${davinci_connection.%[3]s-error.id}\"\n      }\n    ]\n  }\n]"
				},
				"errorDescription": {
				  "value": "[\n  {\n    \"children\": [\n      {\n        \"text\": \"This is an error, really\"\n      }\n    ]\n  }\n]"
				}
			  }
			},
			"position": {
			  "x": 277,
			  "y": 236
			},
			"group": "nodes",
			"removed": false,
			"selected": false,
			"selectable": true,
			"locked": false,
			"grabbable": true,
			"pannable": false,
			"classes": ""
		  },
		  {
			"data": {
			  "id": "66el2f9ywr",
			  "nodeType": "CONNECTION",
			  "connectionId": "9f8f97e94ad87e184960633b424d80b6",
			  "connectorId": "variablesConnector",
			  "name": "abcd123-variables",
			  "label": "abcd123-variables",
			  "status": "configured",
			  "capabilityName": "saveValue",
			  "type": "trigger",
			  "properties": {
				"saveVariables": {
				  "value": [
					{
					  "name": "flowInstanceVariable1",
					  "value": "[\n  {\n    \"children\": [\n      {\n        \"text\": \"5\"\n      }\n    ]\n  }\n]",
					  "key": 0.8117547669856684,
					  "type": "string",
					  "label": "flowInstanceVariable1 (string - flowInstance)"
					}
				  ]
				}
			  }
			},
			"position": {
			  "x": 400,
			  "y": 400
			},
			"group": "nodes",
			"removed": false,
			"selected": false,
			"selectable": true,
			"locked": false,
			"grabbable": true,
			"pannable": false,
			"classes": ""
		  },
		  {
			"data": {
			  "id": "noqw2jmqk5",
			  "nodeType": "CONNECTION",
			  "connectionId": "9f8f97e94ad87e184960633b424d80b6",
			  "connectorId": "variablesConnector",
			  "name": "abcd123-variables",
			  "label": "abcd123-variables",
			  "status": "configured",
			  "capabilityName": "saveFlowValue",
			  "type": "trigger",
			  "properties": {
				"saveFlowVariables": {
				  "value": [
					{
					  "name": "flowVariable123",
					  "value": "[\n  {\n    \"children\": [\n      {\n        \"text\": \"30\"\n      }\n    ]\n  }\n]",
					  "key": 0.9954407178532643,
					  "label": "flowVariable123 (string - flow)",
					  "type": "string"
					}
				  ]
				}
			  }
			},
			"position": {
			  "x": 390,
			  "y": 540
			},
			"group": "nodes",
			"removed": false,
			"selected": false,
			"selectable": true,
			"locked": false,
			"grabbable": true,
			"pannable": false,
			"classes": ""
		  },
		  {
			"data": {
			  "id": "5wdy7ttpyp",
			  "nodeType": "CONNECTION",
			  "connectionId": "9cb5e3fdbbf0eeb602e0ff332ad79e5d",
			  "connectorId": "httpConnector",
			  "name": "abcd123-http",
			  "label": "abcd123-http",
			  "status": "configured",
			  "capabilityName": "makeRestApiCall",
			  "type": "trigger",
			  "properties": {
				"url": {
				  "value": "[\n  {\n    \"children\": [\n      {\n        \"text\": \"https://pingidentity.com\"\n      }\n    ]\n  }\n]"
				},
				"headers": {
				  "value": [
					{
					  "key": "stub-header",
					  "value": "[\n  {\n    \"children\": [\n      {\n        \"text\": \"\"\n      },\n      {\n        \"text\": \"\"\n      },\n      {\n        \"type\": \"link\",\n        \"src\": \"variable.svg\",\n        \"url\": \"testVariable\",\n        \"data\": \"{{global.company.variables.testVariable}}\",\n        \"tooltip\": \"{{global.company.variables.testVariable}}\",\n        \"children\": [\n          {\n            \"text\": \"testVariable\"\n          }\n        ]\n      },\n      {\n        \"text\": \"\"\n      }\n    ]\n  }\n]"
					}
				  ]
				}
			  }
			},
			"position": {
			  "x": 390,
			  "y": 660
			},
			"group": "nodes",
			"removed": false,
			"selected": false,
			"selectable": true,
			"locked": false,
			"grabbable": true,
			"pannable": false,
			"classes": ""
		  }
		]
	  },
	  "data": {},
	  "zoomingEnabled": true,
	  "userZoomingEnabled": true,
	  "zoom": 1,
	  "minZoom": 1e-50,
	  "maxZoom": 1e+50,
	  "panningEnabled": true,
	  "userPanningEnabled": true,
	  "pan": {
		"x": 0,
		"y": 0
	  },
	  "boxSelectionEnabled": true,
	  "renderer": {
		"name": "null"
	  }
	},
	"flowColor": "#FFC8C1",
	"savedDate": 1717075715076,
	"variables": [
	  {
		"context": "company",
		"createdDate": 1716908009814,
		"fields": {
		  "type": "string",
		  "displayName": "Test Variable",
		  "value": "",
		  "min": 0,
		  "max": 2000
		},
		"id": "a8f1f4d0-2b20-462f-ac02-5f9331747dce",
		"type": "property",
		"visibility": "private",
		"name": "testVariable##SK##company",
		"companyId": "942b4724-d83d-418c-966c-ed7d352a985c"
	  },
	  {
		"context": "flowInstance",
		"createdDate": 1717075626055,
		"customerId": "db5f4450b2bd8a56ce076dec0c358a9a",
		"fields": {
		  "type": "string",
		  "displayName": "",
		  "value": "50",
		  "mutable": true,
		  "min": 0,
		  "max": 2000
		},
		"id": "0b7f34a6-4a22-4e2c-b580-5b27f2508665",
		"type": "property",
		"visibility": "private",
		"name": "flowInstanceVariable1##SK##flowInstance",
		"companyId": "942b4724-d83d-418c-966c-ed7d352a985c"
	  },
	  {
		"context": "flow",
		"createdDate": 1717075655455,
		"customerId": "db5f4450b2bd8a56ce076dec0c358a9a",
		"fields": {
		  "type": "string",
		  "displayName": "",
		  "value": "20",
		  "mutable": true,
		  "min": 0,
		  "max": 2000
		},
		"flowId": "0b6f609ba8442a8703aa46ece557ef73",
		"id": "91a17122-a8ab-49a0-8ff2-43515f975956",
		"type": "property",
		"visibility": "private",
		"name": "flowVariable123##SK##flow##SK##0b6f609ba8442a8703aa46ece557ef73",
		"companyId": "942b4724-d83d-418c-966c-ed7d352a985c"
	  }
	],
	"connections": []
  }
EOT

  // Error connector
  connection_link {
    id                           = davinci_connection.%[3]s-error.id
    name                         = davinci_connection.%[3]s-error.name
    replace_import_connection_id = "53ab83a4a4ab919d9f2cb02d9e111ac8"
  }
}`, acctest.PingoneEnvironmentSsoHcl(resourceName, false), commonHcl, resourceName), nil
}

// // tests for changes other than graph data.
// func TestAccResourceFlow_SchemaChanges(t *testing.T) {

// 	resourceBase := "davinci_flow"
// 	resourceName := acctest.ResourceNameGen()
// 	testFlows := acctest.FlowsForTests(resourceName)
// 	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.Simple.Name)

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck: func() {
// 			acctest.PreCheckClient(t)
// 			acctest.PreCheckNewEnvironment(t)
// 		},
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		ExternalProviders: acctest.ExternalProviders,
// 		ErrorCheck:        acctest.ErrorCheck(t),
// 		CheckDestroy:      davinci.Flow_CheckDestroy(),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.Simple.Hcl}),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
// 					resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
// 					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
// 				),
// 			},
// 			//change flow_json.settings only
// 			{
// 				Config: testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.SimpleSettingDrift.Hcl}),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
// 					resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
// 					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
// 					testAccCheckAttributeSimpleFlowSetting(resourceFullName),
// 				),
// 			},
// 			//revert to simple. This likely is not actually reverting because
// 			{
// 				Config: testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.Simple.Hcl}),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
// 					resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
// 					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
// 				),
// 			},
// 			//change flow_json.inputSchema only
// 			{
// 				Config: testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.SimpleInputSchemaDrift.Hcl}),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
// 					resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
// 					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
// 					testAccCheckAttributeSimpleFlowInputSchema(resourceFullName),
// 				),
// 			},
// 			//revert to simple
// 			{
// 				Config: testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.Simple.Hcl}),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
// 					resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
// 					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
// 				),
// 			},
// 			//change flow_json.outputSchema only
// 			{
// 				Config: testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.SimpleOutputSchemaDrift.Hcl}),
// 				Check: resource.ComposeAggregateTestCheckFunc(
// 					resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
// 					resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
// 					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
// 					testAccCheckAttributeSimpleFlowOutputSchema(resourceFullName),
// 				),
// 			},
// 			//revert to simple
// 			{
// 				Config: testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.Simple.Hcl}),
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
// 					resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
// 					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
// 				),
// 			},
// 		},
// 	})
// }

// func testAccCheckAttributeSimpleFlowSetting(resourceFullName string) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		fj, err := acctest.GetAttributeFromState(s, resourceFullName, "flow_json")
// 		if err != nil {
// 			return err
// 		}

// 		var flow dv.Flow
// 		err = json.Unmarshal([]byte(fj), &flow)
// 		if err != nil {
// 			return err
// 		}
// 		flowSettingsMap := flow.Settings.(map[string]interface{})

// 		flowHttpTimeoutInSeconds := acctest.SchemaAttributeFloat64{
// 			AttributeName: "flowHttpTimeoutInSeconds",
// 			ExpectedValue: 300,
// 			ActualValue:   flowSettingsMap["flowHttpTimeoutInSeconds"].(float64),
// 		}

// 		csp := acctest.SchemaAttributeString{
// 			AttributeName: "csp",
// 			ExpectedValue: "worker-src 'self' blob:; script-src 'self' https://cdn.jsdelivr.net https://code.jquery.com https://devsdk.singularkey.com http://cdnjs.cloudflare.com 'unsafe-inline' 'unsafe-eval';",
// 			ActualValue:   flowSettingsMap["csp"].(string),
// 		}

// 		return acctest.ComposeCompare(
// 			flowHttpTimeoutInSeconds.Compare(),
// 			csp.Compare(),
// 		)
// 	}
// }

// func testAccCheckAttributeSimpleFlowInputSchema(resourceFullName string) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		fj, err := acctest.GetAttributeFromState(s, resourceFullName, "flow_json")
// 		if err != nil {
// 			return err
// 		}

// 		var flow dv.Flow
// 		err = json.Unmarshal([]byte(fj), &flow)
// 		if err != nil {
// 			return err
// 		}
// 		//sample inputSchemaJson: {  "isInputSchemaSaved": true,
// 		// "inputSchemaCompiled": {
// 		//   "parameters": {
// 		//     "type": "object",
// 		//     "properties": {
// 		//       "foo": {
// 		//         "description": "fooDesc",
// 		//         "preferredDataType": "string",
// 		//         "isExpanded": true,
// 		//         "type": "string",
// 		//         "name": "foo"
// 		//       }
// 		//     },
// 		//     "additionalProperties": false,
// 		//     "required": [
// 		//       "foo"
// 		//     ]
// 		//   }
// 		// },
// 		// "inputSchema": [
// 		//   {
// 		//     "propertyName": "foo",
// 		//     "description": "fooDesc",
// 		//     "preferredDataType": "string",
// 		//     "preferredControlType": "textField",
// 		//     "isExpanded": true,
// 		//     "required": true
// 		//   }
// 		// ],}

// 		fsMap := flow.InputSchema[0].(map[string]interface{})
// 		inputSchemaPropName := acctest.SchemaAttributeString{
// 			AttributeName: "Input Schema propertyName'foo'",
// 			ExpectedValue: "foo",
// 			ActualValue:   fsMap["propertyName"].(string),
// 		}

// 		isInputSchemaSaved := acctest.SchemaAttributeBoolean{
// 			AttributeName: "isInputSchemaSaved",
// 			ExpectedValue: true,
// 			ActualValue:   flow.IsInputSchemaSaved,
// 		}

// 		var inputSchemaCompiledPropFooDesc string
// 		inputSchemaCompiledMap, ok := flow.InputSchemaCompiled.(map[string]interface{})
// 		if ok {
// 			inputSchemaCompiledMap, ok = inputSchemaCompiledMap["parameters"].(map[string]interface{})
// 			if ok {
// 				inputSchemaCompiledMap, ok = inputSchemaCompiledMap["properties"].(map[string]interface{})
// 				if ok {
// 					inputSchemaCompiledMap, ok = inputSchemaCompiledMap["foo"].(map[string]interface{})
// 					if ok {
// 						inputSchemaCompiledPropFooDesc = inputSchemaCompiledMap["description"].(string)
// 					}
// 				}
// 			}
// 		}

// 		inputSchemaCompiled := acctest.SchemaAttributeString{
// 			AttributeName: "inputSchemaCompiled",
// 			ExpectedValue: "fooDesc",
// 			ActualValue:   inputSchemaCompiledPropFooDesc,
// 		}

// 		return acctest.ComposeCompare(
// 			inputSchemaPropName.Compare(),
// 			isInputSchemaSaved.Compare(),
// 			inputSchemaCompiled.Compare(),
// 		)
// 	}
// }

// func testAccCheckAttributeSimpleFlowOutputSchema(resourceFullName string) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		fj, err := acctest.GetAttributeFromState(s, resourceFullName, "flow_json")
// 		if err != nil {
// 			return err
// 		}

// 		var flow dv.Flow
// 		err = json.Unmarshal([]byte(fj), &flow)
// 		if err != nil {
// 			return err
// 		}
// 		//sample inputSchemaJson: {
// 		// "outputSchemaCompiled": {
// 		//     "output": {
// 		//       "type": "object",
// 		//       "additionalProperties": true,
// 		//       "properties": {
// 		//         "far": {
// 		//           "type": "string"
// 		//         }
// 		//       }
// 		//     }
// 		//   }
// 		// }
// 		var outputMap map[string]interface{}
// 		if flow.OutputSchemaCompiled.Output != nil {
// 			outputMap = flow.OutputSchemaCompiled.Output.(map[string]interface{})
// 			outputPropertiesMap := outputMap["properties"].(map[string]interface{})

// 			outputSchemaCompiled := acctest.SchemaAttributeMap{
// 				AttributeName: "outputSchemaCompiled",
// 				ExpectedValue: map[string]interface{}{"type": "string"},
// 				ActualValue:   outputPropertiesMap["far"].(map[string]interface{}),
// 			}
// 			return outputSchemaCompiled.Compare()
// 		}
// 		return fmt.Errorf("outputSchemaCompiled is nil")
// 	}
// }

// func TestAccResourceFlow_FlowContextVarFlow(t *testing.T) {

// 	resourceBase := "davinci_flow"
// 	resourceName := acctest.ResourceNameGen()
// 	testFlows := acctest.FlowsForTests(resourceName)
// 	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.FlowContextVarFlow.Name)

// 	hcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.FlowContextVarFlow.Hcl})

// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck: func() {
// 			acctest.PreCheckClient(t)
// 			acctest.PreCheckNewEnvironment(t)
// 		},
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		ExternalProviders: acctest.ExternalProviders,
// 		ErrorCheck:        acctest.ErrorCheck(t),
// 		CheckDestroy:      davinci.Flow_CheckDestroy(),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: hcl,
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
// 					resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
// 					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
// 				),
// 			},
// 		},
// 	})
// }

func TestAccResourceFlow_BadParameters(t *testing.T) {

	resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	minimalStepHcl, _, err := testAccResourceFlow_Minimal_WithMappingIDs_HCL(resourceName, name, false)
	if err != nil {
		t.Fatalf("Failed to get HCL: %v", err)
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Flow_CheckDestroy(),
		Steps: []resource.TestStep{
			// Configure
			{
				Config: minimalStepHcl,
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
