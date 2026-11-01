resource "davinci_connection" "socureriskosConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "socureriskosConnector"
  name         = "My awesome socureriskosConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.socureriskosconnector_property_api_key
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.socureriskosconnector_property_base_url
  }
}
