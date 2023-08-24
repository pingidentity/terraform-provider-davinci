package davinci_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
		CheckDestroy:      acctest.CheckResourceDestroy([]string{"davinci_variable"}),
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
		CheckDestroy:      acctest.CheckResourceDestroy([]string{"davinci_variable"}),
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

func testAccGetResourceVariableIDs(resourceName string, environmentID, resourceID *string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return fmt.Errorf("Resource Not found: %s", resourceName)
		}

		*resourceID = rs.Primary.ID
		*environmentID = rs.Primary.Attributes["environment_id"]

		return nil
	}
}

func TestAccResourceVariable_RemovalDrift(t *testing.T) {
	t.Parallel()

	resourceBase := "davinci_variable"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	hcl := testAccResourceVariable_CompanyContext_Hcl(resourceName)

	var resourceID, environmentID string

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		CheckDestroy:      acctest.CheckResourceDestroy([]string{"davinci_variable"}),
		ErrorCheck:        acctest.ErrorCheck(t),
		Steps: []resource.TestStep{
			// Configure
			{
				Config: hcl,
				Check:  testAccGetResourceVariableIDs(resourceFullName, &environmentID, &resourceID),
			},
			// Replan after removal preconfig
			{
				PreConfig: func() {
					c, err := acctest.TestClient()

					if err != nil {
						t.Fatalf("Failed to get API client: %v", err)
					}

					if environmentID == "" || resourceID == "" {
						t.Fatalf("One of environment ID or resource ID cannot be determined. Environment ID: %s, Resource ID: %s", environmentID, resourceID)
					}

					_, err = c.DeleteVariable(&environmentID, resourceID)
					if err != nil {
						t.Fatalf("Failed to delete variable: %v", err)
					}
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
