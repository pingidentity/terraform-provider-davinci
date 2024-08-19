resource "davinci_connection" "connector-oai-launchdarklyrestapi" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-launchdarklyrestapi"
  name         = "My awesome connector-oai-launchdarklyrestapi"

  property {
    name  = "authApiKey"
    type  = "string"
    value = var.connector-oai-launchdarklyrestapi_property_auth_api_key
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-launchdarklyrestapi_property_base_path
  }
}
