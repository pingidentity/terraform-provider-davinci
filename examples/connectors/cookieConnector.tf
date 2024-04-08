resource "davinci_connection" "cookieConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "cookieConnector"
  name         = "My awesome cookieConnector"

  property {
    name  = "hmacSigningKey"
    type  = "string"
    value = var.cookieconnector_property_hmac_signing_key
  }
}
