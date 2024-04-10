resource "davinci_connection" "biocatchConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "biocatchConnector"
  name         = "My awesome biocatchConnector"

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.biocatchconnector_property_api_url
  }

  property {
    name  = "customerId"
    type  = "string"
    value = var.biocatchconnector_property_customer_id
  }

  property {
    name  = "javascriptCdnUrl"
    type  = "string"
    value = var.biocatchconnector_property_javascript_cdn_url
  }

  property {
    name  = "sdkToken"
    type  = "string"
    value = var.biocatchconnector_property_sdk_token
  }

  property {
    name  = "truthApiKey"
    type  = "string"
    value = var.biocatchconnector_property_truth_api_key
  }

  property {
    name  = "truthApiUrl"
    type  = "string"
    value = var.biocatchconnector_property_truth_api_url
  }
}
