resource "davinci_connection" "splunkConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "splunkConnector"
  name         = "My awesome splunkConnector"

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.splunkconnector_property_api_url
  }

  property {
    name  = "port"
    type  = "number"
    value = var.splunkconnector_property_port
  }

  property {
    name  = "token"
    type  = "string"
    value = var.splunkconnector_property_token
  }
}
