resource "davinci_connection" "telesignConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "telesignConnector"
  name         = "My awesome telesignConnector"

  property {
    name  = "authDescription"
    type  = "string"
    value = var.telesignconnector_property_auth_description
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.telesignconnector_property_connector_name
  }

  property {
    name  = "description"
    type  = "string"
    value = var.telesignconnector_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.telesignconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.telesignconnector_property_details2
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.telesignconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.telesignconnector_property_icon_url_png
  }

  property {
    name  = "password"
    type  = "string"
    value = var.telesignconnector_property_password
  }

  property {
    name  = "providerName"
    type  = "string"
    value = var.telesignconnector_property_provider_name
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.telesignconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.telesignconnector_property_show_cred_added_via
  }

  property {
    name  = "title"
    type  = "string"
    value = var.telesignconnector_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.telesignconnector_property_tool_tip
  }

  property {
    name  = "username"
    type  = "string"
    value = var.telesignconnector_property_username
  }
}
