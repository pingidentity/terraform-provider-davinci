resource "davinci_connection" "annotation" {
  name           = "myAnnotationConnector"
  connector_id   = "annotationConnector"
  environment_id = var.pingone_environment_id
}

resource "davinci_connection" "crowd_strike" {
  connector_id   = "crowdStrikeConnector"
  environment_id = var.pingone_environment_id
  name           = "CrowdStrike"
  property {
    name  = "clientId"
    value = "12345678"
  }
  property {
    name  = "clientSecret"
    value = "12345"
  }
}

output "cowd_strike_id" {
  value = resource.davinci_connection.crowd_strike.id
}
