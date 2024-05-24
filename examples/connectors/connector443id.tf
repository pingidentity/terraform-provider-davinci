resource "davinci_connection" "connector443id" {
  environment_id = var.pingone_environment_id

  connector_id = "connector443id"
  name         = "My awesome connector443id"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connector443id_property_api_key
  }
}
