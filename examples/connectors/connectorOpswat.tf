resource "davinci_connection" "connectorOpswat" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorOpswat"
  name         = "My awesome connectorOpswat"

  property {
    name  = "clientID"
    type  = "string"
    value = var.connectoropswat_property_client_i_d
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.connectoropswat_property_client_secret
  }

  property {
    name  = "crossDomainApiPort"
    type  = "string"
    value = var.connectoropswat_property_cross_domain_api_port
  }

  property {
    name  = "maDomain"
    type  = "string"
    value = var.connectoropswat_property_ma_domain
  }
}
