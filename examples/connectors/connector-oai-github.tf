resource "davinci_connection" "connector-oai-github" {
  environment_id = var.pingone_environment_id

  connector_id = "connector-oai-github"
  name         = "My awesome connector-oai-github"

  property {
    name  = "apiVersion"
    type  = "string"
    value = var.connector-oai-github_property_api_version
  }

  property {
    name  = "authBearerToken"
    type  = "string"
    value = var.connector-oai-github_property_auth_bearer_token
  }

  property {
    name  = "basePath"
    type  = "string"
    value = var.connector-oai-github_property_base_path
  }
}
