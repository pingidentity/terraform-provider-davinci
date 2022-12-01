package davinci

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// ReadFlows only accepts Limit as a param
func (c *APIClient) ReadFlows(companyId *string, args *Params) ([]Flow, error) {
	if args.Page != "" {
		log.Println("Param.Page found, not allowed, removing.")
		args.Page = ""
	}
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	cIdString := *cIdPointer
	log.Print(cIdString)
	// req, err := http.NewRequest("GET", fmt.Sprintf("%s/flows", c.HostURL), nil)
	// if err != nil {
	// 	return nil, err
	// }
	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/flows", c.HostURL),
	}
	body, err := c.doRequestRetryable(req, &c.Token, args)
	if err != nil {
		return nil, err
	}

	// Returned flows are an array in top level flowsInfo key
	resp := FlowsInfo{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp.Flow, nil
}

func (c *APIClient) CreateFlowWithJson(companyId *string,
	payloadJson *string) (*Flow, error) {
	if payloadJson == nil {
		return nil, fmt.Errorf("Must provide payloadJson.")
	}
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}
	pfi := FlowImport{}
	pf := Flow{}

	err = json.Unmarshal([]byte(*payloadJson), &pfi)
	if err != nil || pfi.FlowNameMapping == nil {
		log.Printf("Unable to unmarshal json to type FlowImport.\n Will try to unmarshal to type Flow")
		err = json.Unmarshal([]byte(*payloadJson), &pf)
		if err != nil {
			return nil, fmt.Errorf("Unable to unmarshal json to type Flow.")
		}
		pfi = FlowImport{
			Name:            pf.Name,
			Description:     pf.Description,
			FlowNameMapping: map[string]interface{}{pf.FlowID: pf.Name},
			FlowInfo:        pf,
		}
	}
	payload, err := json.Marshal(pfi)

	req := DvHttpRequest{
		Method: "PUT",
		Url:    fmt.Sprintf("%s/flows/import", c.HostURL),
		Body:   strings.NewReader(string(payload)),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	resp := FlowInfo{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Flow, nil
}

// ReadFlows only accepts Limit as a param
func (c *APIClient) ReadFlow(companyId *string, flowId string) (*FlowInfo, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	cIdString := *cIdPointer
	log.Print(cIdString)

	req := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/flows/%s", c.HostURL, flowId),
	}
	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	// Returned flows are an array in top level flowsInfo key
	resp := FlowInfo{}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

// Only specific fields are supported during update:
// - GraphData
// - InputSchema
// - CurrentVersion
// - Name
func (c *APIClient) UpdateFlowWithJson(companyId *string, payloadJson *string, flowId string) (*Flow, error) {
	if payloadJson == nil {
		return nil, fmt.Errorf("Must provide payloadJson.")
	}
	if flowId == "" {
		return nil, fmt.Errorf("Must provide flowId.")
	}
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}
	pfi := FlowImport{}
	pf := Flow{}

	//handle incoming type Flow or Flow Import
	err = json.Unmarshal([]byte(*payloadJson), &pf)
	if err != nil {
		log.Printf("Unable to unmarshal json to type FlowImport.\n Will try to unmarshal to type Flow")
		err = json.Unmarshal([]byte(*payloadJson), &pfi)
		if err != nil {
			return nil, fmt.Errorf("Unable to unmarshal json to type FlowImport.")
		}
		pf = pfi.FlowInfo
	}

	currentFlow, err := c.ReadFlow(cIdPointer, flowId)
	if err != nil {
		return nil, err
	}
	if pf.CurrentVersion > currentFlow.Flow.CurrentVersion {
		pf.CurrentVersion = currentFlow.Flow.CurrentVersion
	}

	pAllowedProps := Flow{
		GraphData:      pf.GraphData,
		InputSchema:    pf.InputSchema,
		CurrentVersion: pf.CurrentVersion,
		Name:           pf.Name,
	}
	payload, err := json.Marshal(pAllowedProps)

	req := DvHttpRequest{
		Method: "PUT",
		Url:    fmt.Sprintf("%s/flows/%s", c.HostURL, flowId),
		Body:   strings.NewReader(string(payload)),
	}

	body, err := c.doRequestRetryable(req, &c.Token, nil)
	if err != nil {
		return nil, err
	}

	resp := FlowInfo{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp.Flow, nil
}

// ReadFlows only accepts Limit as a param
func (c *APIClient) DeleteFlow(companyId *string, flowId string) (*Message, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	cIdString := *cIdPointer
	log.Print(cIdString)

	req := DvHttpRequest{
		Method: "DELETE",
		Url:    fmt.Sprintf("%s/flows/%s", c.HostURL, flowId),
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

// ReadFlows only accepts Limit as a param
func (c *APIClient) DeployFlow(companyId *string, flowId string) (*Message, error) {
	cIdPointer := &c.CompanyID
	if companyId != nil {
		cIdPointer = companyId
	}
	_, err := c.SetEnvironment(cIdPointer)
	if err != nil {
		return nil, err
	}

	cIdString := *cIdPointer
	log.Print(cIdString)
	// req, err := http.NewRequest("PUT", fmt.Sprintf("%s/flows/%s/deploy", c.HostURL, flowId), nil)
	// if err != nil {
	// 	return nil, err
	// }
	req := DvHttpRequest{
		Method: "PUT",
		Url:    fmt.Sprintf("%s/flows/%s/deploy", c.HostURL, flowId),
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
