package davinci_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
)

func TestAccDataSourceConnection_ReadAllAT(t *testing.T) {

	resourceBase := "davinci_connections"
	resourceName := acctest.ResourceNameGen()
	dataSourceFullName := fmt.Sprintf("data.%s.%s", resourceBase, resourceName)

	hcl := testAccDataSourceConnection_Slim(resourceName)
	hcl = fmt.Sprintf(`
data "davinci_connections" "%[1]s" {
	environment_id = var.environment_id
}

resource "davinci_connection" "%[1]s" {
  environment_id = var.environment_id
  connector_id   = "httpConnector"
  name           = "%[1]s"
}

variable "environment_id" {
	default = "089b57b4-70a5-45db-8add-01120f8b0063"
}
`, resourceName)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		// CheckDestroy:      acctest.CheckResourceDestroy([]string{"davinci_connections"}),
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(fmt.Sprintf("%s", dataSourceFullName), "environment_id", "089b57b4-70a5-45db-8add-01120f8b0063"),
					resource.TestCheckResourceAttr(fmt.Sprintf("davinci_connection.%s", resourceName), "environment_id", "089b57b4-70a5-45db-8add-01120f8b0063"),
					// resource.TestCheckResourceAttrSet(dataSourceFullName, "connector_id"),
					// resource.TestCheckResourceAttrSet(dataSourceFullName, "created_date"),
					// resource.TestCheckResourceAttrSet(dataSourceFullName, "environment_id"),
					// resource.TestCheckResourceAttrSet(dataSourceFullName, "id"),
				),
			},
		},
	})
}

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
		CheckDestroy:      acctest.CheckResourceDestroy([]string{"davinci_connection"}),
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
		CheckDestroy:      acctest.CheckResourceDestroy([]string{"davinci_connection"}),
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
