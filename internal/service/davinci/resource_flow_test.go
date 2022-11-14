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
	flow_json = "{\"customerId\":\"dc7918cfa4b50966f8508072c2755c2c\",\"name\":\"tftesting\",\"description\":\"\",\"flowStatus\":\"enabled\",\"createdDate\":1662960509175,\"updatedDate\":1662961640567,\"currentVersion\":0,\"authTokenExpireIds\":[],\"deployedDate\":1662960510106,\"functionConnectionId\":null,\"isOutputSchemaSaved\":false,\"outputSchemaCompiled\":null,\"publishedVersion\":1,\"timeouts\":null,\"flowId\":\"bb45eb4a0e8a5c9d6a21c7ac2d1b3faa\",\"companyId\":\"c431739a-29cd-4d9e-b465-0b37b2c235b1\",\"versionId\":0,\"graphData\":{\"elements\":{\"nodes\":[{\"data\":{\"id\":\"pptape4ac2\",\"nodeType\":\"CONNECTION\",\"connectionId\":\"867ed4363b2bc21c860085ad2baa817d\",\"connectorId\":\"httpConnector\",\"name\":\"Http\",\"label\":\"Http\",\"status\":\"configured\",\"capabilityName\":\"customHtmlMessage\",\"type\":\"trigger\",\"properties\":{\"message\":{\"value\":\"[\\n{\\n\\\"children\\\":[\\n{\\n\\\"text\\\":\\\"hellofoobar\\\"\\n}\\n]\\n}\\n]\"}}},\"position\":{\"x\":570,\"y\":240},\"group\":\"nodes\",\"removed\":false,\"selected\":false,\"selectable\":true,\"locked\":false,\"grabbable\":true,\"pannable\":false,\"classes\":\"\"}]},\"data\":{},\"zoomingEnabled\":true,\"userZoomingEnabled\":true,\"zoom\":1,\"minZoom\":1e-50,\"maxZoom\":1e+50,\"panningEnabled\":true,\"userPanningEnabled\":true,\"pan\":{\"x\":0,\"y\":0},\"boxSelectionEnabled\":true,\"renderer\":{\"name\":\"null\"}},\"flowColor\":\"#AFD5FF\",\"connectorIds\":[\"httpConnector\"],\"savedDate\":1662961640542,\"variables\":[]}"
	deploy = true
	depends_on = [data.davinci_connections.all]
}
`, baseHcl, resourceName)
	return hcl
}

func TestAccResourceFlow_FlowUpdate(t *testing.T) {

	resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)
	flows := acctest.FlowForTests(resourceName)

	beforeHcl := testAccResourceFlow_SimpleFlow_Hcl(resourceName, flows.Simple)
	afterHcl := testAccResourceFlow_SimpleFlow_Hcl(resourceName, flows.Simple)

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
