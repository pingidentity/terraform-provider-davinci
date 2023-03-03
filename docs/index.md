---
page_title: "Provider: PingOne Davinci"
description: |-
  The Davinci provider is used to manage PingOne DaVinci environments.
---

# DaVinci Provider

The "davinci" provider allows operators to manage PingOne DaVinci infrastructure
as code. This provider can be used along with the [PingOne provider](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs)
to stand up fully orchestrated PingOne IAM Experiences.

## Getting Started

This provider will authenticate an API client to PingOne to configure to the DaVinci environment.

### PingOne SSO

You must have a PingOne account and may use multiple environments within that account.

Prerequisites:

> IMPORTANT: To manage multiple Davinci environments with the same credentials, you can add DaVinci to the admin environment. This creates a DaVinci connection and IdP that can be used for SSO into other environments that a user has the `Identity Data Admin` role. The Prerequisites below assume this is the case.

- Have a PingOne environment with DaVinci enabled. 
- Choose or create a user identity in the Administrators account, 
  assign this user the following roles:
    - Environment role and Identity Data Admin role on the Administrators Environment.
    - Organization admin for the organization.
- For each additional DaVinci environment that needs SSO access,
  give this user the Identity Data Admin role for that environment.
  This step can be done in HCL as shown below:

```terraform
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
  name        = "TF Trial"
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
  depends_on     = [data.davinci_connections.read_all]
  environment_id = resource.pingone_role_assignment_user.samir_tf_trial.scope_environment_id
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
    value = "abc"
  }
  property {
    name  = "policyId"
    value = "abc"
  }
  property {
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
  connection_link {
    name = resource.davinci_connection.mfa.name
    // 
    id = resource.davinci_connection.mfa.id
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
    id   = resource.davinci_flow.subflow.id
    name = resource.davinci_flow.subflow.name
  }

  // This depends_on relieves the client from multiple initial authentication attempts
  depends_on = [
    data.davinci_connections.read_all
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) Environment ID PingOne User Login. Default value can be set with the `PINGONE_ENVIRONMENT_ID` environment variable.
- `password` (String, Sensitive) The PingOne password used for SSO into a Davinci tenant.  Default value can be set with the `PINGONE_PASSWORD` environment variable.
- `region` (String) The PingOne region to use.  Options are `AsiaPacific` `Canada` `Europe` and `NorthAmerica`.  Default value can be set with the `PINGONE_REGION` environment variable.
- `username` (String) The PingOne username used for SSO into a Davinci tenant.  Default value can be set with the `PINGONE_USERNAME` environment variable.
