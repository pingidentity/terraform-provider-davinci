//Read all connections - This is a good first call to make
data "davinci_connections" "all" {
}

resource "davinci_connection" "subflow" {
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
  subflows {
    id   = resource.davinci_flow.subflow.id
    name = resource.davinci_flow.subflow.name
  }
  // Dependent connections are defined in conections blocks. 
  // It is best practice to define all connections referenced the flow_json. This prevents a mismatch between the flow_json and the connections

  // This sample references a managed connection
  connections {
    id              = davinci_connection.subflow.id
    connection_name = davinci_connection.subflow.name
  }
  // This sample uses a bootstrapped connection
  connections {
    name = "Http"
    // default connection id for the bootstrapped Http connector
    id = "867ed4363b2bc21c860085ad2baa817d"
  }

}

resource "davinci_flow" "subflow" {
  environment_id = var.environment_id
  flow_json      = file("subflow.json")
  deploy         = true
  connections {
    id              = "867ed4363b2bc21c860085ad2baa817d"
    connection_name = "Http"
  }
  connections {
    id              = davinci_connection.subflow.id
    connection_name = davinci_connection.subflow.name
  }
}
