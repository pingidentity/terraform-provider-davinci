resource "davinci_connection" "fingerprintjsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "fingerprintjsConnector"
  name         = "My awesome fingerprintjsConnector"

  property {
    name  = "apiToken"
    type  = "string"
    value = var.fingerprintjsconnector_property_api_token
  }

  property {
    name  = "javascriptCdnUrl"
    type  = "string"
    value = var.fingerprintjsconnector_property_javascript_cdn_url
  }

  property {
    name  = "token"
    type  = "string"
    value = var.fingerprintjsconnector_property_token
  }
}
