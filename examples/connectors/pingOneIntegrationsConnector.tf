resource "davinci_connection" "pingOneIntegrationsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneIntegrationsConnector"
  name         = "My awesome pingOneIntegrationsConnector"
}
