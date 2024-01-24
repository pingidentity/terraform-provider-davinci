data "davinci_connection" "example_by_name" {
  environment_id = var.pingone_environment_id

  name = "Http"
}

data "davinci_connection" "example_by_id" {
  environment_id = var.pingone_environment_id

  id = "867ed4363b2bc21c860085ad2baa817d"
}