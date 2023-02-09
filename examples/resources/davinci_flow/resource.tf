//Read all connections - This is a good first call to make
data "davinci_connections" "all" {
}

resource "davinci_connection" "flow" {
  name           = "Flow"
  connector_id   = "flowConnector"
  environment_id = var.environment_id
  // Forcing dependency on the inital connection provides better consistency when waiting for bootstrap to complete
  depends_on = [data.davinci_connections.all]
}

resource "davinci_flow" "mainflow" {
  environment_id = var.environment_id
  flow_json      = "{\"customerId\":\"1234\",\"name\":\"mainflow\",\"description\":\"\",\"flowStatus\":\"enabled\",\"createdDate...\"\"connectorIds\":[\"httpConnector\",\"flowConnector\"],\"savedDate\":1662961640542,\"variables\":[]}"
  deploy         = true

  // Dependent subflows are defined in subflows blocks.
  // These should always point to managed subflows
  subflow_link {
    id   = resource.davinci_flow.subflow.id
    name = resource.davinci_flow.subflow.name
  }
  // Dependent connections are defined in conections blocks. 
  // It is best practice to define all connections referenced the flow_json. This prevents a mismatch between the flow_json and the connections

  // This sample references a managed connection
  connection_link {
    id   = davinci_connection.flow.id
    name = davinci_connection.flow.name
  }
  // This sample uses a bootstrapped connection
  connection_link {
    name = "Http"
    // default connection id for the bootstrapped Http connector
    id = "867ed4363b2bc21c860085ad2baa817d"
  }

}

resource "davinci_flow" "subflow" {
  environment_id = var.environment_id
  flow_json      = file("subflow.json")
  deploy         = true
  connection_link {
    id   = "867ed4363b2bc21c860085ad2baa817d"
    name = "Http"
  }
  connection_link {
    id   = davinci_connection.flow.id
    name = davinci_connection.flow.name
  }
}
