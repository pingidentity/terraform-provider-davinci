package sdk

import (
	"context"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

func GetEnvironmentId(c *dv.APIClient, d *schema.ResourceData) {
	if cId, ok := d.GetOk("environment_id"); ok {
		c.CompanyID = cId.(string)
	}
}

// CheckAndRefreshAuth aims to help with SSO to new davinci environments.
// The function checks if the current authentication is authorized for the target environment.
// If the check fails, it attempts to refresh and recheck the client access_token for up to 40s
// This function assumes auth may become valid any time with 40s and bad auth will not be identified until then.
func CheckAndRefreshAuth(ctx context.Context, c *dv.APIClient, d *schema.ResourceData) error {
	if cId, ok := d.GetOk("environment_id"); ok {
		envId := cId.(string)
		// if envId eq cid, assumption is that auth is okay.
		if envId == c.CompanyID {
			return nil
		}
		timeout := 40
		for i := 0; i <= timeout; {
			_, err := c.ReadConnections(&envId, nil)
			if err != nil {
				httpErr, err := c.ParseDvHttpError(err)
				if err == nil && httpErr.Status == 401 && strings.Contains(httpErr.Body, "Authorization failed") && i <= timeout {
					if i == 0 {
						tflog.Info(ctx, "Identified possible need to refresh access_token. Attempting refresh")
					}
					i = i + 4
					time.Sleep(4 * time.Second)
					err = c.InitAuth()
					if err != nil {
						return err
					}
					if i > 0 {
						tflog.Warn(ctx, "Possible fresh DaVinci env. Retrying Auth ... ")
					}
					continue
				}
				return err
			}
			c.CompanyID = envId
			break
		}
	}
	return nil
}
