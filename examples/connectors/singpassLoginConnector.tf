resource "davinci_connection" "singpassLoginConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "singpassLoginConnector"
  name         = "My awesome singpassLoginConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
