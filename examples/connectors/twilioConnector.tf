resource "davinci_connection" "twilioConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "twilioConnector"
  name         = "My awesome twilioConnector"

  property {
    name  = "accountSid"
    type  = "string"
    value = var.twilioconnector_property_account_sid
  }

  property {
    name  = "authDescription"
    type  = "string"
    value = var.twilioconnector_property_auth_description
  }

  property {
    name  = "authMessageTemplate"
    type  = "string"
    value = var.twilioconnector_property_auth_message_template
  }

  property {
    name  = "authToken"
    type  = "string"
    value = var.twilioconnector_property_auth_token
  }

  property {
    name  = "connectorName"
    type  = "string"
    value = var.twilioconnector_property_connector_name
  }

  property {
    name  = "description"
    type  = "string"
    value = var.twilioconnector_property_description
  }

  property {
    name  = "details1"
    type  = "string"
    value = var.twilioconnector_property_details1
  }

  property {
    name  = "details2"
    type  = "string"
    value = var.twilioconnector_property_details2
  }

  property {
    name  = "iconUrl"
    type  = "string"
    value = var.twilioconnector_property_icon_url
  }

  property {
    name  = "iconUrlPng"
    type  = "string"
    value = var.twilioconnector_property_icon_url_png
  }

  property {
    name  = "registerMessageTemplate"
    type  = "string"
    value = var.twilioconnector_property_register_message_template
  }

  property {
    name  = "senderPhoneNumber"
    type  = "string"
    value = var.twilioconnector_property_sender_phone_number
  }

  property {
    name  = "showCredAddedOn"
    type  = "boolean"
    value = var.twilioconnector_property_show_cred_added_on
  }

  property {
    name  = "showCredAddedVia"
    type  = "boolean"
    value = var.twilioconnector_property_show_cred_added_via
  }

  property {
    name  = "title"
    type  = "string"
    value = var.twilioconnector_property_title
  }

  property {
    name  = "toolTip"
    type  = "string"
    value = var.twilioconnector_property_tool_tip
  }
}
