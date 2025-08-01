resource "davinci_connection" "ideemConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "ideemConnector"
  name         = "My awesome ideemConnector"

  property {
    name  = "applicationEnvironment"
    type  = "string"
    value = var.ideemconnector_property_application_environment
  }

  property {
    name  = "applicationId"
    type  = "string"
    value = var.ideemconnector_property_application_id
  }

  property {
    name  = "hostURL"
    type  = "string"
    value = var.ideemconnector_property_host_u_r_l
  }

  property {
    name  = "userIdentifier"
    type  = "string"
    value = var.ideemconnector_property_user_identifier
  }

  property {
    name  = "validateTokenApiKey"
    type  = "string"
    value = var.ideemconnector_property_validate_token_api_key
  }

  property {
    name  = "zsmClientSdkApiKey"
    type  = "string"
    value = var.ideemconnector_property_zsm_client_sdk_api_key
  }
}
