data "davinci_connection" "http_by_name" {
  environment_id = var.pingone_environment_id
  name           = "Http"
}

data "davinci_connection" "http_by_id" {
  environment_id = var.pingone_environment_id
  // This will filter output to only include connections using the "httpConnector" type. 
  // Helpful for validation that only one of a certain type exists.
  id = "867ed4363b2bc21c860085ad2baa817d"
}