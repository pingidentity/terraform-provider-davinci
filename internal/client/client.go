package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	client "github.com/samir-gandhi/davinci-client-go/davinci"
)

func RetryableClient(cInput *client.ClientInput) (*client.APIClient, error) {
	var c *client.APIClient
	var err error
	ctx := context.Background()

	for retries := 0; retries <= 2; retries++ {
		c, err = client.NewClient(cInput)
		if retries == 2 && err != nil {
			return nil, err
		}
		switch {
		case err == nil:
			return c, nil
			// These cases come from the davinci-client-go library and may be subject to change
		case strings.Contains(err.Error(), "Error getting admin callback, got: status: 502, body:"):
			tflog.Info(ctx, "Found retryable error while initializing client. Retrying...")
		case strings.Contains(err.Error(), "Error getting SSO callback, got err: status: 502, body:"):
			tflog.Info(ctx, "Found retryable error while initializing client. Retrying...")
		case strings.Contains(err.Error(), "Auth Token not found, unsuccessful login, got: Found. Redirecting to https://console.pingone.com/davinci/index.html#/sso/callback/?error=AuthenticationFailed&error_description=unknownError2"):
			tflog.Info(ctx, "Found retryable error while initializing client. Retrying...")
		default:
			return nil, err
		}
	}
	return nil, fmt.Errorf("Error initializing client. Please report this as a bug.")
}
