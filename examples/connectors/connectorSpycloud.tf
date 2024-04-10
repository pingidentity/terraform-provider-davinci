resource "davinci_connection" "connectorSpycloud" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorSpycloud"
  name         = "My awesome connectorSpycloud"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectorspycloud_property_api_key
  }
}
