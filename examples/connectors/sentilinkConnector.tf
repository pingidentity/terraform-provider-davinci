resource "davinci_connection" "sentilinkConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "sentilinkConnector"
  name         = "My awesome sentilinkConnector"

  property {
    name  = "account"
    type  = "string"
    value = var.sentilinkconnector_property_account
  }

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.sentilinkconnector_property_api_url
  }

  property {
    name  = "javascriptCdnUrl"
    type  = "string"
    value = var.sentilinkconnector_property_javascript_cdn_url
  }

  property {
    name  = "token"
    type  = "string"
    value = var.sentilinkconnector_property_token
  }
}
