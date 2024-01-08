package acctest

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/patrickcping/pingone-go-sdk-v2/management"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest/pingone"
	"github.com/pingidentity/terraform-provider-davinci/internal/provider"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
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

// TestCheckFunc is used by CheckDestroy to determine for proper resource destruction
var TestCheckFunc func(*terraform.State) error

func init() {
	Provider = provider.New("dev")()

	// Always allocate a new provider instance each invocation, otherwise gRPC
	// ProviderConfigure() can overwrite configuration during concurrent testing.
	ProviderFactories = map[string]func() (*schema.Provider, error){
		"davinci": func() (*schema.Provider, error) {
			testVersion := getProviderTestingVersion()
			if testVersion == "" {
				testVersion = "dev"
			}
			provider := provider.New(testVersion)()

			if provider == nil {
				return nil, fmt.Errorf("Cannot initiate provider factory")
			}
			return provider, nil
		},
	}

	ExternalProviders = map[string]resource.ExternalProvider{
		"pingone": {
			Source:            "pingidentity/pingone",
			VersionConstraint: ">= 0.25, < 1.0",
		},
	}

}

func getProviderTestingVersion() string {
	returnVar := "dev"
	if v := os.Getenv("PINGONE_TESTING_PROVIDER_VERSION"); v != "" {
		returnVar = v
	}
	return returnVar
}

// check required variabes are met for tests
func PreCheckDaVinciClient(t *testing.T) {
	username := os.Getenv("PINGONE_USERNAME")
	password := os.Getenv("PINGONE_PASSWORD")
	accessToken := os.Getenv("PINGONE_DAVINCI_ACCESS_TOKEN")
	if (username == "" || password == "") && accessToken == "" {
		t.Fatal("PINGONE_USERNAME and PINGONE_PASSWORD or PINGONE_DAVINCI_ACCESS_TOKEN are missing and must be set")
	}
	if v := os.Getenv("PINGONE_REGION"); v == "" {
		t.Fatal("PINGONE_REGION is missing and must be set")
	}
	if v := os.Getenv("PINGONE_ENVIRONMENT_ID"); v == "" {
		t.Fatal("PINGONE_ENVIRONMENT_ID is missing and must be set")
	}
}

func PreCheckPingOneClient(t *testing.T) {
	if v := os.Getenv("PINGONE_CLIENT_ID"); v == "" {
		t.Fatal("PINGONE_CLIENT_ID is missing and must be set")
	}
	if v := os.Getenv("PINGONE_CLIENT_SECRET"); v == "" {
		t.Fatal("PINGONE_CLIENT_SECRET is missing and must be set")
	}
	if v := os.Getenv("PINGONE_REGION"); v == "" {
		t.Fatal("PINGONE_REGION is missing and must be set")
	}
	if v := os.Getenv("PINGONE_ENVIRONMENT_ID"); v == "" {
		t.Fatal("PINGONE_ENVIRONMENT_ID is missing and must be set")
	}
}

func PreCheckClient(t *testing.T) {
	PreCheckDaVinciClient(t)
	PreCheckPingOneClient(t)
}

func PreCheckNewEnvironment(t *testing.T) {
	if v := os.Getenv("PINGONE_LICENSE_ID"); v == "" {
		t.Fatal("PINGONE_LICENSE_ID is missing and must be set")
	}
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
	id = davinci_connection.%[1]s.id
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
			id    = resource.davinci_flow.%[1]s.id
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
	return `				
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
  id = davinci_connection.crowd_strike.id
}
	`
}

