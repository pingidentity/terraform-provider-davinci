resource "davinci_connection" "connectorRandomUserMe" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorRandomUserMe"
  name         = "My awesome connectorRandomUserMe"
}
