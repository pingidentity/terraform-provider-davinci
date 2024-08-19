resource "davinci_connection" "idmecommunityConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "idmecommunityConnector"
  name         = "My awesome idmecommunityConnector"

  property {
    name  = "openId"
    type  = "json"
    value = var.idmecommunityconnector_property_open_id
  }
}
