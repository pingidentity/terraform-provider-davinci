resource "davinci_connection" "connectorKeyless" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorKeyless"
  name         = "My awesome connectorKeyless"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
