package davinci_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
)

func TestAccDataSourceConnections_AllConnections_Clean(t *testing.T) {
	testAccDataSourceConnections_AllConnections(t, false)
}

func TestAccDataSourceConnections_AllConnections_WithBootstrap(t *testing.T) {
	testAccDataSourceConnections_AllConnections(t, true)
}

func testAccDataSourceConnections_AllConnections(t *testing.T, withBootstrapConfig bool) {

	resourceBase := "davinci_connections"
	resourceName := acctest.ResourceNameGen()
	dataSourceFullName := fmt.Sprintf("data.%s.%s", resourceBase, resourceName)

	name := resourceName

	hcl := testAccDataSourceConnections_AllConnections_Hcl(resourceName, name, withBootstrapConfig)

	var connectionsCountRegex *regexp.Regexp
	if withBootstrapConfig {
		connectionsCountRegex = regexp.MustCompile(`^1[0-9]{1,2}$`)
	} else {
		connectionsCountRegex = regexp.MustCompile(`^4$`)
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		//CheckDestroy:       acctest.CheckResourceDestroy([]string{"davinci_connection"}),
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(dataSourceFullName, "id", regexp.MustCompile(fmt.Sprintf(`^id-%s-connections$`, verify.P1ResourceIDRegexp.String()))),
					resource.TestMatchResourceAttr(dataSourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
					resource.TestCheckResourceAttr(dataSourceFullName, "connector_ids.#", "0"),
					resource.TestMatchResourceAttr(dataSourceFullName, "connections.#", connectionsCountRegex),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceFullName, "connections.*", map[string]*regexp.Regexp{
						"id":               verify.P1DVResourceIDRegexpFullString,
						"connector_id":     regexp.MustCompile(`^httpConnector$`),
						"company_id":       verify.P1ResourceIDRegexpFullString,
						"customer_id":      verify.P1DVResourceIDRegexpFullString,
						"name":             regexp.MustCompile(fmt.Sprintf(`^%s-1$`, name)),
						"created_date":     verify.EpochDateRegexpFullString,
						"property.#":       regexp.MustCompile(`^2$`),
						"property.0.name":  regexp.MustCompile(`^recaptchaSecretKey$`),
						"property.0.type":  regexp.MustCompile(`^string$`),
						"property.0.value": regexp.MustCompile(`^\*{6}$`),
						"property.1.name":  regexp.MustCompile(`^recaptchaSiteKey$`),
						"property.1.type":  regexp.MustCompile(`^string$`),
						"property.1.value": regexp.MustCompile(`^test2$`),
					}),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceFullName, "connections.*", map[string]*regexp.Regexp{
						"id":           verify.P1DVResourceIDRegexpFullString,
						"connector_id": regexp.MustCompile(`^httpConnector$`),
						"company_id":   verify.P1ResourceIDRegexpFullString,
						"customer_id":  verify.P1DVResourceIDRegexpFullString,
						"name":         regexp.MustCompile(fmt.Sprintf(`^%s-2$`, name)),
						"created_date": verify.EpochDateRegexpFullString,
						"property.#":   regexp.MustCompile(`^0$`),
					}),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceFullName, "connections.*", map[string]*regexp.Regexp{
						"id":           verify.P1DVResourceIDRegexpFullString,
						"connector_id": regexp.MustCompile(`^annotationConnector$`),
						"company_id":   verify.P1ResourceIDRegexpFullString,
						"customer_id":  verify.P1DVResourceIDRegexpFullString,
						"name":         regexp.MustCompile(fmt.Sprintf(`^%s-3$`, name)),
						"created_date": verify.EpochDateRegexpFullString,
						"property.#":   regexp.MustCompile(`^0$`),
					}),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceFullName, "connections.*", map[string]*regexp.Regexp{
						"id":           regexp.MustCompile(`^defaultUserPool$`),
						"connector_id": regexp.MustCompile(`^skUserPool$`),
						"company_id":   verify.P1ResourceIDRegexpFullString,
						"customer_id":  verify.P1DVResourceIDRegexpFullString,
						"name":         regexp.MustCompile(`^User Pool$`),
						"created_date": verify.EpochDateRegexpFullString,
						"property.#":   regexp.MustCompile(`^0$`),
					}),
				),
			},
		},
	})
}

func TestAccDataSourceConnections_FilteredConnections_Clean(t *testing.T) {
	testAccDataSourceConnections_FilteredConnections(t, false)
}

func TestAccDataSourceConnections_FilteredConnections_WithBootstrap(t *testing.T) {
	testAccDataSourceConnections_FilteredConnections(t, true)
}

