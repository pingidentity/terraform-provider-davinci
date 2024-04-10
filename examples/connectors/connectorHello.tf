resource "davinci_connection" "connectorHello" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorHello"
  name         = "My awesome connectorHello"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
