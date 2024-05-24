resource "davinci_connection" "connectorIdMeIdentity" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorIdMeIdentity"
  name         = "My awesome connectorIdMeIdentity"

  property {
    name  = "openId"
    type  = "json"
    value = jsonencode({})
  }
}
