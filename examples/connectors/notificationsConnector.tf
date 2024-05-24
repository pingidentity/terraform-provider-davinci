resource "davinci_connection" "notificationsConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "notificationsConnector"
  name         = "My awesome notificationsConnector"

  property {
    name  = "clientId"
    type  = "string"
    value = var.notificationsconnector_property_client_id
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.notificationsconnector_property_client_secret
  }

  property {
    name  = "envId"
    type  = "string"
    value = var.notificationsconnector_property_env_id
  }

  property {
    name  = "notificationPolicyId"
    type  = "string"
    value = var.notificationsconnector_property_notification_policy_id
  }

  property {
    name  = "region"
    type  = "string"
    value = var.notificationsconnector_property_region
  }
}
