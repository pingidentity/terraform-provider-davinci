terraform {
  required_providers {
    davinci = {
      source = "pingidentity/davinci"
    }
    pingone = {
      source = "pingidentity/pingone"
    }
  }
}

provider "davinci" {
  //Must be Identity Data Admin for Environment
  username = var.pingone_dv_admin_username
  password = var.pingone_dv_admin_password
  region   = var.pingone_region
  // User should exist in Identities of this environment
  environment_id = var.pingone_admin_environment_id
}

provider "pingone" {
  client_id      = var.pingone_admin_client_id
  client_secret  = var.pingone_admin_client_secret
  environment_id = var.pingone_admin_environment_id
  region         = var.pingone_region
}

// Create a new environment using the PingOne provider.  The organization must have the DaVinci feature flag enabled.
resource "pingone_environment" "dv_example" {
  name        = "DaVinci Example"
  description = "A new trial environment for DaVinci configuration-as-code."
  type        = "SANDBOX"
  license_id  = var.license_id

  default_population {
    name        = "My Population"
    description = "My new population for users"
  }

  service {
    type = "SSO"
  }
  service {
    type = "MFA"
  }
  service {
    type = "DaVinci"
  }

}

// Get the roles from the organization
data "pingone_role" "identity_data_admin" {
  name = "Identity Data Admin"
}

data "pingone_role" "environment_admin" {
  name = "Environment Admin"
}

// Get the ID of the DV admin user
data "pingone_user" "dv_admin_user" {
  environment_id = var.pingone_admin_environment_id

  username = var.pingone_dv_admin_username
}

// Assign the "Identity Data Admin" role to the DV admin user
resource "pingone_role_assignment_user" "admin_sso_identity_admin" {
  environment_id       = var.pingone_admin_environment_id
  user_id              = data.pingone_user.dv_admin_user.id
  role_id              = data.pingone_role.identity_data_admin.id
  scope_environment_id = pingone_environment.dv_example.id
}

// Assign the "Environment Admin" role to the DV admin user
resource "pingone_role_assignment_user" "admin_sso_environment_admin" {
  environment_id       = var.pingone_admin_environment_id
  user_id              = data.pingone_user.dv_admin_user.id
  role_id              = data.pingone_role.environment_admin.id
  scope_environment_id = pingone_environment.dv_example.id
}

// This simple read action is used to as a precursor to all other data/resources
// Every other data/resource should have a `depends_on` pointing to this read action
data "davinci_connections" "read_all" {
  // NOTICE: This read action has a dependency is on the role assignment, not environment.
  // Assigning this correctly ensures the role is not destroyed before DaVinci resources on `terraform destroy`.
  depends_on = [
    pingone_role_assignment_user.admin_sso_identity_admin,
    pingone_role_assignment_user.admin_sso_environment_admin
  ]
  environment_id = pingone_environment.dv_example.id
}

// Sample connection  Property names must be discovered by looking at API read response
resource "davinci_connection" "mfa" {
  depends_on = [
    data.davinci_connections.read_all
  ]
  environment_id = pingone_environment.dv_example.id
  connector_id   = "pingOneMfaConnector"
  name           = "PingOne MFA"
  property {
    name  = "clientId"
    value = "abc"
  }
  property {
    name  = "clientSecret"
    value = "abc"
  }
  property {
    name  = "envId"
    value = pingone_environment.dv_example.id
  }
  property {
    name  = "region"
    value = "EU"
  }
}

resource "davinci_connection" "flow_conductor" {
  depends_on = [
    data.davinci_connections.read_all
  ]
  environment_id = pingone_environment.dv_example.id
  connector_id   = "flowConnector"
  name           = "Flow Conductor"
}

resource "davinci_flow" "mainflow" {
  // This depends_on relieves the client from multiple initial authentication attempts
  depends_on = [
    data.davinci_connections.read_all
  ]

  flow_json = file("mainflow.json")
  deploy    = true

  environment_id = pingone_environment.dv_example.id

  // Dependent connections are defined in conections blocks. 
  // It is best practice to define all connections referenced the flow_json. This prevents a mismatch between the flow_json and the connections

  // This sample references a managed connection
  connection_link {
    name = davinci_connection.mfa.name
    id   = davinci_connection.mfa.id
  }
  // This sample references a managed connection, which in the main flow, calls the subflow.
  connection_link {
    name = davinci_connection.flow_conductor.name
    id   = davinci_connection.flow_conductor.id
  }
  // This sample uses a bootstrapped connection
  connection_link {
    name = "Http"
    // default connection id for the bootstrapped Http connector
    id = "867ed4363b2bc21c860085ad2baa817d"
  }

  // Dependent subflows are defined in subflows blocks.
  // These should always point to managed subflows
  subflow_link {
    id   = davinci_flow.subflow.id
    name = davinci_flow.subflow.name
  }
}

resource "davinci_flow" "subflow" {
  // This depends_on relieves the client from multiple initial authentication attempts
  depends_on = [
    data.davinci_connections.read_all
  ]

  flow_json = file("subflow.json")
  deploy    = true

  environment_id = pingone_environment.dv_example.id

  // Dependent connections are defined in conections blocks as with the main flow.

  connection_link {
    name = davinci_connection.mfa.name
    id   = davinci_connection.mfa.id
  }

  connection_link {
    name = "Http"
    id   = "867ed4363b2bc21c860085ad2baa817d"
  }
}
