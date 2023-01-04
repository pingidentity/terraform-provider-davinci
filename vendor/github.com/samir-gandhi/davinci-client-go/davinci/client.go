// Davinci Admin API GO Client
//
// This package is go client to be used for interacting with PingOne DaVinci Administrative APIs.
// Use cases include:
// - Creating Connections
// - Importing Flows
package davinci

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"
)

// baseURL
var baseURL = url.URL{
	Scheme: "https",
	Host:   "orchestrate-api.pingone.com",
	Path:   "/v1",
}

var dvApiHost = map[string]string{
	"NorthAmerica": "orchestrate-api.pingone.com",
	"Europe":       "orchestrate-api.pingone.eu",
	"AsiaPacific":  "orchestrate-api.pingone.asia",
	"Canada":       "orchestrate-api.pingone.ca",
}

// const HostURL string = "https://api.singularkey.com/v1"

func (args Params) QueryParams() url.Values {
	q := make(url.Values)

	if args.Page != "" {
		q.Add("page", args.Page)
	}

	if args.Limit != "" {
		q.Add("limit", args.Limit)
	}
	for i, v := range args.ExtraParams {
		q.Add(i, v)
	}

	return q
}

func NewClient(inputs *ClientInput) (*APIClient, error) {
	// adjust host according to received region
	if inputs.PingOneRegion != "" {
		if dvApiHost[inputs.PingOneRegion] == "" {
			return nil, fmt.Errorf("Invalid region: %v", inputs.PingOneRegion)
		}
		baseURL.Host = dvApiHost[inputs.PingOneRegion]
	}

	hostUrl := baseURL.ResolveReference(&url.URL{}).String()

	fmt.Printf("Using host: %v \n", hostUrl)
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("Got error while creating cookie jar %s", err.Error())
	}
	c := APIClient{
		HTTPClient: &http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
			Timeout: 10 * time.Second,
			Jar:     jar},
		HostURL: hostUrl,
	}

	if inputs.Username == "" || inputs.Password == "" {
		// return nil, fmt.Errorf("User or Password not found")
		return &c, nil
	}

	c.Auth = AuthStruct{
		Username: inputs.Username,
		Password: inputs.Password,
	}

	// Use P1SSO if available
	if inputs.PingOneSSOEnvId != "" {
		c.PingOneSSOEnvId = inputs.PingOneSSOEnvId
	}
	err = c.doSignIn()
	if err != nil {
		return nil, fmt.Errorf("Sign In failed with: %v", err)
	}

	return &c, nil
}

func (c *APIClient) doSignIn() error {
	if c.PingOneSSOEnvId != "" {
		ar, err := c.SignInSSO()
		if err != nil {
			return err
		}
		// if ar.AccessToken == "" {
		// 	// return fmt.Errorf("Sign in failed. No Access Token found %v", ar.)
		// 	return err
		// }
		c.Token = ar.AccessToken
		return nil
	}

	//Default Env User login
	ar, err := c.SignIn()
	if err != nil {
		return err
	}
	c.Token = ar.AccessToken
	return nil
}

func (c *APIClient) InitAuth() error {
	if c.PingOneSSOEnvId != "" {
		ar, err := c.SignInSSO()
		if err != nil {
			return err
		}
		c.Token = ar.AccessToken
		return nil
	}

	//Default Env User login
	ar, err := c.SignIn()
	if err != nil {
		return err
	}
	c.Token = ar.AccessToken
	return nil
}

