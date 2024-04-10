resource "davinci_connection" "connectorZscaler" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorZscaler"
  name         = "My awesome connectorZscaler"

  property {
    name  = "basePath"
    type  = "string"
    value = var.connectorzscaler_property_base_path
  }

  property {
    name  = "baseURL"
    type  = "string"
    value = var.base_url
  }

  property {
    name  = "zscalerAPIkey"
    type  = "string"
    value = var.zscaler_api_key
  }

  property {
    name  = "zscalerPassword"
    type  = "string"
    value = var.connectorzscaler_property_zscaler_password
  }

  property {
    name  = "zscalerUsername"
    type  = "string"
    value = var.connectorzscaler_property_zscaler_username
  }
}
