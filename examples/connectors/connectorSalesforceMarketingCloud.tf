resource "davinci_connection" "connectorSalesforceMarketingCloud" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorSalesforceMarketingCloud"
  name         = "My awesome connectorSalesforceMarketingCloud"

  property {
    name  = "SalesforceMarketingCloudURL"
    type  = "string"
    value = var.salesforce_marketing_cloud_url
  }

  property {
    name  = "accountId"
    type  = "string"
    value = var.connectorsalesforcemarketingcloud_property_account_id
  }

  property {
    name  = "clientId"
    type  = "string"
    value = var.connectorsalesforcemarketingcloud_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.connectorsalesforcemarketingcloud_property_client_secret
  }

  property {
    name  = "scope"
    type  = "string"
    value = var.connectorsalesforcemarketingcloud_property_scope
  }
}
