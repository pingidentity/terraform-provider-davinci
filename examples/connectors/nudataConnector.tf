resource "davinci_connection" "nudataConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "nudataConnector"
  name         = "My awesome nudataConnector"
}
