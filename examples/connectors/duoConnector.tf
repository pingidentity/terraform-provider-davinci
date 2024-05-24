resource "davinci_connection" "duoConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "duoConnector"
  name         = "My awesome duoConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
