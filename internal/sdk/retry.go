package sdk

import (
	"context"
	"fmt"
	// "strings"
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
	if timeElapsed >= *timeout {
		return nil, fmt.Errorf("Operation was not successful within timeout.")
	}
	return nil, fmt.Errorf("Operation was not successful within timeout.")
}

// DoRetryable wraps the given function with a rery loop that doubles the timeout each time.
// This is used to catch and handle common errors that may be due to rate limiting or an issue in the API.
// Error functions in this list should be safe to retry and have a TODO to remove them if the issue is resolved.
func DoRetryable(ctx context.Context, f func() (interface{}, error), timeout *time.Duration) (interface{}, error) {
	defaultTimeout := 20
	timeoutMax := defaultTimeout
	timeElapsed := 1
	var err error
	if timeout != nil {
		timeoutMax = int(*timeout)
	}
	for timeElapsed < timeoutMax {
		if timeElapsed > 1 {
			time.Sleep(time.Duration(timeElapsed) * time.Second)
		}
		res, err := f()
		if err == nil {
			return res, err
		}
		dvErr, parseErr := dv.ParseDvHttpError(err)
		if parseErr != nil {
			return nil, err
		}
		// TODO: Handle other error codes as they are discovered
		switch dvErr.Status {
		case 502:
			switch dvErr.Body {
			case "":
				tflog.Info(ctx, "Rate limit hit, retrying...")
			default:
				return nil, err
			}
		// case 400:
		// 	switch {
		// 	case strings.Contains(dvErr.Body, fmt.Sprintf(`"cause":null,"logLevel":"error","serviceName":null,"message":"Error deleting flow","errorMessage":"Error deleting flow","success":false,"httpResponseCode":400,"code":8032`)):
		// 		tflog.Info(ctx, "Retryable error on connections endpoint, retrying...")
		// 		fmt.Println("Retryable error on connections endpoint: Error deleting flow, retrying...")
		// 	case strings.Contains(dvErr.Body, fmt.Sprintf(`"cause":null,"logLevel":"error","serviceName":null,"message":"Error deleting record","errorMessage":"Error deleting record","success":false,"httpResponseCode":400,"code":6008`)):
		// 		tflog.Info(ctx, "Retryable error while deleting record, retrying...")
		// 		fmt.Println("Retryable error while deleting record, retrying...")
		// 	case strings.Contains(dvErr.Body, fmt.Sprintf(`"cause":null,"logLevel":"error","serviceName":null,"message":"Error updating flow","errorMessage":"Error updating flow","success":false,"httpResponseCode":400,"code":8033`)):
		// 		tflog.Info(ctx, "Retryable error while updating flow, retrying...")
		// 		fmt.Println("Retryable error while updating flow, retrying...")
		// 	// This error occurs occasionally, but is not retryable because it is the same error when reading a connection that does not exist.
		// 	// case strings.Contains(dvErr.Body, fmt.Sprintf(`"cause":null,"logLevel":"error","serviceName":null,"message":"Error retrieving connectors","errorMessage":"Error retrieving connectors","success":false,"httpResponseCode":400,"code":7005`)):
		// 	// 	tflog.Info(ctx, "Retryable error on connections endpoint, retrying...")
		// 	// 	fmt.Println("Retryable error on connections endpoint: Error retrieving connectors, retrying...")
		// 	// This error occurs occasionally, but is not retryable because it is the same error when deleting a connection that does not exist.
		// 	// case strings.Contains(dvErr.Body, fmt.Sprintf(`"cause":null,"logLevel":"error","serviceName":null,"message":"Connectors not found","errorMessage":"Connectors not found","success":false,"httpResponseCode":400,"code":7006`)):
		// 	// 	tflog.Info(ctx, "Retryable error on connections endpoint, retrying...")
		// 	// 	fmt.Println("Retryable error on connections endpoint: Connectors not found, retrying...")
		// 	default:
		// 		return nil, err
		// 	}
		default:
			return nil, err
		}
		timeElapsed = timeElapsed * 2
		continue
	}
	//remove this line
	// panic("Operation was not successful within timeout.")
	return nil, err
}

type RetryableError struct {
	// HTTP Status code to check for
	Status int
	// String to check for within the error body
	Body string
	// Message to log to terraform when retrying
	LogMessage string
}
