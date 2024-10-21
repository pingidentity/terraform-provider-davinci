resource "davinci_connection" "mparticleConnector" {
  environment_id = var.pingone_environment_id

  connector_id = "mparticleConnector"
  name         = "My awesome mparticleConnector"

  property {
    name  = "clientID"
    type  = "string"
    value = var.mparticleconnector_property_client_i_d
  }

  property {
    name  = "clientSecret"
    type  = "string"
    value = var.mparticleconnector_property_client_secret
  }

  property {
    name  = "pod"
    type  = "string"
    value = var.mparticleconnector_property_pod
  }
}
