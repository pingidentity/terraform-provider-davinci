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

- Have a PingOne environment
- Add the DaVinci service to your Administrators environment
- Choose or create a user identity in the Administrators account, 
  give this user the Environment role over the admin Organization.
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
  region   = "NorthAmerica"
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
  environment_id = resource.pingone_role_assignment_user.admin_sso.scope_environment_id
}

resource "davinci_flow" "initial_flow" {
  flow_json = file("default_flow.json")
  deploy    = true
  // NOTICE: this is NOT resource.pingone_environment. Dependency is on the role assignment, not environment.
  environment_id = resource.pingone_role_assignment_user.admin_sso.scope_environment_id

  // This depends_on relieves the client from multiple initial authentication attempts
  depends_on = [
    data.davinci_connections.read_all
  ]
}
```
