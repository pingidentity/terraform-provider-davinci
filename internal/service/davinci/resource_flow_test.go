package davinci_test

import (
	"fmt"
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
					resource.TestCheckResourceAttrSet(resourceFullName, "flow_id"),
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
					resource.TestCheckResourceAttrSet(resourceFullName, "flow_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
				),
			},
			{
				Config: hclDrifted,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "flow_id"),
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
					resource.TestCheckResourceAttrSet(resourceMainflowFullName, "flow_id"),
					resource.TestCheckResourceAttrSet(resourceMainflowFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceMainflowFullName, "deploy"),
					resource.TestCheckResourceAttrSet(resourceSubflowFullName, "flow_id"),
					resource.TestCheckResourceAttrSet(resourceSubflowFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceSubflowFullName, "deploy"),
					resource.TestCheckResourceAttrSet(resourceAnotherMainflowFullName, "flow_id"),
					resource.TestCheckResourceAttrSet(resourceAnotherMainflowFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceAnotherMainflowFullName, "deploy"),
					resource.TestCheckResourceAttrSet(resourceAnotherSubflowFullName, "flow_id"),
					resource.TestCheckResourceAttrSet(resourceAnotherSubflowFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceAnotherSubflowFullName, "deploy"),
				),
			},
			{
				Config: hclDrifted,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceMainflowFullName, "flow_id"),
					resource.TestCheckResourceAttrSet(resourceMainflowFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceMainflowFullName, "deploy"),
					resource.TestCheckResourceAttrSet(resourceSubflowFullName, "flow_id"),
					resource.TestCheckResourceAttrSet(resourceSubflowFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceSubflowFullName, "deploy"),
					resource.TestCheckResourceAttrSet(resourceAnotherMainflowFullName, "flow_id"),
					resource.TestCheckResourceAttrSet(resourceAnotherMainflowFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceAnotherMainflowFullName, "deploy"),
					resource.TestCheckResourceAttrSet(resourceAnotherSubflowFullName, "flow_id"),
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
