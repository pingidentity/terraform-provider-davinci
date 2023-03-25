package acctest

import ()

type flowResource struct {
	Name     string
	FlowJson string
	Subflows []string
	// slice of connectorIds one word lowercased without word "connector". Ex: "httpConnector" -> "http"
	Connections []string
}

type TestFlowsJson struct {
	Simple                       string
	SimpleSettingDrift           string
	SimpleInputSchemaDrift       string
	SimpleOutputSchemaDrift      string
	Drifted                      string
	Mainflow                     string
	MainflowDrifted              string
	Subflow                      string
	AnotherMainflow              string
	AnotherSubflow               string
	WithVariableConnector        string
	BrokenFlow                   string
	PingOneSessionMainFlow       string
	PingOneSessionMainFlowUpdate string
	PingOneSessionSubFlow        string
}

type TestFlowsHcl struct {
	// A simple flow with minimal external dependencies
	Simple                  FlowHcl
	SimpleSettingDrift      FlowHcl
	SimpleInputSchemaDrift  FlowHcl
	SimpleOutputSchemaDrift FlowHcl
	Drifted                 FlowHcl
	//Depends on Subflow and AnotherSubflow
	Mainflow                     FlowHcl
	MainflowDrifted              FlowHcl
	Subflow                      FlowHcl
	AnotherMainflow              FlowHcl
	AnotherSubflow               FlowHcl
	WithVariableConnector        FlowHcl
	BrokenFlow                   FlowHcl
	PingOneSessionMainFlow       FlowHcl
	PingOneSessionMainFlowUpdate FlowHcl
	PingOneSessionSubFlow        FlowHcl
}

type FlowHcl struct {
	Name string
	Hcl  string
}

// Simple connection name/id pair to be converted in TestFlowsHcl.
type flowConnection struct {
	ConnectorId string
	Name        string
	Id          string
}

type TestConnections struct {
	Http        string
	CrowdStrike string
	Flow        string
}
