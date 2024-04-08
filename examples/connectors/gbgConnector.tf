resource "davinci_connection" "gbgConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "gbgConnector"
  name         = "My awesome gbgConnector"

  property {
    name  = "password"
    type  = "string"
    value = var.gbgconnector_property_password
  }

  property {
    name  = "requestUrl"
    type  = "string"
    value = var.gbgconnector_property_request_url
  }

  property {
    name  = "soapAction"
    type  = "string"
    value = var.gbgconnector_property_soap_action
  }

  property {
    name  = "username"
    type  = "string"
    value = var.gbgconnector_property_username
  }
}
