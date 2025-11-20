resource "davinci_connection" "babelStreetConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "babelStreetConnector"
  name         = "My awesome babelStreetConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.babelstreetconnector_property_api_key
  }
}
