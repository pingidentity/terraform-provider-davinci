resource "davinci_connection" "iproovV2Connector" {
  environment_id = var.pingone_environment_id

  connector_id = "iproovV2Connector"
  name         = "My awesome iproovV2Connector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.iproovv2connector_property_api_key
  }

  property {
    name  = "secret"
    type  = "string"
    value = var.iproovv2connector_property_secret
  }

  property {
    name  = "tenant"
    type  = "string"
    value = var.iproovv2connector_property_tenant
  }
}
