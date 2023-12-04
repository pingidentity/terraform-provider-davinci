package davinci_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
)

// TestAccDatasourceApplication_Slim - Depends on testAccResourceApplication_Slim_Hcl
func TestAccDatasourceApplication_SlimByAppId(t *testing.T) {
	resourceBase := "davinci_application"
	resourceName := acctest.ResourceNameGen()
	dataSourceFullName := fmt.Sprintf("data.%s.%s", resourceBase, resourceName)
	hcl := testAccDatasourceApplication_SlimByAppId_Hcl(resourceName)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		CheckDestroy:      acctest.CheckResourceDestroy([]string{"davinci_application"}),
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceFullName, "name"),
					resource.TestCheckResourceAttrSet(dataSourceFullName, "id"),
				),
			},
		},
	})
}

func TestAccDatasourceApplication_SlimById(t *testing.T) {

	resourceBase := "davinci_application"
	resourceName := acctest.ResourceNameGen()
	dataSourceFullName := fmt.Sprintf("data.%s.%s", resourceBase, resourceName)
	hcl := testAccDatasourceApplication_SlimById_Hcl(resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		CheckDestroy:      acctest.CheckResourceDestroy([]string{"davinci_application"}),
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataSourceFullName, "name"),
					resource.TestCheckResourceAttrSet(dataSourceFullName, "id"),
				),
			},
		},
	})
}

func testAccDatasourceApplication_SlimByAppId_Hcl(resourceName string) (hcl string) {
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

func testAccDatasourceApplication_SlimById_Hcl(resourceName string) (hcl string) {
	baseHcl := testAccResourceApplication_Slim_Hcl(resourceName)
	hcl = fmt.Sprintf(`
%[1]s

data "davinci_application" "%[2]s" {
  id             = resource.davinci_application.%[2]s.id
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
}
`, baseHcl, resourceName)
	return hcl
}
