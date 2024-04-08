resource "davinci_connection" "securIdConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "securIdConnector"
  name         = "My awesome securIdConnector"

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.securidconnector_property_api_url
  }

  property {
    name  = "clientKey"
    type  = "string"
    value = var.securidconnector_property_client_key
  }
}
