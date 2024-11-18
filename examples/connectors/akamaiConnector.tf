resource "davinci_connection" "akamaiConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "akamaiConnector"
  name         = "My awesome akamaiConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = var.akamaiconnector_property_custom_auth
  }
}
