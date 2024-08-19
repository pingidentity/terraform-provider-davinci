resource "davinci_connection" "connector-oai-authomizeapireference" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-authomizeapireference"
  name         = "My awesome connector-oai-authomizeapireference"

  property {
    name  = "authApiKey"
    type  = "string"
    value = var.connector-oai-authomizeapireference_property_auth_api_key
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-authomizeapireference_property_base_path
  }
}
