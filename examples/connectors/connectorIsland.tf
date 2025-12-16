resource "davinci_connection" "connectorIsland" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorIsland"
  name         = "My awesome connectorIsland"

  property {
    name  = "customAuth"
    type  = "json"
    value = var.connectorisland_property_custom_auth
  }
}
