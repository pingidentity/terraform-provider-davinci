package davinci_test

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
	"github.com/samir-gandhi/davinci-client-go/davinci"
)

func TestAccResourceFlow_SimpleFlow(t *testing.T) {

	resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	testFlows := acctest.FlowsForTests(resourceName)
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.Simple.Name)

	hcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.Simple.Hcl})

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
				),
			},
		},
	})
}

func TestAccResourceFlow_SimpleFlowUpdate(t *testing.T) {

	resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	testFlows := acctest.FlowsForTests(resourceName)
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.Simple.Name)

	hcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.Simple.Hcl})

	hclDrifted := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.Drifted.Hcl})

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
				),
			},
			{
				Config: hclDrifted,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
				),
			},
		},
	})
}

func TestAccResourceFlow_SubFlows(t *testing.T) {

	resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	testFlows := acctest.FlowsForTests(resourceName)
	resourceMainflowFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.Mainflow.Name)
	resourceSubflowFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.Subflow.Name)
	resourceAnotherMainflowFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.AnotherMainflow.Name)
	resourceAnotherSubflowFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.AnotherSubflow.Name)

	hcl := testAccResourceFlow_SubFlows_Hcl(resourceName, []string{testFlows.Mainflow.Hcl, testFlows.Subflow.Hcl, testFlows.AnotherMainflow.Hcl, testFlows.AnotherSubflow.Hcl})
	hclDrifted := testAccResourceFlow_SubFlows_Hcl(resourceName, []string{testFlows.MainflowDrifted.Hcl, testFlows.Subflow.Hcl, testFlows.AnotherMainflow.Hcl, testFlows.AnotherSubflow.Hcl})

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceMainflowFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceMainflowFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceMainflowFullName, "deploy"),
					resource.TestCheckResourceAttrSet(resourceSubflowFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceSubflowFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceSubflowFullName, "deploy"),
					resource.TestCheckResourceAttrSet(resourceAnotherMainflowFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceAnotherMainflowFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceAnotherMainflowFullName, "deploy"),
					resource.TestCheckResourceAttrSet(resourceAnotherSubflowFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceAnotherSubflowFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceAnotherSubflowFullName, "deploy"),
				),
			},
			{
				Config: hclDrifted,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceMainflowFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceMainflowFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceMainflowFullName, "deploy"),
					resource.TestCheckResourceAttrSet(resourceSubflowFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceSubflowFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceSubflowFullName, "deploy"),
					resource.TestCheckResourceAttrSet(resourceAnotherMainflowFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceAnotherMainflowFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceAnotherMainflowFullName, "deploy"),
					resource.TestCheckResourceAttrSet(resourceAnotherSubflowFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceAnotherSubflowFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceAnotherSubflowFullName, "deploy"),
				),
			},
		},
	})
}

func testAccResourceFlow_SubFlows_Hcl(resourceName string, flowsHcl []string) (hcl string) {
	baseHcl := acctest.BaselineHcl(resourceName)
	hcl = fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s-flow" {
	name = "Flow"
	connector_id = "flowConnector"
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
	depends_on = [data.davinci_connections.read_all]
}

`, baseHcl, resourceName)

	for _, flowHcl := range flowsHcl {
		hcl += flowHcl
	}
	return hcl
}

func testAccResourceFlow_SimpleFlows_Hcl(resourceName string, flowsHcl []string) (hcl string) {
	baseHcl := acctest.BaselineHcl(resourceName)
	hcl = fmt.Sprintf(`
%[1]s

`, baseHcl)

	for _, flowHcl := range flowsHcl {
		hcl += flowHcl
	}
	return hcl
}

func testAccResourceFlow_VariableConnectorFlows_Hcl(resourceName string, flowsHcl []acctest.FlowHcl) (hcl string) {
	baseHcl := acctest.BaselineHcl(resourceName)
	hcl = fmt.Sprintf(`
%[1]s

data "davinci_connections" "variables" {
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
	depends_on = [resource.davinci_flow.%[3]s]
	connector_ids = ["variablesConnector"]
}

