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

	hcl := testAccDataSourceApplications_AllApplications_Hcl(resourceName)

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
	baseHcl := testAccResourceApplication_Slim_Hcl(resourceName)
	hcl = fmt.Sprintf(`
%[1]s

data "davinci_applications" "%[2]s" {
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id

	depends_on = [
		davinci_application.%[2]s
	]
}
`, baseHcl, resourceName)
	return hcl
}
