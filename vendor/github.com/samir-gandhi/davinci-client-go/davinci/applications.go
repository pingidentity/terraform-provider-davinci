package davinci

import (
	"encoding/json"
	"fmt"
	"strings"
)

// ReadFlows only accepts Limit as a param
func (c *APIClient) ReadApplications(companyId *string, args *Params) ([]App, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/apps", c.HostURL),
	}
	body, err := c.doRequestRetryable(req, &c.Token, args)
	if err != nil {
		return nil, err
	}

	// Returned flows are an array in top level flowsInfo key
	resp := Apps{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	// Leaving in case of revert - but this shouldn't be treated as an error.
	// if len(resp.Apps) == 0 {
	// 	return nil, fmt.Errorf("No applications found with given params")
	// }
	return resp.Apps, nil
}

func (c *APIClient) CreateApplication(companyId *string, appName string) (*App, error) {
	if appName == "" {
		return nil, fmt.Errorf("Must provide appName")
	}
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}

	_, err := c.SetEnvironment(cIdPointer)

	if err != nil {
		return nil, err
	}

	p := App{
		Name: appName,
	}

	payload, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	req := DvHttpRequest{
		Method: "POST",
		Url:    fmt.Sprintf("%s/apps", c.HostURL),
		Body:   strings.NewReader(string(payload)),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	r := ReadApp{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	if r.App.Name == "" {
		return nil, fmt.Errorf("Unable to create app")
	}
	return &r.App, nil
}

// UpdateApplication - Update all fields of an application besides Policies. Policies should be updated via UpdatePolicy
func (c *APIClient) UpdateApplication(companyId *string, payload *AppUpdate) (*App, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	if payload == nil || payload.Name == "" || payload.AppID == "" {
		return nil, fmt.Errorf("App Name and ID required in payload")
	}
	appId := payload.AppID
	payload.AppID = ""

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := DvHttpRequest{
		Method: "PUT",
		Url:    fmt.Sprintf("%s/apps/%s", c.HostURL, appId),
		Body:   strings.NewReader(string(reqBody)),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	res := ReadApp{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res.App, nil
}

func (c *APIClient) ReadApplication(companyId *string, appId string) (*App, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}

	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}
	if appId == "" {
		return nil, fmt.Errorf("AppId not provided")
	}

	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/apps/%s", c.HostURL, appId),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	res := ReadApp{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res.App, nil
}

// CreateInitializedApplication is useful when creating an application with flow policy.
// Takes an app payload and calls:
// - CreateApplication
// - UpdateApplication
// - CreateFlowPolicy
func (c *APIClient) CreateInitializedApplication(companyId *string, payload *AppUpdate) (*App, error) {

	//Create Application
	resp, err := c.CreateApplication(companyId, payload.Name)
	if err != nil {
		err = fmt.Errorf("Unable to create application. Error: %v", err)
		return nil, err
	}
	payload.AppID = resp.AppID

	//Create Flow Policies
	for _, v := range payload.Policies {
		resp, err = c.CreateFlowPolicy(companyId, resp.AppID, v)
		if err != nil {
			err = fmt.Errorf("Unable to create application flow policy. Error: %v", err)
			return nil, err
		}
		payload.Policies = resp.Policies
	}

	//Update Application
	resp, err = c.UpdateApplication(companyId, payload)
	if err != nil {
		err = fmt.Errorf("Unable to update created application. Error: %v", err)
		return nil, err
	}

	return resp, nil
}

// Deletes an application based on applicationId
func (c *APIClient) DeleteApplication(companyId *string, appId string) (*Message, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	req := DvHttpRequest{
		Method: "DELETE",
		Url:    fmt.Sprintf("%s/apps/%s", c.HostURL, appId),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	resp := Message{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
