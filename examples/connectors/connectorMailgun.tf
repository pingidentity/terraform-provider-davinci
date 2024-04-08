resource "davinci_connection" "connectorMailgun" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorMailgun"
  name         = "My awesome connectorMailgun"

  property {
    name  = "apiKey"
    type  = "string"
    value = var.connectormailgun_property_api_key
  }

  property {
    name  = "apiVersion"
    type  = "string"
    value = var.connectormailgun_property_api_version
  }

  property {
    name  = "mailgunDomain"
    type  = "string"
    value = var.connectormailgun_property_mailgun_domain
  }
}
