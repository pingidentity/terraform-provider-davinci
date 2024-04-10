resource "davinci_connection" "azureUserManagementConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "azureUserManagementConnector"
  name         = "My awesome azureUserManagementConnector"

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.azureusermanagementconnector_property_base_url
  }

  property {
    name  = "customApiUrl"
    type  = "string"
    value = var.azureusermanagementconnector_property_custom_api_url
  }

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
