package davinci

import (
	"context"
	"fmt"
	// "log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/pingidentity/terraform-provider-davinci/internal/sdk"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

func DataSourceConnection() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceConnectionRead,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "ID of the connection to retrieve. Either id or name must be specified.",
				ExactlyOneOf: []string{"id", "name"},
			},
			"name": {
				Type:         schema.TypeString,
				Optional:     true,
				Computed:     true,
				Description:  "Name of the connection to retrieve. Either id or name must be specified.",
				ExactlyOneOf: []string{"id", "name"},
			},
			"connector_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"environment_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "PingOne environment id",
			},
			"customer_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_date": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"property": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: "Connection configuration",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": {
							Type:      schema.TypeString,
							Computed:  true,
							Sensitive: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

// var dvConnectionRetryableErrors = []sdk.RetryableError{
// 	{
// 		Status: 400,

// 	}
// }

func dataSourceConnectionRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	c := meta.(*dv.APIClient)
	var diags diag.Diagnostics

	err := sdk.CheckAndRefreshAuth(ctx, c, d)
	if err != nil {
		return diag.FromErr(err)
	}
	var connId string
	// prep case where name is provided
	connName, ok := d.GetOk("name")
	if ok {
		connId, err = getConnectionIdByName(ctx, c, connName.(string))
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// prep case where id is provided
	_, ok = d.GetOk("id")
	if ok {
		connId = d.Get("id").(string)
	}

	// get connection by id and parse response
	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.ReadConnection(&c.CompanyID, connId)
	}, nil)

	if err != nil {
		return diag.FromErr(err)
	}
	res, ok := sdkRes.(*dv.Connection)
	if !ok {
		err = fmt.Errorf("Unable to parse response from Davinci API on connection id: %v", connId)
		return diag.FromErr(err)
	}

	d.SetId(res.ConnectionID)

	if err := d.Set("name", res.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("connector_id", res.ConnectorID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("created_date", res.CreatedDate); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("environment_id", res.CompanyID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("customer_id", res.CustomerID); err != nil {
		return diag.FromErr(err)
	}
	props, err := flattenConnectionProperties(&res.Properties)
	if err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("property", props); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(connId)
	return diags
}

func getConnectionIdByName(ctx context.Context, c *dv.APIClient, connName string) (string, error) {
	sdkRes, err := sdk.DoRetryable(ctx, func() (interface{}, error) {
		return c.ReadConnections(&c.CompanyID, nil)
	}, nil)
	if err != nil {
		return "", err
	}

	res, ok := sdkRes.([]dv.Connection)
	if !ok {
		return "", fmt.Errorf("Unable to parse response from Davinci API on connection name: %v", connName)
	}

	for _, conn := range res {
		if conn.Name == connName {
			return conn.ConnectionID, nil
		}
	}
	return "", fmt.Errorf("Unable to find connection with name: %v", connName)
}
