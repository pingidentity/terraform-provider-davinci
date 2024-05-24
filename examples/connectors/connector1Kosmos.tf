resource "davinci_connection" "connector1Kosmos" {
  environment_id = var.pingone_environment_id

  connector_id = "connector1Kosmos"
  name         = "My awesome connector1Kosmos"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
