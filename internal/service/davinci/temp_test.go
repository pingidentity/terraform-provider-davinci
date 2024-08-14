package davinci_test

import (
	"fmt"
	// "regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest/service/davinci"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
)

func testAccResourceFlow_RedactedCompanyVariableFlow_HCL(resourceName, name string, withBootstrapConfig bool) (hcl, mainFlowJson string, err error) {

	mainFlowJson, err = acctest.ReadFlowJsonFile("flows/redactedcompanyvarminimal.json")
	if err != nil {
		return "", "", err
	}

	commonHcl, err := testAccResourceFlow_Common_WithMappingIDs_HCL(resourceName, name)
	if err != nil {
		return "", "", err
	}

	return fmt.Sprintf(`
%[1]s

%[2]s

resource "davinci_flow" "%[3]s" {
  environment_id = pingone_environment.%[3]s.id
	description    = "minimal flow with company variables redacted"

  flow_json = <<EOT
%[4]s
EOT

  // http connector
  connection_link {
    id                           = davinci_connection.%[3]s-http.id
    name                         = davinci_connection.%[3]s-http.name
    replace_import_connection_id = "867ed4363b2bc21c860085ad2baa817d"
  }
}`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), commonHcl, resourceName, mainFlowJson), mainFlowJson, nil
}

func testAccResourceFlow_RedactedCompanyVar(t *testing.T, withBootstrapConfig bool) {

	resourceBase := "davinci_flow"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	redactedVarStepHcl, redactedVarStepJson, err := testAccResourceFlow_RedactedCompanyVariableFlow_HCL(resourceName, name, withBootstrapConfig)
	if err != nil {
		t.Fatalf("Failed to get HCL: %v", err)
	}

	redactedVarStep := resource.TestStep{
		Config: redactedVarStepHcl,
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", "redactedcompanyvarminimal"),
			resource.TestCheckResourceAttr(resourceFullName, "description", "minimal flow with company variables redacted"),
			resource.TestCheckResourceAttr(resourceFullName, "flow_json", fmt.Sprintf("%s\n", redactedVarStepJson)),
			resource.TestCheckResourceAttrSet(resourceFullName, "flow_configuration_json"),
			resource.TestCheckResourceAttrSet(resourceFullName, "flow_export_json"),
			resource.TestCheckResourceAttr(resourceFullName, "connection_link.#", "1"),
			resource.TestCheckResourceAttr(resourceFullName, "deploy", "true"),
			resource.TestCheckResourceAttr(resourceFullName, "subflow_link.#", "0"),
			resource.TestCheckResourceAttr(resourceFullName, "flow_variables.#", "1"),
			// resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "flow_variables.*", map[string]*regexp.Regexp{
			// 	"context": regexp.MustCompile(`^company$`),
			// 	"flow_id": verify.P1DVResourceIDRegexpFullString,
			// 	// "id":      regexp.MustCompile(fmt.Sprintf(`^exampleCompanyVar##SK##company%s$`, verify.P1DVResourceIDRegexp.String())),
			// 	"max":     regexp.MustCompile(`^2000$`),
			// 	"min":     regexp.MustCompile(`^0$`),
			// 	"mutable": regexp.MustCompile(`^true$`),
			// 	// "name":    regexp.MustCompile(`^exampleCompanyVar$`),
			// 	"type": regexp.MustCompile(`^string$`),
			// }),
			// resource.TestMatchTypeSetElemNestedAttrs(resourceFullName, "flow_variables.*", map[string]*regexp.Regexp{
			// 	"context": regexp.MustCompile(`^flow$`),
			// 	"flow_id": verify.P1DVResourceIDRegexpFullString,
			// 	"id":      regexp.MustCompile(fmt.Sprintf(`^test123##SK##flow##SK##%s$`, verify.P1DVResourceIDRegexp.String())),
			// 	"max":     regexp.MustCompile(`^20$`),
			// 	"min":     regexp.MustCompile(`^4$`),
			// 	"mutable": regexp.MustCompile(`^true$`),
			// 	"name":    regexp.MustCompile(`^test123$`),
			// 	"type":    regexp.MustCompile(`^number$`),
			// 	"value":   regexp.MustCompile(`^10$`),
			// }),
		),
	}

	redactedVarStepPlanOnly := redactedVarStep
	redactedVarStepPlanOnly.PlanOnly = true

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)

			if withBootstrapConfig {
				t.Skipf("Skipping test with bootstrap config: https://github.com/pingidentity/terraform-provider-davinci/issues/266")
			}
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Flow_CheckDestroy(),
		Steps: []resource.TestStep{
			// Test redacted company variable flow
			redactedVarStep,
			// run the same config with plan-only:true to simulate a separate, subsequent plan.
			redactedVarStepPlanOnly,
			// Test importing the resource
			{
				ResourceName: resourceFullName,
				ImportStateIdFunc: func() resource.ImportStateIdFunc {
					return func(s *terraform.State) (string, error) {
						rs, ok := s.RootModule().Resources[resourceFullName]
						if !ok {
							return "", fmt.Errorf("Resource Not found: %s", resourceFullName)
						}

						return fmt.Sprintf("%s/%s", rs.Primary.Attributes["environment_id"], rs.Primary.ID), nil
					}
				}(),
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"connection_link",
					"subflow_link",
					"flow_json",
					"flow_export_json",
				},
			},
		},
	})
}

func TestAccResourceFlow_RedactedCompanyVar_Clean(t *testing.T) {
	testAccResourceFlow_RedactedCompanyVar(t, false)
}
