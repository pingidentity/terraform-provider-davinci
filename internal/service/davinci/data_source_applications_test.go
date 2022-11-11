package davinci_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
)

func TestAccDataSourceApplications_AllApplications(t *testing.T) {

	resourceBase := "davinci_applications"
	resourceName := acctest.ResourceNameGen()
	dataSourceFullName := fmt.Sprintf("data.%s.%s", resourceBase, resourceName)

	baseHcl := acctest.PingoneEnvrionmentSsoHcl(resourceName)
	hcl := baseHcl + testAccDataSourceApplications_AllApplications_Hcl(resourceName)

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
					resource.TestCheckResourceAttrSet(dataSourceFullName, "applications.0.environment_id"),
				),
			},
		},
	})
}

func testAccDataSourceApplications_AllApplications_Hcl(resourceName string) (hcl string) {
	hcl = fmt.Sprintf(`
data "davinci_applications" "%[1]s" {
	environment_id = resource.pingone_role_assignment_user.%[1]s.scope_environment_id
}
`, resourceName)
	return hcl
}
