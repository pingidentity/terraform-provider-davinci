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

// func TestAccResourceFlow_RemovalDrift(t *testing.T) {

// 	resourceBase := "davinci_flow"
// 	resourceName := acctest.ResourceNameGen()
// 	testFlows := acctest.FlowsForTests(resourceName)
// 	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.Simple.Name)

// 	hcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.Simple.Hcl})

// 	var flowID, environmentID string

// 	ctx := context.Background()

// 	p1Client, err := acctest.PingOneTestClient(ctx)
// 	if err != nil {
// 		t.Fatalf("Failed to get API client: %v", err)
// 	}

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
// 			// Configure
// 			{
// 				Config: hcl,
// 				Check:  davinci.Flow_GetIDs(resourceFullName, &environmentID, &flowID),
// 			},
// 			// Replan after removal preconfig
// 			{
// 				PreConfig: func() {
// 					davinci.Flow_RemovalDrift_PreConfig(t, environmentID, flowID)
// 				},
// 				RefreshState:       true,
// 				ExpectNonEmptyPlan: true,
// 			},
// 			// Test removal of the environment
// 			{
// 				Config:   hcl,
// 				Check:    davinci.Flow_GetIDs(resourceFullName, &environmentID, &flowID),
// 				SkipFunc: func() (bool, error) { return true, nil },
// 			},
// 			{
// 				PreConfig: func() {
// 					base.Environment_RemovalDrift_PreConfig(ctx, p1Client.API.ManagementAPIClient, t, environmentID)
// 				},
// 				RefreshState:       true,
// 				ExpectNonEmptyPlan: true,
// 				SkipFunc:           func() (bool, error) { return true, nil },
// 			},
// 		},
// 	})
// }

func TestAccResourceFlow_Full_Clean(t *testing.T) {
	testAccResourceFlow_Full(t, false)
}

func TestAccResourceFlow_Full_WithBootstrap(t *testing.T) {
	testAccResourceFlow_Full(t, true)
}

func testAccResourceFlow_Full(t *testing.T, withBootstrapConfig bool) {

	resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	fullStepHcl, fullStepJson, err := testAccResourceFlow_Full_HCL(resourceName, name, withBootstrapConfig)
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
			resource.TestCheckResourceAttr(resourceFullName, "deploy", "true"),
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

	minimalStepHcl, minimalStepJson, err := testAccResourceFlow_Minimal_HCL(resourceName, name, withBootstrapConfig)
	if err != nil {
		t.Fatalf("Failed to get HCL: %v", err)
	}

	minimalStep := resource.TestStep{
		Config: minimalStepHcl,
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", "simple"),
			resource.TestMatchResourceAttr(resourceFullName, "description", regexp.MustCompile(`^$`)),
			resource.TestCheckResourceAttr(resourceFullName, "flow_json", fmt.Sprintf("%s\n", minimalStepJson)),
			resource.TestCheckResourceAttr(resourceFullName, "connection_link.#", "1"),
			resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "connection_link.*", map[string]*regexp.Regexp{
				"id":                           verify.P1DVResourceIDRegexpFullString,
				"replace_import_connection_id": verify.P1DVResourceIDRegexpFullString,
				"name":                         regexp.MustCompile(fmt.Sprintf(`^%s-error$`, name)),
			}),
			resource.TestCheckResourceAttr(resourceFullName, "deploy", "true"),
			resource.TestCheckResourceAttr(resourceFullName, "subflow_link.#", "0"),
			resource.TestCheckResourceAttr(resourceFullName, "flow_variables.#", "0"),
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
			},
		},
	})
}

// func TestAccResourceFlow_SimpleFlowUpdate(t *testing.T) {

// 	resourceBase := "davinci_flow"
// 	resourceName := acctest.ResourceNameGen()
// 	testFlows := acctest.FlowsForTests(resourceName)
// 	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.Simple.Name)

// 	hcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.Simple.Hcl})

// 	hclDrifted := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.Drifted.Hcl})

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
// 			{
// 				Config: hclDrifted,
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
// 					resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
// 					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
// 				),
// 			},
// 		},
// 	})
// }

