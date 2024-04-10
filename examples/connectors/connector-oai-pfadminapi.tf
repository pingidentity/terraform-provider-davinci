resource "davinci_connection" "connector-oai-pfadminapi" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-pfadminapi"
  name         = "My awesome connector-oai-pfadminapi"

  property {
    name  = "authPassword"
    type  = "string"
    value = var.connector-oai-pfadminapi_property_auth_password
  }

  property {
    name  = "authUsername"
    type  = "string"
    value = var.connector-oai-pfadminapi_property_auth_username
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-pfadminapi_property_base_path
  }

  property {
    name  = "sslVerification"
    type  = "string"
    value = var.connector-oai-pfadminapi_property_ssl_verification
  }
}
