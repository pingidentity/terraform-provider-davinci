resource "davinci_connection" "pingOneSnodeRiskConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneSnodeRiskConnector"
  name         = "My awesome pingOneSnodeRiskConnector"
}
