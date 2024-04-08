resource "davinci_connection" "connectorBeyondIdentity" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorBeyondIdentity"
  name         = "My awesome connectorBeyondIdentity"

  property {
    name  = "openId"
    type  = "json"
    value = jsonencode({})
  }
}
