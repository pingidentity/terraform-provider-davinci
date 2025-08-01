resource "davinci_connection" "scrambleIdConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "scrambleIdConnector"
  name         = "My awesome scrambleIdConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = var.scrambleidconnector_property_custom_auth
  }
}
