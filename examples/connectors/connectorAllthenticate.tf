resource "davinci_connection" "connectorAllthenticate" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorAllthenticate"
  name         = "My awesome connectorAllthenticate"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
