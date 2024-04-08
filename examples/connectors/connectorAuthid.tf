resource "davinci_connection" "connectorAuthid" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorAuthid"
  name         = "My awesome connectorAuthid"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
