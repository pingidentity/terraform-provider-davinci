resource "davinci_connection" "microsoftIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "microsoftIdpConnector"
  name         = "My awesome microsoftIdpConnector"

  property {
    name  = "openId"
    type  = "json"
    value = jsonencode({})
  }
}
