//Read all connections - This is a good first call to make
data "davinci_connections" "all" {
  name           = "Flow"
  connector_ids  = "flowConnector"
  environment_id = var.environment_id
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
  subflows {
    subflow_id   = resource.davinci_flow.subflow.flow_id
    subflow_name = resource.davinci_flow.subflow.flow_name
  }
  connections {
    //Bootstrapped connection
    connection_id   = "867ed4363b2bc21c860085ad2baa817d"
    connection_name = "Http"
  }
  connections {
    connection_id   = davinci_connection.subflow.id
    connection_name = davinci_connection.subflow.name
  }
}

resource "davinci_flow" "subflow" {
  environment_id = var.environment_id
  flow_json      = file("subflow.json")
  deploy         = true
  connections {
    connection_id   = "867ed4363b2bc21c860085ad2baa817d"
    connection_name = "Http"
  }
  connections {
    connection_id   = davinci_connection.subflow.id
    connection_name = davinci_connection.subflow.name
  }
}
