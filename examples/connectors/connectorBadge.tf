resource "davinci_connection" "connectorBadge" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorBadge"
  name         = "My awesome connectorBadge"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