func TfHclPingOneDavinci() string {
	return `				
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
  id = davinci_connection.crowd_strike.id
}`
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
// p1services is a list of services besides SSO and DaVinci to enable on the environment
func PingoneEnvironmentServicesSsoHcl(resourceName string, p1Services []string, withBootstrapConfig bool) (hcl string) {
	adminEnvID := os.Getenv("PINGONE_ENVIRONMENT_ID")
	licenseID := os.Getenv("PINGONE_LICENSE_ID")
	username := os.Getenv("PINGONE_USERNAME")

	servicesString := ""
	if len(p1Services) > 0 {
		servicesString = fmt.Sprintf("\"%s\"", strings.Join(p1Services, "\", \""))
	}

	daVinciTags := "null"
	if !withBootstrapConfig {
		daVinciTags = "[\"DAVINCI_MINIMAL\"]"
	}

	return fmt.Sprintf(`
variable "services_%[1]s" {
  type    = list(string)
  default = [%[5]s]
}

resource "pingone_environment" "%[1]s" {
  name       = "tf-testacc-dv-dynamic-%[1]s"
  license_id = "%[2]s"

  service {
    type = "SSO"
  }
  service {
    type = "DaVinci"
	tags = %[6]s
  }

  dynamic "service" {
    for_each = toset(var.services_%[1]s)

    content {
      type = service.key
    }
  }
}
`, resourceName, licenseID, username, adminEnvID, servicesString, daVinciTags)
}

// PingoneEnvironmentSsoHcl returns hcl for a pingone environment and assigns roles used for SSO by davinci
// The following environment vars must be set:
// - PINGONE_ENVIRONMENT_ID
// - PINGONE_LICENSE_ID
// - PINGONE_SSO_USERNAME
// - PINGONE_SSO_PASSWORD
//
// The `resourceName` input can be a random charset and will be used for the name of
// each resource and datasource in the returned hcl.
func PingoneEnvironmentSsoHcl(resourceName string, withBootstrapConfig bool) (hcl string) {
	return PingoneEnvironmentServicesSsoHcl(resourceName, nil, withBootstrapConfig)
}

func BaselineHcl(resourceName string) string {
	pingoneHcl := PingoneEnvironmentSsoHcl(resourceName, true)
	bsConnectionsHcl := BsConnectionsHcl(resourceName)
	return fmt.Sprintf(`
%[1]s

%[2]s
`, pingoneHcl, bsConnectionsHcl)
}

func TestClient() (*dv.APIClient, error) {
	e := ""
	username := os.Getenv("PINGONE_USERNAME")
	if username == "" {
		e = e + "PINGONE_USERNAME "
	}
	password := os.Getenv("PINGONE_PASSWORD")
	if password == "" {
		e = e + "PINGONE_PASSWORD "
	}
	region := os.Getenv("PINGONE_REGION")
	if region == "" {
		e = e + "PINGONE_REGION "
	}
	environment_id := os.Getenv("PINGONE_ENVIRONMENT_ID")
	if environment_id == "" {
		e = e + "PINGONE_ENVIRONMENT_ID "
	}
	host_url := os.Getenv("PINGONE_DAVINCI_HOST_URL")
	if e != "" {
		return nil, fmt.Errorf("missing environment variables: %s", e)
	}

	userAgent := fmt.Sprintf("terraform-provider-davinci/%s", getProviderTestingVersion())

	cInput := dv.ClientInput{
		Username:        username,
		Password:        password,
		PingOneRegion:   region,
		PingOneSSOEnvId: environment_id,
		HostURL:         host_url,
		UserAgent:       userAgent,
	}
	c, err := dv.NewClient(&cInput)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func PingOneTestClient(ctx context.Context) (*pingone.Client, error) {

	config := &pingone.Config{
		ClientID:      os.Getenv("PINGONE_CLIENT_ID"),
		ClientSecret:  os.Getenv("PINGONE_CLIENT_SECRET"),
		EnvironmentID: os.Getenv("PINGONE_ENVIRONMENT_ID"),
		Region:        os.Getenv("PINGONE_REGION"),
		ForceDelete:   false,
	}

	return config.APIClient(ctx, getProviderTestingVersion())

}

func CheckParentEnvironmentDestroy(ctx context.Context, apiClient *management.APIClient, environmentID string) (bool, error) {
	_, r, err := apiClient.EnvironmentsApi.ReadOneEnvironment(ctx, environmentID).Execute()

	return CheckForResourceDestroy(r, err)
}

func CheckForResourceDestroy(r *http.Response, err error) (bool, error) {
	defaultDestroyHttpCode := 404
	return CheckForResourceDestroyCustomHTTPCode(r, err, defaultDestroyHttpCode)
}

func CheckForResourceDestroyCustomHTTPCode(r *http.Response, err error, customHttpCode int) (bool, error) {
	if err != nil {

		if r == nil {
			return false, fmt.Errorf("Response object does not exist and no error detected")
		}

		if r.StatusCode == customHttpCode {
			return true, nil
		}

		return false, err
	}

	return false, nil
}
