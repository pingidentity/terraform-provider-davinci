data "davinci_connections" "all" {
  environment_id = var.pingone_environment_id
}

data "davinci_connections" "http" {
  environment_id = var.pingone_environment_id
  // This will filter output to only include connections using the "httpConnector" type. 
  // Helpful for validation that only one of a certain type exists.
  connector_ids = ["httpConnector"]
}

output "davinci_connection" {
  value = data.davinci_connection.all.connections
}
