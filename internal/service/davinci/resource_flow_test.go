package davinci_test

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
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
	// fmt.Println(hcl)
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
	// fmt.Println(hcl)
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
