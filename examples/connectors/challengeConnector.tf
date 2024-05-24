resource "davinci_connection" "challengeConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "challengeConnector"
  name         = "My awesome challengeConnector"
}
