resource "davinci_connection" "intellicheckConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "intellicheckConnector"
  name         = "My awesome intellicheckConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.intellicheckconnector_property_api_key
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.intellicheckconnector_property_base_url
  }

  property {
    name  = "customerId"
    type  = "string"
    value = var.intellicheckconnector_property_customer_id
  }
}
