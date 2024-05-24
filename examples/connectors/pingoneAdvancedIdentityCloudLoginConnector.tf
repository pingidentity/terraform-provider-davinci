resource "davinci_connection" "pingoneAdvancedIdentityCloudLoginConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingoneAdvancedIdentityCloudLoginConnector"
  name         = "My awesome pingoneAdvancedIdentityCloudLoginConnector"

  property {
    name  = "openId"
    type  = "json"
    value = var.pingoneadvancedidentitycloudloginconnector_property_open_id
  }
}
