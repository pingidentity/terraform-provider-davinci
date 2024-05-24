resource "davinci_connection" "proveConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "proveConnector"
  name         = "My awesome proveConnector"

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.proveconnector_property_base_url
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.proveconnector_property_client_id
  }

  property {
    name  = "grantType"
    type  = "string"
    value = var.proveconnector_property_grant_type
  }

  property {
    name  = "password"
    type  = "string"
    value = var.proveconnector_property_password
  }

  property {
    name  = "username"
    type  = "string"
    value = var.proveconnector_property_username
  }
}
