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
		freshEnvTimeout := 60
		for i := 0; i <= timeout; {
			// initially, test if auth is valid for the target environment
			_, err := c.ReadApplications(&envId, nil)
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
					err = initAuthRetryable(ctx, c)
					if err != nil {
						return err
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

// re-initialize auth, with optionial retry
func initAuthRetryable(ctx context.Context, c *dv.APIClient) error {
	for retries := 0; retries <= 3; retries++ {
		err := c.InitAuth()
		if retries == 3 && err != nil {
			return err
		}
		switch {
		case err == nil:
			return nil
			// These cases come from the davinci-client-go library and may be subject to change
		case strings.Contains(err.Error(), "Error getting admin callback, got: status: 502, body:"):
			tflog.Info(ctx, "Found retryable error while initializing client. Retrying...")
		case strings.Contains(err.Error(), "Error getting SSO callback, got err: status: 502, body:"):
			tflog.Info(ctx, "Found retryable error while initializing client. Retrying...")
		case strings.Contains(err.Error(), "Auth Token not found, unsuccessful login, got: Found. Redirecting to https://console.pingone.com/davinci/index.html#/sso/callback/?error=AuthenticationFailed&error_description=unknownError2"):
			tflog.Info(ctx, "Found retryable error while initializing client. Retrying...")
		default:
			tflog.Info(ctx, "Error re-initializing authorization.")
			return err
		}
	}
	return fmt.Errorf("Error re-initializing authorization. Please report this as a bug.")
}
