resource "davinci_connection" "connectorSaviyntFlow" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorSaviyntFlow"
  name         = "My awesome connectorSaviyntFlow"

  property {
    name  = "domainName"
    type  = "string"
    value = var.connectorsaviyntflow_property_domain_name
  }

  property {
    name  = "path"
    type  = "string"
    value = var.connectorsaviyntflow_property_path
  }

  property {
    name  = "saviyntPassword"
    type  = "string"
    value = var.connectorsaviyntflow_property_saviynt_password
  }

  property {
    name  = "saviyntUserName"
    type  = "string"
    value = var.connectorsaviyntflow_property_saviynt_user_name
  }
}
