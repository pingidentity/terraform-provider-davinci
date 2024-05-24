resource "davinci_connection" "hyprConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "hyprConnector"
  name         = "My awesome hyprConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
