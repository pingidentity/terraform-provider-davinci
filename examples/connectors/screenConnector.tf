resource "davinci_connection" "screenConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "screenConnector"
  name         = "My awesome screenConnector"
}
