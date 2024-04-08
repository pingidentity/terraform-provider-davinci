resource "davinci_connection" "connectorFreshdesk" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorFreshdesk"
  name         = "My awesome connectorFreshdesk"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectorfreshdesk_property_api_key
  }

  property {
    name  = "baseURL"
    type  = "string"
    value = var.base_url
  }

  property {
    name  = "version"
    type  = "string"
    value = var.connectorfreshdesk_property_version
  }
}
