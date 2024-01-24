data "davinci_connections" "all_connections" {
  environment_id = var.pingone_environment_id
}

data "davinci_connections" "all_connections_filtered_by_connector_id" {
  environment_id = var.pingone_environment_id

  connector_ids = ["httpConnector"]
}

