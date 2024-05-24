resource "davinci_connection" "castleConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "castleConnector"
  name         = "My awesome castleConnector"

  property {
    name  = "apiSecret"
    type  = "string"
    value = var.castleconnector_property_api_secret
  }
}
