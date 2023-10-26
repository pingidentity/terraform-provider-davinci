package acctest

import "fmt"

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
	FlowContextVarFlow           string
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
	FlowContextVarFlow           FlowHcl
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

type TestConnection struct {
	ResourcePrefix string
	Name           string
	ConnectorId    string
	Properties     []TestConnectionProperty
}

type TestConnectionProperty struct {
	Name  string
	Value string
	Type  string
}

// Returns resource name for connection that should be used in hcl format: `<ResourceName>_<Name>`
func (tc TestConnection) GetResourceName() (resourceName string) {
	return fmt.Sprintf("%s_%s", tc.ResourcePrefix, tc.Name)
}

func TestAccResourceConnectionHcl(resourceName string, p1Services []string, connections []TestConnection) (hcl string) {
	baseHcl := PingoneEnvrionmentServicesSsoHcl(resourceName, p1Services)
	connectionsHcl := ""
	for _, connection := range connections {
		connectionsHcl += connection.MakeConnectionHcl()
	}

	return baseHcl + connectionsHcl
}

func (tcp TestConnection) MakeConnectionHcl() (hcl string) {
	propertiesHcl := ""
	for _, property := range tcp.Properties {
		propertiesHcl += fmt.Sprintf(`
	property {
		name  = "%s"
		value = "%s"
	}
`, property.Name, property.Value)
	}
	hcl = fmt.Sprintf(`
resource "davinci_connection" "%[2]s" {
	environment_id = resource.pingone_role_assignment_user.%[1]s.scope_environment_id
	depends_on     = [data.davinci_connections.read_all]
	connector_id   = "%[3]s"
	name           = "%[2]s"
	%[4]s
}
`, tcp.ResourcePrefix, tcp.GetResourceName(), tcp.ConnectorId, propertiesHcl)
	return hcl
}
