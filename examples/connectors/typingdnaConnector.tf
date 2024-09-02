resource "davinci_connection" "typingdnaConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "typingdnaConnector"
  name         = "My awesome typingdnaConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = var.typingdnaconnector_property_custom_auth
  }
}
