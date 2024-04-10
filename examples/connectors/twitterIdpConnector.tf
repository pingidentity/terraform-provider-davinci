resource "davinci_connection" "twitterIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "twitterIdpConnector"
  name         = "My awesome twitterIdpConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
