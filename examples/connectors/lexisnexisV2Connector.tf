resource "davinci_connection" "lexisnexisV2Connector" {
  environment_id = var.pingone_environment_id

  connector_id = "lexisnexisV2Connector"
  name         = "My awesome lexisnexisV2Connector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.lexisnexisv2connector_property_api_key
  }

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.lexisnexisv2connector_property_api_url
  }

  property {
    name  = "orgId"
    type  = "string"
    value = var.lexisnexisv2connector_property_org_id
  }

  property {
    name  = "useCustomApiURL"
    type  = "string"
    value = var.use_custom_api_url
  }
}
