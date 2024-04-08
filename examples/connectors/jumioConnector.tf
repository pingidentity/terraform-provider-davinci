resource "davinci_connection" "jumioConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "jumioConnector"
  name         = "My awesome jumioConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.jumioconnector_property_api_key
  }

  property {
    name  = "authDescription"
    type  = "string"
    value = var.jumioconnector_property_auth_description
  }

  property {
    name  = "authUrl"
    type  = "string"
    value = var.jumioconnector_property_auth_url
  }

  property {
    name  = "authorizationTokenLifetime"
    type  = "number"
    value = var.jumioconnector_property_authorization_token_lifetime
  }

  property {
    name  = "baseColor"
    type  = "string"
    value = var.jumioconnector_property_base_color
  }

  property {
    name  = "bgColor"
    type  = "string"
    value = var.jumioconnector_property_bg_color
  }

  property {
    name  = "callbackUrl"
    type  = "string"
    value = var.jumioconnector_property_callback_url
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.jumioconnector_property_client_secret
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.jumioconnector_property_connector_name
  }

  property {
    name  = "description"
    type  = "string"
    value = var.jumioconnector_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.jumioconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.jumioconnector_property_details2
  }

  property {
    name  = "doNotShowInIframe"
    type  = "boolean"
    value = var.jumioconnector_property_do_not_show_in_iframe
  }

  property {
    name  = "docVerificationUrl"
    type  = "string"
    value = var.jumioconnector_property_doc_verification_url
  }

  property {
    name  = "headerImageUrl"
    type  = "string"
    value = var.jumioconnector_property_header_image_url
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.jumioconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.jumioconnector_property_icon_url_png
  }

  property {
    name  = "locale"
    type  = "string"
    value = var.jumioconnector_property_locale
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.jumioconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.jumioconnector_property_show_cred_added_via
  }

  property {
    name  = "title"
    type  = "string"
    value = var.jumioconnector_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.jumioconnector_property_tool_tip
  }
}
