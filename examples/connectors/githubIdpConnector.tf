resource "davinci_connection" "githubIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "githubIdpConnector"
  name         = "My awesome githubIdpConnector"

  property {
    name  = "oauth2"
    type  = "json"
    value = jsonencode({})
  }
}
