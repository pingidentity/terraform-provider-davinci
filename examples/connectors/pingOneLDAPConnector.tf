resource "davinci_connection" "pingOneLDAPConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "pingOneLDAPConnector"
  name         = "My awesome pingOneLDAPConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.pingoneldapconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.pingoneldapconnector_property_client_secret
  }

  property {
    name  = "envId"
    type  = "string"
    value = var.pingoneldapconnector_property_env_id
  }

  property {
    name  = "gatewayId"
    type  = "string"
    value = var.pingoneldapconnector_property_gateway_id
  }

  property {
    name  = "region"
    type  = "string"
    value = var.pingoneldapconnector_property_region
  }
}
