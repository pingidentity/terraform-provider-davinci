resource "davinci_connection" "bitbucketIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "bitbucketIdpConnector"
  name         = "My awesome bitbucketIdpConnector"

  property {
    name  = "oauth2"
    type  = "json"
    value = jsonencode({})
  }
}
