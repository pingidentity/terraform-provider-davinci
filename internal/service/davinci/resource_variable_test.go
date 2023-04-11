package davinci_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
)

func TestAccResourceVariable_CompanyContext(t *testing.T) {

	resourceBase := "davinci_variable"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	hcl := testAccResourceVariable_CompanyContext_Hcl(resourceName)

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
					resource.TestCheckResourceAttrSet(resourceFullName, "name"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "value"),
				),
			},
		},
	})
}

func testAccResourceVariable_CompanyContext_Hcl(resourceName string) (hcl string) {
	baseHcl := acctest.PingoneEnvrionmentSsoHcl(resourceName)
	hcl = fmt.Sprintf(`
%[1]s

resource "davinci_variable" "%[2]s" {
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  name           = "company-%[2]s"
  context        = "company"
  description    = "desc-%[2]s"
  value          = "val-%[2]s"
  type           = "string"
  depends_on     = [data.davinci_connections.read_all]
}
`, baseHcl, resourceName)
	return hcl
}

func TestAccResourceVariable_FlowInstanceContext(t *testing.T) {

	resourceBase := "davinci_variable"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	hcl := testAccResourceVariable_FlowInstanceContext_Hcl(resourceName)

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
					resource.TestCheckResourceAttrSet(resourceFullName, "name"),
					resource.TestCheckResourceAttrSet(resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "value"),
					//Davinci uses 0,2000 as defaults
					resource.TestCheckResourceAttr(resourceFullName, "max", "1000"),
					resource.TestCheckResourceAttr(resourceFullName, "min", "0"),
				),
			},
		},
	})
}

func testAccResourceVariable_FlowInstanceContext_Hcl(resourceName string) (hcl string) {
	baseHcl := acctest.PingoneEnvrionmentSsoHcl(resourceName)
	hcl = fmt.Sprintf(`
%[1]s

resource "davinci_variable" "%[2]s" {
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  name           = "flowInstance-%[2]s"
  context        = "flowInstance"
  description    = "desc-%[2]s"
  value          = "val-%[2]s"
  type           = "string"
  min            = 0
  max            = 1000
  depends_on     = [data.davinci_connections.read_all]
}
`, baseHcl, resourceName)
	return hcl
}
