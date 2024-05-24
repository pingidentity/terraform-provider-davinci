resource "davinci_connection" "errorConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "errorConnector"
  name         = "My awesome errorConnector"
}
