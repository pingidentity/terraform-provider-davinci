resource "davinci_connection" "connector-oai-talendscim" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-talendscim"
  name         = "My awesome connector-oai-talendscim"

  property {
    name  = "authBearerToken"
    type  = "string"
    value = var.connector-oai-talendscim_property_auth_bearer_token
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-talendscim_property_base_path
  }
}
