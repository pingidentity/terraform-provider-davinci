resource "davinci_connection" "bambooConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "bambooConnector"
  name         = "My awesome bambooConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.bambooconnector_property_api_key
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.bambooconnector_property_base_url
  }

  property {
    name  = "companySubDomain"
    type  = "string"
    value = var.bambooconnector_property_company_sub_domain
  }

  property {
    name  = "flowId"
    type  = "string"
    value = var.bambooconnector_property_flow_id
  }

  property {
    name  = "skWebhookUri"
    type  = "string"
    value = var.bambooconnector_property_sk_webhook_uri
  }

  property {
    name  = "webhookToken"
    type  = "string"
    value = var.bambooconnector_property_webhook_token
  }
}
