resource "davinci_connection" "connectorAsignio" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorAsignio"
  name         = "My awesome connectorAsignio"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
