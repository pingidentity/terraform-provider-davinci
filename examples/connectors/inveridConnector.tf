resource "davinci_connection" "inveridConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "inveridConnector"
  name         = "My awesome inveridConnector"

  property {
    name  = "getApiKey"
    type  = "string"
    value = var.inveridconnector_property_get_api_key
  }

  property {
    name  = "host"
    type  = "string"
    value = var.inveridconnector_property_host
  }

  property {
    name  = "postApiKey"
    type  = "string"
    value = var.inveridconnector_property_post_api_key
  }

  property {
    name  = "skWebhookUri"
    type  = "string"
    value = var.inveridconnector_property_sk_webhook_uri
  }

  property {
    name  = "timeToLive"
    type  = "string"
    value = var.inveridconnector_property_time_to_live
  }
}
