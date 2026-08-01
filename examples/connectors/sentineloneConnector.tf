resource "davinci_connection" "sentineloneConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "sentineloneConnector"
  name         = "My awesome sentineloneConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.sentineloneconnector_property_api_key
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.sentineloneconnector_property_base_url
  }
}
