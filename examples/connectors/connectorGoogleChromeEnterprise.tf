resource "davinci_connection" "connectorGoogleChromeEnterprise" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorGoogleChromeEnterprise"
  name         = "My awesome connectorGoogleChromeEnterprise"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
