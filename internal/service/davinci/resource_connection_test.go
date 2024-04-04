package davinci_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest/service/davinci"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
)

func TestAccResourceConnection_RemovalDrift(t *testing.T) {

	resourceBase := "davinci_connection"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	hcl := testAccResourceConnection_Full_HCL1(resourceName, resourceName, true)

	var connectionID, environmentID string

	// ctx := context.Background()

	// p1Client, err := acctest.PingOneTestClient(ctx)
	// if err != nil {
	// 	t.Fatalf("Failed to get API client: %v", err)
	// }

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Connection_CheckDestroy(),
		Steps: []resource.TestStep{
			// Configure
			{
				Config: hcl,
				Check:  davinci.Connection_GetIDs(resourceFullName, &environmentID, &connectionID),
			},
			// Replan after removal preconfig
			{
				PreConfig: func() {
					davinci.Connection_RemovalDrift_PreConfig(t, environmentID, connectionID)
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
			},
			// Test removal of the environment
			{
				Config:   hcl,
				Check:    davinci.Connection_GetIDs(resourceFullName, &environmentID, &connectionID),
				SkipFunc: func() (bool, error) { return true, nil },
			},
			{
				// PreConfig: func() {
				// 	base.Environment_RemovalDrift_PreConfig(ctx, p1Client.API.ManagementAPIClient, t, environmentID)
				// },
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
				SkipFunc:           func() (bool, error) { return true, nil },
			},
		},
	})
}

func TestAccResourceConnection_Full_Clean(t *testing.T) {
	testAccResourceConnection_Full(t, false)
}

func TestAccResourceConnection_Full_WithBootstrap(t *testing.T) {
	testAccResourceConnection_Full(t, true)
}

func testAccResourceConnection_Full(t *testing.T, withBootstrapConfig bool) {
	resourceBase := "davinci_connection"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	fullStep := resource.TestStep{
		Config: testAccResourceConnection_Full_HCL1(resourceName, fmt.Sprintf("%s-1", name), withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", fmt.Sprintf("%s-1", name)),
			resource.TestCheckResourceAttr(resourceFullName, "connector_id", "pingOneMfaConnector"),
			resource.TestCheckResourceAttr(resourceFullName, "property.#", "5"),
			resource.TestMatchResourceAttr(resourceFullName, "created_date", verify.EpochDateRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "customer_id", verify.P1DVResourceIDRegexpFullString),
		),
	}

	minimalStep := resource.TestStep{
		Config: testAccResourceConnection_Minimal_HCL(resourceName, fmt.Sprintf("%s-2", name), withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestMatchResourceAttr(resourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
			resource.TestCheckResourceAttr(resourceFullName, "name", fmt.Sprintf("%s-2", name)),
			resource.TestCheckResourceAttr(resourceFullName, "connector_id", "annotationConnector"),
			resource.TestCheckResourceAttr(resourceFullName, "property.#", "0"),
			resource.TestMatchResourceAttr(resourceFullName, "created_date", verify.EpochDateRegexpFullString),
			resource.TestMatchResourceAttr(resourceFullName, "customer_id", verify.P1DVResourceIDRegexpFullString),
		),
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Connection_CheckDestroy(),
		Steps: []resource.TestStep{
			// Create from scratch
			fullStep,
			{
				Config:  testAccResourceConnection_Full_HCL1(resourceName, fmt.Sprintf("%s-1", name), withBootstrapConfig),
				Destroy: true,
			},
			// Create from scratch
			minimalStep,
			{
				Config:  testAccResourceConnection_Minimal_HCL(resourceName, fmt.Sprintf("%s-2", name), withBootstrapConfig),
				Destroy: true,
			},
			// Test updates
			fullStep,
			minimalStep,
			fullStep,
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
					"property.1.value",
				},
			},
		},
	})
}

