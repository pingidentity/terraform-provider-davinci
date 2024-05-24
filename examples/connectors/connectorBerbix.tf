resource "davinci_connection" "connectorBerbix" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorBerbix"
  name         = "My awesome connectorBerbix"

  property {
    name  = "domainName"
    type  = "string"
    value = var.connectorberbix_property_domain_name
  }

  property {
    name  = "path"
    type  = "string"
    value = var.connectorberbix_property_path
  }

  property {
    name  = "username"
    type  = "string"
    value = var.connectorberbix_property_username
  }
}
