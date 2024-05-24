resource "davinci_connection" "pingOneAuthenticationConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneAuthenticationConnector"
  name         = "My awesome pingOneAuthenticationConnector"
}
