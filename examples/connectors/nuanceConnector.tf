resource "davinci_connection" "nuanceConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "nuanceConnector"
  name         = "My awesome nuanceConnector"

  property {
    name  = "authDescription"
    type  = "string"
    value = var.nuanceconnector_property_auth_description
  }

  property {
    name  = "configSetName"
    type  = "string"
    value = var.nuanceconnector_property_config_set_name
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.nuanceconnector_property_connector_name
  }

  property {
    name  = "description"
    type  = "string"
    value = var.nuanceconnector_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.nuanceconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.nuanceconnector_property_details2
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.nuanceconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.nuanceconnector_property_icon_url_png
  }

  property {
    name  = "passphrase1"
    type  = "string"
    value = var.nuanceconnector_property_passphrase1
  }

  property {
    name  = "passphrase2"
    type  = "string"
    value = var.nuanceconnector_property_passphrase2
  }

  property {
    name  = "passphrase3"
    type  = "string"
    value = var.nuanceconnector_property_passphrase3
  }

  property {
    name  = "passphrase4"
    type  = "string"
    value = var.nuanceconnector_property_passphrase4
  }

  property {
    name  = "passphrase5"
    type  = "string"
    value = var.nuanceconnector_property_passphrase5
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.nuanceconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.nuanceconnector_property_show_cred_added_via
  }

  property {
    name  = "title"
    type  = "string"
    value = var.nuanceconnector_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.nuanceconnector_property_tool_tip
  }
}
