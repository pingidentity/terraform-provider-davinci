resource "davinci_connection" "salesforceConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "salesforceConnector"
  name         = "My awesome salesforceConnector"

  property {
    name  = "adminUsername"
    type  = "string"
    value = var.salesforceconnector_property_admin_username
  }

  property {
    name  = "consumerKey"
    type  = "string"
    value = var.salesforceconnector_property_consumer_key
  }

  property {
    name  = "domainName"
    type  = "string"
    value = var.salesforceconnector_property_domain_name
  }

  property {
    name  = "environment"
    type  = "string"
    value = var.salesforceconnector_property_environment
  }

  property {
    name  = "privateKey"
    type  = "string"
    value = var.salesforceconnector_property_private_key
  }
}
