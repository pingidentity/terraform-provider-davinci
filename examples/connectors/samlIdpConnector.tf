resource "davinci_connection" "samlIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "samlIdpConnector"
  name         = "My awesome samlIdpConnector"

  property {
    name  = "saml"
    type  = "json"
    value = jsonencode({})
  }
}
