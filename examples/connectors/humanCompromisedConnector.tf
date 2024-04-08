resource "davinci_connection" "humanCompromisedConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "humanCompromisedConnector"
  name         = "My awesome humanCompromisedConnector"

  property {
    name  = "appId"
    type  = "string"
    value = var.humancompromisedconnector_property_app_id
  }

  property {
    name  = "authToken"
    type  = "string"
    value = var.humancompromisedconnector_property_auth_token
  }
}
