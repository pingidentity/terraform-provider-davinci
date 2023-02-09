package davinci_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
)

func TestAccResourceApplication_Slim(t *testing.T) {

	resourceBase := "davinci_application"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	hcl := testAccResourceApplication_Slim_Hcl(resourceName)

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
					//TODO - test attributes in TypeSet.
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					// resource.TestCheckNoResourceAttr(resourceFullName, "application_id"),
					// TODO - use this on integrated acc test
					// resource.TestCheckTypeSetElemNestedAttrs(resourceFullName,
					// 	"policies.0.policy_flows.*",
					// 	map[string]string{
					// 		"version_id": "-1",
					// 		"weight":     "100",
					// 	}),
				),
			},
		},
	})
}

func testAccResourceApplication_Slim_Hcl(resourceName string) (hcl string) {
	baseHcl := acctest.BaselineHcl(resourceName)
	hcl = fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
	name = "TF ACC Test"
	depends_on = [ data.davinci_connections.read_all]
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
`, baseHcl, resourceName)
	return hcl
}

func TestAccResourceApplication_WithFlowPolicy(t *testing.T) {

	resourceBase := "davinci_application"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		// CheckDestroy: testAccCheckExampleResourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApplication_WithFlowPolicy_Hcl(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "policy.0.policy_id"),
					resource.TestCheckNoResourceAttr(resourceFullName, "policy.1.policy_id"),
				),
			},
			{
				Config: testAccResourceApplication_WithFlowPolicyUpdate_Hcl(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "policy.1.policy_id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "policy.0.policy_id"),
				),
			},
		},
	})
}

func testAccResourceApplication_WithFlowPolicy_Hcl(resourceName string) (hcl string) {
	flows := acctest.FlowsForTests(resourceName)

	baseHcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{flows.Simple.Hcl})
	hcl = fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
	name = "TF ACC Test"
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
	depends_on = [ data.davinci_connections.read_all]
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
  policy {
    name = "simpleflow"
    policy_flow {
      flow_id    = resource.davinci_flow.%[3]s.id
      version_id = -1
      weight     = 100
    }
		status = "enabled"
  }
}
`, baseHcl, resourceName, flows.Simple.Name)
	return hcl
}

func testAccResourceApplication_WithFlowPolicyUpdate_Hcl(resourceName string) (hcl string) {
	flows := acctest.FlowsForTests(resourceName)

	baseHcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{flows.Simple.Hcl})
	hcl = fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
	name = "TF ACC Test"
	environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
	depends_on = [ data.davinci_connections.read_all]
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
  policy {
    name = "simpleflow"
    policy_flow {
      flow_id    = resource.davinci_flow.%[3]s.id
      version_id = -1
      weight     = 100
    }
		status = "enabled"
  }
  policy {
    name = "subsequentPolicy"
    policy_flow {
      flow_id    = resource.davinci_flow.%[3]s.id
      version_id = -1
      weight     = 100
    }
		status = "enabled"
  }
}
`, baseHcl, resourceName, flows.Simple.Name)
	return hcl
}
