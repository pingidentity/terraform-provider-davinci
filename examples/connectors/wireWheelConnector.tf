resource "davinci_connection" "wireWheelConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "wireWheelConnector"
  name         = "My awesome wireWheelConnector"

  property {
    name  = "baseURL"
    type  = "string"
    value = var.base_url
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.wirewheelconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.wirewheelconnector_property_client_secret
  }

  property {
    name  = "issuerId"
    type  = "string"
    value = var.wirewheelconnector_property_issuer_id
  }
}
