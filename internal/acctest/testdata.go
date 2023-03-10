package acctest

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
)

func makeSubflowsHcl(resourceName string, subflows []string) (subflowsHcl string) {
	for _, subflowName := range subflows {
		subflowsHcl += fmt.Sprintf(`
	subflow_link {
		id = davinci_flow.%[1]s-%[2]s.id
		name = davinci_flow.%[1]s-%[2]s.name
	}
		
`, resourceName, subflowName)
	}
	return subflowsHcl
}

func makeFlowConnectionsHcl(resourceName string, connections []string) (connectionsHcl string) {
	for _, connName := range connections {
		rName := fmt.Sprintf("davinci_connection.%s-%s", resourceName, connName)
		if _, ok := bsConnections[connName]; ok {
			rName = "data." + rName
		}

		connectionsHcl += fmt.Sprintf(`
	connection_link {
		id = %[1]s.id
		name = %[1]s.name
	}
	
`,
			rName)
	}
	return connectionsHcl
}

func makeFlowHcl(resourceName string, flow flowResource) FlowHcl {
	subflowsHcl := makeSubflowsHcl(resourceName, flow.Subflows)
	connectionsHcl := makeFlowConnectionsHcl(resourceName, flow.Connections)
	hcl := fmt.Sprintf(`
resource "davinci_flow" "%[1]s-%[2]s" {
	flow_json = %[3]s
	environment_id = resource.pingone_role_assignment_user.%[1]s.scope_environment_id
	depends_on = [data.davinci_connections.read_all]

	deploy = true
	
	%[4]s

	%[5]s
}
`, resourceName, flow.Name, flow.FlowJson, subflowsHcl, connectionsHcl)
	flowHcl := FlowHcl{
		Name: resourceName + "-" + flow.Name,
		Hcl:  hcl,
	}
	return flowHcl
}

func FlowsForTests(resourceName string) TestFlowsHcl {
	flowJsons := getFlowJsons(resourceName)
	return TestFlowsHcl{
		Simple: makeFlowHcl(resourceName, flowResource{
			Name:        "simple",
			FlowJson:    flowJsons.Simple,
			Connections: []string{"http", "functions"},
		}),
		Drifted: makeFlowHcl(resourceName, flowResource{
			Name:        "simple",
			FlowJson:    flowJsons.Drifted,
			Connections: []string{"http"},
		}),
		Mainflow: makeFlowHcl(resourceName, flowResource{
			Name:        "mainflow",
			FlowJson:    flowJsons.Mainflow,
			Subflows:    []string{"subflow", "another_subflow"},
			Connections: []string{"http", "flow"},
		}),
		MainflowDrifted: makeFlowHcl(resourceName, flowResource{
			Name:        "mainflow",
			FlowJson:    flowJsons.MainflowDrifted,
			Subflows:    []string{"subflow", "another_subflow"},
			Connections: []string{"http", "flow"},
		}),
		AnotherMainflow: makeFlowHcl(resourceName, flowResource{
			Name:        "another_mainflow",
			FlowJson:    flowJsons.AnotherMainflow,
			Subflows:    []string{"subflow"},
			Connections: []string{"http", "flow"},
		}),
		Subflow: makeFlowHcl(resourceName, flowResource{
			Name:        "subflow",
			FlowJson:    flowJsons.Subflow,
			Subflows:    []string{"another_subflow"},
			Connections: []string{"http", "flow"},
		}),
		AnotherSubflow: makeFlowHcl(resourceName, flowResource{
			Name:        "another_subflow",
			FlowJson:    flowJsons.AnotherSubflow,
			Connections: []string{"http"},
		}),
		WithVariableConnector: makeFlowHcl(resourceName, flowResource{
			Name:        "with_variable_connector",
			FlowJson:    flowJsons.WithVariableConnector,
			Connections: []string{"http", "variables"},
		}),
		BrokenFlow: makeFlowHcl(resourceName, flowResource{
			Name:        "broken_flow",
			FlowJson:    flowJsons.BrokenFlow,
			Connections: []string{"errorcustomize"},
		}),
		//notimplemented
		// PingOneSessionFlow: makeFlowHcl(resourceName, flowResource{
		// 	Name:        "pingone_session_flow",
		// 	FlowJson:    flowJsons.PingOneSessionFlow,
		// 	Connections: []string{"annotation", "flow", "variables", "pingoneauthentication", "node"},
		// }),
	}
}

func FlowsForTestsMap(resourceName string) map[string]string {
	var flowsMap map[string]string
	testFlows := FlowsForTests(resourceName)
	mapstructure.Decode(testFlows, &flowsMap)
	return flowsMap
}

// Data read of bootstrapped connections to be used in davinci_flow connections blocks
func BsConnectionsHcl(resourceName string) string {
	var tc string
	for i, conn := range bsConnections {
		hcl := fmt.Sprintf(`
data "davinci_connection" "%[1]s-%[2]s" {
	environment_id = resource.pingone_role_assignment_user.%[1]s.scope_environment_id
	id = "%[3]s"
	depends_on = [data.davinci_connections.read_all]
}

`, resourceName, i, conn.Id)
		tc = tc + hcl
	}
	return tc
}

func KitchenSink(resourceName string) string {
	var hcl string
	baseHcl := PingoneEnvrionmentSsoHcl(resourceName)
	flows := addAllTestFlows(resourceName, baseHcl)

	return hcl + baseHcl + flows
}

func addAllTestFlows(resourceName string, hcl string) string {
	testFlows := FlowsForTestsMap(resourceName)
	var flowHcl string
	for i, flow := range testFlows {
		flowHcl += fmt.Sprintf(`
resource "davinci_flow" "%[2]s-%[1]s" {
	environment_id = resource.pingone_role_assignment_user.%[1]s.scope_environment_id
	depends_on = [data.davinci_connections.read_all]
	flow_json = %[3]s
	deploy = true
}

`, resourceName, i, flow)
	}
	hcl = hcl + flowHcl
	return hcl

}
