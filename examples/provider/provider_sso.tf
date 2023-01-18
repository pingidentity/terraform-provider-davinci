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
  username = var.pingone_username
  password = var.pingone_password
  region   = var.pingone_region
  // User should exist in Identities of this environment
  environment_id = var.pingone_environment_id
}

provider "pingone" {
  client_id      = var.pingone_client_id
  client_secret  = var.pingone_client_secret
  environment_id = var.pingone_environment_id
  region         = var.pingone_region

  force_delete_production_type = false

}

resource "pingone_environment" "tf_trial" {
  name        = "Temp TF Trial"
  description = "My new environment"
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
    type = "Risk"
  }
  service {
    type = "DaVinci"
  }

}

data "pingone_role" "identity_data_admin" {
  name = "Identity Data Admin"
  depends_on = [
    pingone_environment.tf_trial
  ]
}

resource "pingone_role_assignment_user" "admin_sso" {
  environment_id       = var.environment_id
  user_id              = var.admin_user_id
  role_id              = data.pingone_role.identity_data_admin.id
  scope_environment_id = resource.pingone_environment.tf_trial.id
}

// This simple read action is used to as a precursor to all other data/resources
// Every other data/resource should have a `depends_on` pointing to this read action
data "davinci_connections" "read_all" {
  // NOTICE: this is NOT resource.pingone_environment. Dependency is on the role assignment, not environment.
  // Assigning this correctly ensures the role is not destroyed before davinci resources.
  environment_id = resource.pingone_role_assignment_user.admin_sso.scope_environment_id
}

// Sample connection resource. Property names must be discovered by looking at API read response
resource "davinci_connection" "mfa" {
  depends_on     = [data.davinci_connections.all]
  environment_id = resource.pingone_role_assignment_user.samir_tf_trial.scope_environment_id
  connector_id   = "pingOneMfaConnector"
  name           = "PingOne MFA"
  properties {
    name  = "clientId"
    value = "abc"
  }
  properties {
    name  = "clientSecret"
    value = "abc"
  }
  properties {
    name  = "envId"
    value = "abc"
  }
  properties {
    name  = "policyId"
    value = "abc"
  }
  properties {
    name  = "region"
    value = "EU"
  }
}

resource "davinci_flow" "mainflow" {
  flow_json = file("mainflow.json")
  deploy    = true
  // NOTICE: this is NOT resource.pingone_environment. Dependency is on the role assignment, not environment.
  environment_id = resource.pingone_role_assignment_user.admin_sso.scope_environment_id

  // Dependent connections are defined in conections blocks. 
  // It is best practice to define all connections referenced the flow_json. This prevents a mismatch between the flow_json and the connections

  // This sample references a managed connection
  connections {
    name = resource.davinci_connection.mfa.name
    // 
    id = resource.davinci_connection.mfa.id
  }
  // This sample uses a bootstrapped connection
  connections {
    name = "Http"
    // default connection id for the bootstrapped Http connector
    id = "867ed4363b2bc21c860085ad2baa817d"
  }

  // Dependent subflows are defined in subflows blocks.
  // These should always point to managed subflows
  subflows {
    subflow_id   = resource.davinci_flow.subflow.flow_id
    subflow_name = resource.davinci_flow.subflow.flow_name
  }

  // This depends_on relieves the client from multiple initial authentication attempts
  depends_on = [
    data.davinci_connections.read_all
  ]
}
