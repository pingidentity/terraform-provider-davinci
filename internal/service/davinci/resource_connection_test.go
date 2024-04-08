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

	jsonCustomAttributesType := resource.TestStep{
		Config: testAccResourceConnection_PropertyDataTypesJsonCustomAttributes_HCL(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "property.#", "1"),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "customAttributes",
				"value": "{\"preferredControlType\":\"tableViewAttributes\",\"sections\":[\"connectorAttributes\"],\"type\":\"array\",\"value\":[{\"attributeType\":\"sk\",\"description\":\"Username\",\"maxLength\":\"300\",\"minLength\":\"1\",\"name\":\"username\",\"required\":true,\"type\":\"string\",\"value\":null},{\"attributeType\":\"sk\",\"description\":\"First Name\",\"maxLength\":\"100\",\"minLength\":\"1\",\"name\":\"firstName\",\"required\":false,\"type\":\"string\",\"value\":null},{\"attributeType\":\"sk\",\"description\":\"Last Name\",\"maxLength\":\"100\",\"minLength\":\"1\",\"name\":\"lastName\",\"required\":false,\"type\":\"string\",\"value\":null},{\"attributeType\":\"sk\",\"description\":\"Display Name\",\"maxLength\":\"250\",\"minLength\":\"1\",\"name\":\"name\",\"required\":false,\"type\":\"string\",\"value\":null},{\"attributeType\":\"sk\",\"description\":\"Email\",\"maxLength\":\"250\",\"minLength\":\"1\",\"name\":\"email\",\"required\":false,\"type\":\"string\",\"value\":null}]}",
				"type":  "json",
			}),
		),
	}

	jsonOpenIDType := resource.TestStep{
		Config: testAccResourceConnection_PropertyDataTypesJsonOpenID_HCL(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "property.#", "1"),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "openId",
				"value": "{\"properties\":{\"clientId\":{\"displayName\":\"Client ID\",\"placeholder\":\"\",\"preferredControlType\":\"textField\",\"required\":true,\"type\":\"string\",\"value\":\"test\"},\"clientSecret\":{\"displayName\":\"Client Secret\",\"preferredControlType\":\"textField\",\"required\":true,\"secure\":true,\"type\":\"string\",\"value\":\"test\"},\"issuerUrl\":{\"displayName\":\"Base URL\",\"preferredControlType\":\"textField\",\"required\":true,\"type\":\"string\",\"value\":\"https://ping-eng.com\"},\"returnToUrl\":{\"displayName\":\"Application Return To URL\",\"info\":\"When using the embedded flow player widget and an IDP/Social Login connector, provide a callback URL to return back to the application.\",\"preferredControlType\":\"textField\",\"value\":\"https://ping-eng.com/callback\"},\"scope\":{\"displayName\":\"Scope\",\"preferredControlType\":\"textField\",\"required\":true,\"requiredValue\":\"openid\",\"type\":\"string\",\"value\":\"openid\"},\"skRedirectUri\":{\"copyToClip\":true,\"disabled\":true,\"displayName\":\"Redirect URL\",\"info\":\"Enter this in your identity provider configuration to allow it to redirect the browser back to DaVinci. If you use a custom PingOne domain, modify the URL accordingly.\",\"initializeValue\":\"SINGULARKEY_REDIRECT_URI\",\"preferredControlType\":\"textField\",\"type\":\"string\"}}}",
				"type":  "json",
			}),
		),
	}

	jsonCustomAuthType := resource.TestStep{
		Config: testAccResourceConnection_PropertyDataTypesJsonCustomAuth_HCL(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "property.#", "1"),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "customAuth",
				"value": "{\"properties\":{\"authTypeDropdown\":{\"displayName\":\"Auth Type\",\"enum\":[\"oauth2\",\"openId\"],\"options\":[{\"name\":\"Oauth2\",\"value\":\"oauth2\"},{\"name\":\"OpenId\",\"value\":\"openId\"}],\"preferredControlType\":\"dropDown\",\"required\":true,\"value\":\"oauth2\"},\"authorizationEndpoint\":{\"displayName\":\"Authorization Endpoint\",\"preferredControlType\":\"textField\",\"required\":true,\"value\":\"fdsfs\"},\"bearerToken\":{\"displayName\":\"Token Attachment\",\"info\":\"Optional field. Prepend token with this value. Example: Bearer or Token\",\"preferredControlType\":\"textField\",\"type\":\"boolean\"},\"clientId\":{\"displayName\":\"App ID\",\"preferredControlType\":\"textField\",\"required\":true,\"value\":\"fdsfs\"},\"clientSecret\":{\"displayName\":\"Client Secret\",\"preferredControlType\":\"textField\",\"required\":true,\"secure\":true,\"value\":\"testDummySecret\"},\"code\":{\"displayName\":\"User Info Post Process\",\"info\":\"This code will run to simplify the response from the connector while logging in.\",\"language\":\"javascript\",\"preferredControlType\":\"codeEditor\",\"value\":\"test\"},\"customAttributes\":{\"displayName\":\"Connector Attributes\",\"info\":\"These attributes will be available in User Connector Attribute Mapping.\",\"preferredControlType\":\"tableViewAttributes\",\"sections\":[\"connectorAttributes\"],\"type\":\"array\",\"value\":[{\"attributeType\":\"sk\",\"description\":\"ID\",\"maxLength\":\"300\",\"minLength\":\"1\",\"name\":\"id\",\"required\":true,\"type\":\"string\",\"value\":null},{\"attributeType\":\"sk\",\"description\":\"Display Name\",\"maxLength\":\"250\",\"minLength\":\"1\",\"name\":\"name\",\"required\":false,\"type\":\"string\",\"value\":null},{\"attributeType\":\"sk\",\"description\":\"Email\",\"maxLength\":\"250\",\"minLength\":\"1\",\"name\":\"email\",\"required\":false,\"type\":\"string\",\"value\":null}]},\"issuerUrl\":{\"displayName\":\"Issuer URL\",\"info\":\"Required if auth type is OpenID\",\"preferredControlType\":\"textField\",\"value\":\"fdsfs\"},\"providerName\":{\"displayName\":\"Provider Name\",\"preferredControlType\":\"textField\",\"required\":true,\"value\":\"fdfs\"},\"returnToUrl\":{\"displayName\":\"Application Return To URL\",\"info\":\"When using the embedded flow player widget and an IdP/Social Login connector, provide a callback URL to return back to the application.\",\"preferredControlType\":\"textField\",\"value\":\"test\"},\"scope\":{\"displayName\":\"Scope\",\"preferredControlType\":\"textField\",\"required\":true,\"value\":\"myscope\"},\"skRedirectUri\":{\"copyToClip\":true,\"disabled\":true,\"displayName\":\"Redirect URL\",\"initializeValue\":\"SINGULARKEY_REDIRECT_URI\",\"preferredControlType\":\"textField\"},\"tokenEndpoint\":{\"displayName\":\"Token Endpoint\",\"preferredControlType\":\"textField\",\"required\":true,\"value\":\"fdsfs\"},\"userConnectorAttributeMapping\":{\"newMappingAllowed\":true,\"preferredControlType\":\"userConnectorAttributeMapping\",\"sections\":[\"attributeMapping\"],\"title1\":null,\"title2\":null,\"type\":\"object\"},\"userInfoEndpoint\":{\"displayName\":\"User Info Endpoint\",\"preferredControlType\":\"textFieldArrayView\",\"required\":true,\"value\":[\"fdsdsfs\"]}}}",
				"type":  "json",
			}),
		),
	}

	jsonOAuth2Type := resource.TestStep{
		Config: testAccResourceConnection_PropertyDataTypesJsonOAuth2_HCL(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "property.#", "1"),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "oauth2",
				"value": "{\"properties\":{\"clientId\":{\"displayName\":\"Application ID\",\"preferredControlType\":\"textField\",\"required\":true,\"type\":\"string\",\"value\":\"applicationID\"},\"clientSecret\":{\"displayName\":\"Client Secret\",\"preferredControlType\":\"textField\",\"required\":true,\"secure\":true,\"type\":\"string\",\"value\":\"dummyClinetSecret\"},\"customAttributes\":{\"displayName\":\"Connector Attributes\",\"info\":\"These attributes will be available in User Connector Attribute Mapping.\",\"preferredControlType\":\"tableViewAttributes\",\"sections\":[\"connectorAttributes\"],\"type\":\"array\",\"value\":[{\"attributeType\":\"sk\",\"description\":\"ID\",\"maxLength\":\"300\",\"minLength\":\"1\",\"name\":\"id\",\"required\":true,\"type\":\"string\",\"value\":null},{\"attributeType\":\"sk\",\"description\":\"Display Name\",\"maxLength\":\"250\",\"minLength\":\"1\",\"name\":\"name\",\"required\":false,\"type\":\"string\",\"value\":null},{\"attributeType\":\"sk\",\"description\":\"Email\",\"maxLength\":\"250\",\"minLength\":\"1\",\"name\":\"email\",\"required\":false,\"type\":\"string\",\"value\":null}]},\"disableCreateUser\":{\"displayName\":\"Disable Shadow User Creation\",\"info\":\"A shadow user is implicitly created, unless disabled.\",\"preferredControlType\":\"toggleSwitch\",\"value\":true},\"providerName\":{\"displayName\":\"Provider Name\",\"preferredControlType\":\"textField\",\"type\":\"string\",\"value\":\"Login with GitHub\"},\"returnToUrl\":{\"displayName\":\"Application Return To URL\",\"info\":\"When using the embedded flow player widget and an IDP/Social Login connector, provide a callback URL to return back to the application.\",\"preferredControlType\":\"textField\",\"value\":\"https://ping-eng.com/callback\"},\"scope\":{\"displayName\":\"Scope\",\"preferredControlType\":\"textField\",\"required\":true,\"requiredValue\":\"email\",\"type\":\"string\",\"value\":\"myscope\"},\"skRedirectUri\":{\"copyToClip\":true,\"disabled\":true,\"displayName\":\"DaVinci Redirect URL\",\"info\":\"Enter this in your identity provider configuration to allow it to redirect the browser back to DaVinci. If you use a custom PingOne domain, modify the URL accordingly.\",\"initializeValue\":\"SINGULARKEY_REDIRECT_URI\",\"preferredControlType\":\"textField\",\"type\":\"string\"},\"userConnectorAttributeMapping\":{\"newMappingAllowed\":true,\"preferredControlType\":\"userConnectorAttributeMapping\",\"sections\":[\"attributeMapping\"],\"title1\":null,\"title2\":null,\"type\":\"object\",\"value\":{\"mapping\":{\"email\":{\"value1\":\"email\"},\"name\":{\"value1\":\"name\"},\"username\":{\"value1\":\"id\"}},\"userPoolConnectionId\":\"defaultUserPool\"}}}}",
				"type":  "json",
			}),
		),
	}

	jsonSAMLType := resource.TestStep{
		Config: testAccResourceConnection_PropertyDataTypesJsonSAML_HCL(resourceName, name, withBootstrapConfig),
		Check: resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttr(resourceFullName, "property.#", "1"),
			resource.TestCheckTypeSetElemNestedAttrs(resourceFullName, "property.*", map[string]string{
				"name":  "saml",
				"value": "{\"properties\":{\"customAttributes\":{\"displayName\":\"Connector Attributes\",\"info\":\"Add the attributes that you expect to receive from the identity provider. This allows you to map them on the Attribute Mapping tab.\",\"preferredControlType\":\"tableViewAttributes\",\"sections\":[\"connectorAttributes\"],\"type\":\"array\",\"value\":[{\"attributeType\":\"sk\",\"description\":\"Subject\",\"maxLength\":\"300\",\"minLength\":\"1\",\"name\":\"saml_subject\",\"required\":true,\"type\":\"string\",\"value\":null}]},\"dvSamlSpMetadataUrl\":{\"copyToClip\":true,\"disabled\":true,\"displayName\":\"DaVinci SAML SP Metadata URL\",\"info\":\"Your DaVinci SAML SP Metadata URL. This allows an identity provider to redirect the browser back to DaVinci.\",\"initializeValue\":\"DAVINCI_SAML_SP_METADATA_URI\",\"preferredControlType\":\"textField\"},\"metadataXml\":{\"displayName\":\"Identity Provider SAML Metadata\",\"info\":\"Paste the SAML metadata provided by the identity provider.\",\"preferredControlType\":\"textArea\",\"type\":\"string\",\"value\":\"metadata\"},\"providerName\":{\"displayName\":\"Provider Name\",\"type\":\"string\"},\"returnToUrl\":{\"displayName\":\"Application Redirect URL\",\"info\":\"Your application's redirect URL, such as \\\"https://app.yourorganization.com/\\\". Enter this URL if you embed the DaVinci widget in your application. This allows DaVinci to redirect the browser back to your application.\",\"preferredControlType\":\"textField\",\"value\":\"https://ping-eng.com/callback\"},\"userConnectorAttributeMapping\":{\"newMappingAllowed\":true,\"preferredControlType\":\"userConnectorAttributeMapping\",\"sections\":[\"attributeMapping\"],\"title1\":\"Identity Provider Attributes\",\"title2\":null,\"type\":\"object\",\"value\":{\"mapping\":{\"username\":{\"value1\":\"saml_subject\"}},\"userPoolConnectionId\":\"defaultUserPool\"}}}}",
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
			// JSON custom attributes type
			jsonCustomAttributesType,
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
				Config:  testAccResourceConnection_PropertyDataTypesJsonCustomAttributes_HCL(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// JSON OpenID type
			jsonOpenIDType,
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
				Config:  testAccResourceConnection_PropertyDataTypesJsonOpenID_HCL(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// JSON custom auth type
			jsonCustomAuthType,
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
				Config:  testAccResourceConnection_PropertyDataTypesJsonCustomAuth_HCL(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// JSON oauth2 type
			jsonOAuth2Type,
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
				Config:  testAccResourceConnection_PropertyDataTypesJsonOAuth2_HCL(resourceName, name, withBootstrapConfig),
				Destroy: true,
			},
			// JSON saml type
			jsonSAMLType,
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
				Config:  testAccResourceConnection_PropertyDataTypesJsonSAML_HCL(resourceName, name, withBootstrapConfig),
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

func testAccResourceConnection_PropertyDataTypesJsonCustomAttributes_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "skUserPool"
  name           = "%[3]s"

  property {
    name = "customAttributes"
    value = jsonencode({
      "type" : "array",
      "preferredControlType" : "tableViewAttributes",
      "sections" : [
        "connectorAttributes"
      ],
      "value" : [
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

func testAccResourceConnection_PropertyDataTypesJsonOpenID_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "pingFederateConnectorV2"
  name           = "%[3]s"

  property {
    name = "openId"
    value = jsonencode({
      "properties" : {
        "skRedirectUri" : {
          "type" : "string",
          "displayName" : "Redirect URL",
          "info" : "Enter this in your identity provider configuration to allow it to redirect the browser back to DaVinci. If you use a custom PingOne domain, modify the URL accordingly.",
          "preferredControlType" : "textField",
          "disabled" : true,
          "initializeValue" : "SINGULARKEY_REDIRECT_URI",
          "copyToClip" : true
        },
        "clientId" : {
          "type" : "string",
          "displayName" : "Client ID",
          "placeholder" : "",
          "preferredControlType" : "textField",
          "required" : true,
          "value" : "test"
        },
        "clientSecret" : {
          "type" : "string",
          "displayName" : "Client Secret",
          "preferredControlType" : "textField",
          "secure" : true,
          "required" : true,
          "value" : "test"
        },
        "scope" : {
          "type" : "string",
          "displayName" : "Scope",
          "preferredControlType" : "textField",
          "requiredValue" : "openid",
          "value" : "openid",
          "required" : true
        },
        "issuerUrl" : {
          "type" : "string",
          "displayName" : "Base URL",
          "preferredControlType" : "textField",
          "value" : "https://ping-eng.com",
          "required" : true
        },
        "returnToUrl" : {
          "displayName" : "Application Return To URL",
          "preferredControlType" : "textField",
          "info" : "When using the embedded flow player widget and an IDP/Social Login connector, provide a callback URL to return back to the application.",
          "value" : "https://ping-eng.com/callback"
        }
      }
    })
    type = "json"
  }
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceConnection_PropertyDataTypesJsonCustomAuth_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "genericConnector"
  name           = "%[3]s"

  property {
    name = "customAuth"
    value = jsonencode({
      "properties" : {
        "customAttributes" : {
          "type" : "array",
          "displayName" : "Connector Attributes",
          "preferredControlType" : "tableViewAttributes",
          "info" : "These attributes will be available in User Connector Attribute Mapping.",
          "sections" : [
            "connectorAttributes"
          ],
          "value" : [
            {
              "name" : "id",
              "description" : "ID",
              "type" : "string",
              "value" : null,
              "minLength" : "1",
              "maxLength" : "300",
              "required" : true,
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
        },
        "userConnectorAttributeMapping" : {
          "type" : "object",
          "preferredControlType" : "userConnectorAttributeMapping",
          "newMappingAllowed" : true,
          "title1" : null,
          "title2" : null,
          "sections" : [
            "attributeMapping"
          ]
        },
        "providerName" : {
          "displayName" : "Provider Name",
          "preferredControlType" : "textField",
          "required" : true,
          "value" : "fdfs"
        },
        "authTypeDropdown" : {
          "displayName" : "Auth Type",
          "preferredControlType" : "dropDown",
          "required" : true,
          "options" : [
            {
              "name" : "Oauth2",
              "value" : "oauth2"
            },
            {
              "name" : "OpenId",
              "value" : "openId"
            }
          ],
          "enum" : [
            "oauth2",
            "openId"
          ],
          "value" : "oauth2"
        },
        "issuerUrl" : {
          "preferredControlType" : "textField",
          "displayName" : "Issuer URL",
          "info" : "Required if auth type is OpenID",
          "value" : "fdsfs"
        },
        "skRedirectUri" : {
          "displayName" : "Redirect URL",
          "preferredControlType" : "textField",
          "disabled" : true,
          "initializeValue" : "SINGULARKEY_REDIRECT_URI",
          "copyToClip" : true
        },
        "clientId" : {
          "displayName" : "App ID",
          "preferredControlType" : "textField",
          "required" : true,
          "value" : "fdsfs"
        },
        "clientSecret" : {
          "displayName" : "Client Secret",
          "preferredControlType" : "textField",
          "secure" : true,
          "required" : true,
          "value" : "testDummySecret"
        },
        "scope" : {
          "displayName" : "Scope",
          "preferredControlType" : "textField",
          "required" : true,
          "value" : "myscope"
        },
        "code" : {
          "displayName" : "User Info Post Process",
          "info" : "This code will run to simplify the response from the connector while logging in.",
          "preferredControlType" : "codeEditor",
          "language" : "javascript",
          "value" : "test"
        },
        "authorizationEndpoint" : {
          "displayName" : "Authorization Endpoint",
          "preferredControlType" : "textField",
          "required" : true,
          "value" : "fdsfs"
        },
        "tokenEndpoint" : {
          "displayName" : "Token Endpoint",
          "preferredControlType" : "textField",
          "required" : true,
          "value" : "fdsfs"
        },
        "bearerToken" : {
          "preferredControlType" : "textField",
          "type" : "boolean",
          "displayName" : "Token Attachment",
          "info" : "Optional field. Prepend token with this value. Example: Bearer or Token"
        },
        "userInfoEndpoint" : {
          "displayName" : "User Info Endpoint",
          "preferredControlType" : "textFieldArrayView",
          "required" : true,
          "value" : [
            "fdsdsfs"
          ]
        },
        "returnToUrl" : {
          "displayName" : "Application Return To URL",
          "preferredControlType" : "textField",
          "info" : "When using the embedded flow player widget and an IdP/Social Login connector, provide a callback URL to return back to the application.",
          "value" : "test"
        }
      }
    })
    type = "json"
  }
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceConnection_PropertyDataTypesJsonOAuth2_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "githubIdpConnector"
  name           = "%[3]s"

  property {
    name = "oauth2"
    value = jsonencode({
      "properties" : {
        "providerName" : {
          "type" : "string",
          "displayName" : "Provider Name",
          "preferredControlType" : "textField",
          "value" : "Login with GitHub"
        },
        "skRedirectUri" : {
          "type" : "string",
          "displayName" : "DaVinci Redirect URL",
          "info" : "Enter this in your identity provider configuration to allow it to redirect the browser back to DaVinci. If you use a custom PingOne domain, modify the URL accordingly.",
          "preferredControlType" : "textField",
          "disabled" : true,
          "initializeValue" : "SINGULARKEY_REDIRECT_URI",
          "copyToClip" : true
        },
        "clientId" : {
          "type" : "string",
          "displayName" : "Application ID",
          "preferredControlType" : "textField",
          "required" : true,
          "value" : "applicationID"
        },
        "clientSecret" : {
          "type" : "string",
          "displayName" : "Client Secret",
          "preferredControlType" : "textField",
          "secure" : true,
          "required" : true,
          "value" : "dummyClinetSecret"
        },
        "scope" : {
          "type" : "string",
          "displayName" : "Scope",
          "preferredControlType" : "textField",
          "requiredValue" : "email",
          "required" : true,
          "value" : "myscope"
        },
        "customAttributes" : {
          "type" : "array",
          "displayName" : "Connector Attributes",
          "preferredControlType" : "tableViewAttributes",
          "info" : "These attributes will be available in User Connector Attribute Mapping.",
          "sections" : [
            "connectorAttributes"
          ],
          "value" : [
            {
              "name" : "id",
              "description" : "ID",
              "type" : "string",
              "value" : null,
              "minLength" : "1",
              "maxLength" : "300",
              "required" : true,
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
        },
        "userConnectorAttributeMapping" : {
          "type" : "object",
          "preferredControlType" : "userConnectorAttributeMapping",
          "newMappingAllowed" : true,
          "title1" : null,
          "title2" : null,
          "sections" : [
            "attributeMapping"
          ],
          "value" : {
            "userPoolConnectionId" : "defaultUserPool",
            "mapping" : {
              "username" : {
                "value1" : "id"
              },
              "name" : {
                "value1" : "name"
              },
              "email" : {
                "value1" : "email"
              }
            }
          }
        },
        "disableCreateUser" : {
          "displayName" : "Disable Shadow User Creation",
          "preferredControlType" : "toggleSwitch",
          "value" : true,
          "info" : "A shadow user is implicitly created, unless disabled."
        },
        "returnToUrl" : {
          "displayName" : "Application Return To URL",
          "preferredControlType" : "textField",
          "info" : "When using the embedded flow player widget and an IDP/Social Login connector, provide a callback URL to return back to the application.",
          "value" : "https://ping-eng.com/callback"
        }
      }
    })
    type = "json"
  }
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccResourceConnection_PropertyDataTypesJsonSAML_HCL(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "samlIdpConnector"
  name           = "%[3]s"

  property {
    name = "saml"
    value = jsonencode({
      "properties" : {
        "dvSamlSpMetadataUrl" : {
          "displayName" : "DaVinci SAML SP Metadata URL",
          "info" : "Your DaVinci SAML SP Metadata URL. This allows an identity provider to redirect the browser back to DaVinci.",
          "preferredControlType" : "textField",
          "disabled" : true,
          "initializeValue" : "DAVINCI_SAML_SP_METADATA_URI",
          "copyToClip" : true
        },
        "providerName" : {
          "type" : "string",
          "displayName" : "Provider Name"
        },
        "metadataXml" : {
          "type" : "string",
          "displayName" : "Identity Provider SAML Metadata",
          "info" : "Paste the SAML metadata provided by the identity provider.",
          "preferredControlType" : "textArea",
          "value" : "metadata"
        },
        "returnToUrl" : {
          "displayName" : "Application Redirect URL",
          "preferredControlType" : "textField",
          "info" : "Your application's redirect URL, such as \"https://app.yourorganization.com/\". Enter this URL if you embed the DaVinci widget in your application. This allows DaVinci to redirect the browser back to your application.",
          "value" : "https://ping-eng.com/callback"
        },
        "userConnectorAttributeMapping" : {
          "type" : "object",
          "preferredControlType" : "userConnectorAttributeMapping",
          "newMappingAllowed" : true,
          "title1" : "Identity Provider Attributes",
          "title2" : null,
          "sections" : [
            "attributeMapping"
          ],
          "value" : {
            "userPoolConnectionId" : "defaultUserPool",
            "mapping" : {
              "username" : {
                "value1" : "saml_subject"
              }
            }
          }
        },
        "customAttributes" : {
          "type" : "array",
          "displayName" : "Connector Attributes",
          "preferredControlType" : "tableViewAttributes",
          "info" : "Add the attributes that you expect to receive from the identity provider. This allows you to map them on the Attribute Mapping tab.",
          "sections" : [
            "connectorAttributes"
          ],
          "value" : [
            {
              "name" : "saml_subject",
              "description" : "Subject",
              "type" : "string",
              "value" : null,
              "minLength" : "1",
              "maxLength" : "300",
              "required" : true,
              "attributeType" : "sk"
            }
          ]
        }
      }
    })
    type = "json"
  }
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}
