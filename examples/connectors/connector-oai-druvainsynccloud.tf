resource "davinci_connection" "connector-oai-druvainsynccloud" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-druvainsynccloud"
  name         = "My awesome connector-oai-druvainsynccloud"

  property {
    name  = "authClientId"
    type  = "string"
    value = var.connector-oai-druvainsynccloud_property_auth_client_id
  }

  property {
    name  = "authClientSecret"
    type  = "string"
    value = var.connector-oai-druvainsynccloud_property_auth_client_secret
  }

  property {
    name  = "authTokenUrl"
    type  = "string"
    value = var.connector-oai-druvainsynccloud_property_auth_token_url
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-druvainsynccloud_property_base_path
  }
}
