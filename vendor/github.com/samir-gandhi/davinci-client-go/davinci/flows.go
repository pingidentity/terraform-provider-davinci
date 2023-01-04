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

func ParseFlowImportJson(payload *string) (*FlowImport, error) {
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

func ParseFlowJson(payload *string) (*Flow, error) {
	fi := FlowImport{}
	flow := Flow{}
	// is FlowImport or Flow
	err := json.Unmarshal([]byte(*payload), &flow)
	if err == nil && flow.GraphData.Elements.Nodes != nil {
		return &flow, nil
	}
	err = json.Unmarshal([]byte(*payload), &fi)
	if err == nil {
		pfi := fi.FlowInfo
		return &pfi, nil
	}
	return nil, fmt.Errorf("Unable to parse payload to type FlowImport")
}

func ParseFlowsImportJson(payload *string) (*FlowsImport, error) {
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

// MakeFlowPayload accepts
// payload: string of format Flows, FlowImport, or Flow
// output: desired type of FlowsImport, FlowImport, or Flow
// Payloads can only be converted to matching plurality
func MakeFlowPayload(payload *string, output string) (*string, error) {
	if output == "" {
		output = "FlowImport"
		fis := Flows{}
		err := json.Unmarshal([]byte(*payload), &fis)
		if err == nil && len(fis.Flow) > 0 {
			output = "FlowsImport"
		}
	}
	switch output {
	case "FlowsImport":
		pfis, _ := ParseFlowsImportJson(payload)
		fjBytes, err := json.Marshal(pfis)
		if err != nil {
			return nil, fmt.Errorf("Unable to marshal payload.")
		}
		fjString := string(fjBytes)
		payloadString := &fjString
		return payloadString, nil
	case "Flow":
		pfi, _ := ParseFlowJson(payload)
		fjBytes, err := json.Marshal(pfi)
		if err != nil {
			return nil, fmt.Errorf("Unable to unmarshal json to type Flow.")
		}
		fjString := string(fjBytes)
		payload = &fjString
		return payload, nil
	case "FlowImport":
		pfi, _ := ParseFlowImportJson(payload)
		fjBytes, err := json.Marshal(pfi)
		if err != nil {
			return nil, fmt.Errorf("Unable to unmarshal json to type FlowImport.")
		}
		fjString := string(fjBytes)
		payload = &fjString
		return payload, nil
	default:
		return nil, fmt.Errorf("Output must be one of: FlowsImport, FlowImport, or Flow.")
	}
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

	payload, err := MakeFlowPayload(payloadJson, "")
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

	// Call orx, this replicates the GET made from UI which seems to trigger some database function:
	reqOrx := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/flows/%s", c.HostURL, resp.Flow.FlowID),
	}
	params := Params{
		ExtraParams: map[string]string{
			"attributes": "orx",
		},
	}
	_, err = c.doRequestRetryable(reqOrx, &c.Token, &params)
	if err != nil {
		return nil, err
	}

	// Call apps, this replicates the GET made from UI which seems to trigger some database function:
	reqApps := DvHttpRequest{
		Method: "GET",
		Url:    fmt.Sprintf("%s/flows/%s", c.HostURL, resp.Flow.FlowID),
	}
	_, err = c.doRequestRetryable(reqApps, &c.Token, nil)
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
	pf := Flow{}
	flow, err := MakeFlowPayload(payloadJson, "Flow")
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(*flow), &pf)

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

	resp := Flow{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
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
