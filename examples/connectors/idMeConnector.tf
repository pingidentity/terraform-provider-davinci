resource "davinci_connection" "idMeConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "idMeConnector"
  name         = "My awesome idMeConnector"

  property {
    name  = "oauth2"
    type  = "json"
    value = jsonencode({})
  }
}
