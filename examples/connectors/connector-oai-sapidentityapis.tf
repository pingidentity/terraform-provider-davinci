resource "davinci_connection" "connector-oai-sapidentityapis" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-sapidentityapis"
  name         = "My awesome connector-oai-sapidentityapis"

  property {
    name  = "authApiKey"
    type  = "string"
    value = var.connector-oai-sapidentityapis_property_auth_api_key
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-sapidentityapis_property_base_path
  }
}
