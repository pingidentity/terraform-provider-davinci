resource "davinci_connection" "webhookConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "webhookConnector"
  name         = "My awesome webhookConnector"

  property {
    name  = "urls"
    type  = "string"
    value = var.webhookconnector_property_urls
  }
}
