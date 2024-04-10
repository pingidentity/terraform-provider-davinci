resource "davinci_connection" "connectorJiraServiceDesk" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorJiraServiceDesk"
  name         = "My awesome connectorJiraServiceDesk"

  property {
    name  = "JIRAServiceDeskAuth"
    type  = "string"
    value = var.jira_service_desk_auth
  }

  property {
    name  = "JIRAServiceDeskCreateData"
    type  = "string"
    value = var.jira_service_desk_create_data
  }

  property {
    name  = "JIRAServiceDeskURL"
    type  = "string"
    value = var.jira_service_desk_url
  }

  property {
    name  = "JIRAServiceDeskUpdateData"
    type  = "string"
    value = var.jira_service_desk_update_data
  }

  property {
    name  = "method"
    type  = "string"
    value = var.connectorjiraservicedesk_property_method
  }
}
