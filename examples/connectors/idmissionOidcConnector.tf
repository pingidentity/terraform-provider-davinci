resource "davinci_connection" "idmissionOidcConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "idmissionOidcConnector"
  name         = "My awesome idmissionOidcConnector"

  property {
    name  = "customAuth"
    type  = "json"
    value = var.idmissionoidcconnector_property_custom_auth
  }
}
