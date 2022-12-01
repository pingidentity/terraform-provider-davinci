package davinci_test

import (
	// "fmt"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
)

func TestAccDataSourceConnection_ReadSingle(t *testing.T) {

	resourceBase := "davinci_connection"
	resourceName := acctest.ResourceNameGen()
	dataSourceFullName := fmt.Sprintf("data.%s.%s", resourceBase, resourceName)
	hcl := testAccDataSourceConnection_Slim(resourceName)

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
					resource.TestCheckResourceAttrSet(dataSourceFullName, "name"),
					resource.TestCheckResourceAttrSet(dataSourceFullName, "connector_id"),
					resource.TestCheckResourceAttrSet(dataSourceFullName, "created_date"),
					resource.TestCheckResourceAttrSet(dataSourceFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(dataSourceFullName, "connection_id"),
				),
			},
		},
	})
}

func testAccDataSourceConnection_Slim(resourceName string) (hcl string) {
	baseHcl := testAccDataSourceConnection_AllConections_Hcl(resourceName)
	hcl = fmt.Sprintf(`
%[1]s

data "davinci_connection" "%[2]s"{
	environment_id = resource.pingone_environment.%[2]s.id
	connection_id = tolist(data.davinci_connections.%[2]s.connections)[0].connection_id
}
`, baseHcl, resourceName)
	return hcl
}
