package sdk

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"

	// "strings"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

// DoRetryable wraps the given function with a rery loop that doubles the timeout each time.
// This is used to catch and handle common errors that may be due to rate limiting or an issue in the API.
// Error functions in this list should be safe to retry and have a TODO to remove them if the issue is resolved.
func DoRetryable(ctx context.Context, c *dv.APIClient, environmentID string, f func() (interface{}, *http.Response, error)) (interface{}, error) {
	defaultTimeout := 10
	return DoRetryableWithCustomTimeout(ctx, c, environmentID, f, time.Duration(defaultTimeout)*time.Minute)
}

func DoRetryableWithCustomTimeout(ctx context.Context, c *dv.APIClient, environmentID string, f func() (interface{}, *http.Response, error), timeout time.Duration) (interface{}, error) {

	var res interface{}
	authRetryLimit := 10
	authRetryCount := 1

	err := retry.RetryContext(ctx, timeout, func() *retry.RetryError {

		var r *http.Response
		var err error
		res, r, err = f()
		if err != nil || r.StatusCode >= 300 {
			switch t := err.(type) {

			case dv.ErrorResponse:
				tflog.Warn(ctx, "Detected DaVinci application error.", map[string]interface{}{
					"message":            t.Message,
					"code":               t.Code,
					"http_response_code": t.HttpResponseCode,
					"environmentID":      environmentID,
				})

				if authRetryCount <= authRetryLimit && t.HttpResponseCode == http.StatusUnauthorized && t.Code == dv.DV_ERROR_CODE_INVALID_TOKEN_FOR_ENVIRONMENT {
					log.Printf("Client unauthorized for the environment, retrying auth (%d/%d)...", authRetryCount, authRetryLimit)

					signOnErr := c.DoSignIn(&environmentID)
					if signOnErr != nil {
						log.Printf("Sign in failed...%s..", signOnErr)
						return retry.NonRetryableError(signOnErr)
					}
					log.Printf("Sign in success.  Handing back to retrier..")

					authRetryCount++

					return retry.RetryableError(err)
				}

			case *url.Error:
				tflog.Warn(ctx, "Detected HTTP error", map[string]interface{}{
					"error": t.Err.Error(),
				})

			default:
				tflog.Warn(ctx, "Detected unknown error (retry)", map[string]interface{}{
					"type": t,
				})

				if res1, matchErr := regexp.MatchString(`^http: ContentLength=[0-9]+ with Body length [0-9]+$`, t.Error()); matchErr == nil && res1 {
					return retry.RetryableError(err)
				}
			}

			if err == nil {
				err = fmt.Errorf("HTTP error %d", r.StatusCode)
			}

			return retry.NonRetryableError(err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
