resource "davinci_connection" "flowConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "flowConnector"
  name         = "My awesome flowConnector"

  property {
    name  = "enforcedSignedToken"
    type  = "boolean"
    value = var.flowconnector_property_enforced_signed_token
  }

  property {
    name  = "inputSchema"
    type  = "string"
    value = var.flowconnector_property_input_schema
  }

  property {
    name  = "pemPublicKey"
    type  = "string"
    value = var.flowconnector_property_pem_public_key
  }
}
