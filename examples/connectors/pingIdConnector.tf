resource "davinci_connection" "pingIdConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingIdConnector"
  name         = "My awesome pingIdConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
