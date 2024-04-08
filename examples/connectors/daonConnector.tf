resource "davinci_connection" "daonConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "daonConnector"
  name         = "My awesome daonConnector"

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.daonconnector_property_api_url
  }

  property {
    name  = "password"
    type  = "string"
    value = var.daonconnector_property_password
  }

  property {
    name  = "username"
    type  = "string"
    value = var.daonconnector_property_username
  }
}
