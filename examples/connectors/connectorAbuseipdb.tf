resource "davinci_connection" "connectorAbuseipdb" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorAbuseipdb"
  name         = "My awesome connectorAbuseipdb"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectorabuseipdb_property_api_key
  }
}