func TestAccResourceConnection_Properties(t *testing.T) {
	resourceBase := "davinci_connection"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	withBootstrapConfig := false

	fullStep1 := resource.TestStep{
		Config: testAccResourceConnection_Full_HCL1(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "property.#", "5"),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "region",
				"value": "NA",
				"type":  "string",
			}),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "envId",
				"value": "env-abc-123",
				"type":  "string",
			}),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "clientId",
				"value": "env-client-id",
				"type":  "string",
			}),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "clientSecret",
				"value": "env-client-secret",
				"type":  "string",
			}),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "policyId",
				"value": "policy-abc-123",
				"type":  "string",
			}),
		),
	}

	fullStep2 := resource.TestStep{
		Config: testAccResourceConnection_Full_HCL2(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "property.#", "4"),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "region",
				"value": "EU",
				"type":  "string",
			}),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "envId",
				"value": "env-abc-123",
				"type":  "string",
			}),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "clientId",
				"value": "env-client-id",
				"type":  "string",
			}),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "clientSecret",
				"value": "env-client-secret",
				"type":  "string",
			}),
		),
	}

	var connectionID, environmentID string

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Connection_CheckDestroy(),
		Steps: []resource.TestStep{
			// Test updates
			fullStep2, // without an optional property
			fullStep1, // with an optional property
			fullStep2, // without an optional property
			{
				Config: testAccResourceConnection_Full_HCL1(resourceName, name, withBootstrapConfig), // with an optional property to remove
				Check:  davinci.Connection_GetIDs(resourceFullName, &environmentID, &connectionID),
			},
			// Remove policyId via api and check for non-empty plan
			{
				PreConfig: func() {
					c, err := acctest.TestClient()
					if err != nil {
						t.Fatalf("Failed to get API client: %v", err)
					}

					if environmentID == "" || connectionID == "" {
						t.Fatalf("One of environment ID or connection ID cannot be determined. Environment ID: %s, Resource ID: %s", environmentID, connectionID)
					}

					connection, err := c.ReadConnection(environmentID, connectionID)
					if err != nil {
						t.Fatalf("Failed to read connection: - wut %v", err)
					}

					if _, ok := connection.Properties["policyId"]; ok {
						//remove policyId property
						delete(connection.Properties, "policyId")
					} else {
						t.Fatalf("Failed to read connection property: policyId")
					}

					_, err = c.UpdateConnection(environmentID, connection)
					if err != nil {
						t.Fatalf("Failed to update connection: %v", err)
					}
				},
				RefreshState:       true,
				ExpectNonEmptyPlan: true,
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
					"property.1.value",
				},
			},
		},
	})
}

