resource "davinci_connection" "connectorSmarty" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorSmarty"
  name         = "My awesome connectorSmarty"

  property {
    name  = "authId"
    type  = "string"
    value = var.connectorsmarty_property_auth_id
  }

  property {
    name  = "authToken"
    type  = "string"
    value = var.connectorsmarty_property_auth_token
  }

  property {
    name  = "license"
    type  = "string"
    value = var.connectorsmarty_property_license
  }
}
