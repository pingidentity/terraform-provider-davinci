resource "davinci_connection" "connectorTrulioo" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorTrulioo"
  name         = "My awesome connectorTrulioo"

  property {
    name  = "clientID"
    type  = "string"
    value = var.connectortrulioo_property_client_i_d
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.connectortrulioo_property_client_secret
  }
}
