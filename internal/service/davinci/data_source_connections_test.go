package davinci_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
)

func TestAccDataSourceConnection_AllConections(t *testing.T) {

	// resourceName := "foo"
	// resourceFullName := fmt.Sprintf("davinci_connection.%s", resourceName)
	// dataSourceFullName := fmt.Sprintf("data.%s", resourceFullName)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
				data "pingone_password_policy" "example_by_name" {
					environment_id = var.environment_id
				
					name = "Standard"
				}
				variable "environment_id" {}
				data "davinci_connections" "all" {
				}
				output "davinci_connections" {
					value = data.davinci_connections.all.connections
				}
				output "pingone_password_pol" {
					value = data.pingone_password_policy.example_by_name.description
				}
				`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.davinci_connections.all", "connections.0.name"),
					resource.TestCheckResourceAttrSet("data.davinci_connections.all", "environment_id"),
				),
			},
		},
	})
}
