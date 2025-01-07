resource "davinci_connection" "connectorClear" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorClear"
  name         = "My awesome connectorClear"

  property {
    name  = "customAuth"
    type  = "json"
    value = var.connectorclear_property_custom_auth
  }
}
