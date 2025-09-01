resource "davinci_connection" "akamaiApConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "akamaiApConnector"
  name         = "My awesome akamaiApConnector"
}
