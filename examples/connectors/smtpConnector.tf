resource "davinci_connection" "smtpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "smtpConnector"
  name         = "My awesome smtpConnector"

  property {
    name  = "hostname"
    type  = "string"
    value = var.smtpconnector_property_hostname
  }

  property {
    name  = "name"
    type  = "string"
    value = var.smtpconnector_property_name
  }

  property {
    name  = "password"
    type  = "string"
    value = var.smtpconnector_property_password
  }

  property {
    name  = "port"
    type  = "number"
    value = var.smtpconnector_property_port
  }

  property {
    name  = "secureFlag"
    type  = "boolean"
    value = var.smtpconnector_property_secure_flag
  }

  property {
    name  = "username"
    type  = "string"
    value = var.smtpconnector_property_username
  }
}
