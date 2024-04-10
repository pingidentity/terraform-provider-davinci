resource "davinci_connection" "adobemarketoConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "adobemarketoConnector"
  name         = "My awesome adobemarketoConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.adobemarketoconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.adobemarketoconnector_property_client_secret
  }

  property {
    name  = "endpoint"
    type  = "string"
    value = var.adobemarketoconnector_property_endpoint
  }
}
