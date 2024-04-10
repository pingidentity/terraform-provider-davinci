resource "davinci_connection" "credovaConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "credovaConnector"
  name         = "My awesome credovaConnector"

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.credovaconnector_property_base_url
  }

  property {
    name  = "password"
    type  = "string"
    value = var.credovaconnector_property_password
  }

  property {
    name  = "username"
    type  = "string"
    value = var.credovaconnector_property_username
  }
}
