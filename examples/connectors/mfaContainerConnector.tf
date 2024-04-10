resource "davinci_connection" "mfaContainerConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "mfaContainerConnector"
  name         = "My awesome mfaContainerConnector"
}
