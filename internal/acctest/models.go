package acctest

import (
	"encoding/json"
	"fmt"

	dv "github.com/samir-gandhi/davinci-client-go/davinci"
)

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

type TestApplicationFlowPolicy struct {
	FlowPolicyResourceName string
	Name                   string
	EnvironmentID          string
	ApplicationID          string
	ID                     string
}

type TestApplication struct {
	ApplicationResourceName string
	Name                    string
	ID                      string
	EnvironmentID           string
}

func (ta TestApplication) GetResourceFullName() (resourceName string) {
	return fmt.Sprintf("davinci_application.%s", ta.ApplicationResourceName)
}

func (ta *TestApplication) SetName(name string) {
	ta.Name = name
}

func (ta *TestApplication) SetID(id string) {
	ta.ID = id
}

func (ta *TestApplication) SetEnvironmentID(environmentID string) {
	ta.EnvironmentID = environmentID
}

func (tafp TestApplicationFlowPolicy) GetResourceFullName() (resourceName string) {
	return fmt.Sprintf("davinci_application_flow_policy.%s", tafp.FlowPolicyResourceName)
}

func (tafp *TestApplicationFlowPolicy) SetName(name string) {
	tafp.Name = name
}

func (tafp *TestApplicationFlowPolicy) SetEnvironmentID(environmentID string) {
	tafp.EnvironmentID = environmentID
}

func (tafp *TestApplicationFlowPolicy) SetApplicationID(applicationID string) {
	tafp.ApplicationID = applicationID
}

func (tafp *TestApplicationFlowPolicy) SetID(id string) {
	tafp.ID = id
}

// Returns resource name for connection that should be used in hcl format: `<ResourceName>_<Name>`
func (tc TestConnection) GetResourceName() (resourceName string) {
	return fmt.Sprintf("%s_%s", tc.ResourcePrefix, tc.Name)
}

func TestAccResourceConnectionHcl(resourceName string, p1Services []string, connections []TestConnection) (hcl string) {
	baseHcl := PingoneEnvironmentServicesSsoHcl(resourceName, p1Services, true)
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
  connector_id   = "%[3]s"
  name           = "%[2]s"
	%[4]s
}
`, tcp.ResourcePrefix, tcp.GetResourceName(), tcp.ConnectorId, propertiesHcl)
	return hcl
}

func DeepCloneFlow(src, dst *dv.Flow) error {
	bytes, err := json.Marshal(src)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, dst)
	if err != nil {
		return err
	}

	return nil
}
