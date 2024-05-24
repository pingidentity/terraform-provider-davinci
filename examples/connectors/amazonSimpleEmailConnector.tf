resource "davinci_connection" "amazonSimpleEmailConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "amazonSimpleEmailConnector"
  name         = "My awesome amazonSimpleEmailConnector"

  property {
    name  = "awsAccessKey"
    type  = "string"
    value = var.amazonsimpleemailconnector_property_aws_access_key
  }

  property {
    name  = "awsAccessSecret"
    type  = "string"
    value = var.amazonsimpleemailconnector_property_aws_access_secret
  }

  property {
    name  = "awsRegion"
    type  = "string"
    value = "eu-west-1"
  }

  property {
    name  = "from"
    type  = "string"
    value = "support@bxretail.org"
  }
}
