resource "davinci_connection" "idDatawebConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "idDatawebConnector"
  name         = "My awesome idDatawebConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
