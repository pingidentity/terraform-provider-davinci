resource "davinci_connection" "googleWorkSpaceAdminConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "googleWorkSpaceAdminConnector"
  name         = "My awesome googleWorkSpaceAdminConnector"

  property {
    name  = "iss"
    type  = "string"
    value = var.googleworkspaceadminconnector_property_iss
  }

  property {
    name  = "privateKey"
    type  = "string"
    value = var.googleworkspaceadminconnector_property_private_key
  }

  property {
    name  = "sub"
    type  = "string"
    value = var.googleworkspaceadminconnector_property_sub
  }
}
