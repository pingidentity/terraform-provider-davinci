resource "davinci_connection" "scopesConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "scopesConnector"
  name         = "My awesome scopesConnector"

  property {
    name  = "scopes"
    type  = "json"
    value = var.scopesconnector_property_scopes
  }
}
