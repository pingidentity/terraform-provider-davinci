resource "davinci_connection" "dataZooConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "dataZooConnector"
  name         = "My awesome dataZooConnector"

  property {
    name  = "password"
    type  = "string"
    value = var.datazooconnector_property_password
  }

  property {
    name  = "username"
    type  = "string"
    value = var.datazooconnector_property_username
  }
}
