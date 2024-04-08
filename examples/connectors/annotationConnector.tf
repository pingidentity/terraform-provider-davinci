resource "davinci_connection" "annotationConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "annotationConnector"
  name         = "My awesome annotationConnector"
}
