resource "davinci_connection" "connectorHuman" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorHuman"
  name         = "My awesome connectorHuman"

  property {
    name  = "humanAuthenticationToken"
    type  = "string"
    value = var.connectorhuman_property_human_authentication_token
  }

  property {
    name  = "humanCustomerID"
    type  = "string"
    value = var.human_customer_id
  }

  property {
    name  = "humanPolicyName"
    type  = "string"
    value = var.connectorhuman_property_human_policy_name
  }
}
