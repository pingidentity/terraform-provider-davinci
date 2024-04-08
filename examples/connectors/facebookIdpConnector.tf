resource "davinci_connection" "facebookIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "facebookIdpConnector"
  name         = "My awesome facebookIdpConnector"

  property {
    name  = "oauth2"
    type  = "json"
    value = jsonencode({})
  }
}
