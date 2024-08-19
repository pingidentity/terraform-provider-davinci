resource "davinci_connection" "siftConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "siftConnector"
  name         = "My awesome siftConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.siftconnector_property_api_key
  }
}
