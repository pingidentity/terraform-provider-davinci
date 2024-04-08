resource "davinci_connection" "entrustConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "entrustConnector"
  name         = "My awesome entrustConnector"

  property {
    name  = "applicationId"
    type  = "string"
    value = var.entrustconnector_property_application_id
  }

  property {
    name  = "serviceDomain"
    type  = "string"
    value = var.entrustconnector_property_service_domain
  }
}
