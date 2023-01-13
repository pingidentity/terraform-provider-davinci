package sdk

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

// DoRetryable wraps the given function in a retry loop on passed in retryable errors
func DoRetryableError(ctx context.Context, f func() (interface{}, error), timeout *time.Duration, retryableErrors []RetryableError) (interface{}, error) {
	defaultTimeout := 20 * time.Second
	timeElapsed := 0 * time.Second
	defaultRetryableErrors := []RetryableError{{
		Status:     502,
		Body:       "",
		LogMessage: "Rate limit hit, retrying...",
	}}
	retryableErrors = append(retryableErrors, defaultRetryableErrors...)
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
			for _, retryableError := range retryableErrors {
				// TODO: Handle other error codes as they are discovered
				if dvErr.Status == retryableError.Status && dvErr.Body == retryableError.Body {
					tflog.Info(ctx, retryableError.LogMessage)
					timeElapsed += 2 * time.Second
					time.Sleep(2 * time.Second)
					continue
				}
			}
		}
		return res, err
	}
	return nil, fmt.Errorf("Operation was not successful within timeout.")
}

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
			// TODO: Handle other error codes as they are discovered
			if dvErr.Status == 502 && dvErr.Body == "" {
				tflog.Info(ctx, "Rate limit hit, retrying...")
				timeElapsed += 2 * time.Second
				time.Sleep(2 * time.Second)
				continue
			}
		}
		return res, err
	}
	return nil, fmt.Errorf("Operation was not successful within timeout.")
}

type RetryableError struct {
	// HTTP Status code to check for
	Status int
	// String to check for within the error body
	Body string
	// Message to log to terraform when retrying
	LogMessage string
}
