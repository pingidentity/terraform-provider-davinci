resource "davinci_connection" "connectorAmazonDynamoDB" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorAmazonDynamoDB"
  name         = "My awesome connectorAmazonDynamoDB"

  property {
    name  = "awsAccessKey"
    type  = "string"
    value = var.connectoramazondynamodb_property_aws_access_key
  }

  property {
    name  = "awsAccessSecret"
    type  = "string"
    value = var.connectoramazondynamodb_property_aws_access_secret
  }

  property {
    name  = "awsRegion"
    type  = "string"
    value = "eu-west-1"
  }
}
