resource "davinci_connection" "servicenowConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "servicenowConnector"
  name         = "My awesome servicenowConnector"

  property {
    name  = "adminUsername"
    type  = "string"
    value = var.servicenowconnector_property_admin_username
  }

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.servicenowconnector_property_api_url
  }

  property {
    name  = "password"
    type  = "string"
    value = var.servicenowconnector_property_password
  }
}
