resource "davinci_variable" "my_awesome_region_variable" {
  environment_id = var.environment_id

  context = "company"

  name        = "region"
  description = "a company variable referenced in the main flow"
  value       = "northamerica"
  type        = "string"
}

resource "davinci_variable" "my_awesome_language_variable" {
  environment_id = var.environment_id

  context = "flowInstance"

  name        = "language"
  description = "a flow instance variable referenced in the sub flow and the main flow"
  value       = "en"
  type        = "string"
}

resource "davinci_connection" "my_awesome_flow_connector" {
  environment_id = var.environment_id

  name         = "Flow"
  connector_id = "flowConnector"
}

resource "davinci_flow" "my_awesome_subflow" {
  depends_on = [
    davinci_variable.my_awesome_language_variable,
  ]

  environment_id = var.environment_id

  name      = "My Awesome Subflow"
  flow_json = file("./path/to/example-subflow.json")

  connection_link {
    id                           = davinci_connection.my_awesome_flow_connector.id
    name                         = davinci_connection.my_awesome_flow_connector.name
    replace_import_connection_id = "33329a264e268ab31fb19637debf1ea3"
  }
}

resource "davinci_flow" "my_awesome_main_flow" {
  depends_on = [
    davinci_variable.my_awesome_region_variable,
    davinci_variable.my_awesome_language_variable,
  ]

  environment_id = var.environment_id

  name      = "My Awesome Main Flow"
  flow_json = file("./path/to/example-mainflow.json")

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

resource "davinci_variable" "my_awesome_context_variable" {
  environment_id = var.environment_id

  context = "flow"
  flow_id = davinci_flow.my_awesome_main_flow.id

  name        = "userContextCode"
  description = "a flow variable used in the main flow"
  type        = "number"
  min         = "10"
}