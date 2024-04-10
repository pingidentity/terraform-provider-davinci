resource "davinci_connection" "slackConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "slackConnector"
  name         = "My awesome slackConnector"

  property {
    name  = "oauth2"
    type  = "json"
    value = jsonencode({})
  }
}
