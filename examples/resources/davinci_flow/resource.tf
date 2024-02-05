resource "davinci_connection" "my_awesome_flow_connector" {
  environment_id = var.environment_id

  name         = "Flow"
  connector_id = "flowConnector"
}

resource "davinci_flow" "my_awesome_subflow" {
  environment_id = var.environment_id

  flow_json = jsondecode(file("./path/to/example-subflow.json"))
  name      = "My Awesome Subflow"
  deploy    = true

  connection_link {
    id                           = davinci_connection.my_awesome_flow_connector.id
    name                         = davinci_connection.my_awesome_flow_connector.name
    replace_import_connection_id = "33329a264e268ab31fb19637debf1ea3"
  }
}

resource "davinci_flow" "my_awesome_main_flow" {
  environment_id = var.environment_id

  flow_json = jsondecode(file("./path/to/example-mainflow.json"))
  name      = "My Awesome Main Flow"
  deploy    = true

  subflow_link {
    id                        = resource.davinci_flow.my_awesome_subflow.id
    name                      = resource.davinci_flow.my_awesome_subflow.name
    replace_import_subflow_id = "07503fed5c02849dbbd5ee932da654b2"
  }

  connection_link {
    id                           = davinci_connection.my_awesome_flow_connector.id
    name                         = davinci_connection.my_awesome_flow_connector.name
    replace_import_connection_id = "33329a264e268ab31fb19637debf1ea3"
  }
}