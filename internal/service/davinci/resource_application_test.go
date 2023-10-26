package davinci_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
)

func TestAccResourceApplication_Slim(t *testing.T) {

	resourceBase := "davinci_application"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	hcl := testAccResourceApplication_Slim_Hcl(resourceName)

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
  name           = "TF ACC Test"
  depends_on     = [data.davinci_connections.read_all]
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

	resource.ParallelTest(t, resource.TestCase{
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
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{},
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

func TestAccResourceApplication_P1SessionFlowPolicy(t *testing.T) {

	resourceBase := "davinci_application"
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
				Config: testAccResourceApplication_P1SessionFlowPolicy_Hcl(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "policy.0.policy_id"),
					resource.TestCheckNoResourceAttr(resourceFullName, "policy.1.policy_id"),
				),
			},
			{
				Config: testAccResourceApplication_P1SessionFlowPolicyUpdate_Hcl(resourceName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resourceFullName, "id"),
					resource.TestCheckResourceAttrSet(resourceFullName, "policy.0.policy_id"),
					resource.TestCheckNoResourceAttr(resourceFullName, "policy.1.policy_id"),
				),
			},
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
					"context", // This shouldn't be ignored, can be solved on migration to the plugin framework
				},
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
`, baseHcl, resourceName, flows.PingOneSessionMainFlow.Name)
	return hcl
}

func TestAccResourceApplication_BadParameters(t *testing.T) {

	resourceBase := "davinci_application"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	hcl := testAccResourceApplication_Slim_Hcl(resourceName)

	resource.Test(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		CheckDestroy:      acctest.CheckResourceDestroy([]string{"davinci_application"}),
		Steps: []resource.TestStep{
			// Configure
			{
				Config: hcl,
			},
			// Errors
			{
				ResourceName: resourceFullName,
				ImportState:  true,
				ExpectError:  regexp.MustCompile(`Invalid import ID specified \(".*"\).  The ID should be in the format "environment_id/davinci_application_id" and must match regex: .*`),
			},
			{
				ResourceName:  resourceFullName,
				ImportStateId: "/",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`Invalid import ID specified \(".*"\).  The ID should be in the format "environment_id/davinci_application_id" and must match regex: .*`),
			},
			{
				ResourceName:  resourceFullName,
				ImportStateId: "badformat/badformat",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`Invalid import ID specified \(".*"\).  The ID should be in the format "environment_id/davinci_application_id" and must match regex: .*`),
			},
		},
	})
}

func testAccGetResourceApplicationIDs(resourceName string, environmentID, resourceID *string) resource.TestCheckFunc {
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

func TestAccResourceApplication_RemovalDrift(t *testing.T) {

	resourceBase := "davinci_application"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	hcl := testAccResourceApplication_Slim_Hcl(resourceName)

	var resourceID, environmentID string

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:          func() { acctest.PreCheckPingOneAndTfVars(t) },
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		CheckDestroy:      acctest.CheckResourceDestroy([]string{"davinci_application"}),
		Steps: []resource.TestStep{
			// Configure
			{
				Config: hcl,
				Check:  testAccGetResourceApplicationIDs(resourceFullName, &environmentID, &resourceID),
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

					_, err = c.DeleteApplication(&environmentID, resourceID)
					if err != nil {
						t.Fatalf("Failed to delete application: %v", err)
					}
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
