resource "davinci_connection" "idranddConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "idranddConnector"
  name         = "My awesome idranddConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.idranddconnector_property_api_key
  }

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.idranddconnector_property_api_url
  }
}
