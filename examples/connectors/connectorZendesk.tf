resource "davinci_connection" "connectorZendesk" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorZendesk"
  name         = "My awesome connectorZendesk"

  property {
    name  = "apiToken"
    type  = "string"
    value = var.connectorzendesk_property_api_token
  }

  property {
    name  = "emailUsername"
    type  = "string"
    value = var.connectorzendesk_property_email_username
  }

  property {
    name  = "subdomain"
    type  = "string"
    value = var.connectorzendesk_property_subdomain
  }
}