func TestAccResourceConnection_ComplexProperties(t *testing.T) {
	resourceBase := "davinci_connection"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	name := resourceName

	withBootstrapConfig := false

	mixedType := resource.TestStep{
		Config: testAccResourceConnection_PropertyDataTypesMixed_HCL(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "property.#", "6"),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "hostname",
				"value": "localhost",
				"type":  "string",
			}),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "port",
				"value": "2525",
				"type":  "number",
			}),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "secureFlag",
				"value": "true",
				"type":  "boolean",
			}),
		),
	}

	jsonArrayType := resource.TestStep{
		Config: testAccResourceConnection_PropertyDataTypesJsonArray_HCL(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "property.#", "1"),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "customAttributes",
				"value": "{\"preferredControlType\":\"tableViewAttributes\",\"sections\":[\"connectorAttributes\"],\"type\":\"array\",\"value\":[{\"attributeType\":\"sk\",\"description\":\"Username\",\"maxLength\":\"300\",\"minLength\":\"1\",\"name\":\"username\",\"required\":true,\"type\":\"string\",\"value\":null},{\"attributeType\":\"sk\",\"description\":\"First Name\",\"maxLength\":\"100\",\"minLength\":\"1\",\"name\":\"firstName\",\"required\":false,\"type\":\"string\",\"value\":null},{\"attributeType\":\"sk\",\"description\":\"Last Name\",\"maxLength\":\"100\",\"minLength\":\"1\",\"name\":\"lastName\",\"required\":false,\"type\":\"string\",\"value\":null},{\"attributeType\":\"sk\",\"description\":\"Display Name\",\"maxLength\":\"250\",\"minLength\":\"1\",\"name\":\"name\",\"required\":false,\"type\":\"string\",\"value\":null},{\"attributeType\":\"sk\",\"description\":\"Email\",\"maxLength\":\"250\",\"minLength\":\"1\",\"name\":\"email\",\"required\":false,\"type\":\"string\",\"value\":null}]}",
				"type":  "json",
			}),
		),
	}

	jsonObjectType := resource.TestStep{
		Config: testAccResourceConnection_PropertyDataTypesJsonObject_HCL(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "property.#", "1"),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "openId",
				"value": "{\"properties\":{\"clientId\":{\"displayName\":\"Client ID\",\"placeholder\":\"\",\"preferredControlType\":\"textField\",\"required\":true,\"type\":\"string\",\"value\":\"test\"},\"clientSecret\":{\"displayName\":\"Client Secret\",\"preferredControlType\":\"textField\",\"required\":true,\"secure\":true,\"type\":\"string\",\"value\":\"test\"},\"issuerUrl\":{\"displayName\":\"Base URL\",\"preferredControlType\":\"textField\",\"required\":true,\"type\":\"string\",\"value\":\"https://ping-eng.com\"},\"returnToUrl\":{\"displayName\":\"Application Return To URL\",\"info\":\"When using the embedded flow player widget and an IDP/Social Login connector, provide a callback URL to return back to the application.\",\"preferredControlType\":\"textField\",\"value\":\"https://ping-eng.com/callback\"},\"scope\":{\"displayName\":\"Scope\",\"preferredControlType\":\"textField\",\"required\":true,\"requiredValue\":\"openid\",\"type\":\"string\",\"value\":\"openid\"},\"skRedirectUri\":{\"copyToClip\":true,\"disabled\":true,\"displayName\":\"Redirect URL\",\"info\":\"Enter this in your identity provider configuration to allow it to redirect the browser back to DaVinci. If you use a custom PingOne domain, modify the URL accordingly.\",\"initializeValue\":\"SINGULARKEY_REDIRECT_URI\",\"preferredControlType\":\"textField\",\"type\":\"string\"}}}",
				"type":  "json",
			}),
		),
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Connection_CheckDestroy(),
		Steps: []resource.TestStep{
			// mixed types (string, number, boolean)
			mixedType,
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
			},
			{
				Config:  testAccResourceConnection_PropertyDataTypesMixed_HCL(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// JSON array type
			jsonArrayType,
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
			},
			{
				Config:  testAccResourceConnection_PropertyDataTypesJsonArray_HCL(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// JSON object type
			jsonObjectType,
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
					"property.0.value",
				},
			},
			{
				Config:  testAccResourceConnection_PropertyDataTypesJsonObject_HCL(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
		},
	})
}

func TestAccResourceConnection_BadParameters(t *testing.T) {

	resourceBase := "davinci_connection"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)

	hcl := testAccResourceConnection_Full_HCL1(resourceName, resourceName, false)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProtoV6ProviderFactories: acctest.ProtoV6ProviderFactories,
		ExternalProviders:        acctest.ExternalProviders,
		ErrorCheck:               acctest.ErrorCheck(t),
		CheckDestroy:             davinci.Connection_CheckDestroy(),
		Steps: []resource.TestStep{
			// Configure
			{
				Config: hcl,
			},
			// Errors
			{
				ResourceName: resourceFullName,
				ImportState:  true,
				ExpectError:  regexp.MustCompile(`Invalid import ID specified \(".*"\).  The ID should be in the format "environment_id/davinci_connection_id" and must match regex: .*`),
			},
			{
				ResourceName:  resourceFullName,
				ImportStateId: "/",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`Invalid import ID specified \(".*"\).  The ID should be in the format "environment_id/davinci_connection_id" and must match regex: .*`),
			},
			{
				ResourceName:  resourceFullName,
				ImportStateId: "badformat/badformat",
				ImportState:   true,
				ExpectError:   regexp.MustCompile(`Invalid import ID specified \(".*"\).  The ID should be in the format "environment_id/davinci_connection_id" and must match regex: .*`),
			},
		},
	})
}

