resource "davinci_connection" "idemiaConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "idemiaConnector"
  name         = "My awesome idemiaConnector"

  property {
    name  = "apikey"
    type  = "string"
    value = var.idemiaconnector_property_apikey
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.idemiaconnector_property_base_url
  }
}
