resource "davinci_connection" "kbaConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "kbaConnector"
  name         = "My awesome kbaConnector"

  property {
    name  = "authDescription"
    type  = "string"
    value = var.kbaconnector_property_auth_description
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.kbaconnector_property_connector_name
  }

  property {
    name  = "description"
    type  = "string"
    value = var.kbaconnector_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.kbaconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.kbaconnector_property_details2
  }

  property {
    name  = "formFieldsList"
    type  = "json"
    value = var.kbaconnector_property_form_fields_list
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.kbaconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.kbaconnector_property_icon_url_png
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.kbaconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.kbaconnector_property_show_cred_added_via
  }

  property {
    name  = "title"
    type  = "string"
    value = var.kbaconnector_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.kbaconnector_property_tool_tip
  }
}
