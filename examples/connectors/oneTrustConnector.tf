resource "davinci_connection" "oneTrustConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "oneTrustConnector"
  name         = "My awesome oneTrustConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.onetrustconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.onetrustconnector_property_client_secret
  }
}
