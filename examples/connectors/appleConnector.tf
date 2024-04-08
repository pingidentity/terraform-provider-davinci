resource "davinci_connection" "appleConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "appleConnector"
  name         = "My awesome appleConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
