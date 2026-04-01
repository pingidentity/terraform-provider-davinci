resource "davinci_connection" "connectorAmazonAwsSecretsManager" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorAmazonAwsSecretsManager"
  name         = "My awesome connectorAmazonAwsSecretsManager"

  property {
    name  = "accessKeyId"
    type  = "string"
    value = var.connectoramazonawssecretsmanager_property_access_key_id
  }

  property {
    name  = "region"
    type  = "string"
    value = "eu-west-1"
  }

  property {
    name  = "secondaryAccessKeyId"
    type  = "string"
    value = var.connectoramazonawssecretsmanager_property_secondary_access_key_id
  }

  property {
    name  = "secondaryRegion"
    type  = "string"
    value = var.connectoramazonawssecretsmanager_property_secondary_region
  }

  property {
    name  = "secondarySecretAccessKey"
    type  = "string"
    value = var.connectoramazonawssecretsmanager_property_secondary_secret_access_key
  }

  property {
    name  = "secretAccessKey"
    type  = "string"
    value = var.connectoramazonawssecretsmanager_property_secret_access_key
  }
}
