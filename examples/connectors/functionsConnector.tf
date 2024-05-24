resource "davinci_connection" "functionsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "functionsConnector"
  name         = "My awesome functionsConnector"
}
