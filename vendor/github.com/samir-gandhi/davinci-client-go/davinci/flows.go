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

type flowJson struct {
	payload *string
}

func ParseFlowJson(payload *string) (*FlowImport, error) {
	fi := FlowImport{}
	flow := Flow{}
	// is FlowImport or Flow
	err := json.Unmarshal([]byte(*payload), &fi)
	if err == nil && fi.FlowNameMapping != nil {
		return &fi, nil
	}
	err = json.Unmarshal([]byte(*payload), &flow)
	if err == nil {
		pfi := FlowImport{
			Name:            flow.Name,
			Description:     flow.Description,
			FlowNameMapping: map[string]string{flow.FlowID: flow.Name},
			FlowInfo:        flow,
		}
		return &pfi, nil
	}
	return nil, fmt.Errorf("Unable to parse payload to type FlowImport")
}

func ParseFlowsJson(payload *string) (*FlowsImport, error) {
	fis := Flows{}
	//is Flows
	err := json.Unmarshal([]byte(*payload), &fis)
	if err == nil && len(fis.Flow) > 0 {
		pfis := FlowsImport{
			Name:            "",
			Description:     "",
			FlowInfo:        fis,
			FlowNameMapping: map[string]string{},
		}
		for _, v := range fis.Flow {
			pfis.FlowNameMapping[v.FlowID] = v.Name
		}
		return &pfis, nil
	}
	return nil, fmt.Errorf("Unable parse payload to type Flows")
}

func MakeFlowPayload(payload *string) (*string, error) {
	//is Flows
	pfis, _ := ParseFlowsJson(payload)
	if pfis != nil {
		fjBytes, err := json.Marshal(pfis)
		if err != nil {
			return nil, fmt.Errorf("Unable to marshal payload.")
		}
		fjString := string(fjBytes)
		payloadString := &fjString
		return payloadString, nil
	}

	// is FlowImport or Flow
	pfi, _ := ParseFlowJson(payload)
	if pfi != nil {
		fjBytes, err := json.Marshal(pfi)
		if err != nil {
			return nil, fmt.Errorf("Unable to unmarshal json to type Flow.")
		}
		fjString := string(fjBytes)
		payload = &fjString
		return payload, nil
	}
	return nil, fmt.Errorf("Invalid or unsupported flow payload.")
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

	payload, err := MakeFlowPayload(payloadJson)
	if err != nil {
		return nil, err
	}

	req := DvHttpRequest{
		Method: "PUT",
		Url:    fmt.Sprintf("%s/flows/import", c.HostURL),
		Body:   strings.NewReader(*payload),
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
