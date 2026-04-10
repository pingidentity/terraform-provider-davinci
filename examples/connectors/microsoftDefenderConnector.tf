resource "davinci_connection" "microsoftDefenderConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "microsoftDefenderConnector"
  name         = "My awesome microsoftDefenderConnector"

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.microsoftdefenderconnector_property_base_url
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.microsoftdefenderconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.microsoftdefenderconnector_property_client_secret
  }

  property {
    name  = "tenantId"
    type  = "string"
    value = var.microsoftdefenderconnector_property_tenant_id
  }
}
