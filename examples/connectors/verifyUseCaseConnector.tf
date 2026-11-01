resource "davinci_connection" "verifyUseCaseConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "verifyUseCaseConnector"
  name         = "My awesome verifyUseCaseConnector"
}
