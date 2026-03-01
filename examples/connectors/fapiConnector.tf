resource "davinci_connection" "fapiConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "fapiConnector"
  name         = "My awesome fapiConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = var.fapiconnector_property_custom_auth
  }
}
