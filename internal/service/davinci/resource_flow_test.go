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
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)
	flows := acctest.FlowForTests(resourceName)

	hcl := testAccResourceFlow_SimpleFlow_Hcl(resourceName, flows.Simple)

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

func TestAccResourceFlow_FlowUpdate(t *testing.T) {

	resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)
	flows := acctest.FlowForTests(resourceName)

	beforeHcl := testAccResourceFlow_SimpleFlow_Hcl(resourceName, flows.Simple)
	afterHcl := testAccResourceFlow_SimpleFlow_Hcl(resourceName, flows.Drifted)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: beforeHcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "flow_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
				),
			},
			{
				Config: afterHcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "flow_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "deploy"),
				),
			},
		},
	})
}

func testAccResourceFlow_SimpleFlow_Hcl(resourceName, flowJson string) (hcl string) {
	baseHcl := acctest.PingoneEnvrionmentSsoHcl(resourceName)
	hcl = fmt.Sprintf(`
%[1]s

data "davinci_connections" "all" {
	connector_ids = ["httpConnector"]
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
}

output "dv_connections" { 
	value = data.davinci_connections.all
}

resource "davinci_flow" "%[2]s" {
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
	flow_json = %[3]s
	deploy = true
	depends_on = [data.davinci_connections.all]
}
`, baseHcl, resourceName, flowJson)
	return hcl
}

func TestAccResourceFlow_SubFlows(t *testing.T) {

	resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	resourceMainflowFullName := fmt.Sprintf("%s.%s", resourceBase, "mainflow-"+resourceName)
	resourceSubflowFullName := fmt.Sprintf("%s.%s", resourceBase, "subflow-"+resourceName)
	resourceAnotherMainflowFullName := fmt.Sprintf("%s.%s", resourceBase, "anothermainflow-"+resourceName)
	resourceAnotherSubflowFullName := fmt.Sprintf("%s.%s", resourceBase, "anothersubflow-"+resourceName)
	// flows := acctest.FlowForTests(resourceName)

	hcl := testAccResourceFlow_SubFlows_Hcl(resourceName)
	hclDrifted := testAccResourceFlow_SubFlowsDrift_Hcl(resourceName)

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

func testAccResourceFlow_SubFlows_Hcl(resourceName string) (hcl string) {
	baseHcl := acctest.PingoneEnvrionmentSsoHcl(resourceName)
	flows := acctest.FlowForTests(resourceName)
	hcl = fmt.Sprintf(`
%[1]s

data "davinci_connections" "all" {
	connector_ids = ["httpConnector"]
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
}

resource "davinci_connection" "subflow" {
	name = "Flow"
	connector_id = "flowConnector"
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
}

resource "davinci_flow" "mainflow-%[2]s" {
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
	flow_json = %[3]s
	deploy = true
	subflows {
		subflow_id = resource.davinci_flow.subflow-%[2]s.flow_id
		subflow_name = resource.davinci_flow.subflow-%[2]s.flow_name
	}
	subflows {
		subflow_id = resource.davinci_flow.anothersubflow-%[2]s.flow_id
		subflow_name = resource.davinci_flow.anothersubflow-%[2]s.flow_name
	}
	depends_on = [data.davinci_connections.all, davinci_connection.subflow]
}

resource "davinci_flow" "anothermainflow-%[2]s" {
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
	flow_json = %[5]s
	deploy = true
	subflows {
		subflow_id = resource.davinci_flow.subflow-%[2]s.flow_id
		subflow_name = resource.davinci_flow.subflow-%[2]s.flow_name
	}
	depends_on = [data.davinci_connections.all, davinci_connection.subflow]
}

resource "davinci_flow" "subflow-%[2]s" {
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
	flow_json = %[4]s
	deploy = true
	subflows {
		subflow_id = resource.davinci_flow.anothersubflow-%[2]s.flow_id
		subflow_name = resource.davinci_flow.anothersubflow-%[2]s.flow_name
	}
	depends_on = [data.davinci_connections.all, davinci_connection.subflow]
}

resource "davinci_flow" "anothersubflow-%[2]s" {
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
	flow_json = %[6]s
	deploy = true
	depends_on = [data.davinci_connections.all, davinci_connection.subflow]
}
`, baseHcl, resourceName, flows.MainFlow, flows.Subflow, flows.AnotherMainflow, flows.AnotherSubflow)
	return hcl
}

// TODO: Tech debt: This test is not working as expected. It is not detecting the drift.
func testAccResourceFlow_SubFlowsDrift_Hcl(resourceName string) (hcl string) {
	baseHcl := acctest.PingoneEnvrionmentSsoHcl(resourceName)
	flows := acctest.FlowForTests(resourceName)
	hcl = fmt.Sprintf(`
%[1]s

data "davinci_connections" "all" {
	connector_ids = ["httpConnector"]
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
}

resource "davinci_connection" "subflow" {
	name = "Flow"
	connector_id = "flowConnector"
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
}

resource "davinci_flow" "mainflow-%[2]s" {
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
	flow_json = %[3]s
	deploy = true
	subflows {
		subflow_id = resource.davinci_flow.subflow-%[2]s.flow_id
		subflow_name = resource.davinci_flow.subflow-%[2]s.flow_name
	}
	subflows {
		subflow_id = resource.davinci_flow.anothersubflow-%[2]s.flow_id
		subflow_name = resource.davinci_flow.anothersubflow-%[2]s.flow_name
	}
	depends_on = [data.davinci_connections.all, davinci_connection.subflow]
}

resource "davinci_flow" "anothermainflow-%[2]s" {
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
	flow_json = %[5]s
	deploy = true
	subflows {
		subflow_id = resource.davinci_flow.subflow-%[2]s.flow_id
		subflow_name = resource.davinci_flow.subflow-%[2]s.flow_name
	}
	depends_on = [data.davinci_connections.all, davinci_connection.subflow]
}

resource "davinci_flow" "subflow-%[2]s" {
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
	flow_json = %[4]s
	deploy = true
	subflows {
		subflow_id = resource.davinci_flow.anothersubflow-%[2]s.flow_id
		subflow_name = resource.davinci_flow.anothersubflow-%[2]s.flow_name
	}
	depends_on = [data.davinci_connections.all, davinci_connection.subflow]
}

resource "davinci_flow" "anothersubflow-%[2]s" {
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
	flow_json = %[6]s
	deploy = true
	depends_on = [data.davinci_connections.all, davinci_connection.subflow]
}
`, baseHcl, resourceName, flows.MainFlowDrifted, flows.Subflow, flows.AnotherMainflow, flows.AnotherSubflow)
	return hcl
}
