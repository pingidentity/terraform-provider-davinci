resource "davinci_connection" "pingOneVerifyConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneVerifyConnector"
  name         = "My awesome pingOneVerifyConnector"

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
    value = var.pingoneverifyconnector_property_region
  }
}
