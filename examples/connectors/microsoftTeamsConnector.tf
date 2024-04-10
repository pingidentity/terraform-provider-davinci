resource "davinci_connection" "microsoftTeamsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "microsoftTeamsConnector"
  name         = "My awesome microsoftTeamsConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
