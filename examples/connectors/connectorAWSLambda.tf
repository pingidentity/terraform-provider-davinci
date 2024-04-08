resource "davinci_connection" "connectorAWSLambda" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorAWSLambda"
  name         = "My awesome connectorAWSLambda"

  property {
    name  = "accessKeyId"
    type  = "string"
    value = var.connectorawslambda_property_access_key_id
  }

  property {
    name  = "region"
    type  = "string"
    value = "eu-west-1"
  }

  property {
    name  = "secretAccessKey"
    type  = "string"
    value = var.connectorawslambda_property_secret_access_key
  }
}
