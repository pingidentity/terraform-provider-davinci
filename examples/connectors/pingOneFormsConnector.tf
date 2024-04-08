resource "davinci_connection" "pingOneFormsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneFormsConnector"
  name         = "My awesome pingOneFormsConnector"
}
