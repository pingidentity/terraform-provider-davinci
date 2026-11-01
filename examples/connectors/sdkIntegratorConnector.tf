resource "davinci_connection" "sdkIntegratorConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "sdkIntegratorConnector"
  name         = "My awesome sdkIntegratorConnector"
}
