package davinci_test

import (
	"fmt"
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
					resource.TestCheckResourceAttrSet(resourceFullName, "connection_id"),
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
					resource.TestCheckResourceAttrSet(resourceFullName, "connection_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttr(resourceFullName, "name", resourceName),
				),
			},
			{
				Config: afterHcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "connection_id"),
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
	connector_id = "crowdStrikeConnector"
	name         = "%[2]s"
	properties {
		name  = "clientId"
		value = "%[3]s"
	}
	properties {
		name  = "clientSecret"
		value = "%[3]s"
	}
}
`, baseHcl, resourceName, clientId, clientSecret)

	return hcl
}
