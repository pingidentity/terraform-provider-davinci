resource "davinci_connection" "jiraConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "jiraConnector"
  name         = "My awesome jiraConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.jiraconnector_property_api_key
  }

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.jiraconnector_property_api_url
  }

  property {
    name  = "email"
    type  = "string"
    value = var.jiraconnector_property_email
  }
}
