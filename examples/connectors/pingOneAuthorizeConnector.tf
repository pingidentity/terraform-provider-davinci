resource "davinci_connection" "pingOneAuthorizeConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneAuthorizeConnector"
  name         = "My awesome pingOneAuthorizeConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.pingoneauthorizeconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.pingoneauthorizeconnector_property_client_secret
  }

  property {
    name  = "endpointURL"
    type  = "string"
    value = var.endpoint_url
  }
}
