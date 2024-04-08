resource "davinci_connection" "pingOneCredentialsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneCredentialsConnector"
  name         = "My awesome pingOneCredentialsConnector"

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
    name  = "digitalWalletApplicationId"
    type  = "string"
    value = var.pingonecredentialsconnector_property_digital_wallet_application_id
  }

  property {
    name  = "envId"
    type  = "string"
    value = var.pingone_worker_app_environment_id
  }

  property {
    name  = "region"
    type  = "string"
    value = var.pingonecredentialsconnector_property_region
  }
}
