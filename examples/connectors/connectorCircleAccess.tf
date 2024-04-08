resource "davinci_connection" "connectorCircleAccess" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorCircleAccess"
  name         = "My awesome connectorCircleAccess"

  property {
    name  = "appKey"
    type  = "string"
    value = var.connectorcircleaccess_property_app_key
  }

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }

  property {
    name  = "loginUrl"
    type  = "string"
    value = var.connectorcircleaccess_property_login_url
  }

  property {
    name  = "readKey"
    type  = "string"
    value = var.connectorcircleaccess_property_read_key
  }

  property {
    name  = "returnToUrl"
    type  = "string"
    value = var.connectorcircleaccess_property_return_to_url
  }

  property {
    name  = "writeKey"
    type  = "string"
    value = var.connectorcircleaccess_property_write_key
  }
}
