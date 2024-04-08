resource "davinci_connection" "connectorSignicat" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorSignicat"
  name         = "My awesome connectorSignicat"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
