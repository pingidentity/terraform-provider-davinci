package davinci_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
)

func TestAccResourceApplicationFlowPolicy_Slim(t *testing.T) {

	resourceBase := "davinci_application_flow_policy"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		CheckDestroy:      acctest.CheckResourceDestroy([]string{"davinci_application"}),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApplicationFlowPolicy_Base_Hcl(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "name"),
					resource.TestCheckResourceAttrSet(resourceFullName, "status"),
				),
			},
		},
	})
}

func TestAccResourceApplicationFlowPolicy_Base(t *testing.T) {

	resourceBase := "davinci_application_flow_policy"
	resourceName := acctest.ResourceNameGen()
	resourceNameB := fmt.Sprintf("%s-b", resourceName)
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)
	resourceFullNameB := fmt.Sprintf("%s.%s", resourceBase, resourceNameB)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		CheckDestroy:      acctest.CheckResourceDestroy([]string{"davinci_application"}),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApplicationFlowPolicy_Base_Hcl(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "name"),
					resource.TestCheckResourceAttrSet(resourceFullName, "status"),
					testAccCheckApplicationFlowPolicy(acctest.TestApplicationFlowPolicy{
						FlowPolicyResourceName: resourceName,
					}),
				),
			},
			{
				Config: testAccResourceApplicationFlowPolicy_BaseUpdate_Hcl(resourceName, resourceNameB),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "name"),
					resource.TestCheckResourceAttrSet(resourceFullName, "status"),
					resource.TestCheckResourceAttrSet(resourceFullNameB, "id"),
					resource.TestCheckResourceAttrSet(resourceFullNameB, "name"),
					resource.TestCheckResourceAttrSet(resourceFullNameB, "status"),
					testAccCheckApplicationFlowPolicy(acctest.TestApplicationFlowPolicy{
						FlowPolicyResourceName: resourceName,
					}),
				),
			},
		},
	})
}

// ensure policy was created
func testAccCheckApplicationFlowPolicy(polcheck acctest.TestApplicationFlowPolicy) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := acctest.TestClient()
		if err != nil {
			return err
		}
		resourceFullName := polcheck.GetResourceFullName()
		rs, ok := s.RootModule().Resources[resourceFullName]
		if !ok {
			return fmt.Errorf("Resource Not found: %s", resourceFullName)
		}

		polcheck.SetID(rs.Primary.ID)
		polcheck.SetName(rs.Primary.Attributes["name"])
		polcheck.SetEnvironmentID(rs.Primary.Attributes["environment_id"])
		polcheck.SetApplicationID(rs.Primary.Attributes["application_id"])

		res, err := client.ReadApplication(&polcheck.EnvironmentID, polcheck.ApplicationID)
		if err != nil {
			return err
		}
		found := fmt.Errorf("Unable to find created policy in response from Davinci API")
		for _, v := range res.Policies {
			if v.Name == polcheck.Name && v.PolicyID == polcheck.ID {
				found = nil
				break
			}
		}
		return found
	}
}

func testAccResourceApplicationFlowPolicy_Base_Hcl(resourceName string) (hcl string) {
	flows := acctest.FlowsForTests(resourceName)

	baseHcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{flows.Simple.Hcl})
	hcl = fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
  name           = "TF ACC Test"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  oauth {
    enabled = true
    values {
      allowed_grants                = ["authorizationCode"]
      allowed_scopes                = ["openid", "profile"]
      enabled                       = true
      enforce_signed_request_openid = false
      redirect_uris                 = ["https://auth.pingone.com/env-id/rp/callback/openid_connect"]
    }
  }
  saml {
    values {
      enabled                = false
      enforce_signed_request = false
    }
  }
}

resource "davinci_application_flow_policy" "%[2]s" {
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  application_id = davinci_application.%[2]s.id
  name           = "simpleflow"
  policy_flow {
    flow_id    = resource.davinci_flow.%[3]s.id
    version_id = -1
    weight     = 100
  }
  status = "enabled"
}
`, baseHcl, resourceName, flows.Simple.Name)
	return hcl
}

func testAccResourceApplicationFlowPolicy_BaseUpdate_Hcl(resourceName string, resourceNameB string) (hcl string) {
	flows := acctest.FlowsForTests(resourceName)

	baseHcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{flows.Simple.Hcl})
	hcl = fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
  name           = "TF ACC Test"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  oauth {
    enabled = true
    values {
      allowed_grants                = ["authorizationCode"]
      allowed_scopes                = ["openid", "profile"]
      enabled                       = true
      enforce_signed_request_openid = false
      redirect_uris                 = ["https://auth.pingone.com/env-id/rp/callback/openid_connect"]
    }
  }
  saml {
    values {
      enabled                = false
      enforce_signed_request = false
    }
  }
}
resource "davinci_application_flow_policy" "%[2]s" {
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  application_id = davinci_application.%[2]s.id
  name           = "simpleflow"
  policy_flow {
    flow_id    = resource.davinci_flow.%[3]s.id
    version_id = -1
    weight     = 100
  }
  status = "enabled"
}
resource "davinci_application_flow_policy" "%[4]s" {
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  application_id = davinci_application.%[2]s.id
  name           = "simpleflow-another"
  policy_flow {
    flow_id    = resource.davinci_flow.%[3]s.id
    version_id = -1
    weight     = 100
  }
  status = "enabled"
}
`, baseHcl, resourceName, flows.Simple.Name, resourceNameB)
	return hcl
}

