resource "davinci_connection" "iovationConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "iovationConnector"
  name         = "My awesome iovationConnector"

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.iovationconnector_property_api_url
  }

  property {
    name  = "javascriptCdnUrl"
    type  = "string"
    value = var.iovationconnector_property_javascript_cdn_url
  }

  property {
    name  = "subKey"
    type  = "string"
    value = var.iovationconnector_property_sub_key
  }

  property {
    name  = "subscriberAccount"
    type  = "string"
    value = var.iovationconnector_property_subscriber_account
  }

  property {
    name  = "subscriberId"
    type  = "string"
    value = var.iovationconnector_property_subscriber_id
  }

  property {
    name  = "subscriberPasscode"
    type  = "string"
    value = var.iovationconnector_property_subscriber_passcode
  }

  property {
    name  = "version"
    type  = "string"
    value = var.iovationconnector_property_version
  }
}
