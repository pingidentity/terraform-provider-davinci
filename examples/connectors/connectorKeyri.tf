resource "davinci_connection" "connectorKeyri" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorKeyri"
  name         = "My awesome connectorKeyri"
}
