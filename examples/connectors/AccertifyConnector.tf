resource "davinci_connection" "AccertifyConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "AccertifyConnector"
  name         = "My awesome AccertifyConnector"

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.accertifyconnector_property_base_url
  }

  property {
    name  = "basicAuthPassword"
    type  = "string"
    value = var.accertifyconnector_property_basic_auth_password
  }

  property {
    name  = "basicAuthUsername"
    type  = "string"
    value = var.accertifyconnector_property_basic_auth_username
  }
}
