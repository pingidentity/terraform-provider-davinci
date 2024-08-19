resource "davinci_connection" "privateidConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "privateidConnector"
  name         = "My awesome privateidConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = var.privateidconnector_property_custom_auth
  }
}