`, baseHcl, resourceName, flowsHcl[0].Name)

	for _, flowHcl := range flowsHcl {
		hcl += flowHcl.Hcl
	}
	return hcl
}

func TestAccResourceFlow_VariableConnectorFlow(t *testing.T) {

	resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	testFlows := acctest.FlowsForTests(resourceName)
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.WithVariableConnector.Name)

	hcl := testAccResourceFlow_VariableConnectorFlows_Hcl(resourceName, []acctest.FlowHcl{testFlows.WithVariableConnector})
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
					resource.TestCheckResourceAttrWith("data.davinci_connections.variables", "connections.#", func(value string) error {
						v, err := strconv.Atoi(value)
						if err != nil {
							return err
						}
						if v > 1 {
							return fmt.Errorf("Flow Import created additional variables connection")
						}
						return nil
					}),
					resource.TestCheckResourceAttrWith("data.davinci_connections.read_all", "connections.#", func(value string) error {
						v, err := strconv.Atoi(value)
						if err != nil {
							return err
						}
						if v != 9 {
							return fmt.Errorf("Bootstrap not completed, or has changed. Expected 9 connections, got %d", v)
						}
						return nil
					}),
				),
			},
		},
	})
}

func TestAccResourceFlow_BrokenFlow(t *testing.T) {

	// resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	testFlows := acctest.FlowsForTests(resourceName)
	// resourceFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.BrokenFlow.Name)

	hcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.BrokenFlow.Hcl})
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config:      hcl,
				ExpectError: regexp.MustCompile(`Error: status: 400`),
			},
		},
	})
}

// tests for changes other than graph data.
func TestAccResourceFlow_SchemaChanges(t *testing.T) {

	resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	testFlows := acctest.FlowsForTests(resourceName)
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.Simple.Name)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.Simple.Hcl}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
				),
			},
			//change flow_json.settings only
			{
				Config: testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.SimpleSettingDrift.Hcl}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
					testAccCheckAttributeSimpleFlowSetting(resourceFullName),
				),
			},
			//revert to simple. This likely is not actually reverting because
			{
				Config: testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.Simple.Hcl}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
				),
			},
			//change flow_json.inputSchema only
			{
				Config: testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.SimpleInputSchemaDrift.Hcl}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
					testAccCheckAttributeSimpleFlowInputSchema(resourceFullName),
				),
			},
			//revert to simple
			{
				Config: testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.Simple.Hcl}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
				),
			},
			//change flow_json.outputSchema only
			{
				Config: testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.SimpleOutputSchemaDrift.Hcl}),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
					testAccCheckAttributeSimpleFlowOutputSchema(resourceFullName),
				),
			},
			//revert to simple
			{
				Config: testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.Simple.Hcl}),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
				),
			},
		},
	})
}

func testAccCheckAttributeSimpleFlowSetting(resourceFullName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		fj, err := acctest.GetAttributeFromState(s, resourceFullName, "flow_json")
		if err != nil {
			return err
		}

		var flow davinci.Flow
		err = json.Unmarshal([]byte(fj), &flow)
		if err != nil {
			return err
		}
		flowSettingsMap := flow.Settings.(map[string]interface{})

		flowHttpTimeoutInSeconds := acctest.SchemaAttributeFloat64{
			AttributeName: "flowHttpTimeoutInSeconds",
			ExpectedValue: 300,
			ActualValue:   flowSettingsMap["flowHttpTimeoutInSeconds"].(float64),
		}

		csp := acctest.SchemaAttributeString{
			AttributeName: "csp",
			ExpectedValue: "worker-src 'self' blob:; script-src 'self' https://cdn.jsdelivr.net https://code.jquery.com https://devsdk.singularkey.com http://cdnjs.cloudflare.com 'unsafe-inline' 'unsafe-eval';",
			ActualValue:   flowSettingsMap["csp"].(string),
		}

		return acctest.ComposeCompare(
			flowHttpTimeoutInSeconds.Compare(),
			csp.Compare(),
		)
	}
}

func testAccCheckAttributeSimpleFlowInputSchema(resourceFullName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		fj, err := acctest.GetAttributeFromState(s, resourceFullName, "flow_json")
		if err != nil {
			return err
		}

		var flow davinci.Flow
		err = json.Unmarshal([]byte(fj), &flow)
		if err != nil {
			return err
		}
		//sample inputSchemaJson: {  "isInputSchemaSaved": true,
		// "inputSchemaCompiled": {
		//   "parameters": {
		//     "type": "object",
		//     "properties": {
		//       "foo": {
		//         "description": "fooDesc",
		//         "preferredDataType": "string",
		//         "isExpanded": true,
		//         "type": "string",
		//         "name": "foo"
		//       }
		//     },
		//     "additionalProperties": false,
		//     "required": [
		//       "foo"
		//     ]
		//   }
		// },
		// "inputSchema": [
		//   {
		//     "propertyName": "foo",
		//     "description": "fooDesc",
		//     "preferredDataType": "string",
		//     "preferredControlType": "textField",
		//     "isExpanded": true,
		//     "required": true
		//   }
		// ],}

		fsMap := flow.InputSchema[0].(map[string]interface{})
		inputSchemaPropName := acctest.SchemaAttributeString{
			AttributeName: "Input Schema propertyName'foo'",
			ExpectedValue: "foo",
			ActualValue:   fsMap["propertyName"].(string),
		}

		isInputSchemaSaved := acctest.SchemaAttributeBoolean{
			AttributeName: "isInputSchemaSaved",
			ExpectedValue: true,
			ActualValue:   flow.IsInputSchemaSaved,
		}

		var inputSchemaCompiledPropFooDesc string
		inputSchemaCompiledMap, ok := flow.InputSchemaCompiled.(map[string]interface{})
		if ok {
			inputSchemaCompiledMap, ok = inputSchemaCompiledMap["parameters"].(map[string]interface{})
			if ok {
				inputSchemaCompiledMap, ok = inputSchemaCompiledMap["properties"].(map[string]interface{})
				if ok {
					inputSchemaCompiledMap, ok = inputSchemaCompiledMap["foo"].(map[string]interface{})
					if ok {
						inputSchemaCompiledPropFooDesc = inputSchemaCompiledMap["description"].(string)
					}
				}
			}
		}

		inputSchemaCompiled := acctest.SchemaAttributeString{
			AttributeName: "inputSchemaCompiled",
			ExpectedValue: "fooDesc",
			ActualValue:   inputSchemaCompiledPropFooDesc,
		}

		return acctest.ComposeCompare(
			inputSchemaPropName.Compare(),
			isInputSchemaSaved.Compare(),
			inputSchemaCompiled.Compare(),
		)
	}
}

func testAccCheckAttributeSimpleFlowOutputSchema(resourceFullName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		fj, err := acctest.GetAttributeFromState(s, resourceFullName, "flow_json")
		if err != nil {
			return err
		}

		var flow davinci.Flow
		err = json.Unmarshal([]byte(fj), &flow)
		if err != nil {
			return err
		}
		//sample inputSchemaJson: {
		// "outputSchemaCompiled": {
		//     "output": {
		//       "type": "object",
		//       "additionalProperties": true,
		//       "properties": {
		//         "far": {
		//           "type": "string"
		//         }
		//       }
		//     }
		//   }
		// }
		var outputMap map[string]interface{}
		if flow.OutputSchemaCompiled.Output != nil {
			outputMap = flow.OutputSchemaCompiled.Output.(map[string]interface{})
			outputPropertiesMap := outputMap["properties"].(map[string]interface{})

			outputSchemaCompiled := acctest.SchemaAttributeMap{
				AttributeName: "outputSchemaCompiled",
				ExpectedValue: map[string]interface{}{"type": "string"},
				ActualValue:   outputPropertiesMap["far"].(map[string]interface{}),
			}
			return outputSchemaCompiled.Compare()
		}
		return fmt.Errorf("outputSchemaCompiled is nil")
	}
}

func TestAccResourceFlow_FlowContextVarFlow(t *testing.T) {

	resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	testFlows := acctest.FlowsForTests(resourceName)
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, testFlows.FlowContextVarFlow.Name)

	hcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{testFlows.FlowContextVarFlow.Hcl})

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
				),
			},
		},
	})
}
