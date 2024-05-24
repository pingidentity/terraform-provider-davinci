resource "davinci_connection" "connectIdConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "connectIdConnector"
  name         = "My awesome connectIdConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
