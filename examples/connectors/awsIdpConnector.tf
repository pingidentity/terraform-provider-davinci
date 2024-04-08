resource "davinci_connection" "awsIdpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "awsIdpConnector"
  name         = "My awesome awsIdpConnector"

  property {
    name  = "openId"
    type  = "json"
    value = jsonencode({})
  }
}
