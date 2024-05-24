resource "davinci_connection" "connectorGoogleanalyticsUA" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorGoogleanalyticsUA"
  name         = "My awesome connectorGoogleanalyticsUA"

  property {
    name  = "trackingID"
    type  = "string"
    value = var.tracking_id
  }

  property {
    name  = "version"
    type  = "string"
    value = var.connectorgoogleanalyticsua_property_version
  }
}
