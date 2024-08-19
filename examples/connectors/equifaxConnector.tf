resource "davinci_connection" "equifaxConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "equifaxConnector"
  name         = "My awesome equifaxConnector"

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.equifaxconnector_property_base_url
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.equifaxconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.equifaxconnector_property_client_secret
  }

  property {
    name  = "equifaxSoapApiEnvironment"
    type  = "string"
    value = var.equifaxconnector_property_equifax_soap_api_environment
  }

  property {
    name  = "memberNumber"
    type  = "string"
    value = var.equifaxconnector_property_member_number
  }

  property {
    name  = "password"
    type  = "string"
    value = var.equifaxconnector_property_password
  }

  property {
    name  = "username"
    type  = "string"
    value = var.equifaxconnector_property_username
  }
}
