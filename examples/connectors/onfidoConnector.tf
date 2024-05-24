resource "davinci_connection" "onfidoConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "onfidoConnector"
  name         = "My awesome onfidoConnector"

  property {
    name  = "androidPackageName"
    type  = "string"
    value = var.onfidoconnector_property_android_package_name
  }

  property {
    name  = "apiKey"
    type  = "string"
    value = var.onfidoconnector_property_api_key
  }

  property {
    name  = "authDescription"
    type  = "string"
    value = var.onfidoconnector_property_auth_description
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.onfidoconnector_property_base_url
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.onfidoconnector_property_connector_name
  }

  property {
    name  = "customizeSteps"
    type  = "boolean"
    value = var.onfidoconnector_property_customize_steps
  }

  property {
    name  = "description"
    type  = "string"
    value = var.onfidoconnector_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.onfidoconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.onfidoconnector_property_details2
  }

  property {
    name  = "iOSBundleId"
    type  = "string"
    value = var.onfidoconnector_property_i_o_s_bundle_id
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.onfidoconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.onfidoconnector_property_icon_url_png
  }

  property {
    name  = "javascriptCSSUrl"
    type  = "string"
    value = var.javascript_css_url
  }

  property {
    name  = "javascriptCdnUrl"
    type  = "string"
    value = var.onfidoconnector_property_javascript_cdn_url
  }

  property {
    name  = "language"
    type  = "string"
    value = var.onfidoconnector_property_language
  }

  property {
    name  = "referenceStepsList"
    type  = "json"
    value = var.onfidoconnector_property_reference_steps_list
  }

  property {
    name  = "referrerUrl"
    type  = "string"
    value = var.onfidoconnector_property_referrer_url
  }

  property {
    name  = "retrieveReports"
    type  = "boolean"
    value = var.onfidoconnector_property_retrieve_reports
  }

  property {
    name  = "shouldCloseOnOverlayClick"
    type  = "boolean"
    value = var.onfidoconnector_property_should_close_on_overlay_click
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.onfidoconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.onfidoconnector_property_show_cred_added_via
  }

  property {
    name  = "stepsList"
    type  = "boolean"
    value = var.onfidoconnector_property_steps_list
  }

  property {
    name  = "title"
    type  = "string"
    value = var.onfidoconnector_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.onfidoconnector_property_tool_tip
  }

  property {
    name  = "useLanguage"
    type  = "boolean"
    value = var.onfidoconnector_property_use_language
  }

  property {
    name  = "useModal"
    type  = "boolean"
    value = var.onfidoconnector_property_use_modal
  }

  property {
    name  = "viewDescriptions"
    type  = "string"
    value = var.onfidoconnector_property_view_descriptions
  }

  property {
    name  = "viewTitle"
    type  = "string"
    value = var.onfidoconnector_property_view_title
  }
}