func (c *APIClient) doRequestVerbose(req *http.Request, authToken *string, args *Params) (*DvHttpResponse, error) {
	token := c.Token

	if authToken != nil {
		token = *authToken
		var bearer = "Bearer " + token
		req.Header.Add("Authorization", bearer)
	}
	if req.Header.Get("Content-Type") == "" {
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
	if args != nil {
		req.URL.RawQuery = args.QueryParams().Encode()
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	rbody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusFound {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, rbody)
	}
	resp := DvHttpResponse{
		Body:       rbody,
		Headers:    res.Header,
		StatusCode: res.StatusCode,
	}

	if res.StatusCode == http.StatusFound && res.Header["Location"] != nil {
		resp.Location, _ = url.Parse(res.Header["Location"][0])
		resp.LocationParams, _ = url.ParseQuery(resp.Location.RawQuery)
		// Handle wepbage hash value strangeness
		if resp.Location.Fragment != "" {
			_, v, ok := strings.Cut(resp.Location.Fragment, "?")
			if ok {
				resp.LocationParams, _ = url.ParseQuery(v)
			}
		}
	}
	if res.Header["Set-Cookie"] != nil {
		c.HTTPClient.Jar.SetCookies(req.URL, res.Cookies())
	}

	return &resp, err
}

func (c *APIClient) doRequest(req *http.Request, authToken *string, args *Params) ([]byte, *http.Response, error) {
	token := c.Token
	if authToken != nil {
		token = *authToken
	}
	if token != "" {
		var bearer = "Bearer " + token
		req.Header.Add("Authorization", bearer)
		req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	}
	if args != nil {
		req.URL.RawQuery = args.QueryParams().Encode()
	}
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, nil, err
	}
	statusOk := res.StatusCode >= 200 && res.StatusCode < 300
	if !statusOk {
		return body, res, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
	return body, res, err
}

func (c *APIClient) doRequestRetryable(req DvHttpRequest, authToken *string, args *Params) ([]byte, error) {
	reqInit, err := http.NewRequest(req.Method, req.Url, req.Body)
	if err != nil {
		return nil, err
	}
	reqRetry, err := http.NewRequest(req.Method, req.Url, req.Body)
	if err != nil {
		return nil, err
	}
	body, res, err := c.doRequest(reqInit, authToken, args)
	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusUnauthorized && c.AuthRefresh == false {
		if err != nil {
			return nil, err
		}
		err = c.refreshAuth()
		if err != nil {
			return nil, err
		}
		_, err := c.SetEnvironment(&c.CompanyID)
		if err != nil {
			return nil, err
		}

		var resRetry *http.Response
		var bodyRetry []byte
		bodyRetry, resRetry, err = c.doRequest(reqRetry, authToken, args)
		if err != nil {
			return nil, err
		}
		res = resRetry
		body = bodyRetry
	}
	if err != nil {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}
	return body, err
}

// refreshAuth is used to rerun the sign-on process.
// This is useful when the client's initial access_token was made before
// the target environment was created. (common in Terraform)
func (c *APIClient) refreshAuth() error {
	c.AuthRefresh = true
	// c.HTTPClient.Jar = nil
	// jar, err := cookiejar.New(nil)
	// if err != nil {
	// 	return fmt.Errorf("Got error while creating cookie jar %s", err.Error())
	// }
	// c.HTTPClient.Jar = jar
	err := c.doSignIn()
	if err != nil {
		return fmt.Errorf("Refreshing Sign In failed with: %v", err)
	}
	return nil
}

// sample incoming must be formatted as similar to:
// fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
func (c *APIClient) ParseDvHttpError(e error) (*DvHttpError, error) {
	eBefore, eBody, ok := strings.Cut(e.Error(), ", body: ")
	_, eStatus, ok := strings.Cut(eBefore, "status: ")
	eStatusInt, err := strconv.Atoi(eStatus)
	if ok != true || err != nil {
		return nil, fmt.Errorf("Invalid error parameter. ")
	}
	return &DvHttpError{
		Status: eStatusInt,
		Body:   eBody,
	}, nil
}

// sample incoming must be formatted as similar to:
// fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
func ParseDvHttpError(e error) (*DvHttpError, error) {
	eBefore, eBody, ok := strings.Cut(e.Error(), ", body: ")
	_, eStatus, ok := strings.Cut(eBefore, "status: ")
	eStatusInt, err := strconv.Atoi(eStatus)
	if ok != true || err != nil {
		return nil, fmt.Errorf("Invalid error parameter. ")
	}
	return &DvHttpError{
		Status: eStatusInt,
		Body:   eBody,
	}, nil
}
