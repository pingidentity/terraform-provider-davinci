resource "davinci_connection" "complyAdvatangeConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "complyAdvatangeConnector"
  name         = "My awesome complyAdvatangeConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.complyadvatangeconnector_property_api_key
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.complyadvatangeconnector_property_base_url
  }
}
