resource "davinci_connection" "haveIBeenPwnedConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "haveIBeenPwnedConnector"
  name         = "My awesome haveIBeenPwnedConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.haveibeenpwnedconnector_property_api_key
  }

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.haveibeenpwnedconnector_property_api_url
  }

  property {
    name  = "userAgent"
    type  = "string"
    value = var.haveibeenpwnedconnector_property_user_agent
  }
}