func testAccResourceConnection_Full_HCL1(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "pingOneMfaConnector"
  name           = "%[3]s"
  property {
    name  = "region"
    value = "NA"
  }
  property {
    name  = "envId"
    value = "env-abc-123"
  }
  property {
    name  = "clientId"
    value = "env-client-id"
  }
  property {
    name  = "clientSecret"
    value = "env-client-secret"
  }
  property {
    name  = "policyId"
    value = "policy-abc-123"
  }
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceConnection_Full_HCL2(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "pingOneMfaConnector"
  name           = "%[3]s"
  property {
    name  = "region"
    value = "EU"
  }
  property {
    name  = "envId"
    value = "env-abc-123"
  }
  property {
    name  = "clientId"
    value = "env-client-id"
  }
  property {
    name  = "clientSecret"
    value = "env-client-secret"
  }
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceConnection_Minimal_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "annotationConnector"
  name           = "%[3]s"
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceConnection_PropertyDataTypesMixed_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "smtpConnector"
  name           = "%[3]s"

  property {
    name  = "name"
    value = "test"
    type  = "string"
  }

  property {
    name  = "hostname"
    value = "localhost"
    type  = "string"
  }

  property {
    name  = "port"
    value = "2525"
    type  = "number"
  }

  property {
    name  = "secureFlag"
    value = "true"
    type  = "boolean"
  }

  property {
    name  = "username"
    value = "test"
    type  = "string"
  }

  property {
    name  = "password"
    value = "test"
    type  = "string"
  }
}`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceConnection_PropertyDataTypesJsonArray_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "skUserPool"
  name           = "%[3]s"

  property {
    name = "customAttributes"
    value = jsonencode({
		"type": "array",
		"preferredControlType": "tableViewAttributes",
		"sections": [
			"connectorAttributes"
		],
		"value": [
      {
        "name" : "username",
        "description" : "Username",
        "type" : "string",
        "value" : null,
        "minLength" : "1",
        "maxLength" : "300",
        "required" : true,
        "attributeType" : "sk"
      },
      {
        "name" : "firstName",
        "description" : "First Name",
        "type" : "string",
        "value" : null,
        "minLength" : "1",
        "maxLength" : "100",
        "required" : false,
        "attributeType" : "sk"
      },
      {
        "name" : "lastName",
        "description" : "Last Name",
        "type" : "string",
        "value" : null,
        "minLength" : "1",
        "maxLength" : "100",
        "required" : false,
        "attributeType" : "sk"
      },
      {
        "name" : "name",
        "description" : "Display Name",
        "type" : "string",
        "value" : null,
        "minLength" : "1",
        "maxLength" : "250",
        "required" : false,
        "attributeType" : "sk"
      },
      {
        "name" : "email",
        "description" : "Email",
        "type" : "string",
        "value" : null,
        "minLength" : "1",
        "maxLength" : "250",
        "required" : false,
        "attributeType" : "sk"
      }
    ]
})
    type = "json"
  }
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceConnection_PropertyDataTypesJsonObject_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "pingFederateConnectorV2"
  name           = "%[3]s"

  property {
    name = "openId"
    value = jsonencode({
		"properties": {
		  "skRedirectUri": {
			"type": "string",
			"displayName": "Redirect URL",
			"info": "Enter this in your identity provider configuration to allow it to redirect the browser back to DaVinci. If you use a custom PingOne domain, modify the URL accordingly.",
			"preferredControlType": "textField",
			"disabled": true,
			"initializeValue": "SINGULARKEY_REDIRECT_URI",
			"copyToClip": true
		  },
		  "clientId": {
			"type": "string",
			"displayName": "Client ID",
			"placeholder": "",
			"preferredControlType": "textField",
			"required": true,
			"value": "test"
		  },
		  "clientSecret": {
			"type": "string",
			"displayName": "Client Secret",
			"preferredControlType": "textField",
			"secure": true,
			"required": true,
			"value": "test"
		  },
		  "scope": {
			"type": "string",
			"displayName": "Scope",
			"preferredControlType": "textField",
			"requiredValue": "openid",
			"value": "openid",
			"required": true
		  },
		  "issuerUrl": {
			"type": "string",
			"displayName": "Base URL",
			"preferredControlType": "textField",
			"value": "https://ping-eng.com",
			"required": true
		  },
		  "returnToUrl": {
			"displayName": "Application Return To URL",
			"preferredControlType": "textField",
			"info": "When using the embedded flow player widget and an IDP/Social Login connector, provide a callback URL to return back to the application.",
			"value": "https://ping-eng.com/callback"
		  }
		}
	  })
    type = "json"
  }
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}
