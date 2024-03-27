resource "davinci_connection" "my_awesome_annotation_connection" {
  environment_id = var.pingone_environment_id

  name         = "myAnnotationConnector"
  connector_id = "annotationConnector"
}

resource "davinci_connection" "my_awesome_crowdstrike_connection" {
  environment_id = var.pingone_environment_id

  connector_id = "crowdStrikeConnector"
  name         = "CrowdStrike"

  property {
    name  = "clientId"
    value = var.crowdstrike_client_id
  }

  property {
    name  = "clientSecret"
    value = var.crowdstrike_client_secret
  }
}
