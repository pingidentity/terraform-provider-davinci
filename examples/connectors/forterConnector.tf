resource "davinci_connection" "forterConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "forterConnector"
  name         = "My awesome forterConnector"

  property {
    name  = "apiVersion"
    type  = "string"
    value = var.forterconnector_property_api_version
  }

  property {
    name  = "secretKey"
    type  = "string"
    value = var.forterconnector_property_secret_key
  }

  property {
    name  = "siteId"
    type  = "string"
    value = var.forterconnector_property_site_id
  }
}
