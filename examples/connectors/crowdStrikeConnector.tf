resource "davinci_connection" "crowdStrikeConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "crowdStrikeConnector"
  name         = "My awesome crowdStrikeConnector"

  property {
    name  = "baseURL"
    type  = "string"
    value = var.base_url
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.crowdstrikeconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.crowdstrikeconnector_property_client_secret
  }
}
