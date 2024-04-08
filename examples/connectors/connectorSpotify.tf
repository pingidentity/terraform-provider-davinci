resource "davinci_connection" "connectorSpotify" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorSpotify"
  name         = "My awesome connectorSpotify"

  property {
    name  = "oauth2"
    type  = "json"
    value = jsonencode({})
  }
}
