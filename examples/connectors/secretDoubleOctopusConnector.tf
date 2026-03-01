resource "davinci_connection" "secretDoubleOctopusConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "secretDoubleOctopusConnector"
  name         = "My awesome secretDoubleOctopusConnector"

  property {
    name  = "apiToken"
    type  = "string"
    value = var.secretdoubleoctopusconnector_property_api_token
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.secretdoubleoctopusconnector_property_base_url
  }

  property {
    name  = "serviceId"
    type  = "string"
    value = var.secretdoubleoctopusconnector_property_service_id
  }

  property {
    name  = "x509Certificate"
    type  = "string"
    value = var.secretdoubleoctopusconnector_property_x509_certificate
  }
}
