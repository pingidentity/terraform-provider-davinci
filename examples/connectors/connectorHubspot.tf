resource "davinci_connection" "connectorHubspot" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorHubspot"
  name         = "My awesome connectorHubspot"

  property {
    name  = "bearerToken"
    type  = "string"
    value = var.connectorhubspot_property_bearer_token
  }
}
