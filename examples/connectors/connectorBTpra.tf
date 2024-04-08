resource "davinci_connection" "connectorBTpra" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorBTpra"
  name         = "My awesome connectorBTpra"

  property {
    name  = "clientID"
    type  = "string"
    value = var.connectorbtpra_property_client_i_d
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.connectorbtpra_property_client_secret
  }

  property {
    name  = "praAPIurl"
    type  = "string"
    value = var.pra_api_url
  }
}
