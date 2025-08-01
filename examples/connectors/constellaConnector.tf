resource "davinci_connection" "constellaConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "constellaConnector"
  name         = "My awesome constellaConnector"

  property {
    name  = "appToken"
    type  = "string"
    value = var.constellaconnector_property_app_token
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.constellaconnector_property_base_url
  }

  property {
    name  = "token"
    type  = "string"
    value = var.constellaconnector_property_token
  }

  property {
    name  = "username"
    type  = "string"
    value = var.constellaconnector_property_username
  }
}
