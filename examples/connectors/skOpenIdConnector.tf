resource "davinci_connection" "skOpenIdConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "skOpenIdConnector"
  name         = "My awesome skOpenIdConnector"
}
