resource "davinci_connection" "authenticationUseCaseConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "authenticationUseCaseConnector"
  name         = "My awesome authenticationUseCaseConnector"
}
