resource "davinci_connection" "sinchConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "sinchConnector"
  name         = "My awesome sinchConnector"

  property {
    name  = "acceptLanguage"
    type  = "string"
    value = var.sinchconnector_property_accept_language
  }

  property {
    name  = "applicationKey"
    type  = "string"
    value = var.sinchconnector_property_application_key
  }

  property {
    name  = "secretKey"
    type  = "string"
    value = var.sinchconnector_property_secret_key
  }
}
