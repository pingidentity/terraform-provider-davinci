resource "davinci_connection" "connectorJamf" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorJamf"
  name         = "My awesome connectorJamf"

  property {
    name  = "jamfPassword"
    type  = "string"
    value = var.connectorjamf_property_jamf_password
  }

  property {
    name  = "jamfUsername"
    type  = "string"
    value = var.connectorjamf_property_jamf_username
  }

  property {
    name  = "serverName"
    type  = "string"
    value = var.connectorjamf_property_server_name
  }
}
