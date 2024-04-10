resource "davinci_connection" "pingauthadapter" {
  environment_id = var.pingone_environment_id

  connector_id = "pingauthadapter"
  name         = "My awesome pingauthadapter"
}
