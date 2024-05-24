resource "davinci_connection" "connectorValidsoft" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorValidsoft"
  name         = "My awesome connectorValidsoft"

  property {
    name  = "customAuth"
    type  = "json"
    value = jsonencode({})
  }
}
