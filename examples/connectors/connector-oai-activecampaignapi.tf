resource "davinci_connection" "connector-oai-activecampaignapi" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-activecampaignapi"
  name         = "My awesome connector-oai-activecampaignapi"

  property {
    name  = "authApiKey"
    type  = "string"
    value = var.connector-oai-activecampaignapi_property_auth_api_key
  }

  property {
    name  = "authApiVersion"
    type  = "string"
    value = var.connector-oai-activecampaignapi_property_auth_api_version
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-activecampaignapi_property_base_path
  }
}
