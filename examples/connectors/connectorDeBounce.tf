resource "davinci_connection" "connectorDeBounce" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorDeBounce"
  name         = "My awesome connectorDeBounce"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectordebounce_property_api_key
  }
}
