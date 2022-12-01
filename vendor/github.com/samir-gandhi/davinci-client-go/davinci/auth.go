package davinci

import (
	"encoding/json"
	// "errors"
	"bytes"
	"fmt"
	"net/http"
	"strings"
)

// // Sign up - Create new user, return user token upon successful creation
// func (c *APIClient) SignUp(auth AuthStruct) (*AuthResponse, error) {
// 	if auth.Username == "" || auth.Password == "" {
// 		return nil, fmt.Errorf("define username and password")
// 	}
// 	rb, err := json.Marshal(auth)
// 	if err != nil {
// 		return nil, err
// 	}

// 	req, err := http.NewRequest("POST", fmt.Sprintf("%s/signup", c.HostURL), strings.NewReader(string(rb)))
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	ar := AuthResponse{}
// 	err = json.Unmarshal(body, &ar)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &ar, nil
// }

// SignIn - Get a new token for user
func (c *APIClient) SignIn() (*AuthResponse, error) {
	// For Prod Getting an Access Token takes multiple steps:
	// 1. Login with User/PW - get access_token
	// 2. Start Auth Flow - Get json response
	// 3. Post response to skCallback

	// Login
	if c.Auth.Username == "" || c.Auth.Password == "" {
		return nil, fmt.Errorf("define username and password")
	}
	lReqBody, err := json.Marshal(c.Auth)
	if err != nil {
		return nil, err
	}

	lreq, err := http.NewRequest("POST", fmt.Sprintf("%s/customers/login", c.HostURL), strings.NewReader(string(lReqBody)))
	if err != nil {
		return nil, err
	}

	lbody, _, err := c.doRequest(lreq, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("Error on User Login, got: %v", err)
	}

	lr := LoginResponse{}
	err = json.Unmarshal(lbody, &lr)

	if err != nil || lr.AccessToken == "" {
		return nil, fmt.Errorf("Error on User Login: %v", string(lbody))
	}

	// Start Auth
	var sreq *http.Request
	if c.HostURL == "https://orchestrate-api.pingone.com/v1" {
		sreq, err = http.NewRequest("POST", fmt.Sprintf("https://auth.pingone.com/%s/davinci/policy/%s/start", lr.CompanyID, lr.FlowPolicyID), nil)
	} else {
		sreq, err = http.NewRequest("POST", fmt.Sprintf("%s/auth/%s/policy/%s/start", c.HostURL, lr.CompanyID, lr.FlowPolicyID), nil)
	}
	if err != nil {
		return nil, err
	}

	sbody, _, err := c.doRequest(sreq, &lr.SkSdkToken.AccessToken, nil)
	if err != nil {
		return nil, fmt.Errorf("Error on Start Auth, got: %v", err)
	}

	sr := Callback{}
	err = json.Unmarshal(sbody, &sr)
	if err != nil {
		return nil, err
	}

	// Callback
	cReqBody, err := json.Marshal(sr)
	if err != nil {
		return nil, err
	}
	areq, err := http.NewRequest("POST", fmt.Sprintf("%s/customers/skcallback", c.HostURL), strings.NewReader(string(cReqBody)))
	if err != nil {
		return nil, err
	}
	abody, _, err := c.doRequest(areq, &lr.AccessToken, nil)
	if err != nil {
		return nil, fmt.Errorf("Error on Callback, got: %v", err)
	}

	ar := AuthResponse{}
	err = json.Unmarshal(abody, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}

func (c *APIClient) SignInSSO() (*AuthResponse, error) {
	// For Prod an accessToken is aquired by providing an authToken takes multiple steps:
	// 1. Generate SSO Url and refresh state (a)
	// 2. Authorize - provides get code or FlowId (b)
	// 3a. Got FlowId, Log in with admin (c)
	// 3b. Refresh FlowId for Code (d)
	// 4. Send SSO code or state for callback to get authToken (e)
	// 5. Use authToken for DV accessToken (f)

	// Login
	var dvSsoCode, dvFlowId, dvSsoState, dvSsoAuthToken string
	if c.Auth.Username == "" || c.Auth.Password == "" {
		return nil, fmt.Errorf("define username and password")
	}
	if c.PingOneSSOEnvId == "" {
		return nil, fmt.Errorf("define PingOne Admin and Target EnvId")
	}

	// step 1 - Start SSO, refresh state and generate callback
	areq, err := http.NewRequest("GET", fmt.Sprintf("%s/customers/pingone/sso", c.HostURL), nil)
	if err != nil {
		return nil, err
	}
	aParams := Params{
		"", "", map[string]string{
			"env": c.PingOneSSOEnvId,
		},
	}
	ares, err := c.doRequestVerbose(areq, nil, &aParams)
	if err != nil || ares.StatusCode != 302 {
		return nil, fmt.Errorf("Error getting SSO callback, got err: %v\n", err)
	}
	if ares.StatusCode != 302 {
		return nil, fmt.Errorf("Error getting SSO callback, got err: %v", string(ares.Body))
	}
	if ares.LocationParams.Get("state") == "" {
		return nil, fmt.Errorf("Error Parsing SSO State not found, got: %s", ares.Location)
	}
	dvSsoState = ares.LocationParams["state"][0]

	//step 2 - Directly Execute Callback from step 1
	// Receive cookies and code or flowid in response
	breq, err := http.NewRequest("GET", fmt.Sprintf("%s://%s%s", ares.Location.Scheme, ares.Location.Host, ares.Location.Path), nil)
	if err != nil {
		return nil, err
	}
	bParams := Params{
		"", "", map[string]string{},
	}
	for i, v := range ares.LocationParams {
		bParams.ExtraParams[i] = v[0]
	}

	bres, err := c.doRequestVerbose(breq, nil, &bParams)
	if err != nil {
		return nil, fmt.Errorf("Error following SSO callback, got error: %v\n", err)
	}
	if bres.StatusCode != 302 {
		return nil, fmt.Errorf("Error following SSO callback, got: %v\n", string(bres.Body))
	}
	if bres.LocationParams.Get("flowId") != "" {
		dvFlowId = bres.LocationParams["flowId"][0]
	}
	if bres.LocationParams.Get("code") != "" {
		dvSsoCode = bres.LocationParams["code"][0]
	}
	if dvFlowId == "" && dvSsoCode == "" {
		return nil, fmt.Errorf("Error: SSO Location header did not provide Code or FlowId: %s", bres.Location)
	}

	if dvFlowId != "" {
		// step 3 Refresh FlowID to retrieve dvSsoCode

		//step 3a Log in Admin
		// Assumption that this refreshes backend SSO state..
		crb := map[string]string{
			"username": c.Auth.Username,
			"password": c.Auth.Password}
		cReqBody, err := json.Marshal(crb)
		creq, err := http.NewRequest("POST", fmt.Sprintf("%s://%s/%s/flows/%s", ares.Location.Scheme, ares.Location.Host, c.PingOneSSOEnvId, dvFlowId), bytes.NewBuffer(cReqBody))
		if err != nil {
			return nil, err
		}
		// PingOne Auth Specific Header
		creq.Header.Set("Content-Type", "application/vnd.pingidentity.usernamePassword.check+json; charset=UTF-8")

		_, err = c.doRequestVerbose(creq, nil, nil)
		if err != nil {
			return nil, fmt.Errorf("Error Authenticating PingOne Admin: %v", err)
		}
		//step 3b Retrieve dvSsoCode with refreshed Auth
		dreq, err := http.NewRequest("GET", fmt.Sprintf("%s://%s/%s/as/resume", ares.Location.Scheme, ares.Location.Host, c.PingOneSSOEnvId), nil)
		if err != nil {
			return nil, err
		}
		dParams := Params{
			"", "", map[string]string{
				"flowId": dvFlowId,
			},
		}
		dres, err := c.doRequestVerbose(dreq, nil, &dParams)
		if err != nil {
			return nil, fmt.Errorf("Error resuming auth, got error: %v\n", err)
		}
		if dres.StatusCode != 302 {
			return nil, fmt.Errorf("Error resuming auth, got: %v\n", string(dres.Body))
		}
		if dres.LocationParams.Get("code") == "" {
			return nil, fmt.Errorf("Error Parsing SSO Location, dvSsoCode not found: %v", dres.Location)
		}
		dvSsoCode = dres.LocationParams["code"][0]
	}
	//step 4 use dvSsoCode and dvSsoState to get dvAuthToken
	ereq, err := http.NewRequest("GET", fmt.Sprintf("%s/customers/pingone/callback", c.HostURL), nil)
	if err != nil {
		return nil, err
	}
	eParams := Params{
		"", "", map[string]string{},
	}
	if dvSsoCode != "" {
		eParams.ExtraParams["code"] = dvSsoCode
	}
	if dvSsoState != "" {
		eParams.ExtraParams["state"] = dvSsoState
	}
	eres, err := c.doRequestVerbose(ereq, nil, &eParams)
	if err != nil {
		return nil, fmt.Errorf("Error getting admin callback, got: %v\n", err)
	}
	if eres.StatusCode != 302 {
		return nil, fmt.Errorf("Error getting admin callback, got: %v\n", string(eres.Body))
	}
	if eres.LocationParams.Get("authToken") == "" {
		return nil, fmt.Errorf("Auth Token not found, unsuccessful login, got: %v", string(eres.Body))
	}
	dvSsoAuthToken = eres.LocationParams["authToken"][0]

	//step 5 Swap dvSsoAuthToken for access_token
	frb := map[string]string{
		"authToken": dvSsoAuthToken}
	fReqBody, err := json.Marshal(frb)

	freq, err := http.NewRequest("POST", fmt.Sprintf("%s/customers/sso/auth", c.HostURL), strings.NewReader(string(fReqBody)))
	if err != nil {
		return nil, err
	}
	fres, err := c.doRequestVerbose(freq, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("Error getting admin callback, got: %v", err)
	}

	ar := AuthResponse{}
	err = json.Unmarshal(fres.Body, &ar)
	if err != nil {
		return nil, err
	}

	return &ar, nil
}
