resource "davinci_connection" "connectorShopify" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorShopify"
  name         = "My awesome connectorShopify"

  property {
    name  = "accessToken"
    type  = "string"
    value = var.connectorshopify_property_access_token
  }

  property {
    name  = "apiVersion"
    type  = "string"
    value = var.connectorshopify_property_api_version
  }

  property {
    name  = "multipassSecret"
    type  = "string"
    value = var.connectorshopify_property_multipass_secret
  }

  property {
    name  = "multipassStoreDomain"
    type  = "string"
    value = var.connectorshopify_property_multipass_store_domain
  }

  property {
    name  = "yourStoreName"
    type  = "string"
    value = var.connectorshopify_property_your_store_name
  }
}
