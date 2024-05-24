resource "davinci_connection" "httpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "httpConnector"
  name         = "My awesome httpConnector"

  property {
    name  = "connectionId"
    type  = "string"
    value = var.httpconnector_property_connection_id
  }

  property {
    name  = "recaptchaSecretKey"
    type  = "string"
    value = var.httpconnector_property_recaptcha_secret_key
  }

  property {
    name  = "recaptchaSiteKey"
    type  = "string"
    value = var.httpconnector_property_recaptcha_site_key
  }

  property {
    name  = "whiteList"
    type  = "string"
    value = var.httpconnector_property_white_list
  }
}
