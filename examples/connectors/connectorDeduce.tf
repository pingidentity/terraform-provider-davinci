resource "davinci_connection" "connectorDeduce" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorDeduce"
  name         = "My awesome connectorDeduce"

  property {
    name  = "apikey"
    type  = "string"
    value = var.connectordeduce_property_apikey
  }

  property {
    name  = "siteId"
    type  = "string"
    value = var.connectordeduce_property_site_id
  }
}
