resource "davinci_connection" "connectorInfinipoint" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorInfinipoint"
  name         = "My awesome connectorInfinipoint"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
