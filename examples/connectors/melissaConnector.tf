resource "davinci_connection" "melissaConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "melissaConnector"
  name         = "My awesome melissaConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.melissaconnector_property_api_key
  }
}
