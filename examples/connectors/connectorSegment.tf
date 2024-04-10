resource "davinci_connection" "connectorSegment" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorSegment"
  name         = "My awesome connectorSegment"

  property {
    name  = "version"
    type  = "string"
    value = var.connectorsegment_property_version
  }

  property {
    name  = "writeKey"
    type  = "string"
    value = var.connectorsegment_property_write_key
  }
}
