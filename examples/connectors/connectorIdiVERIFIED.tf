resource "davinci_connection" "connectorIdiVERIFIED" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorIdiVERIFIED"
  name         = "My awesome connectorIdiVERIFIED"

  property {
    name  = "apiSecret"
    type  = "string"
    value = var.connectoridiverified_property_api_secret
  }

  property {
    name  = "companyKey"
    type  = "string"
    value = var.connectoridiverified_property_company_key
  }

  property {
    name  = "idiEnv"
    type  = "string"
    value = var.connectoridiverified_property_idi_env
  }

  property {
    name  = "siteKey"
    type  = "string"
    value = var.connectoridiverified_property_site_key
  }

  property {
    name  = "uniqueUrl"
    type  = "string"
    value = var.connectoridiverified_property_unique_url
  }
}
