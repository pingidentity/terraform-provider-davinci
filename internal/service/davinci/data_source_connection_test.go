package davinci_test

import (
	"fmt"
	"regexp"

	// "os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/pingidentity/terraform-provider-davinci/internal/acctest"
	"github.com/pingidentity/terraform-provider-davinci/internal/verify"
)

func TestAccDataSourceConnection_DefinedConnection_ByID_Clean(t *testing.T) {
	testAccDataSourceConnection_DefinedConnection_ByID(t, false)
}

func TestAccDataSourceConnection_DefinedConnection_ByID_WithBootstrap(t *testing.T) {
	testAccDataSourceConnection_DefinedConnection_ByID(t, true)
}

func testAccDataSourceConnection_DefinedConnection_ByID(t *testing.T, withBootstrapConfig bool) {

	resourceBase := "davinci_connection"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)
	dataSourceFullName := fmt.Sprintf("data.%s", resourceFullName)

	name := resourceName

	hcl := testAccDataSourceConnection_DefinedConnection_ByID_Hcl(resourceName, name, withBootstrapConfig)

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
					resource.TestCheckResourceAttrPair(dataSourceFullName, "id", resourceFullName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "environment_id", resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "name", resourceFullName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "connector_id", resourceFullName, "connector_id"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "customer_id", resourceFullName, "customer_id"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "created_date", resourceFullName, "created_date"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "property.#", resourceFullName, "property.#"),
				),
			},
		},
	})
}

func TestAccDataSourceConnection_DefinedConnection_ByName_Clean(t *testing.T) {
	testAccDataSourceConnection_DefinedConnection_ByName(t, false)
}

func TestAccDataSourceConnection_DefinedConnection_ByName_WithBootstrap(t *testing.T) {
	testAccDataSourceConnection_DefinedConnection_ByName(t, true)
}

func testAccDataSourceConnection_DefinedConnection_ByName(t *testing.T, withBootstrapConfig bool) {

	resourceBase := "davinci_connection"
	resourceName := acctest.ResourceNameGen()
	resourceFullName := fmt.Sprintf("%s.%s", resourceBase, resourceName)
	dataSourceFullName := fmt.Sprintf("data.%s", resourceFullName)

	name := resourceName

	hcl := testAccDataSourceConnection_DefinedConnection_ByName_Hcl(resourceName, name, withBootstrapConfig)

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
					resource.TestCheckResourceAttrPair(dataSourceFullName, "id", resourceFullName, "id"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "environment_id", resourceFullName, "environment_id"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "name", resourceFullName, "name"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "connector_id", resourceFullName, "connector_id"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "customer_id", resourceFullName, "customer_id"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "created_date", resourceFullName, "created_date"),
					resource.TestCheckResourceAttrPair(dataSourceFullName, "property.#", resourceFullName, "property.#"),
				),
			},
		},
	})
}

func TestAccDataSourceConnection_BootstrapConnection_ByID_Clean(t *testing.T) {

	withBootstrapConfig := false

	resourceName := acctest.ResourceNameGen()

	name := resourceName

	hcl := testAccDataSourceConnection_BootstrapConnection_ByID_Hcl(resourceName, name, withBootstrapConfig)

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
				Config:      hcl,
				ExpectError: regexp.MustCompile(`Connection not found`),
			},
		},
	})
}

func TestAccDataSourceConnection_BootstrapConnection_ByID_WithBootstrap(t *testing.T) {

	withBootstrapConfig := true

	resourceBase := "davinci_connection"
	resourceName := acctest.ResourceNameGen()
	dataSourceFullName := fmt.Sprintf("data.%s.%s", resourceBase, resourceName)

	name := resourceName

	hcl := testAccDataSourceConnection_BootstrapConnection_ByID_Hcl(resourceName, name, withBootstrapConfig)

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
					resource.TestMatchResourceAttr(dataSourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
					resource.TestMatchResourceAttr(dataSourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
					resource.TestCheckResourceAttr(dataSourceFullName, "name", "PingOne"),
					resource.TestCheckResourceAttr(dataSourceFullName, "connector_id", "pingOneSSOConnector"),
					resource.TestCheckResourceAttr(dataSourceFullName, "property.#", "4"),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceFullName, "property.*", map[string]*regexp.Regexp{
						"name":  regexp.MustCompile(`^region$`),
						"value": regexp.MustCompile(`^[A-Z]{2}$`),
						"type":  regexp.MustCompile(`^string$`),
					}),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceFullName, "property.*", map[string]*regexp.Regexp{
						"name":  regexp.MustCompile(`^envId$`),
						"value": verify.P1ResourceIDRegexpFullString,
						"type":  regexp.MustCompile(`^string$`),
					}),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceFullName, "property.*", map[string]*regexp.Regexp{
						"name":  regexp.MustCompile(`^clientId$`),
						"value": verify.P1ResourceIDRegexpFullString,
						"type":  regexp.MustCompile(`^string$`),
					}),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceFullName, "property.*", map[string]*regexp.Regexp{
						"name":  regexp.MustCompile(`^clientSecret$`),
						"value": regexp.MustCompile(`^\*{6}$`),
						"type":  regexp.MustCompile(`^string$`),
					}),
					resource.TestMatchResourceAttr(dataSourceFullName, "created_date", verify.EpochDateRegexpFullString),
					resource.TestMatchResourceAttr(dataSourceFullName, "customer_id", verify.P1DVResourceIDRegexpFullString),
				),
			},
		},
	})
}

