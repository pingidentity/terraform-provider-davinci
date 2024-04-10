resource "davinci_connection" "connectorSecuronix" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorSecuronix"
  name         = "My awesome connectorSecuronix"

  property {
    name  = "domainName"
    type  = "string"
    value = var.connectorsecuronix_property_domain_name
  }

  property {
    name  = "token"
    type  = "string"
    value = var.connectorsecuronix_property_token
  }
}
