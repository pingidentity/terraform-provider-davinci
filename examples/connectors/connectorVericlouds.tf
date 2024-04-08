resource "davinci_connection" "connectorVericlouds" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorVericlouds"
  name         = "My awesome connectorVericlouds"

  property {
    name  = "apiSecret"
    type  = "string"
    value = var.connectorvericlouds_property_api_secret
  }

  property {
    name  = "apikey"
    type  = "string"
    value = var.connectorvericlouds_property_apikey
  }
}
