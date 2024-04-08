resource "davinci_connection" "kaizenVoizConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "kaizenVoizConnector"
  name         = "My awesome kaizenVoizConnector"

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.kaizenvoizconnector_property_api_url
  }

  property {
    name  = "applicationName"
    type  = "string"
    value = var.kaizenvoizconnector_property_application_name
  }

  property {
    name  = "authDescription"
    type  = "string"
    value = var.kaizenvoizconnector_property_auth_description
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.kaizenvoizconnector_property_connector_name
  }

  property {
    name  = "description"
    type  = "string"
    value = var.kaizenvoizconnector_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.kaizenvoizconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.kaizenvoizconnector_property_details2
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.kaizenvoizconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.kaizenvoizconnector_property_icon_url_png
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.kaizenvoizconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.kaizenvoizconnector_property_show_cred_added_via
  }

  property {
    name  = "title"
    type  = "string"
    value = var.kaizenvoizconnector_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.kaizenvoizconnector_property_tool_tip
  }
}
