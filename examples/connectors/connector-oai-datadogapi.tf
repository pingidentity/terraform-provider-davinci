resource "davinci_connection" "connector-oai-datadogapi" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-datadogapi"
  name         = "My awesome connector-oai-datadogapi"

  property {
    name  = "authApiKey"
    type  = "string"
    value = var.connector-oai-datadogapi_property_auth_api_key
  }

  property {
    name  = "authApplicationKey"
    type  = "string"
    value = var.connector-oai-datadogapi_property_auth_application_key
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-datadogapi_property_base_path
  }
}
