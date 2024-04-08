resource "davinci_connection" "connectorIPGeolocationio" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorIPGeolocationio"
  name         = "My awesome connectorIPGeolocationio"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectoripgeolocationio_property_api_key
  }
}
