package davinci_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
)

// TestAccDatasourceApplication_Slim - Depends on testAccResourceApplication_Slim_Hcl
func TestAccDatasourceApplication_Slim(t *testing.T) {

	resourceBase := "davinci_application"
	resourceName := acctest.ResourceNameGen()
	dataSourceFullName := fmt.Sprintf("data.%s.%s", resourceBase, resourceName)
	hcl := testAccDatasourceApplication_Slim_Hcl(resourceName)

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
				),
			},
		},
	})
}

func testAccDatasourceApplication_Slim_Hcl(resourceName string) (hcl string) {
	baseHcl := testAccResourceApplication_Slim_Hcl(resourceName)
	hcl = fmt.Sprintf(`
%[1]s

data "davinci_application" "%[2]s" {
	application_id = resource.davinci_application.%[2]s.id
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
}
`, baseHcl, resourceName)
	return hcl
}
