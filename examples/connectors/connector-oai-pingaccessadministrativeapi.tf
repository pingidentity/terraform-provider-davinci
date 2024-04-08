resource "davinci_connection" "connector-oai-pingaccessadministrativeapi" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-pingaccessadministrativeapi"
  name         = "My awesome connector-oai-pingaccessadministrativeapi"

  property {
    name  = "authPassword"
    type  = "string"
    value = var.connector-oai-pingaccessadministrativeapi_property_auth_password
  }

  property {
    name  = "authUsername"
    type  = "string"
    value = var.connector-oai-pingaccessadministrativeapi_property_auth_username
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-pingaccessadministrativeapi_property_base_path
  }

  property {
    name  = "sslVerification"
    type  = "string"
    value = var.connector-oai-pingaccessadministrativeapi_property_ssl_verification
  }
}
