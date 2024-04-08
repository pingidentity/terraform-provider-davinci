resource "davinci_connection" "pingOneMfaConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneMfaConnector"
  name         = "My awesome pingOneMfaConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.pingone_worker_app_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.pingone_worker_app_client_secret
  }

  property {
    name  = "envId"
    type  = "string"
    value = var.pingone_worker_app_environment_id
  }

  property {
    name  = "policyId"
    type  = "string"
    value = var.pingonemfaconnector_property_policy_id
  }

  property {
    name  = "region"
    type  = "string"
    value = var.pingonemfaconnector_property_region
  }
}
