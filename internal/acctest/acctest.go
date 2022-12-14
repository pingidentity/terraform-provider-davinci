package acctest

import (
	// "context"
	// "encoding/json"
	"fmt"
	// "io"
	// "log"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pingidentity/terraform-provider-davinci/internal/provider"
)

// ProviderFactories is a static map containing only the main provider instance
//
// Use other ProviderFactories functions, such as FactoriesAlternate,
// for tests requiring special provider configurations.
var ProviderFactories map[string]func() (*schema.Provider, error)

// Provider is the "main" provider instance
//
// This Provider can be used in testing code for API calls without requiring
// the use of saving and referencing specific ProviderFactories instances.
//
// PreCheck(t) must be called before using this provider instance.
var Provider *schema.Provider

// providerFactories are used to instantiate a provider during acceptance testing.
// The factory function will be invoked for every Terraform CLI command executed
// to create a provider server to which the CLI can reattach.

// ExternalProviders is a map of additional providers that may be used during
// testing.
var ExternalProviders map[string]resource.ExternalProvider

func init() {
	Provider = provider.New("dev")()

	// Always allocate a new provider instance each invocation, otherwise gRPC
	// ProviderConfigure() can overwrite configuration during concurrent testing.
	ProviderFactories = map[string]func() (*schema.Provider, error){
		"davinci": func() (*schema.Provider, error) {
			provider := provider.New("dev")()

			if provider == nil {
				return nil, fmt.Errorf("Cannot initiate provider factory")
			}
			return provider, nil
		},
	}

	ExternalProviders = map[string]resource.ExternalProvider{
		"pingone": {
			Source:            "pingidentity/pingone",
			VersionConstraint: "0.7.0",
		},
	}

}

// check required variabes are met for tests
func PreCheck(t *testing.T) {
	if v := os.Getenv("PINGONE_USERNAME"); v == "" {
		t.Fatal("PINGONE_USERNAME is missing and must be set")
	}
	if v := os.Getenv("PINGONE_PASSWORD"); v == "" {
		t.Fatal("PINGONE_PASSWORD is missing and must be set")
	}
	if v := os.Getenv("PINGONE_REGION"); v == "" {
		t.Fatal("PINGONE_REGION is missing and must be set")
	}
	if v := os.Getenv("PINGONE_ENVIRONMENT_ID"); v == "" {
		t.Fatal("PINGONE_ENVIRONMENT_ID is missing and must be set")
	}
}

func PreCheckPingOne(t *testing.T) {
	PreCheck(t)
	if v := os.Getenv("PINGONE_LICENSE_ID"); v == "" {
		t.Fatal("PINGONE_LICENSE_ID is missing and must be set")
	}
	if v := os.Getenv("PINGONE_CLIENT_ID"); v == "" {
		t.Fatal("PINGONE_CLIENT_ID is missing and must be set")
	}
	if v := os.Getenv("PINGONE_CLIENT_SECRET"); v == "" {
		t.Fatal("PINGONE_CLIENT_SECRET is missing and must be set")
	}
}

func PreCheckPingOneAndTfVars(t *testing.T) {
	PreCheckPingOne(t)
	// if v := os.Getenv("TF_VAR_environment_id"); v == "" {
	// 	t.Fatal("TF_VAR_environment_id is missing and must be set")
	// }
}

// func TestClient(ctx context.Context) (*client.APIClient, error) {

// 	cInput := client.ClientInput{
// 		Username: username,
// 		Password: password,
// 	}
// 	client, err := client.NewClient(&cInput)
// 	if companyid != "" {
// 		client.CompanyID = companyid
// 	}
// 	if err != nil {
// 		log.Fatalf("failed to make client %v: ", err)
// 	}
// 	return client, nil

// }

func ErrorCheck(t *testing.T) resource.ErrorCheckFunc {
	return func(err error) error {
		if err == nil {
			return nil
		}
		return err
	}
}

func ResourceNameGen() string {
	strlen := 10
	return acctest.RandStringFromCharSet(strlen, acctest.CharSetAlpha)
}

func RandStringFieldGen() string {
	strlen := 10
	return acctest.RandString(strlen)
}

func RandStringWithPrefix(prefix string) string {

	return acctest.RandomWithPrefix(prefix)
}

