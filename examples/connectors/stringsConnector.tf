resource "davinci_connection" "stringsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "stringsConnector"
  name         = "My awesome stringsConnector"
}
