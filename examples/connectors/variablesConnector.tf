resource "davinci_connection" "variablesConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "variablesConnector"
  name         = "My awesome variablesConnector"
}
