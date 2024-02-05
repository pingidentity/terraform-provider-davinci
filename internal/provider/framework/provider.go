package framework

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/patrickcping/pingone-go-sdk-v2/pingone/model"
	"github.com/pingidentity/terraform-provider-davinci/internal/client"
	"github.com/pingidentity/terraform-provider-davinci/internal/framework"
	"github.com/pingidentity/terraform-provider-davinci/internal/service/davinci"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

// Ensure DaVinciProvider satisfies various provider interfaces.
var _ provider.Provider = &davinciProvider{}

// DaVinciProvider defines the provider implementation.
type davinciProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// davinciProviderModel describes the provider data model.
type davinciProviderModel struct {
	Username      types.String `tfsdk:"username"`
	Password      types.String `tfsdk:"password"`
	Region        types.String `tfsdk:"region"`
	EnvironmentID types.String `tfsdk:"environment_id"`
	HostURL       types.String `tfsdk:"host_url"`
	AccessToken   types.String `tfsdk:"access_token"`
}

func (p *davinciProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "davinci"
	resp.Version = p.version
}

func (p *davinciProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {

	usernameDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"The PingOne username used for SSO into a Davinci tenant.  Default value can be set with the `PINGONE_USERNAME` environment variable. Must provide username and password, or access_token.",
	)

	passwordDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"The PingOne password used for SSO into a Davinci tenant.  Default value can be set with the `PINGONE_PASSWORD` environment variable. Must provide username and password, or access_token.",
	)

	regionDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"The PingOne region to use.  Options are `AsiaPacific` `Canada` `Europe` and `NorthAmerica`.  Default value can be set with the `PINGONE_REGION` environment variable.",
	)

	environmentIDDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"Environment ID PingOne User Login. Default value can be set with the `PINGONE_ENVIRONMENT_ID` environment variable.",
	)

	hostURLDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"To override the default region-based url, provide a PingOne DaVinci API host url. Default value can be set with the `PINGONE_DAVINCI_HOST_URL` environment variable.",
	)

	accessTokenDescription := framework.SchemaAttributeDescriptionFromMarkdown(
		"PingOne DaVinci specific access token. Must be authorized for environment_id.  Default value can be set with the `PINGONE_DAVINCI_ACCESS_TOKEN` environment variable. Must provide username and password, or access_token.",
	)

	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"username": schema.StringAttribute{
				Description:         usernameDescription.Description,
				MarkdownDescription: usernameDescription.MarkdownDescription,
				Optional:            true,
			},

			"password": schema.StringAttribute{
				Description:         passwordDescription.Description,
				MarkdownDescription: passwordDescription.MarkdownDescription,
				Optional:            true,
				Sensitive:           true,
			},

			"region": schema.StringAttribute{
				Description:         regionDescription.Description,
				MarkdownDescription: regionDescription.MarkdownDescription,
				Optional:            true,

				Validators: []validator.String{
					stringvalidator.OneOf(model.RegionsAvailableList()...),
				},
			},

			"environment_id": schema.StringAttribute{
				Description:         environmentIDDescription.Description,
				MarkdownDescription: environmentIDDescription.MarkdownDescription,
				Optional:            true,
			},

			"host_url": schema.StringAttribute{
				Description:         hostURLDescription.Description,
				MarkdownDescription: hostURLDescription.MarkdownDescription,
				Optional:            true,
			},

			"access_token": schema.StringAttribute{
				Description:         accessTokenDescription.Description,
				MarkdownDescription: accessTokenDescription.MarkdownDescription,
				Optional:            true,
			},
		},
	}
}

func (p *davinciProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Debug(ctx, "[v6] Provider configure start")
	var data davinciProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Set the defaults
	tflog.Info(ctx, "[v6] Provider setting defaults..")
	debugLogMessage := "[v6] Provider parameter %s missing, defaulting to environment variable"

	if data.Username.IsNull() {
		v := os.Getenv("PINGONE_USERNAME")
		tflog.Debug(ctx, fmt.Sprintf(debugLogMessage, "username"), map[string]interface{}{
			"env_var":       "PINGONE_USERNAME",
			"env_var_value": v,
		})
		data.Username = basetypes.NewStringValue(v)
	}

	if data.Password.IsNull() {
		v := os.Getenv("PINGONE_PASSWORD")
		tflog.Debug(ctx, fmt.Sprintf(debugLogMessage, "password"), map[string]interface{}{
			"env_var":       "PINGONE_PASSWORD",
			"env_var_value": v,
		})
		data.Password = basetypes.NewStringValue(v)
	}

	if data.Region.IsNull() {
		v := os.Getenv("PINGONE_REGION")
		tflog.Debug(ctx, fmt.Sprintf(debugLogMessage, "region"), map[string]interface{}{
			"env_var":       "PINGONE_REGION",
			"env_var_value": v,
		})
		data.Region = basetypes.NewStringValue(v)
	}

	if data.EnvironmentID.IsNull() {
		v := os.Getenv("PINGONE_ENVIRONMENT_ID")
		tflog.Debug(ctx, fmt.Sprintf(debugLogMessage, "environment_id"), map[string]interface{}{
			"env_var":       "PINGONE_ENVIRONMENT_ID",
			"env_var_value": v,
		})
		data.EnvironmentID = basetypes.NewStringValue(v)
	}

	if data.HostURL.IsNull() {
		v := os.Getenv("PINGONE_DAVINCI_HOST_URL")
		tflog.Debug(ctx, fmt.Sprintf(debugLogMessage, "host_url"), map[string]interface{}{
			"env_var":       "PINGONE_DAVINCI_HOST_URL",
			"env_var_value": v,
		})
		data.HostURL = basetypes.NewStringValue(v)
	}

	if data.AccessToken.IsNull() {
		v := os.Getenv("PINGONE_DAVINCI_ACCESS_TOKEN")
		tflog.Debug(ctx, fmt.Sprintf(debugLogMessage, "access_token"), map[string]interface{}{
			"env_var":       "PINGONE_DAVINCI_ACCESS_TOKEN",
			"env_var_value": v,
		})
		data.AccessToken = basetypes.NewStringValue(v)
	}

	userAgent := fmt.Sprintf("terraform-provider-davinci/%s", p.version)

	if v := strings.TrimSpace(os.Getenv("DAVINCI_TF_APPEND_USER_AGENT")); v != "" {
		userAgent += fmt.Sprintf(" %s", v)
	}

	cInput := dv.ClientInput{
		Username:        data.Username.ValueString(),
		Password:        data.Password.ValueString(),
		PingOneRegion:   data.Region.ValueString(),
		PingOneSSOEnvId: data.EnvironmentID.ValueString(),
		HostURL:         data.HostURL.ValueString(),
		AccessToken:     data.AccessToken.ValueString(),
		UserAgent:       userAgent,
	}
	apiClient, err := client.RetryableClient(&cInput)
	if err != nil {
		resp.Diagnostics.AddError(
			"Failed to create client",
			fmt.Sprintf("Error when creating client: %s", err),
		)
		return
	}

	var resourceConfig framework.ResourceType
	resourceConfig.Client = apiClient
	tflog.Info(ctx, "[v6] Provider initialized client")

	resp.ResourceData = resourceConfig
	resp.DataSourceData = resourceConfig

}

func (p *davinciProvider) Resources(ctx context.Context) []func() resource.Resource {
	v := make([]func() resource.Resource, 0)
	v = append(v, davinci.Resources()...)
	return v
}

func (p *davinciProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	v := make([]func() datasource.DataSource, 0)
	v = append(v, davinci.DataSources()...)
	return v
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &davinciProvider{
			version: version,
		}
	}
}
