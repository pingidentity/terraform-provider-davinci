package davinci_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
)

func TestAccResourceApplicationPolicy_Base(t *testing.T) {

	resourceBase := "davinci_application_flow_policy"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		CheckDestroy:      acctest.CheckResourceDestroy([]string{"davinci_application"}),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApplicationPolicy_Base_Hcl(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
				),
			},
		},
	})
}

func testAccResourceApplicationPolicy_Base_Hcl(resourceName string) (hcl string) {
	flows := acctest.FlowsForTests(resourceName)

	baseHcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{flows.Simple.Hcl})
	hcl = fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
  name           = "TF ACC Test"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  depends_on     = [data.davinci_connections.read_all]
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
  name = "simpleflow"
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

func TestAccResourceApplication_WithFlowPolicy(t *testing.T) {

	resourceBase := "davinci_application_flow_policy"
	resourceName := acctest.ResourceNameGen()
	resourceNameB := fmt.Sprintf("%s-b", resourceName)
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)
	resourceFullNameB := fmt.Sprintf("%s.%s", resourceBase, resourceNameB)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		CheckDestroy:      acctest.CheckResourceDestroy([]string{"davinci_application"}),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApplication_WithFlowPolicy_Hcl(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "name"),
					resource.TestCheckResourceAttrSet(resourceFullName, "status"),
				),
			},
			{
				Config: testAccResourceApplication_WithFlowPolicyUpdate_Hcl(resourceName, resourceNameB),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "name"),
					resource.TestCheckResourceAttrSet(resourceFullName, "status"),
					resource.TestCheckResourceAttrSet(resourceFullNameB, "id"),
					resource.TestCheckResourceAttrSet(resourceFullNameB, "name"),
					resource.TestCheckResourceAttrSet(resourceFullNameB, "status"),
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
  name           = "TF ACC Test"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  depends_on     = [data.davinci_connections.read_all]
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
  name = "simpleflow"
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

func testAccResourceApplication_WithFlowPolicyUpdate_Hcl(resourceName string, resourceNameB string) (hcl string) {
	flows := acctest.FlowsForTests(resourceName)

	baseHcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{flows.Simple.Hcl})
	hcl = fmt.Sprintf(`
%[1]s

resource "davinci_application" "%[2]s" {
  name           = "TF ACC Test"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  depends_on     = [data.davinci_connections.read_all]
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
  name = "simpleflow"
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
  name = "simpleflow-another"
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

func TestAccResourceApplication_P1SessionFlowPolicy(t *testing.T) {
	resourceAppBase := "davinci_application"
	resourceAppFlowPolicyBase := "davinci_application_flow_policy"
	resourceName := acctest.ResourceNameGen()
	resourceAppFullName := fmt.Sprintf("%s.%s", resourceAppBase, resourceName)
	resourceAppFlowPolicyFullName := fmt.Sprintf("%s.%s", resourceAppFlowPolicyBase, resourceName)
	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		CheckDestroy:      acctest.CheckResourceDestroy([]string{"davinci_application"}),
		Steps: []resource.TestStep{
			{
				Config: testAccResourceApplication_P1SessionFlowPolicy_Hcl(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceAppFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceAppFullName, "name"),
					resource.TestCheckNoResourceAttr(resourceAppFullName, "policy"),
					resource.TestCheckResourceAttrSet(resourceAppFlowPolicyFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceAppFlowPolicyFullName, "name"),
					resource.TestCheckResourceAttrSet(resourceAppFlowPolicyFullName, "status"),
				),
			},
			{
				Config: testAccResourceApplication_P1SessionFlowPolicyUpdate_Hcl(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceAppFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceAppFullName, "name"),
					resource.TestCheckNoResourceAttr(resourceAppFullName, "policy"),
					resource.TestCheckResourceAttrSet(resourceAppFlowPolicyFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceAppFlowPolicyFullName, "name"),
					resource.TestCheckResourceAttrSet(resourceAppFlowPolicyFullName, "status"),
				),
			},
		},
	})
}

func testAccResourceApplication_P1SessionFlowPolicy_Hcl(resourceName string) (hcl string) {
	flows := acctest.FlowsForTests(resourceName)

	baseHcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{flows.PingOneSessionMainFlow.Hcl, flows.PingOneSessionSubFlow.Hcl})
	hcl = fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s-pingoneauthentication" {
  name           = "Ping One Authentication"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  connector_id   = "pingOneAuthenticationConnector"
  depends_on     = [data.davinci_connections.read_all]
}
resource "davinci_connection" "%[2]s-node" {
  name           = "Node"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  connector_id   = "nodeConnector"
  depends_on     = [data.davinci_connections.read_all]
}

resource "davinci_connection" "%[2]s-flow" {
  name           = "Flow"
  connector_id   = "flowConnector"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  depends_on     = [data.davinci_connections.read_all]
}

resource "davinci_application" "%[2]s" {
  name           = "TF ACC Test-%[2]s"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  depends_on     = [data.davinci_connections.read_all]
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
  name = "simpleflow"
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

func testAccResourceApplication_P1SessionFlowPolicyUpdate_Hcl(resourceName string) (hcl string) {
	flows := acctest.FlowsForTests(resourceName)

	baseHcl := testAccResourceFlow_SimpleFlows_Hcl(resourceName, []string{flows.PingOneSessionMainFlowUpdate.Hcl, flows.PingOneSessionSubFlow.Hcl})
	hcl = fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s-pingoneauthentication" {
  name           = "Ping One Authentication"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  connector_id   = "pingOneAuthenticationConnector"
  depends_on     = [data.davinci_connections.read_all]
}
resource "davinci_connection" "%[2]s-node" {
  name           = "Node"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  connector_id   = "nodeConnector"
  depends_on     = [data.davinci_connections.read_all]
}

resource "davinci_connection" "%[2]s-flow" {
  name           = "Flow"
  connector_id   = "flowConnector"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  depends_on     = [data.davinci_connections.read_all]
}

resource "davinci_application" "%[2]s" {
  name           = "TF ACC Test-%[2]s"
  environment_id = resource.pingone_role_assignment_user.%[2]s.scope_environment_id
  depends_on     = [data.davinci_connections.read_all]
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
  name = "simpleflow"
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
