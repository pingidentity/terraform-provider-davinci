resource "davinci_connection" "finicityConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "finicityConnector"
  name         = "My awesome finicityConnector"

  property {
    name  = "appKey"
    type  = "string"
    value = var.finicityconnector_property_app_key
  }

  property {
    name  = "baseUrl"
    type  = "string"
    value = var.finicityconnector_property_base_url
  }

  property {
    name  = "partnerId"
    type  = "string"
    value = var.finicityconnector_property_partner_id
  }

  property {
    name  = "partnerSecret"
    type  = "string"
    value = var.finicityconnector_property_partner_secret
  }
}
