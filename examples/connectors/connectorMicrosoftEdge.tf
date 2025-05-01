resource "davinci_connection" "connectorMicrosoftEdge" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorMicrosoftEdge"
  name         = "My awesome connectorMicrosoftEdge"

  property {
    name  = "customAuth"
    type  = "json"
    value = var.connectormicrosoftedge_property_custom_auth
  }
}
