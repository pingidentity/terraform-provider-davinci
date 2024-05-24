resource "davinci_connection" "connectorPaloAltoPrisma" {
  environment_id = var.pingone_environment_id

  connector_id = "connectorPaloAltoPrisma"
  name         = "My awesome connectorPaloAltoPrisma"

  property {
    name  = "baseURL"
    type  = "string"
    value = var.base_url
  }

  property {
    name  = "prismaPassword"
    type  = "string"
    value = var.connectorpaloaltoprisma_property_prisma_password
  }

  property {
    name  = "prismaUsername"
    type  = "string"
    value = var.connectorpaloaltoprisma_property_prisma_username
  }
}
