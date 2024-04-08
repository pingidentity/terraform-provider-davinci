resource "davinci_connection" "googleConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "googleConnector"
  name         = "My awesome googleConnector"

  property {
    name  = "openId"
    type  = "json"
    value = jsonencode({})
  }
}
