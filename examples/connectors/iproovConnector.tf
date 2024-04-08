resource "davinci_connection" "iproovConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "iproovConnector"
  name         = "My awesome iproovConnector"

  property {
    name  = "allowLandscape"
    type  = "boolean"
    value = var.iproovconnector_property_allow_landscape
  }

  property {
    name  = "apiKey"
    type  = "string"
    value = var.iproovconnector_property_api_key
  }

  property {
    name  = "authDescription"
    type  = "string"
    value = var.iproovconnector_property_auth_description
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.iproovconnector_property_base_url
  }

  property {
    name  = "color1"
    type  = "string"
    value = var.iproovconnector_property_color1
  }

  property {
    name  = "color2"
    type  = "string"
    value = var.iproovconnector_property_color2
  }

  property {
    name  = "color3"
    type  = "string"
    value = var.iproovconnector_property_color3
  }

  property {
    name  = "color4"
    type  = "string"
    value = var.iproovconnector_property_color4
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.iproovconnector_property_connector_name
  }

  property {
    name  = "customTitle"
    type  = "string"
    value = var.iproovconnector_property_custom_title
  }

  property {
    name  = "description"
    type  = "string"
    value = var.iproovconnector_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.iproovconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.iproovconnector_property_details2
  }

  property {
    name  = "enableCameraSelector"
    type  = "boolean"
    value = var.iproovconnector_property_enable_camera_selector
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.iproovconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.iproovconnector_property_icon_url_png
  }

  property {
    name  = "javascriptCSSUrl"
    type  = "string"
    value = var.javascript_css_url
  }

  property {
    name  = "javascriptCdnUrl"
    type  = "string"
    value = var.iproovconnector_property_javascript_cdn_url
  }

  property {
    name  = "kioskMode"
    type  = "boolean"
    value = var.iproovconnector_property_kiosk_mode
  }

  property {
    name  = "logo"
    type  = "string"
    value = var.iproovconnector_property_logo
  }

  property {
    name  = "password"
    type  = "string"
    value = var.iproovconnector_property_password
  }

  property {
    name  = "secret"
    type  = "string"
    value = var.iproovconnector_property_secret
  }

  property {
    name  = "showCountdown"
    type  = "boolean"
    value = var.iproovconnector_property_show_countdown
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.iproovconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.iproovconnector_property_show_cred_added_via
  }

  property {
    name  = "startScreenTitle"
    type  = "string"
    value = var.iproovconnector_property_start_screen_title
  }

  property {
    name  = "title"
    type  = "string"
    value = var.iproovconnector_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.iproovconnector_property_tool_tip
  }

  property {
    name  = "username"
    type  = "string"
    value = var.iproovconnector_property_username
  }
}
