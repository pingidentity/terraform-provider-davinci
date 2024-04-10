resource "davinci_connection" "skPeopleIntelligenceConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "skPeopleIntelligenceConnector"
  name         = "My awesome skPeopleIntelligenceConnector"

  property {
    name  = "authUrl"
    type  = "string"
    value = var.skpeopleintelligenceconnector_property_auth_url
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.skpeopleintelligenceconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.skpeopleintelligenceconnector_property_client_secret
  }

  property {
    name  = "dppa"
    type  = "string"
    value = var.skpeopleintelligenceconnector_property_dppa
  }

  property {
    name  = "glba"
    type  = "string"
    value = var.skpeopleintelligenceconnector_property_glba
  }

  property {
    name  = "searchUrl"
    type  = "string"
    value = var.skpeopleintelligenceconnector_property_search_url
  }
}
