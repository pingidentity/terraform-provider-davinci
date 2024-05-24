resource "davinci_connection" "connectorAcuant" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorAcuant"
  name         = "My awesome connectorAcuant"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
