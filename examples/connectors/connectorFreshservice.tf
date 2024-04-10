resource "davinci_connection" "connectorFreshservice" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorFreshservice"
  name         = "My awesome connectorFreshservice"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectorfreshservice_property_api_key
  }

  property {
    name  = "domain"
    type  = "string"
    value = var.connectorfreshservice_property_domain
  }
}
