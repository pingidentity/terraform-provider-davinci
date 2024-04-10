resource "davinci_connection" "connectorAuthomize" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorAuthomize"
  name         = "My awesome connectorAuthomize"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectorauthomize_property_api_key
  }
}
