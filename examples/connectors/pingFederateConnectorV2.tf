resource "davinci_connection" "pingFederateConnectorV2" {
  environment_id = var.pingone_environment_id

  connector_id = "pingFederateConnectorV2"
  name         = "My awesome pingFederateConnectorV2"

  property {
    name  = "openId"
    type  = "json"
    value = jsonencode({})
  }
}