func TestAccResourceApplicationFlowPolicy_P1SessionFlowPolicy(t *testing.T) {
	resourceAppBase := "davinci_application"
	resourceAppFlowPolicyBase := "davinci_application_flow_policy"
	resourceName := acctest.ResourceNameGen()
	resourceAppFullName := fmt.Sprintf("%s.%s", resourceAppBase, resourceName)
	resourceAppFlowPolicyFullName := fmt.Sprintf("%s.%s", resourceAppFlowPolicyBase, resourceName)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		CheckDestroy:      acctest.CheckResourceDestroy([]string{"davinci_application"}),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApplicationFlowPolicy_P1SessionFlowPolicy_Hcl(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceAppFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceAppFullName, "name"),
					resource.TestCheckNoResourceAttr(resourceAppFullName, "policy.0.name"),
					resource.TestCheckResourceAttrSet(resourceAppFlowPolicyFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceAppFlowPolicyFullName, "name"),
					resource.TestCheckResourceAttrSet(resourceAppFlowPolicyFullName, "status"),
				),
			},
			{
				Config: testAccResourceApplicationFlowPolicy_P1SessionFlowPolicyUpdate_Hcl(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceAppFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceAppFullName, "name"),
					resource.TestCheckResourceAttrSet(resourceAppFullName, "policy.0.name"),
					resource.TestCheckResourceAttrSet(resourceAppFlowPolicyFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceAppFlowPolicyFullName, "name"),
					resource.TestCheckResourceAttrSet(resourceAppFlowPolicyFullName, "status"),
				),
			},
		},
	})
}

func testAccResourceApplicationFlowPolicy_P1SessionFlowPolicy_Hcl(resourceName string) (hcl string) {
	flows := acctest.FlowsForTests(resourceName)

	baseHcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{flows.PingOneSessionMainFlow.Hcl, flows.PingOneSessionSubFlow.Hcl})
	hcl = fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s-flow" {
  name           = "Flow"
  connector_id   = "flowConnector"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
}

resource "davinci_application" "%[2]s" {
  name           = "TF ACC Test-%[2]s"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  oauth {
    enabled = true
    values {
      allowed_grants                = ["authorizationCode"]
      allowed_scopes                = ["openid", "profile"]
      enabled                       = true
      enforce_signed_request_openid = false
      redirect_uris                 = ["https://auth.pingone.com/env-id/rp/callback/openid_connect"]
    }
  }
  saml {
    values {
      enabled                = false
      enforce_signed_request = false
    }
  }
}

resource "davinci_application_flow_policy" "%[2]s" {
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  application_id = davinci_application.%[2]s.id
  name           = "simpleflow"
  policy_flow {
    flow_id    = resource.davinci_flow.%[3]s.id
    version_id = -1
    weight     = 100
  }
  status = "enabled"
}
`, baseHcl, resourceName, flows.PingOneSessionMainFlow.Name)
	return hcl
}

func testAccResourceApplicationFlowPolicy_P1SessionFlowPolicyUpdate_Hcl(resourceName string) (hcl string) {
	flows := acctest.FlowsForTests(resourceName)

	baseHcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{flows.PingOneSessionMainFlowUpdate.Hcl, flows.PingOneSessionSubFlow.Hcl})
	hcl = fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s-flow" {
  name           = "Flow"
  connector_id   = "flowConnector"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
}

resource "davinci_application" "%[2]s" {
  name           = "TF ACC Test-%[2]s"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  oauth {
    enabled = true
    values {
      allowed_grants                = ["authorizationCode"]
      allowed_scopes                = ["openid", "profile"]
      enabled                       = true
      enforce_signed_request_openid = false
      redirect_uris                 = ["https://auth.pingone.com/env-id/rp/callback/openid_connect"]
    }
  }
  saml {
    values {
      enabled                = false
      enforce_signed_request = false
    }
  }
}
resource "davinci_application_flow_policy" "%[2]s" {
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  application_id = davinci_application.%[2]s.id
  name           = "simpleflow"
  policy_flow {
    flow_id    = resource.davinci_flow.%[3]s.id
    version_id = -1
    weight     = 100
  }
  status = "enabled"
}
`, baseHcl, resourceName, flows.PingOneSessionMainFlowUpdate.Name)
	return hcl
}
