resource "davinci_connection" "pingOneSSOConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneSSOConnector"
  name         = "My awesome pingOneSSOConnector"

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
    name  = "region"
    type  = "string"
    value = var.pingonessoconnector_property_region
  }
}
