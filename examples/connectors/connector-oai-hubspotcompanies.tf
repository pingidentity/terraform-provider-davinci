resource "davinci_connection" "connector-oai-hubspotcompanies" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-hubspotcompanies"
  name         = "My awesome connector-oai-hubspotcompanies"

  property {
    name  = "authBearerToken"
    type  = "string"
    value = var.connector-oai-hubspotcompanies_property_auth_bearer_token
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-hubspotcompanies_property_base_path
  }
}
