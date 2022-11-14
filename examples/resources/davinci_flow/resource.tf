resource "davinci_flow" "my_flow" {
  flow_json      = "{\"customerId\":\"1234\",\"name\":\"tftesting\",\"description\":\"\",\"flowStatus\":\"enabled\",\"createdDate...\"\"connectorIds\":[\"httpConnector\"],\"savedDate\":1662961640542,\"variables\":[]}"
  deploy         = false
  environment_id = var.pingone_environment_id
}

resource "davinci_flow" "file_flow" {
  flow_json      = file("flow.json")
  deploy         = true
  environment_id = var.pingone_environment_id
}
