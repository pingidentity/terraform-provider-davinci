resource "davinci_connection" "pingOneRecognizeConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneRecognizeConnector"
  name         = "My awesome pingOneRecognizeConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.pingonerecognizeconnector_property_api_key
  }

  property {
    name  = "customerName"
    type  = "string"
    value = var.pingonerecognizeconnector_property_customer_name
  }

  property {
    name  = "region"
    type  = "string"
    value = var.pingonerecognizeconnector_property_region
  }
}
