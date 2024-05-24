resource "davinci_connection" "argyleConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "argyleConnector"
  name         = "My awesome argyleConnector"

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.argyleconnector_property_api_url
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.argyleconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.argyleconnector_property_client_secret
  }

  property {
    name  = "javascriptWebUrl"
    type  = "string"
    value = var.argyleconnector_property_javascript_web_url
  }

  property {
    name  = "pluginKey"
    type  = "string"
    value = var.argyleconnector_property_plugin_key
  }
}
