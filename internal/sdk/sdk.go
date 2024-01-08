package sdk

// re-initialize auth, with retryable error response
// func initAuthRetryable(ctx context.Context, c *dv.APIClient) *retry.RetryError {
// 	err := c.InitAuth()

// 	switch {
// 	case err == nil:
// 		return nil
// 		// These cases come from the davinci-client-go library and may be subject to change
// 	case strings.Contains(err.Error(), "Error getting admin callback, got: status: 502, body:"):
// 		tflog.Warn(ctx, "Status code 502 on admin callback. Retrying...")
// 		retry.RetryableError(err)
// 	case strings.Contains(err.Error(), "Error getting SSO callback, got err: status: 502, body:"):
// 		tflog.Warn(ctx, "Status code 502 on SSO callback. Retrying...")
// 		retry.RetryableError(err)
// 	case strings.Contains(err.Error(), "Auth Token not found, unsuccessful login, got: Found. Redirecting to https://console.pingone.com/davinci/index.html#/sso/callback/?error=AuthenticationFailed&error_description=unknownError2"):
// 		tflog.Warn(ctx, "Auth token not found on callback. Retrying...")
// 		retry.RetryableError(err)
// 	default:
// 		tflog.Warn(ctx, "Error re-initializing authorization.")
// 		return retry.NonRetryableError(err)
// 	}

// 	return retry.NonRetryableError(fmt.Errorf("Error re-initializing authorization. Please report this as a bug."))
// }
