resource "davinci_connection" "adobeExperienceManagerConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "adobeExperienceManagerConnector"
  name         = "My awesome adobeExperienceManagerConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.adobeexperiencemanagerconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.adobeexperiencemanagerconnector_property_client_secret
  }

  property {
    name  = "orgId"
    type  = "string"
    value = var.adobeexperiencemanagerconnector_property_org_id
  }
}
