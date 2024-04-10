resource "davinci_connection" "payfoneConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "payfoneConnector"
  name         = "My awesome payfoneConnector"

  property {
    name  = "appClientId"
    type  = "string"
    value = var.payfoneconnector_property_app_client_id
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.payfoneconnector_property_base_url
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.payfoneconnector_property_client_id
  }

  property {
    name  = "password"
    type  = "string"
    value = var.payfoneconnector_property_password
  }

  property {
    name  = "simulatorMode"
    type  = "boolean"
    value = var.payfoneconnector_property_simulator_mode
  }

  property {
    name  = "simulatorPhoneNumber"
    type  = "string"
    value = var.payfoneconnector_property_simulator_phone_number
  }

  property {
    name  = "skCallbackBaseUrl"
    type  = "string"
    value = var.payfoneconnector_property_sk_callback_base_url
  }

  property {
    name  = "username"
    type  = "string"
    value = var.payfoneconnector_property_username
  }
}
