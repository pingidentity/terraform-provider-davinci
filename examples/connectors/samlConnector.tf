resource "davinci_connection" "samlConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "samlConnector"
  name         = "My awesome samlConnector"
}