// func TestAccResourceFlow_SubFlows(t *testing.T) {

// 	resourceBase := "davinci_flow"
// 	resourceName := acctest.ResourceNameGen()
// 	testFlows := acctest.FlowsForTests(resourceName)
// 	resourceMainflowFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.Mainflow.Name)
// 	resourceSubflowFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.Subflow.Name)
// 	resourceAnotherMainflowFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.AnotherMainflow.Name)
// 	resourceAnotherSubflowFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.AnotherSubflow.Name)

// 	hcl := testAccResourceFlow_SubFlows_Hcl(resourceName, []string{testFlows.Mainflow.Hcl, testFlows.Subflow.Hcl, testFlows.AnotherMainflow.Hcl, testFlows.AnotherSubflow.Hcl})
// 	hclDrifted := testAccResourceFlow_SubFlows_Hcl(resourceName, []string{testFlows.MainflowDrifted.Hcl, testFlows.Subflow.Hcl, testFlows.AnotherMainflow.Hcl, testFlows.AnotherSubflow.Hcl})

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
// 					resource.TestCheckResourceAttrSet(resourceMainflowFullName, "id"),
// 					resource.TestCheckResourceAttrSet(resourceMainflowFullName, "environment_id"),
// 					resource.TestCheckResourceAttrSet(resourceMainflowFullName, "deploy"),
// 					resource.TestCheckResourceAttrSet(resourceSubflowFullName, "id"),
// 					resource.TestCheckResourceAttrSet(resourceSubflowFullName, "environment_id"),
// 					resource.TestCheckResourceAttrSet(resourceSubflowFullName, "deploy"),
// 					resource.TestCheckResourceAttrSet(resourceAnotherMainflowFullName, "id"),
// 					resource.TestCheckResourceAttrSet(resourceAnotherMainflowFullName, "environment_id"),
// 					resource.TestCheckResourceAttrSet(resourceAnotherMainflowFullName, "deploy"),
// 					resource.TestCheckResourceAttrSet(resourceAnotherSubflowFullName, "id"),
// 					resource.TestCheckResourceAttrSet(resourceAnotherSubflowFullName, "environment_id"),
// 					resource.TestCheckResourceAttrSet(resourceAnotherSubflowFullName, "deploy"),
// 				),
// 			},
// 			{
// 				Config: hclDrifted,
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestCheckResourceAttrSet(resourceMainflowFullName, "id"),
// 					resource.TestCheckResourceAttrSet(resourceMainflowFullName, "environment_id"),
// 					resource.TestCheckResourceAttrSet(resourceMainflowFullName, "deploy"),
// 					resource.TestCheckResourceAttrSet(resourceSubflowFullName, "id"),
// 					resource.TestCheckResourceAttrSet(resourceSubflowFullName, "environment_id"),
// 					resource.TestCheckResourceAttrSet(resourceSubflowFullName, "deploy"),
// 					resource.TestCheckResourceAttrSet(resourceAnotherMainflowFullName, "id"),
// 					resource.TestCheckResourceAttrSet(resourceAnotherMainflowFullName, "environment_id"),
// 					resource.TestCheckResourceAttrSet(resourceAnotherMainflowFullName, "deploy"),
// 					resource.TestCheckResourceAttrSet(resourceAnotherSubflowFullName, "id"),
// 					resource.TestCheckResourceAttrSet(resourceAnotherSubflowFullName, "environment_id"),
// 					resource.TestCheckResourceAttrSet(resourceAnotherSubflowFullName, "deploy"),
// 				),
// 			},
// 			// Test importing the resource
// 			{
// 				ResourceName: resourceMainflowFullName,
// 				ImportStateIdFunc: func() resource.ImportStateIdFunc {
// 					return func(s *terraform.State) (string, error) {
// 						rs, ok := s.RootModule().Resources[resourceMainflowFullName]
// 						if !ok {
// 							return "", fmt.Errorf("Resource Not found: %s", resourceMainflowFullName)
// 						}

