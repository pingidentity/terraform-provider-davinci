resource "davinci_connection" "incodeConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "incodeConnector"
  name         = "My awesome incodeConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
