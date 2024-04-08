resource "davinci_connection" "digilockerConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "digilockerConnector"
  name         = "My awesome digilockerConnector"

  property {
    name  = "oauth2"
    type  = "json"
    value = jsonencode({})
  }
}
