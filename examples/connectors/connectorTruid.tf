resource "davinci_connection" "connectorTruid" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorTruid"
  name         = "My awesome connectorTruid"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
