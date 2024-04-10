resource "davinci_connection" "yotiConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "yotiConnector"
  name         = "My awesome yotiConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
