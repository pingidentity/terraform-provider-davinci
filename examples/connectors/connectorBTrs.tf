resource "davinci_connection" "connectorBTrs" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorBTrs"
  name         = "My awesome connectorBTrs"

  property {
    name  = "clientID"
    type  = "string"
    value = var.connectorbtrs_property_client_i_d
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.connectorbtrs_property_client_secret
  }

  property {
    name  = "rsAPIurl"
    type  = "string"
    value = var.rs_api_url
  }
}