func TestAccDataSourceConnection_BootstrapConnection_ByName_Clean(t *testing.T) {

	withBootstrapConfig := false

	resourceName := acctest.ResourceNameGen()

	name := resourceName

	hcl := testAccDataSourceConnection_BootstrapConnection_ByName_Hcl(resourceName, name, withBootstrapConfig)

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
				Config:      hcl,
				ExpectError: regexp.MustCompile(`Connection not found`),
			},
		},
	})
}

func TestAccDataSourceConnection_BootstrapConnection_ByName_WithBootstrap(t *testing.T) {

	withBootstrapConfig := true

	resourceBase := "davinci_connection"
	resourceName := acctest.ResourceNameGen()
	dataSourceFullName := fmt.Sprintf("data.%s.%s", resourceBase, resourceName)

	name := resourceName

	hcl := testAccDataSourceConnection_BootstrapConnection_ByName_Hcl(resourceName, name, withBootstrapConfig)

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
					resource.TestMatchResourceAttr(dataSourceFullName, "id", verify.P1DVResourceIDRegexpFullString),
					resource.TestMatchResourceAttr(dataSourceFullName, "environment_id", verify.P1ResourceIDRegexpFullString),
					resource.TestCheckResourceAttr(dataSourceFullName, "name", "PingOne"),
					resource.TestCheckResourceAttr(dataSourceFullName, "connector_id", "pingOneSSOConnector"),
					resource.TestCheckResourceAttr(dataSourceFullName, "property.#", "4"),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceFullName, "property.*", map[string]*regexp.Regexp{
						"name":  regexp.MustCompile(`^region$`),
						"value": regexp.MustCompile(`^[A-Z]{2}$`),
						"type":  regexp.MustCompile(`^string$`),
					}),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceFullName, "property.*", map[string]*regexp.Regexp{
						"name":  regexp.MustCompile(`^envId$`),
						"value": verify.P1ResourceIDRegexpFullString,
						"type":  regexp.MustCompile(`^string$`),
					}),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceFullName, "property.*", map[string]*regexp.Regexp{
						"name":  regexp.MustCompile(`^clientId$`),
						"value": verify.P1ResourceIDRegexpFullString,
						"type":  regexp.MustCompile(`^string$`),
					}),
					resource.TestMatchTypeSetElemNestedAttrs(dataSourceFullName, "property.*", map[string]*regexp.Regexp{
						"name":  regexp.MustCompile(`^clientSecret$`),
						"value": regexp.MustCompile(`^\*{6}$`),
						"type":  regexp.MustCompile(`^string$`),
					}),
					resource.TestMatchResourceAttr(dataSourceFullName, "created_date", verify.EpochDateRegexpFullString),
					resource.TestMatchResourceAttr(dataSourceFullName, "customer_id", verify.P1DVResourceIDRegexpFullString),
				),
			},
		},
	})
}

func testAccDataSourceConnection_DefinedConnection_ByID_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
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

data "davinci_connection" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  id             = davinci_connection.%[2]s.id
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccDataSourceConnection_DefinedConnection_ByName_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
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

data "davinci_connection" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  name           = davinci_connection.%[2]s.name
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccDataSourceConnection_BootstrapConnection_ByID_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

data "davinci_connection" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  id             = "94141bf2f1b9b59a5f5365ff135e02bb" // the PingOne connector
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}

func testAccDataSourceConnection_BootstrapConnection_ByName_Hcl(resourceName, name string, withBootstrapConfig bool) (hcl string) {
	return fmt.Sprintf(`
%[1]s

data "davinci_connection" "%[2]s" {
  environment_id = pingone_environment.%[2]s.id
  name           = "PingOne" // the PingOne connector
}
`, acctest.PingoneEnvironmentSsoHcl(resourceName, withBootstrapConfig), resourceName, name)
}