func MainTfHcl(resourceName string) (hcl string) {
	return fmt.Sprintf(`				
resource "davinci_connection" "%[1]s" {
	connector_id = "crowdStrikeConnector"
	name         = "CrowdStrike"
	properties {
		name  = "clientId"
		value = "1234"
	}
	properties {
		name  = "clientSecret"
		value = "12345"
	}
}
data "davinci_connection" %[1]s" {
	connection_id = davinci_connection.%[1]s.connection_id
}
resource "davinci_flow" "%[1]s" {
	flow_json = "{\"customerId\":\"dc7918cfa4b50966f8508072c2755c2c\",\"name\":\"tf testing-changed\",\"description\":\"\",\"flowStatus\":\"enabled\",\"createdDate\":1662960509175,\"updatedDate\":1662961640567,\"currentVersion\":0,\"authTokenExpireIds\":[],\"deployedDate\":1662960510106,\"functionConnectionId\":null,\"isOutputSchemaSaved\":false,\"outputSchemaCompiled\":null,\"publishedVersion\":1,\"timeouts\":null,\"flowId\":\"bb45eb4a0e8a5c9d6a21c7ac2d1b3faa\",\"companyId\":\"c431739a-29cd-4d9e-b465-0b37b2c235b1\",\"versionId\":0,\"graphData\":{\"elements\":{\"nodes\":[{\"data\":{\"id\":\"pptape4ac2\",\"nodeType\":\"CONNECTION\",\"connectionId\":\"867ed4363b2bc21c860085ad2baa817d\",\"connectorId\":\"httpConnector\",\"name\":\"Http\",\"label\":\"Http\",\"status\":\"configured\",\"capabilityName\":\"customHtmlMessage\",\"type\":\"trigger\",\"properties\":{\"message\":{\"value\":\"[\\n  {\\n    \\\"children\\\": [\\n      {\\n        \\\"text\\\": \\\"hello foobar\\\"\\n      }\\n    ]\\n  }\\n]\"}}},\"position\":{\"x\":570,\"y\":240},\"group\":\"nodes\",\"removed\":false,\"selected\":false,\"selectable\":true,\"locked\":false,\"grabbable\":true,\"pannable\":false,\"classes\":\"\"}]},\"data\":{},\"zoomingEnabled\":true,\"userZoomingEnabled\":true,\"zoom\":1,\"minZoom\":1e-50,\"maxZoom\":1e+50,\"panningEnabled\":true,\"userPanningEnabled\":true,\"pan\":{\"x\":0,\"y\":0},\"boxSelectionEnabled\":true,\"renderer\":{\"name\":\"null\"}},\"flowColor\":\"#AFD5FF\",\"connectorIds\":[\"httpConnector\"],\"savedDate\":1662961640542,\"variables\":[]}"
	deploy = true
}
resource "davinci_application" "%[1]s" {
	name = "tf-acc-test-%[1]s"
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
	policies {
		name = "tf-test-acc-%[1]s"
		policy_flows {
			flow_id    = resource.davinci_flow.%[1]s.flow_id
			version_id = -1
			weight     = 100
		}
	}
	saml {
		values {
			enabled                = false
			enforce_signed_request = false
		}
	}
}

data "davinci_application" "%[1]s" {
	application_id = resource.davinci_application.use_default_flow.application_id
}
	`, resourceName)
}

func MainTfHclUpdate() string {
	return fmt.Sprintf(`				
	resource "davinci_connection" "crowd_strike" {
		connector_id = "crowdStrikeConnector"
		name         = "CrowdStrike"
		properties {
			name  = "clientId"
			value = "9876"
		}
		properties {
			name  = "clientSecret"
			value = "9876"
		}
	}
	data "davinci_connection" "crowd_strike" {
		connection_id = davinci_connection.crowd_strike.connection_id
	}
	`)
}

func TfHclPingOneDavinci() string {
	return fmt.Sprintf(`				
resource "davinci_connection" "crowd_strike" {
	connector_id = "crowdStrikeConnector"
	name         = "CrowdStrike"
	properties {
		name  = "clientId"
		value = "9876"
	}
	properties {
		name  = "clientSecret"
		value = "9876"
	}
}
data "davinci_connection" "crowd_strike" {
	connection_id = davinci_connection.crowd_strike.connection_id
}

	`)
}

// PingoneEnvrionmentSsoHcl returns hcl for a pingone environment and assigns roles used for SSO by davinci
// The following environment vars must be set:
// - PINGONE_ENVIRONMENT_ID
// - PINGONE_LICENSE_ID
// - PINGONE_SSO_USERNAME
// - PINGONE_SSO_PASSWORD
//
// The `resourceName` input can be a random charset and will be used for the name of
// each resource and datasource in the returned hcl.
func PingoneEnvrionmentSsoHcl(resourceName string) (hcl string) {
	adminEnvID := os.Getenv("PINGONE_ENVIRONMENT_ID")
	licenseID := os.Getenv("PINGONE_LICENSE_ID")
	username := os.Getenv("PINGONE_USERNAME")
	return fmt.Sprintf(`
resource "pingone_environment" "%[1]s" {
	name = "tf-testacc-dynamic-%[1]s"
	type = "SANDBOX"
	license_id = "%[2]s"
	default_population {
	}
	service {
		type = "SSO"
	}
	service {
		type = "DaVinci"
	}
}

data "pingone_role" "%[1]s" {
	name = "Identity Data Admin"
}

data "pingone_user" "%[1]s"{
	username             = "%[3]s"
	environment_id       = "%[4]s"
}

resource "pingone_role_assignment_user" "%[1]s" {
	environment_id       = "%[4]s"
	user_id              = data.pingone_user.%[1]s.id
	role_id              = data.pingone_role.%[1]s.id
	scope_environment_id = resource.pingone_environment.%[1]s.id
}

`, resourceName, licenseID, username, adminEnvID)
}

func BaselineHcl(resourceName string) string {
	pingoneHcl := PingoneEnvrionmentSsoHcl(resourceName)
	bsConnectionsHcl := BsConnectionsHcl(resourceName)
	return fmt.Sprintf(`
%[1]s
data "davinci_connections" "read_all" {
	environment_id = resource.pingone_role_assignment_user.%[3]s.scope_environment_id
}

%[2]s
`, pingoneHcl, bsConnectionsHcl, resourceName)
}
