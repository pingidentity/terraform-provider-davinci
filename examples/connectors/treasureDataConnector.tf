resource "davinci_connection" "treasureDataConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "treasureDataConnector"
  name         = "My awesome treasureDataConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.treasuredataconnector_property_api_key
  }
}
