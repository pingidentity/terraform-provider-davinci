resource "davinci_connection" "connectorIdentityNow" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorIdentityNow"
  name         = "My awesome connectorIdentityNow"

  property {
    name  = "clientId"
    type  = "string"
    value = var.connectoridentitynow_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.connectoridentitynow_property_client_secret
  }

  property {
    name  = "tenant"
    type  = "string"
    value = var.connectoridentitynow_property_tenant
  }
}
