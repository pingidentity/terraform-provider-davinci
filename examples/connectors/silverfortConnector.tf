resource "davinci_connection" "silverfortConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "silverfortConnector"
  name         = "My awesome silverfortConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.silverfortconnector_property_api_key
  }

  property {
    name  = "appUserSecret"
    type  = "string"
    value = var.silverfortconnector_property_app_user_secret
  }

  property {
    name  = "consoleApi"
    type  = "string"
    value = var.silverfortconnector_property_console_api
  }
}
