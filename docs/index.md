---
page_title: "Provider: PingOne DaVinci"
description: |-
  The DaVinci provider is used to manage PingOne DaVinci environments.
---

# DaVinci Provider

The "davinci" provider allows operators to manage PingOne DaVinci infrastructure
as code. This provider can be used along with the [PingOne provider](https://registry.terraform.io/providers/pingidentity/pingone/latest/docs)
to stand up fully orchestrated PingOne IAM Experiences.

## Getting Started

This provider will authenticate an API client via username and password of an Identity in PingOne to configure to the DaVinci environment.

### PingOne SSO

To manage multiple PingOne DaVinci environments with the same credentials, identify and treat one user in one environment as a service account. This is commonly a user in the `Administrators` environment. The environment in which the user exists must default to the `Single_Factor` Authentication Policy. Thus, it may be more appropriate to create this user in an alternate environment.

This user must: 
- Be in a PingOne Environment with DaVinci enabled. 
  - The PingOne environment must default to the  `Single_Factor` Authentication Policy
- Have at least the following roles in the environment in which it exists:
  - `Environment Admin` for it's own environment
  - `Identity Data Admin` for it's own environment
  - `Organization Admin` at the organization level. 
- For each additional DaVinci environment that needs SSO access, give this user the `Identity Data Admin` and `Environment Admin` roles for that environment.  See the **[Full deployment example with a PingOne environment](#full-deployment-example-with-a-pingone-environment)** example below for how to do this in HCL.

## Example Usage

### Authenticate using static credentials

```terraform
terraform {
  required_providers {
    davinci = {
      source = "pingidentity/davinci"
    }
  }
}

provider "davinci" {
  username       = var.pingone_dv_admin_username
  password       = var.pingone_dv_admin_password
  environment_id = var.pingone_environment_id
  region         = var.pingone_region
}
```

### Authenticate using credentials as environment variables

```terraform
terraform {
  required_providers {
    davinci = {
      source = "pingidentity/davinci"
    }
  }
}

provider "davinci" {
}
```

```shell
$ export PINGONE_USERNAME="admin-username"
$ export PINGONE_PASSWORD="admin-user-password"
$ export PINGONE_ENVIRONMENT_ID="dv-env-id"
$ export PINGONE_REGION="NorthAmerica"
$ terraform plan
```

### Full deployment example with a PingOne environment
The following assumes that the DV admin user has been provided the `Environment Admin` and `Identity Data Admin` roles scoped to the PingOne environment containing the Identity, and the `Organization Admin` role scoped to the organization.

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
  // NOTICE: This read action has a dependency on the role assignment, not environment.
  // Assigning this correctly ensures the role is not destroyed before DaVinci resources during `terraform destroy`.
  depends_on = [
    pingone_role_assignment_user.admin_sso_identity_admin,
    pingone_role_assignment_user.admin_sso_environment_admin
  ]
  environment_id = pingone_environment.dv_example.id
}

resource "davinci_flow" "mainflow" {
  // This depends_on relieves the client from multiple initial authentication attempts
  depends_on = [
    data.davinci_connections.read_all
  ]

  // pingone_sign_on_and_password_reset.json represents an export of the sample flow in a new environment
  flow_json = file("pingone_sign_on_and_password_reset.json")
  deploy    = true

  environment_id = pingone_environment.dv_example.id

  // Dependent connections are defined in conection_link blocks. 
  // It is typically required to define all connections referenced the flow_json. This prevents a mismatch between the flow_json and the connection ids

  // This sample uses bootstrapped connections and references the hardcoded default values of those connections.
  connection_link {
    id   = "6d8f6f706c45fd459a86b3f092602544"
    name = "Error"
  }

  connection_link {
    id   = "de650ca45593b82c49064ead10b9fe17"
    name = "Functions"
  }

  connection_link {
    id   = "867ed4363b2bc21c860085ad2baa817d"
    name = "Http"
  }

  connection_link {
    id   = "94141bf2f1b9b59a5f5365ff135e02bb"
    name = "PingOne SSO"
  }

  connection_link {
    id   = "3b55f2ca6689560c64cb5bed5afbe40f"
    name = "Token Management"
  }

  connection_link {
    id   = "4cb5e3465d718a84087ec9b6a6251e52"
    name = "User Policy"
  }

  connection_link {
    id   = "06922a684039827499bdbdd97f49827b"
    name = "Variables"
  }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `environment_id` (String) Environment ID PingOne User Login. Default value can be set with the `PINGONE_ENVIRONMENT_ID` environment variable.
- `password` (String, Sensitive) The PingOne password used for SSO into a Davinci tenant.  Default value can be set with the `PINGONE_PASSWORD` environment variable.
- `region` (String) The PingOne region to use.  Options are `AsiaPacific` `Canada` `Europe` and `NorthAmerica`.  Default value can be set with the `PINGONE_REGION` environment variable.
- `username` (String) The PingOne username used for SSO into a Davinci tenant.  Default value can be set with the `PINGONE_USERNAME` environment variable.
