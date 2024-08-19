resource "davinci_connection" "connector-oai-talendim" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-talendim"
  name         = "My awesome connector-oai-talendim"

  property {
    name  = "authBearerToken"
    type  = "string"
    value = var.connector-oai-talendim_property_auth_bearer_token
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-talendim_property_base_path
  }
}
