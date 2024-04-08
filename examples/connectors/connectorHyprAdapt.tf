resource "davinci_connection" "connectorHyprAdapt" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorHyprAdapt"
  name         = "My awesome connectorHyprAdapt"

  property {
    name  = "accessToken"
    type  = "string"
    value = var.connectorhypradapt_property_access_token
  }
}
