resource "davinci_connection" "accessRequestConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "accessRequestConnector"
  name         = "My awesome accessRequestConnector"

  property {
    name  = "baseURL"
    type  = "string"
    value = var.accessrequestconnector_property_base_u_r_l
  }

  property {
    name  = "endUserClientId"
    type  = "string"
    value = var.accessrequestconnector_property_end_user_client_id
  }

  property {
    name  = "endUserClientPrivateKey"
    type  = "string"
    value = var.accessrequestconnector_property_end_user_client_private_key
  }

  property {
    name  = "realm"
    type  = "string"
    value = var.accessrequestconnector_property_realm
  }

  property {
    name  = "serviceAccountId"
    type  = "string"
    value = var.accessrequestconnector_property_service_account_id
  }

  property {
    name  = "serviceAccountPrivateKey"
    type  = "string"
    value = var.accessrequestconnector_property_service_account_private_key
  }
}
