resource "davinci_connection" "connectorIdmelon" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorIdmelon"
  name         = "My awesome connectorIdmelon"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
