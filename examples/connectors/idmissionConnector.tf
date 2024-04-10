resource "davinci_connection" "idmissionConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "idmissionConnector"
  name         = "My awesome idmissionConnector"

  property {
    name  = "authDescription"
    type  = "string"
    value = var.idmissionconnector_property_auth_description
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.idmissionconnector_property_connector_name
  }

  property {
    name  = "description"
    type  = "string"
    value = var.idmissionconnector_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.idmissionconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.idmissionconnector_property_details2
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.idmissionconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.idmissionconnector_property_icon_url_png
  }

  property {
    name  = "loginId"
    type  = "string"
    value = var.idmissionconnector_property_login_id
  }

  property {
    name  = "merchantId"
    type  = "string"
    value = var.idmissionconnector_property_merchant_id
  }

  property {
    name  = "password"
    type  = "string"
    value = var.idmissionconnector_property_password
  }

  property {
    name  = "productId"
    type  = "string"
    value = var.idmissionconnector_property_product_id
  }

  property {
    name  = "productName"
    type  = "string"
    value = var.idmissionconnector_property_product_name
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.idmissionconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.idmissionconnector_property_show_cred_added_via
  }

  property {
    name  = "title"
    type  = "string"
    value = var.idmissionconnector_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.idmissionconnector_property_tool_tip
  }

  property {
    name  = "url"
    type  = "string"
    value = var.idmissionconnector_property_url
  }
}
