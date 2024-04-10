resource "davinci_connection" "tmtConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "tmtConnector"
  name         = "My awesome tmtConnector"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.tmtconnector_property_api_key
  }

  property {
    name  = "apiSecret"
    type  = "string"
    value = var.tmtconnector_property_api_secret
  }

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.tmtconnector_property_api_url
  }
}
