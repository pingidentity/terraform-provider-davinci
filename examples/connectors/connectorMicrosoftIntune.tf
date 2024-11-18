resource "davinci_connection" "connectorMicrosoftIntune" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorMicrosoftIntune"
  name         = "My awesome connectorMicrosoftIntune"

  property {
    name  = "clientId"
    type  = "string"
    value = var.connectormicrosoftintune_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.connectormicrosoftintune_property_client_secret
  }

  property {
    name  = "grantType"
    type  = "string"
    value = var.connectormicrosoftintune_property_grant_type
  }

  property {
    name  = "scope"
    type  = "string"
    value = var.connectormicrosoftintune_property_scope
  }

  property {
    name  = "tenant"
    type  = "string"
    value = var.connectormicrosoftintune_property_tenant
  }
}
