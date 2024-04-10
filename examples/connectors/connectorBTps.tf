resource "davinci_connection" "connectorBTps" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorBTps"
  name         = "My awesome connectorBTps"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectorbtps_property_api_key
  }

  property {
    name  = "apiUser"
    type  = "string"
    value = var.connectorbtps_property_api_user
  }

  property {
    name  = "domain"
    type  = "string"
    value = var.connectorbtps_property_domain
  }
}
