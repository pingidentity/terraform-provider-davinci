package sdk

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

// DoRetryable wraps the given function with a retry loop
// This is used to catch and handle common Terraform-level functional errors such as race conditions.
func DoRetryable(ctx context.Context, c *dv.APIClient, environmentID string, f func() (interface{}, *http.Response, error)) (interface{}, error) {
	defaultTimeout := 10
	return DoRetryableWithCustomTimeout(ctx, c, environmentID, f, time.Duration(defaultTimeout)*time.Minute)
}

func DoRetryableWithCustomTimeout(ctx context.Context, c *dv.APIClient, environmentID string, f func() (interface{}, *http.Response, error), timeout time.Duration) (interface{}, error) {

	var res interface{}

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

			case *url.Error:
				tflog.Warn(ctx, "Detected HTTP error", map[string]interface{}{
					"error": t.Err.Error(),
				})

			default:
				tflog.Warn(ctx, "Detected unknown error (retry)", map[string]interface{}{
					"type": t,
				})
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
