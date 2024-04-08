resource "davinci_connection" "nodeConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "nodeConnector"
  name         = "My awesome nodeConnector"
}
