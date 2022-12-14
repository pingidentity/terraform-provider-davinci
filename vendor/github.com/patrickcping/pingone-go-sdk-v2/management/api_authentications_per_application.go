/*
PingOne Platform API - SSO and Base

The PingOne Platform API covering the base and SSO services (otherwise known as the Management APIs)

API version: 2022-08-02
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package management

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// AuthenticationsPerApplicationApiService AuthenticationsPerApplicationApi service
type AuthenticationsPerApplicationApiService service

type ApiV1EnvironmentsEnvironmentIDApplicationSignonsGetRequest struct {
	ctx context.Context
	ApiService *AuthenticationsPerApplicationApiService
	environmentID string
	limit *int32
	samplePeriod *int32
	samplePeriodCount *int32
	filter *string
}

func (r ApiV1EnvironmentsEnvironmentIDApplicationSignonsGetRequest) Limit(limit int32) ApiV1EnvironmentsEnvironmentIDApplicationSignonsGetRequest {
	r.limit = &limit
	return r
}

func (r ApiV1EnvironmentsEnvironmentIDApplicationSignonsGetRequest) SamplePeriod(samplePeriod int32) ApiV1EnvironmentsEnvironmentIDApplicationSignonsGetRequest {
	r.samplePeriod = &samplePeriod
	return r
}

func (r ApiV1EnvironmentsEnvironmentIDApplicationSignonsGetRequest) SamplePeriodCount(samplePeriodCount int32) ApiV1EnvironmentsEnvironmentIDApplicationSignonsGetRequest {
	r.samplePeriodCount = &samplePeriodCount
	return r
}

func (r ApiV1EnvironmentsEnvironmentIDApplicationSignonsGetRequest) Filter(filter string) ApiV1EnvironmentsEnvironmentIDApplicationSignonsGetRequest {
	r.filter = &filter
	return r
}

func (r ApiV1EnvironmentsEnvironmentIDApplicationSignonsGetRequest) Execute() (*http.Response, error) {
	return r.ApiService.V1EnvironmentsEnvironmentIDApplicationSignonsGetExecute(r)
}

/*
V1EnvironmentsEnvironmentIDApplicationSignonsGet READ Authentications Per Application (Partial)

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param environmentID
 @return ApiV1EnvironmentsEnvironmentIDApplicationSignonsGetRequest
*/
func (a *AuthenticationsPerApplicationApiService) V1EnvironmentsEnvironmentIDApplicationSignonsGet(ctx context.Context, environmentID string) ApiV1EnvironmentsEnvironmentIDApplicationSignonsGetRequest {
	return ApiV1EnvironmentsEnvironmentIDApplicationSignonsGetRequest{
		ApiService: a,
		ctx: ctx,
		environmentID: environmentID,
	}
}

// Execute executes the request
func (a *AuthenticationsPerApplicationApiService) V1EnvironmentsEnvironmentIDApplicationSignonsGetExecute(r ApiV1EnvironmentsEnvironmentIDApplicationSignonsGetRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthenticationsPerApplicationApiService.V1EnvironmentsEnvironmentIDApplicationSignonsGet")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/environments/{environmentID}/applicationSignons"
	localVarPath = strings.Replace(localVarPath, "{"+"environmentID"+"}", url.PathEscape(parameterToString(r.environmentID, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.limit != nil {
		localVarQueryParams.Add("limit", parameterToString(*r.limit, ""))
	}
	if r.samplePeriod != nil {
		localVarQueryParams.Add("samplePeriod", parameterToString(*r.samplePeriod, ""))
	}
	if r.samplePeriodCount != nil {
		localVarQueryParams.Add("samplePeriodCount", parameterToString(*r.samplePeriodCount, ""))
	}
	if r.filter != nil {
		localVarQueryParams.Add("filter", parameterToString(*r.filter, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v P1Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
