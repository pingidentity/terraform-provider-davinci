resource "davinci_connection" "connectorCloudflare" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorCloudflare"
  name         = "My awesome connectorCloudflare"

  property {
    name  = "accountId"
    type  = "string"
    value = var.connectorcloudflare_property_account_id
  }

  property {
    name  = "apiToken"
    type  = "string"
    value = var.connectorcloudflare_property_api_token
  }
}
