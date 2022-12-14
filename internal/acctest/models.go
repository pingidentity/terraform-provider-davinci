package acctest

type flowResource struct {
	Name     string
	FlowJson string
	Subflows []string
	// slice of connectorIds one word lowercased without word "connector". Ex: "httpConnector" -> "http"
	Connections []string
}

type TestFlowsJson struct {
	Simple          string
	Drifted         string
	Mainflow        string
	MainflowDrifted string
	Subflow         string
	AnotherMainflow string
	AnotherSubflow  string
}

type TestFlowsHcl struct {
	// A simple flow with minimal external dependencies
	Simple  FlowHcl
	Drifted FlowHcl
	//Depends on Subflow and AnotherSubflow
	Mainflow        FlowHcl
	MainflowDrifted FlowHcl
	Subflow         FlowHcl
	AnotherMainflow FlowHcl
	AnotherSubflow  FlowHcl
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

// Hardcoded map of bootstrapped connections that come with new DV tenant.
// keys are connectorIds are one word lowercased
var bsConnections = map[string]flowConnection{
	"annotation": {
		ConnectorId: "annotationConnector",
		Name:        "Annotation",
		Id:          "921bfae85c38ed45045e07be703d86b8",
	},
	"errorcustomize": {
		ConnectorId: "errorConnector",
		Name:        "Error Customize",
		Id:          "6d8f6f706c45fd459a86b3f092602544",
	},
	"functions": {
		ConnectorId: "functionsConnector",
		Name:        "Function",
		Id:          "de650ca45593b82c49064ead10b9fe17",
	},
	"http": {
		ConnectorId: "httpConnector",
		Name:        "Http",
		Id:          "867ed4363b2bc21c860085ad2baa817d",
	},
	"pingonesso": {
		ConnectorId: "pingOneSSOConnector",
		Name:        "PingOne",
		Id:          "94141bf2f1b9b59a5f5365ff135e02bb",
	},
	"skopenid": {
		ConnectorId: "skOpenIdConnector",
		Name:        "Token Management",
		Id:          "3b55f2ca6689560c64cb5bed5afbe40f",
	},
	"userpolicy": {
		ConnectorId: "userPolicyConnector",
		Name:        "User Policy",
		Id:          "4cb5e3465d718a84087ec9b6a6251e52",
	},
	"skuserpool": {
		Name:        "Default User Pool",
		ConnectorId: "skUserPool",
		Id:          "defaultUserPool",
	},
	"variables": {
		ConnectorId: "variablesConnector",
		Name:        "Variables",
		Id:          "06922a684039827499bdbdd97f49827b",
	},
}
