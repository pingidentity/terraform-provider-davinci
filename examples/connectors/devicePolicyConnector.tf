resource "davinci_connection" "devicePolicyConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "devicePolicyConnector"
  name         = "My awesome devicePolicyConnector"
}
