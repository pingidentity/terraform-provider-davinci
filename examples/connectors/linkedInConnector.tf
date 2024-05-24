resource "davinci_connection" "linkedInConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "linkedInConnector"
  name         = "My awesome linkedInConnector"

  property {
    name  = "oauth2"
    type  = "json"
    value = jsonencode({})
  }
}
