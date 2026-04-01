resource "davinci_connection" "mitekConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "mitekConnector"
  name         = "My awesome mitekConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.mitekconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.mitekconnector_property_client_secret
  }

  property {
    name  = "hostURL"
    type  = "string"
    value = var.mitekconnector_property_host_u_r_l
  }

  property {
    name  = "requstAPIVersion"
    type  = "string"
    value = var.mitekconnector_property_requst_a_p_i_version
  }

  property {
    name  = "skWebhookUri"
    type  = "string"
    value = var.mitekconnector_property_sk_webhook_uri
  }
}
