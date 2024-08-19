resource "davinci_connection" "microsoftDynamicsCustomerInsightsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "microsoftDynamicsCustomerInsightsConnector"
  name         = "My awesome microsoftDynamicsCustomerInsightsConnector"

  property {
    name  = "baseURL"
    type  = "string"
    value = var.microsoftdynamicscustomerinsightsconnector_property_base_u_r_l
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.microsoftdynamicscustomerinsightsconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.microsoftdynamicscustomerinsightsconnector_property_client_secret
  }

  property {
    name  = "environmentName"
    type  = "string"
    value = var.microsoftdynamicscustomerinsightsconnector_property_environment_name
  }

  property {
    name  = "grantType"
    type  = "string"
    value = var.microsoftdynamicscustomerinsightsconnector_property_grant_type
  }

  property {
    name  = "tenant"
    type  = "string"
    value = var.microsoftdynamicscustomerinsightsconnector_property_tenant
  }

  property {
    name  = "version"
    type  = "string"
    value = var.microsoftdynamicscustomerinsightsconnector_property_version
  }
}
