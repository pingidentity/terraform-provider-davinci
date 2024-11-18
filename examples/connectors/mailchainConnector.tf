resource "davinci_connection" "mailchainConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "mailchainConnector"
  name         = "My awesome mailchainConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.mailchainconnector_property_api_key
  }

  property {
    name  = "version"
    type  = "string"
    value = var.mailchainconnector_property_version
  }
}
