resource "davinci_connection" "connector-oai-mailjetapi" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-mailjetapi"
  name         = "My awesome connector-oai-mailjetapi"

  property {
    name  = "authPassword"
    type  = "string"
    value = var.connector-oai-mailjetapi_property_auth_password
  }

  property {
    name  = "authUsername"
    type  = "string"
    value = var.connector-oai-mailjetapi_property_auth_username
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-mailjetapi_property_base_path
  }
}
