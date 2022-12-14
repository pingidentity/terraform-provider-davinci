package sdk

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

// RetryableError is an error that can be retried
func DoRetryable(ctx context.Context, f func() (interface{}, error), timeout *time.Duration) (interface{}, error) {
	defaultTimeout := 20 * time.Second
	timeElapsed := 0 * time.Second
	if timeout == nil {
		timeout = &defaultTimeout
	}
	for timeElapsed < *timeout {
		res, err := f()
		if err != nil {
			dvErr, err := dv.ParseDvHttpError(err)
			if err != nil {
				return nil, err
			}
			if dvErr.Status == 400 {
				tflog.Info(ctx, "Operation failed with 400, retrying...")
				timeElapsed += 2 * time.Second
				time.Sleep(2 * time.Second)
				continue
			}
		}
		return res, err
	}
	return nil, fmt.Errorf("Operation was not successful within timeout.")
}
