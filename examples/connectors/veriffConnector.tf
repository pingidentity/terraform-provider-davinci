resource "davinci_connection" "veriffConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "veriffConnector"
  name         = "My awesome veriffConnector"

  property {
    name  = "access_token"
    type  = "string"
    value = var.veriffconnector_property_access_token
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.veriffconnector_property_base_url
  }

  property {
    name  = "password"
    type  = "string"
    value = var.veriffconnector_property_password
  }
}
