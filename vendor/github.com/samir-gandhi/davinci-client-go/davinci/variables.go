package davinci

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"net/url"
	"strings"
)

func (c *APIClient) ReadVariables(companyId *string, args *Params) (map[string]Variable, error) {
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
		Url:    fmt.Sprintf("%s/constructs", c.HostURL),
	}
	body, err := c.doRequestRetryable(req, &c.Token, args)
	if err != nil {
		return nil, err
	}

	// Vars are returned as map
	resp := map[string]Variable{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *APIClient) ReadVariable(companyId *string, variableName string) (map[string]Variable, error) {
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
		Url:    fmt.Sprintf("%s/constructs/%s", c.HostURL, url.PathEscape(variableName)),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	// Vars are returned as map
	resp := map[string]Variable{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	if len(resp) != 1 {
		return nil, fmt.Errorf("status:404 body: Variable not found or invalid data returned")
	}

	return resp, nil
}

func (c *APIClient) CreateVariable(companyId *string, variable *VariablePayload) (map[string]Variable, error) {
	validate := validator.New()
	if err := validate.Struct(variable); err != nil {
		return nil, err
	}

	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	reqBody, err := json.Marshal(variable)
	if err != nil {
		return nil, err
	}

	req := DvHttpRequest{
		Method: "POST",
		Url:    fmt.Sprintf("%s/constructs", c.HostURL),
		Body:   strings.NewReader(string(reqBody)),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}
	var resp map[string]Variable
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// UpdateVariable can update fields besides Name and Context
func (c *APIClient) UpdateVariable(companyId *string, variable *VariablePayload) (map[string]Variable, error) {
	validate := validator.New()
	if err := validate.Struct(variable); err != nil {
		return nil, err
	}

	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	vName := variable.Name
	variable.Name = ""
	reqBody, err := json.Marshal(variable)
	if err != nil {
		return nil, err
	}

	req := DvHttpRequest{
		Method: "PUT",
		Url:    fmt.Sprintf("%s/constructs/%s", c.HostURL, url.PathEscape(fmt.Sprintf(`%s##SK##%s`, vName, variable.Context))),
		Body:   strings.NewReader(string(reqBody)),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}
	var resp map[string]Variable
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *APIClient) DeleteVariable(companyId *string, variableName string) (*Message, error) {
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
		Url:    fmt.Sprintf("%s/constructs/%s", c.HostURL, url.PathEscape(variableName)),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	// Vars are returned as map
	resp := Message{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}
