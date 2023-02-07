resource "davinci_connection" "annotation" {
  name           = "myAnnotationConnector"
  connector_id   = "annotationConnector"
  environment_id = var.pingone_environment_id
}

resource "davinci_connection" "crowd_strike" {
  connector_id   = "crowdStrikeConnector"
  environment_id = var.pingone_environment_id
  name           = "CrowdStrike"
  properties {
    name  = "clientId"
    value = "12345678"
  }
  properties {
    name  = "clientSecret"
    value = "12345"
  }
}

output "cowd_strike_id" {
  value = resource.davinci_connection.crowd_strike.id
}
