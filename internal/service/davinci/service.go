package davinci

import (
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

type serviceClientType struct {
	Client *dv.APIClient
}

func Resources() []func() resource.Resource {
	return []func() resource.Resource{
		NewFlowResource,
	}
}

func DataSources() []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}
