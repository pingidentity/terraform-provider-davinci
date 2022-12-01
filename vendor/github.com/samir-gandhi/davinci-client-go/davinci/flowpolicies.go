package davinci

import (
	"encoding/json"
	"fmt"
	"strings"
)

func (c *APIClient) CreateFlowPolicy(companyId *string, appId string, policy Policy) (*App, error) {
	if appId == "" {
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
	payload := policy
	payload.PolicyID = ""
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := DvHttpRequest{
		Method: "POST",
		Url:    fmt.Sprintf("%s/apps/%s/policy", c.HostURL, appId),
		Body:   strings.NewReader(string(reqBody)),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	r := App{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, err
	}

	if len(r.Policies) == 0 {
		return nil, fmt.Errorf("Unable to create FlowPolicy")
	}
	return &r, nil
}

func (c *APIClient) UpdateFlowPolicy(companyId *string, appId string, policy Policy) (*App, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	if appId == "" || policy.PolicyID == "" {
		return nil, fmt.Errorf("Missing appId or policy.PolicyID")
	}
	payload := policy
	payload.PolicyID = ""
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req := DvHttpRequest{
		Method: "PUT",
		Url:    fmt.Sprintf("%s/apps/%s/policy/%s", c.HostURL, appId, policy.PolicyID),
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

// Deletes an application based on applicationId
func (c *APIClient) DeleteFlowPolicy(companyId *string, appId string, policyId string) (*Message, error) {
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
		Url:    fmt.Sprintf("%s/apps/%s/policy/%s", c.HostURL, appId, policyId),
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
