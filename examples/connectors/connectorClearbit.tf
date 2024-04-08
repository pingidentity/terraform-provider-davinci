resource "davinci_connection" "connectorClearbit" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorClearbit"
  name         = "My awesome connectorClearbit"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectorclearbit_property_api_key
  }

  property {
    name  = "riskApiVersion"
    type  = "string"
    value = var.connectorclearbit_property_risk_api_version
  }

  property {
    name  = "version"
    type  = "string"
    value = var.connectorclearbit_property_version
  }
}