// 						return fmt.Sprintf("%s/%s", rs.Primary.Attributes["environment_id"], rs.Primary.ID), nil
// 					}
// 				}(),
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 				ImportStateVerifyIgnore: []string{
// 					"deploy",
// 				},
// 			},
// 		},
// 	})
// }

// func testAccResourceFlow_SubFlows_Hcl(resourceName string, flowsHcl []string) (hcl string) {
// 	baseHcl := acctest.BaselineHcl(resourceName)
// 	hcl = fmt.Sprintf(`
// %[1]s

// resource "davinci_connection" "%[2]s-flow" {
//   name           = "Flow"
//   connector_id   = "flowConnector"
//   environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
// }

// `, baseHcl, resourceName)

// 	for _, flowHcl := range flowsHcl {
// 		hcl += flowHcl
// 	}
// 	return hcl
// }

func testAccResourceFlow_Common_HCL(resourceName, name string) (hcl string, err error) {

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

func testAccResourceFlow_Full_HCL(resourceName, name string, withBootstrapConfig bool) (hcl, mainFlowJson string, err error) {

	mainFlowJson, err = acctest.ReadFlowJsonFile("flows/full-basic.json")
	if err != nil {
		return "", "", err
	}

	commonHcl, err := testAccResourceFlow_Common_HCL(resourceName, name)

	return fmt.Sprintf(`
%[1]s

%[2]s

resource "davinci_flow" "%[3]s" {
  environment_id = pingone_environment.%[3]s.id

  name = "my awesome flow"
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
    replace_import_connection_id = "33329a264e268ab31fb19637debf1ea3"
  }

  // Error connector
  connection_link {
    id                           = davinci_connection.%[3]s-error.id
    name                         = davinci_connection.%[3]s-error.name
    replace_import_connection_id = "53ab83a4a4ab919d9f2cb02d9e111ac8"
  }

  // Subflow 2
  subflow_link {
    id   = davinci_flow.%[3]s-subflow-2.id
    name = davinci_flow.%[3]s-subflow-2.name
	replace_import_subflow_id = "07503fed5c02849dbbd5ee932da654b2"
  }

  // Subflow 1
  subflow_link {
    id   = davinci_flow.%[3]s-subflow-1.id
    name = davinci_flow.%[3]s-subflow-1.name
	replace_import_subflow_id = "00f66e8926ced6ef5b83619fde4a314a"
  }
}`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), commonHcl, resourceName, mainFlowJson), mainFlowJson, nil
}

func testAccResourceFlow_Minimal_HCL(resourceName, name string, withBootstrapConfig bool) (hcl, mainFlowJson string, err error) {

	mainFlowJson, err = acctest.ReadFlowJsonFile("flows/full-minimal.json")
	if err != nil {
		return "", "", err
	}

	commonHcl, err := testAccResourceFlow_Common_HCL(resourceName, name)

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

// func testAccResourceFlow_SimpleFlows_Hcl(resourceName string, flowsHcl []string) (hcl string) {
// 	baseHcl := acctest.BaselineHcl(resourceName)
// 	hcl = fmt.Sprintf(`
// %[1]s

// `, baseHcl)

// 	for _, flowHcl := range flowsHcl {
// 		hcl += flowHcl
// 	}
// 	return hcl
// }

// func testAccResourceFlow_VariableConnectorFlows_Hcl(resourceName string, flowsHcl []acctest.FlowHcl) (hcl string) {
// 	baseHcl := acctest.BaselineHcl(resourceName)
// 	hcl = fmt.Sprintf(`
// %[1]s

// data "davinci_connections" "variables" {
//   environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
//   depends_on     = [resource.davinci_flow.%[3]s]
//   connector_ids  = ["variablesConnector"]
// }

// `, baseHcl, resourceName, flowsHcl[0].Name)

// 	for _, flowHcl := range flowsHcl {
// 		hcl += flowHcl.Hcl
// 	}
// 	return hcl
// }

// func TestAccResourceFlow_VariableConnectorFlow(t *testing.T) {

// 	resourceBase := "davinci_flow"
// 	resourceName := acctest.ResourceNameGen()
// 	testFlows := acctest.FlowsForTests(resourceName)
// 	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.WithVariableConnector.Name)

// 	hcl := testAccResourceFlow_VariableConnectorFlows_Hcl(resourceName, []acctest.FlowHcl{testFlows.WithVariableConnector})
// 	resource.ParallelTest(t, resource.TestCase{
// 		PreCheck: func() {
// 			acctest.PreCheckClient(t)
// 			acctest.PreCheckNewEnvironment(t)
// 		},
// 		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
// 		ExternalProviders: acctest.ExternalProviders,
// 		CheckDestroy:      davinci.Flow_CheckDestroy(),
// 		ErrorCheck:        acctest.ErrorCheck(t),
// 		Steps: []resource.TestStep{
// 			{
// 				Config: hcl,
// 				Check: resource.ComposeTestCheckFunc(
// 					resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
// 					resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
// 					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
// 					resource.TestCheckResourceAttrWith("data.davinci_connections.variables", "connections.#", func(value string) error {
// 						v, err := strconv.Atoi(value)
// 						if err != nil {
// 							return err
// 						}
// 						if v > 1 {
// 							return fmt.Errorf("Flow Import created additional variables connection")
// 						}
// 						return nil
// 					}),
// 					resource.TestCheckResourceAttrWith("data.davinci_connections.read_all", "connections.#", func(value string) error {
// 						v, err := strconv.Atoi(value)
// 						if err != nil {
// 							return err
// 						}
// 						if v != acctest.BsConnectionsCount {
// 							return fmt.Errorf("Bootstrap not completed, or has changed. Expected %d connections, got %d", acctest.BsConnectionsCount, v)
// 						}
// 						return nil
// 					}),
// 				),
// 			},
// 			// Test importing the resource
// 			{
// 				ResourceName: resourceFullName,
// 				ImportStateIdFunc: func() resource.ImportStateIdFunc {
// 					return func(s *terraform.State) (string, error) {
// 						rs, ok := s.RootModule().Resources[resourceFullName]
// 						if !ok {
// 							return "", fmt.Errorf("Resource Not found: %s", resourceFullName)
// 						}

// 						return fmt.Sprintf("%s/%s", rs.Primary.Attributes["environment_id"], rs.Primary.ID), nil
// 					}
// 				}(),
// 				ImportState:       true,
// 				ImportStateVerify: true,
// 				ImportStateVerifyIgnore: []string{
// 					"deploy",
// 				},
// 			},
// 		},
// 	})
// }

// func TestAccResourceFlow_BrokenFlow(t *testing.T) {

// 	// resourceBase := "davinci_flow"
// 	resourceName := acctest.ResourceNameGen()
// 	testFlows := acctest.FlowsForTests(resourceName)
// 	// resourceFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.BrokenFlow.Name)

// 	hcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.BrokenFlow.Hcl})
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
// 				Config:      hcl,
// 				ExpectError: regexp.MustCompile(`Error: status: 400`),
// 			},
// 		},
// 	})
// }

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

// func TestAccResourceFlow_BadParameters(t *testing.T) {

// 	resourceBase := "davinci_flow"
// 	resourceName := acctest.ResourceNameGen()
// 	testFlows := acctest.FlowsForTests(resourceName)
// 	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.Simple.Name)

// 	hcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.Simple.Hcl})

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
// 			// Configure
// 			{
// 				Config: hcl,
// 			},
// 			// Errors
// 			{
// 				ResourceName: resourceFullName,
// 				ImportState:  true,
// 				ExpectError:  regexp.MustCompile(`Invalid import ID specified \(".*"\).  The ID should be in the format "environment_id/davinci_flow_id" and must match regex: .*`),
// 			},
// 			{
// 				ResourceName:  resourceFullName,
// 				ImportStateId: "/",
// 				ImportState:   true,
// 				ExpectError:   regexp.MustCompile(`Invalid import ID specified \(".*"\).  The ID should be in the format "environment_id/davinci_flow_id" and must match regex: .*`),
// 			},
// 			{
// 				ResourceName:  resourceFullName,
// 				ImportStateId: "badformat/badformat",
// 				ImportState:   true,
// 				ExpectError:   regexp.MustCompile(`Invalid import ID specified \(".*"\).  The ID should be in the format "environment_id/davinci_flow_id" and must match regex: .*`),
// 			},
// 		},
// 	})
// }
