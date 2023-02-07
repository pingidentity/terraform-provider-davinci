package sdk

import (
	"context"
	"fmt"
	"strconv"
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
		timeout := 100
		freshEnvTimeout := 50
		for i := 0; i <= timeout; {
			// initially, test if auth is valid for the target environment
			apps, err := c.ReadApplications(&envId, nil)
			if err != nil {
				httpErr, err := c.ParseDvHttpError(err)
				if err == nil && httpErr.Status == 401 && strings.Contains(httpErr.Body, "Authorization failed") && i <= timeout {
					if i == 0 {
						tflog.Info(ctx, "Identified possible need to refresh access_token. Attempting refresh")
						i = i + 1
						time.Sleep(1 * time.Second)
					} else {
						i = i + 4
						time.Sleep(4 * time.Second)
						if i > 0 {
							if i > freshEnvTimeout {
								return fmt.Errorf("Unable to retrieve access_token within %ss for environment %s. Please check your credentials", strconv.Itoa(freshEnvTimeout), envId)
							}
							tflog.Warn(ctx, "Possible fresh DaVinci env. Retrying Auth ... ")
						}
					}
					err = c.InitAuth()
					if err != nil {
						return err
					}
					continue
				}
				return err
			}
			// If auth had to be initialized, check if the target environment is ready.
			if i >= 1 {
				// For new environments, need to wait for bootstrapping to complete.
				// The final step is creation of the app.
				if len(apps) == 0 {
					tflog.Warn(ctx, "Waiting for bootstrap to complete... ")
					i = i + 5
					time.Sleep(5 * time.Second)
					continue
				}
			}
			c.CompanyID = envId
			break
		}
	}
	return nil
}
