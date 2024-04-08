resource "davinci_connection" "tutloxpConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "tutloxpConnector"
  name         = "My awesome tutloxpConnector"

  property {
    name  = "apiUrl"
    type  = "string"
    value = var.tutloxpconnector_property_api_url
  }

  property {
    name  = "dppaCode"
    type  = "string"
    value = var.tutloxpconnector_property_dppa_code
  }

  property {
    name  = "glbCode"
    type  = "string"
    value = var.tutloxpconnector_property_glb_code
  }

  property {
    name  = "password"
    type  = "string"
    value = var.tutloxpconnector_property_password
  }

  property {
    name  = "username"
    type  = "string"
    value = var.tutloxpconnector_property_username
  }
}
