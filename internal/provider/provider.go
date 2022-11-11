package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/patrickcping/pingone-go-sdk-v2/pingone/model"
	"github.com/pingidentity/terraform-provider-davinci/internal/service/davinci"
	client "github.com/samir-gandhi/davinci-client-go/davinci"
)

func init() {
	// Docs will use markdown
	schema.DescriptionKind = schema.StringMarkdown
	schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
		desc := s.Description
		if s.Default != nil {
			desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
		}
		return strings.TrimSpace(desc)
	}
}

// Provider -
func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"username": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("PINGONE_USERNAME", nil),
					Description: "The PingOne username used for SSO into a Davinci tenant.",
				},
				"password": {
					Type:        schema.TypeString,
					Required:    true,
					Sensitive:   true,
					DefaultFunc: schema.EnvDefaultFunc("PINGONE_PASSWORD", nil),
					Description: "The PingOne password used for SSO into a Davinci tenant.",
				},
				"region": {
					Type:             schema.TypeString,
					Required:         true,
					DefaultFunc:      schema.EnvDefaultFunc("PINGONE_REGION", nil),
					Description:      "The PingOne region to use.  Options are `AsiaPacific` `Canada` `Europe` and `NorthAmerica`.  Default value can be set with the `PINGONE_REGION` environment variable.",
					ValidateDiagFunc: validation.ToDiagFunc(validation.StringInSlice(model.RegionsAvailableList(), false)),
				},
				"environment_id": {
					Type:        schema.TypeString,
					Required:    true,
					DefaultFunc: schema.EnvDefaultFunc("PINGONE_ENVIRONMENT_ID", nil),
					Description: "Environment ID PingOne User Login. Default value can be set with the `PINGONE_ENVIRONMENT_ID` environment variable.",
				},
			},
			ResourcesMap: map[string]*schema.Resource{
				// "davinci_connection":  davinci.ResourceConnection(),
				// "davinci_flow":        davinci.ResourceFlow(),
				// "davinci_application": davinci.ResourceApplication(),
			},
			DataSourcesMap: map[string]*schema.Resource{
				"davinci_connections": davinci.DataSourceConnections(),
				// "davinci_connection":   davinci.DataSourceConnection(),
				// "davinci_applications": davinci.DataSourceApplications(),
				// "davinci_application":  davinci.DataSourceApplication(),
			},
		}
		p.ConfigureContextFunc = configure(version, p)
		return p
	}
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		username := d.Get("username").(string)
		password := d.Get("password").(string)
		region := d.Get("region").(string)
		var environment_id string
		if env, ok := d.GetOk("environment_id"); ok {
			environment_id = env.(string)
		}

		var diags diag.Diagnostics

		cInput := client.ClientInput{
			Username:        username,
			Password:        password,
			PingOneRegion:   region,
			PingOneSSOEnvId: environment_id,
		}
		c, err := client.NewClient(&cInput)
		if err != nil {
			return nil, diag.FromErr(err)
		}
		if environment_id != "" {
			c.CompanyID = environment_id
		}
		//In case non-sso is desired in the future
		// c, err := client.NewClient(nil)
		// if err != nil {
		// 	return nil, diag.FromErr(err)
		// }
		return c, diags
	}
}
