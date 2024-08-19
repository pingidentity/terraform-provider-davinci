resource "davinci_connection" "connector-oai-venafi" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-venafi"
  name         = "My awesome connector-oai-venafi"

  property {
    name  = "authApiKey"
    type  = "string"
    value = var.connector-oai-venafi_property_auth_api_key
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-venafi_property_base_path
  }
}
