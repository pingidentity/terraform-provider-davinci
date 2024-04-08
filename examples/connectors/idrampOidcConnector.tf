resource "davinci_connection" "idrampOidcConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "idrampOidcConnector"
  name         = "My awesome idrampOidcConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
