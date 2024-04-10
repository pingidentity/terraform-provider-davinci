resource "davinci_connection" "locationPolicyConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "locationPolicyConnector"
  name         = "My awesome locationPolicyConnector"
}
