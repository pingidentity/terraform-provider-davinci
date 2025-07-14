resource "davinci_connection" "kfKerberosConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "kfKerberosConnector"
  name         = "My awesome kfKerberosConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.kfkerberosconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.kfkerberosconnector_property_client_secret
  }

  property {
    name  = "envId"
    type  = "string"
    value = var.kfkerberosconnector_property_env_id
  }

  property {
    name  = "region"
    type  = "string"
    value = var.kfkerberosconnector_property_region
  }
}
