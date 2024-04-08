resource "davinci_connection" "connectorWhatsAppBusiness" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorWhatsAppBusiness"
  name         = "My awesome connectorWhatsAppBusiness"

  property {
    name  = "accessToken"
    type  = "string"
    value = var.connectorwhatsappbusiness_property_access_token
  }

  property {
    name  = "appSecret"
    type  = "string"
    value = var.connectorwhatsappbusiness_property_app_secret
  }

  property {
    name  = "skWebhookUri"
    type  = "string"
    value = var.connectorwhatsappbusiness_property_sk_webhook_uri
  }

  property {
    name  = "verifyToken"
    type  = "string"
    value = var.connectorwhatsappbusiness_property_verify_token
  }

  property {
    name  = "version"
    type  = "string"
    value = var.connectorwhatsappbusiness_property_version
  }
}
