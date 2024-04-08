resource "davinci_connection" "connectorWinmagic" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorWinmagic"
  name         = "My awesome connectorWinmagic"

  property {
    name  = "openId"
    type  = "json"
    value = jsonencode({})
  }
}
