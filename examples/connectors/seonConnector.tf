resource "davinci_connection" "seonConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "seonConnector"
  name         = "My awesome seonConnector"

  property {
    name  = "baseURL"
    type  = "string"
    value = var.base_url
  }

  property {
    name  = "licenseKey"
    type  = "string"
    value = var.seonconnector_property_license_key
  }
}
