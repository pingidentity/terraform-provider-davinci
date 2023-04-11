package davinci_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
)

func TestAccResourceConnection_Slim(t *testing.T) {

	resourceBase := "davinci_connection"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	hcl := testAccResourceConnection_Slim_Hcl(resourceName, "slim")

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
					resource.TestCheckResourceAttr(resourceFullName, "name", resourceName),
				),
			},
		},
	})
}

func TestAccResourceConnection_SlimWithUpdate(t *testing.T) {

	resourceBase := "davinci_connection"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	beforeHcl := testAccResourceConnection_Slim_Hcl(resourceName, "before")
	afterHcl := testAccResourceConnection_Slim_Hcl(resourceName, "after")

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
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttr(resourceFullName, "name", resourceName),
				),
			},
			{
				Config: afterHcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttr(resourceFullName, "name", resourceName),
				),
			},
		},
	})
}

func testAccResourceConnection_Slim_Hcl(resourceName, valuePrefix string) (hcl string) {
	baseHcl := acctest.PingoneEnvrionmentSsoHcl(resourceName)
	clientId := acctest.RandStringWithPrefix(valuePrefix)
	clientSecret := acctest.RandStringWithPrefix(valuePrefix)
	hcl = fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s" {
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  depends_on     = [data.davinci_connections.read_all]
  connector_id   = "crowdStrikeConnector"
  name           = "%[2]s"
  property {
    name  = "clientId"
    value = "%[3]s"
  }
  property {
    name  = "clientSecret"
    value = "%[3]s"
  }
}
`, baseHcl, resourceName, clientId, clientSecret)

	return hcl
}

// Test to try to hit API rate Limit
func TestAccReadResources_Heavy(t *testing.T) {

	resourceBase := "davinci_connection"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	hcl := testAccReadResources_Heavy_Hcl(resourceName, "heavy")
	// fmt.Printf(`HCL: \n %s \n`, hcl)

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
					resource.TestCheckResourceAttr(resourceFullName, "name", resourceName),
				),
			},
		},
	})
}

func testAccReadResources_Heavy_Hcl(resourceName, valuePrefix string) (hcl string) {
	baseHcl := acctest.BaselineHcl(resourceName)
	clientId := acctest.RandStringWithPrefix(valuePrefix)
	clientSecret := acctest.RandStringWithPrefix(valuePrefix)
	hcl = fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s" {
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  connector_id   = "crowdStrikeConnector"
  name           = "%[2]s"
  property {
    name  = "clientId"
    value = "%[3]s"
  }
  property {
    name  = "clientSecret"
    value = "%[3]s"
  }
  depends_on = [data.davinci_connections.read_all]
}

resource "davinci_flow" "%[2]s_simple_flow" {
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  flow_json      = "{\"customerId\":\"dc7918cfa4b50966f8508072c2755c2c\",\"name\":\"tftesting-%[2]s-flow\",\"description\":\"\",\"flowStatus\":\"enabled\",\"createdDate\":1662960509175,\"updatedDate\":1662961640567,\"currentVersion\":0,\"authTokenExpireIds\":[],\"deployedDate\":1662960510106,\"functionConnectionId\":null,\"isOutputSchemaSaved\":false,\"outputSchemaCompiled\":null,\"publishedVersion\":1,\"timeouts\":null,\"flowId\":\"bb45eb4a0e8a5c9d6a21c7ac2d1b3faa\",\"companyId\":\"c431739a-29cd-4d9e-b465-0b37b2c235b1\",\"versionId\":0,\"graphData\":{\"elements\":{\"nodes\":[{\"data\":{\"id\":\"pptape4ac2\",\"nodeType\":\"CONNECTION\",\"connectionId\":\"867ed4363b2bc21c860085ad2baa817d\",\"connectorId\":\"httpConnector\",\"name\":\"Http\",\"label\":\"Http\",\"status\":\"configured\",\"capabilityName\":\"customHtmlMessage\",\"type\":\"trigger\",\"properties\":{\"message\":{\"value\":\"[\\n{\\n\\\"children\\\":[\\n{\\n\\\"text\\\":\\\"This is a simple test flow\\\"\\n}\\n]\\n}\\n]\"}}},\"position\":{\"x\":570,\"y\":240},\"group\":\"nodes\",\"removed\":false,\"selected\":false,\"selectable\":true,\"locked\":false,\"grabbable\":true,\"pannable\":false,\"classes\":\"\"}]},\"data\":{},\"zoomingEnabled\":true,\"userZoomingEnabled\":true,\"zoom\":1,\"minZoom\":1e-50,\"maxZoom\":1e+50,\"panningEnabled\":true,\"userPanningEnabled\":true,\"pan\":{\"x\":0,\"y\":0},\"boxSelectionEnabled\":true,\"renderer\":{\"name\":\"null\"}},\"flowColor\":\"#AFD5FF\",\"connectorIds\":[\"httpConnector\"],\"savedDate\":1662961640542,\"variables\":[]}"
  depends_on     = [data.davinci_connections.read_all]
}

resource "davinci_application" "%[2]s_simple_flow_app" {
  name           = "simple-%[2]s"
  depends_on     = [data.davinci_connections.read_all]
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  oauth {
    enabled = true
    values {
      allowed_grants                = ["authorizationCode"]
      allowed_scopes                = ["openid", "profile"]
      enabled                       = true
      enforce_signed_request_openid = false
      redirect_uris                 = ["https://auth.pingone.com/0000-0000-000/rp/callback/openid_connect"]
    }
  }
  policy {
    name = "PingOne - Sign On and Password Reset"
    policy_flow {
      flow_id    = resource.davinci_flow.%[2]s_simple_flow.id
      version_id = -1
      weight     = 100
    }
  }
  saml {
    values {
      enabled                = false
      enforce_signed_request = false
    }
  }
}
`, baseHcl, resourceName, clientId, clientSecret)

	for i := 0; i < 600; i++ {
		connHcl := fmt.Sprintf(`
data "davinci_connection" "http_%[2]s_%[1]s" {
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  id             = davinci_connection.%[2]s.id
  depends_on     = [data.davinci_connections.read_all]
}

data "davinci_application" "http_%[2]s_%[1]s" {
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  application_id = davinci_application.%[2]s_simple_flow_app.id
  depends_on     = [data.davinci_connections.read_all]
}


`, strconv.Itoa(i), resourceName, clientId, clientSecret)
		hcl += connHcl
	}

	return hcl
}
