package davinci

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Gets array of all connections for the provided company
func (c *APIClient) ReadConnections(companyId *string, args *Params) ([]Connection, error) {
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
		Url:    fmt.Sprintf("%s/connections", c.HostURL),
	}

	body, err := c.doRequestRetryable(req, &c.Token, args)
	if err != nil {
		return nil, err
	}

	connections := []Connection{}
	err = json.Unmarshal(body, &connections)
	if err != nil {
		return nil, err
	}

	return connections, nil
}

// Gets single connections based on ConnectionId
func (c *APIClient) ReadConnection(companyId *string, connectionId string) (*Connection, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}

	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}
	if connectionId == "" {
		return nil, fmt.Errorf("connectionId not provided")
	}

	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/connections/%s", c.HostURL, connectionId),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	connection := Connection{}
	err = json.Unmarshal(body, &connection)
	if err != nil {
		return nil, err
	}

	return &connection, nil
}

// Create a bare connection, properties can be added _after_ creation
func (c *APIClient) CreateConnection(companyId *string, payload *Connection) (*Connection, error) {
	if companyId != nil {
		c.CompanyID = *companyId
	}
	_, err := c.SetEnvironment(&c.CompanyID)
	if err != nil {
		return nil, err
	}

	if payload == nil || payload.Name == "" || payload.ConnectorID == "" {
		return nil, fmt.Errorf("Empty or invalid payload")
	}
	connectionCreateBody := Connection{
		Name:        payload.Name,
		ConnectorID: payload.ConnectorID,
	}
	reqBody, err := json.Marshal(connectionCreateBody)
	if err != nil {
		return nil, err
	}
	req := DvHttpRequest{
		Method: "POST",
		Url:    fmt.Sprintf("%s/connections", c.HostURL),
		Body:   strings.NewReader(string(reqBody)),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	connResponse := Connection{}
	err = json.Unmarshal(body, &connResponse)
	if err != nil {
		return nil, err
	}

	return &connResponse, nil
}

// Update existing connection properties.
//
/* Sample minimal payload:
&Connection{
	ConnectionID: "foo-123"
	Properties: Properties{
		"foo": struct {
			Value string `json:"value"`
		}{"bar"}
	}
}
*/
func (c *APIClient) UpdateConnection(companyId *string, payload *Connection) (*Connection, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	if payload == nil || payload.Name == "" || payload.ConnectorID == "" {
		return nil, fmt.Errorf("Empty or invalid payload")
	}

	//Update connection ONLY allows properties
	propsOnly := Connection{
		Properties: payload.Properties,
	}

	reqBody, err := json.Marshal(propsOnly)
	if err != nil {
		return nil, err
	}

	req := DvHttpRequest{
		Method: "PUT",
		Url:    fmt.Sprintf("%s/connections/%s", c.HostURL, payload.ConnectionID),
		Body:   strings.NewReader(string(reqBody)),
	}
	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	res := Connection{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// Create a connection and fill connection properties
/* Sample minimal payload:
&Connection{
	ConnectorID: "fooConnector"
	Name: "Foo Connector"
	Properties: Properties{
		"foo": struct {
			Value string `json:"value"`
		}{"bar"}
	}
}
*/
func (c *APIClient) CreateInitializedConnection(companyId *string, payload *Connection) (*Connection, error) {
	if companyId != nil {
		c.CompanyID = *companyId
	}
	connCreatePayload := Connection{
		Name:        payload.Name,
		ConnectorID: payload.ConnectorID,
	}

	resp, err := c.CreateConnection(companyId, &connCreatePayload)
	if err != nil {
		err = fmt.Errorf("Unable to create connection. Error: %v", err)
		return nil, err
	}
	payload.ConnectionID = resp.ConnectionID

	res, err := c.UpdateConnection(companyId, payload)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Deletes a connection based on ConnectionId
func (c *APIClient) DeleteConnection(companyId *string, connectionId string) (*Message, error) {
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
		Url:    fmt.Sprintf("%s/connections/%s", c.HostURL, connectionId),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	connection := Message{}
	err = json.Unmarshal(body, &connection)
	if err != nil {
		return nil, err
	}

	return &connection, nil
}
