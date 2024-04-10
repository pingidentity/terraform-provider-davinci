resource "davinci_connection" "connectorSvipe" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorSvipe"
  name         = "My awesome connectorSvipe"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
