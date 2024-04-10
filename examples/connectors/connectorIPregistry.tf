resource "davinci_connection" "connectorIPregistry" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorIPregistry"
  name         = "My awesome connectorIPregistry"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectoripregistry_property_api_key
  }
}
