resource "davinci_connection" "imageConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "imageConnector"
  name         = "My awesome imageConnector"
}
