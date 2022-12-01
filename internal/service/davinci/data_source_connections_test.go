package davinci_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
)

func TestAccDataSourceConnection_AllConections(t *testing.T) {

	resourceBase := "davinci_connections"
	resourceName := acctest.ResourceNameGen()
	dataSourceFullName := fmt.Sprintf("data.%s.%s", resourceBase, resourceName)
	hcl := testAccDataSourceConnection_AllConections_Hcl(resourceName)

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
					resource.TestCheckResourceAttrSet(dataSourceFullName, "connections.0.name"),
					resource.TestCheckResourceAttrSet(dataSourceFullName, "environment_id"),
				),
			},
		},
	})
}

func testAccDataSourceConnection_AllConections_Hcl(resourceName string) (hcl string) {
	baseHcl := acctest.PingoneEnvrionmentSsoHcl(resourceName)
	hcl = fmt.Sprintf(`
%[1]s

data "davinci_connections" "%[2]s" {
	environment_id = resource.pingone_environment.%[2]s.id
}
`, baseHcl, resourceName)
	return hcl
}

func TestAccDataSourceConnection_FilteredConnections(t *testing.T) {

	resourceBase := "davinci_connections"
	resourceName := acctest.ResourceNameGen()
	dataSourceFullName := fmt.Sprintf("data.%s.%s", resourceBase, resourceName)
	hcl := testAccDataSourceConnection_FilteredConnections_Hcl(resourceName)

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
					resource.TestCheckResourceAttrSet(dataSourceFullName, "connections.1.name"),
					resource.TestCheckNoResourceAttr(dataSourceFullName, "connections.2.name"),
					resource.TestCheckResourceAttrSet(dataSourceFullName, "environment_id"),
				),
			},
		},
	})
}

func testAccDataSourceConnection_FilteredConnections_Hcl(resourceName string) (hcl string) {
	baseHcl := acctest.PingoneEnvrionmentSsoHcl(resourceName)
	hcl = fmt.Sprintf(`
%[1]s

data "davinci_connections" "%[2]s" {
	connector_ids = [ "httpConnector", "functionsConnector" ]
	environment_id = resource.pingone_environment.%[2]s.id
}
`, baseHcl, resourceName)
	return hcl
}
