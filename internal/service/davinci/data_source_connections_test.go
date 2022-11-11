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

	baseHcl := acctest.PingoneEnvrionmentSsoHcl(resourceName)
	hcl := baseHcl + testAccDataSourceConnection_AllConections_Hcl(resourceName)

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
	hcl = fmt.Sprintf(`
data "davinci_connections" "%[1]s" {
	environment_id = resource.pingone_environment.%[1]s.id
}
`, resourceName)
	return hcl
}
