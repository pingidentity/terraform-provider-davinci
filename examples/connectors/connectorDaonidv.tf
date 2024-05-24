resource "davinci_connection" "connectorDaonidv" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorDaonidv"
  name         = "My awesome connectorDaonidv"

  property {
    name  = "openId"
    type  = "json"
    value = jsonencode({})
  }
}
