resource "davinci_connection" "digidentityConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "digidentityConnector"
  name         = "My awesome digidentityConnector"

  property {
    name  = "oauth2"
    type  = "json"
    value = jsonencode({})
  }
}
