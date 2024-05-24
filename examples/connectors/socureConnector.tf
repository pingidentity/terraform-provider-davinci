resource "davinci_connection" "socureConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "socureConnector"
  name         = "My awesome socureConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.socureconnector_property_api_key
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.socureconnector_property_base_url
  }

  property {
    name  = "customApiUrl"
    type  = "string"
    value = var.socureconnector_property_custom_api_url
  }

  property {
    name  = "sdkKey"
    type  = "string"
    value = var.socureconnector_property_sdk_key
  }

  property {
    name  = "skWebhookUri"
    type  = "string"
    value = var.socureconnector_property_sk_webhook_uri
  }
}
