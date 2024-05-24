resource "davinci_connection" "genericConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "genericConnector"
  name         = "My awesome genericConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
