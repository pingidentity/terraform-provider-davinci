resource "davinci_connection" "symc" {
  environment_id = var.pingone_environment_id

  connector_id = "symc"
  name         = "My awesome symc"

  property {
    name  = "authDescription"
    type  = "string"
    value = var.symc_property_auth_description
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.symc_property_connector_name
  }

  property {
    name  = "description"
    type  = "string"
    value = var.symc_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.symc_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.symc_property_details2
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.symc_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.symc_property_icon_url_png
  }

  property {
    name  = "pfxBase64"
    type  = "string"
    value = var.symc_property_pfx_base64
  }

  property {
    name  = "pfxPassword"
    type  = "string"
    value = var.symc_property_pfx_password
  }

  property {
    name  = "pushLoginEnabled"
    type  = "boolean"
    value = var.symc_property_push_login_enabled
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.symc_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.symc_property_show_cred_added_via
  }

  property {
    name  = "title"
    type  = "string"
    value = var.symc_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.symc_property_tool_tip
  }
}
