resource "davinci_connection" "connectorMailchimp" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorMailchimp"
  name         = "My awesome connectorMailchimp"

  property {
    name  = "transactionalApiKey"
    type  = "string"
    value = var.connectormailchimp_property_transactional_api_key
  }

  property {
    name  = "transactionalApiVersion"
    type  = "string"
    value = var.connectormailchimp_property_transactional_api_version
  }
}
