resource "davinci_connection" "analyticsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "analyticsConnector"
  name         = "My awesome analyticsConnector"
}
