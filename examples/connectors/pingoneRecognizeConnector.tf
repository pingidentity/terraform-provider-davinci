resource "davinci_connection" "pingoneRecognizeConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingoneRecognizeConnector"
  name         = "My awesome pingoneRecognizeConnector"
}
