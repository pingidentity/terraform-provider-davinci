package davinci_test

import (
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
					resource.TestCheckResourceAttrSet(dataSourceFullName, "id"),
				),
			},
		},
	})
}

func TestAccDataSourceConnection_ReadSingleByName(t *testing.T) {

	resourceBase := "davinci_connection"
	resourceName := acctest.ResourceNameGen()
	dataSourceFullName := fmt.Sprintf("data.%s.%s", resourceBase, resourceName)
	hcl := testAccDataSourceConnection_SlimByName(resourceName)

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
					resource.TestCheckResourceAttrSet(dataSourceFullName, "id"),
				),
			},
		},
	})
}

func testAccDataSourceConnection_Slim(resourceName string) (hcl string) {
	baseHcl := acctest.BaselineHcl(resourceName)
	hcl = fmt.Sprintf(`
%[1]s

data "davinci_connection" "%[2]s"{
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
	id = tolist(data.davinci_connections.read_all.connections)[0].id
	depends_on = [data.davinci_connections.read_all]
}
`, baseHcl, resourceName)
	return hcl
}

func testAccDataSourceConnection_SlimByName(resourceName string) (hcl string) {
	baseHcl := acctest.BaselineHcl(resourceName)
	hcl = fmt.Sprintf(`
	%[1]s
	
data "davinci_connection" "%[2]s"{
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
	name = "Http"
	depends_on = [data.davinci_connections.read_all]
}
`, baseHcl, resourceName)
	return hcl
}
