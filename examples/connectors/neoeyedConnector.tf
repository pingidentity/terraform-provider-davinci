resource "davinci_connection" "neoeyedConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "neoeyedConnector"
  name         = "My awesome neoeyedConnector"

  property {
    name  = "appKey"
    type  = "string"
    value = var.neoeyedconnector_property_app_key
  }

  property {
    name  = "javascriptCdnUrl"
    type  = "string"
    value = var.neoeyedconnector_property_javascript_cdn_url
  }
}
