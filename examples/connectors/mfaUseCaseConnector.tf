resource "davinci_connection" "mfaUseCaseConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "mfaUseCaseConnector"
  name         = "My awesome mfaUseCaseConnector"
}
