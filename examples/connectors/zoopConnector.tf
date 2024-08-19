resource "davinci_connection" "zoopConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "zoopConnector"
  name         = "My awesome zoopConnector"

  property {
    name  = "agencyId"
    type  = "string"
    value = var.zoopconnector_property_agency_id
  }

  property {
    name  = "apiKey"
    type  = "string"
    value = var.zoopconnector_property_api_key
  }

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.zoopconnector_property_api_url
  }
}