func testAccDataSourceConnections_FilteredConnections(t *testing.T, withBootstrapConfig bool) {

	resourceBase := "davinci_connections"
	resourceName := acctest.ResourceNameGen()
	dataSourceFullName := fmt.Sprintf("data.%s.%s", resourceBase, resourceName)

	name := resourceName

	hcl := testAccDataSourceConnections_FilteredConnections_Hcl(resourceName, name, withBootstrapConfig)

	var connectionsCountRegex *regexp.Regexp
	if withBootstrapConfig {
		connectionsCountRegex = regexp.MustCompile(`^3$`)
	} else {
		connectionsCountRegex = regexp.MustCompile(`^2$`)
	}

	resource.ParallelTest(t, resource.TestCase{
		PreCheck: func() {
			acctest.PreCheckClient(t)
			acctest.PreCheckNewEnvironment(t)
		},
		ProviderFactories: acctest.ProviderFactories,
		ExternalProviders: acctest.ExternalProviders,
		ErrorCheck:        acctest.ErrorCheck(t),
		//CheckDestroy:       acctest.CheckResourceDestroy([]string{"davinci_connection"}),
		Steps: []resource.TestStep{
			{
				Config: hcl,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr(dataSourceFullName, "id", regexp.MustCompile(fmt.Sprintf(`^id-%s-connections$`, verify.P1ResourceIDRegexp.String()))),
					resource.TestMatchResourceAttr(dataSourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
					resource.TestCheckResourceAttr(dataSourceFullName, "connector_ids.#", "1"),
					resource.TestMatchResourceAttr(dataSourceFullName, "connections.#", connectionsCountRegex),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceFullName, "connections.*", map[string]*regexp.Regexp{
						"id":               verify.P1DVResourceIDRegexpFullString,
						"connector_id":     regexp.MustCompile(`^httpConnector$`),
						"company_id":       verify.P1ResourceIDRegexpFullString,
						"customer_id":      verify.P1DVResourceIDRegexpFullString,
						"name":             regexp.MustCompile(fmt.Sprintf(`^%s-1$`, name)),
						"created_date":     verify.EpochDateRegexpFullString,
						"property.#":       regexp.MustCompile(`^2$`),
						"property.0.name":  regexp.MustCompile(`^recaptchaSecretKey$`),
						"property.0.type":  regexp.MustCompile(`^string$`),
						"property.0.value": regexp.MustCompile(`^\*{6}$`),
						"property.1.name":  regexp.MustCompile(`^recaptchaSiteKey$`),
						"property.1.type":  regexp.MustCompile(`^string$`),
						"property.1.value": regexp.MustCompile(`^test2$`),
					}),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceFullName, "connections.*", map[string]*regexp.Regexp{
						"id":           verify.P1DVResourceIDRegexpFullString,
						"connector_id": regexp.MustCompile(`^httpConnector$`),
						"company_id":   verify.P1ResourceIDRegexpFullString,
						"customer_id":  verify.P1DVResourceIDRegexpFullString,
						"name":         regexp.MustCompile(fmt.Sprintf(`^%s-2$`, name)),
						"created_date": verify.EpochDateRegexpFullString,
						"property.#":   regexp.MustCompile(`^0$`),
					}),
				),
			},
		},
	})
}

func testAccDataSourceConnections_AllConnections_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s-1" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "httpConnector"
  name           = "%[3]s-1"
  property {
    name  = "recaptchaSecretKey"
    value = "test"
  }
  property {
    name  = "recaptchaSiteKey"
    value = "test2"
  }
}

resource "davinci_connection" "%[2]s-2" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "httpConnector"
  name           = "%[3]s-2"
}

resource "davinci_connection" "%[2]s-3" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "annotationConnector"
  name           = "%[3]s-3"
}

data "davinci_connections" "%[2]s" {
  environment_id = resource.pingone_environment.%[2]s.id

  depends_on = [
    davinci_connection.%[2]s-1,
    davinci_connection.%[2]s-2,
    davinci_connection.%[2]s-3,
  ]
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccDataSourceConnections_FilteredConnections_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

resource "davinci_connection" "%[2]s-1" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "httpConnector"
  name           = "%[3]s-1"
  property {
    name  = "recaptchaSecretKey"
    value = "test"
  }
  property {
    name  = "recaptchaSiteKey"
    value = "test2"
  }
}

resource "davinci_connection" "%[2]s-2" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "httpConnector"
  name           = "%[3]s-2"
}

resource "davinci_connection" "%[2]s-3" {
  environment_id = pingone_environment.%[2]s.id
  connector_id   = "annotationConnector"
  name           = "%[3]s-3"
}

data "davinci_connections" "%[2]s" {
  environment_id = resource.pingone_environment.%[2]s.id

  connector_ids = [
    "httpConnector",
  ]

  depends_on = [
    davinci_connection.%[2]s-1,
    davinci_connection.%[2]s-2,
    davinci_connection.%[2]s-3,
  ]
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}
