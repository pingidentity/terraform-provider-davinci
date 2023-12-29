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

To get started using the PingOne and DaVinci Terraform providers, first you'll need an active PingOne cloud subscription with the DaVinci service license additions.  Get instant access with a [PingOne trial account](https://www.pingidentity.com/en/try-ping.html), or read more about Ping Identity at [pingidentity.com](https://www.pingidentity.com).

Further information about enabing the DaVinci service license can be found [here](https://terraform.pingidentity.com/getting-started/davinci/#the-pingone-davinci-service-license).

### Configure PingOne and DaVinci for Terraform access

For detailed instructions on how to prepare PingOne and DaVinci for Terraform access, see the [DaVinci provider getting started guide](https://terraform.pingidentity.com/getting-started/davinci/#configure-pingone-for-terraform-access) at [terraform.pingidentity.com](https://terraform.pingidentity.com).

## Example Usage

The following examples demonstrate how to configure the DaVinci provider. 

If the PingOne Environment and Permissions are managed in:
  - **separate module** - use [static credentials](#authenticate-using-static-credentials) or [environment variable credentials](#authenticate-using-credentials-as-environment-variables).
  - **same module as davinci configuration** - use the [full deployment example](#full-deployment-example-with-a-pingone-environment).\

For a more thorough example on how to set up PingOne for DaVinci API authentication see [getting started}(https://terraform.pingidentity.com/getting-started/davinci/)

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
$ # Optional if using PingOne TF provider
$ export PINGONE_CLIENT_ID="client-id"
$ export PINGONE_CLIENT_SECRET="client-secret"
$ terraform plan
```

### Full deployment example with a PingOne environment

The following assumes that the PingOne worker app has been provided the `Environment Admin` and `Identity Data Admin` roles scoped to the PingOne environment containing the Identity, and the `Organization Admin` role scoped to the organization.

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
data "pingone_role" "davinci_admin" {
  name = "DaVinci Admin"
}

// Get the ID of the DV admin user
data "pingone_user" "dv_admin_user" {
  environment_id = var.pingone_admin_environment_id
  username       = var.pingone_dv_admin_username
}

// Assign the "DaVinci" role to the DV admin user
resource "pingone_role_assignment_user" "davinci_admin_sso" {
  environment_id       = var.pingone_admin_environment_id
  user_id              = data.pingone_user.dv_admin_user.id
  role_id              = data.pingone_role.davinci_admin.id
  scope_environment_id = pingone_environment.dv_example.id
}

// This simple read action is used to as a precursor to all other data/resources
// Every other data/resource should have a `depends_on` pointing to this read action
data "davinci_connections" "read_all" {
  // NOTICE: This read action has a dependency on the role assignment, not environment.
  // Assigning this correctly ensures the role is not destroyed before DaVinci resources during `terraform destroy`.
  depends_on = [
    pingone_role_assignment_user.davinci_admin_sso,
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
    name = "Error Customize"
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
    name = "PingOne"
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

> This above example was updated to reflect new DaVinci roles added on August 5th, 2023. The `DaVinci Admin` role is now required for user SSO and interaction with DaVinci APIs.

### Custom User Agent information

The DaVinci provider allows custom information to be appended to the default user agent string (that includes Terraform provider version information) by setting the `DAVINCI_TF_APPEND_USER_AGENT` environment variable.  This can be useful when troubleshooting issues with Ping Identity Support, or adding context to HTTP requests.

```shell
$ export DAVINCI_TF_APPEND_USER_AGENT="Jenkins/2.426.2"
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `access_token` (String) PingOne DaVinci specific access token. Must be authorized for environment_id.  Default value can be set with the `PINGONE_DAVINCI_ACCESS_TOKEN` environment variable. Must provide username and password, or access_token.
- `environment_id` (String) Environment ID PingOne User Login. Default value can be set with the `PINGONE_ENVIRONMENT_ID` environment variable.
- `host_url` (String) To override the default region-based url, provide a PingOne DaVinci API host url. Default value can be set with the `PINGONE_DAVINCI_HOST_URL` environment variable.
- `password` (String, Sensitive) The PingOne password used for SSO into a Davinci tenant.  Default value can be set with the `PINGONE_PASSWORD` environment variable. Must provide username and password, or access_token.
- `region` (String) The PingOne region to use.  Options are `AsiaPacific` `Canada` `Europe` and `NorthAmerica`.  Default value can be set with the `PINGONE_REGION` environment variable.
- `username` (String) The PingOne username used for SSO into a Davinci tenant.  Default value can be set with the `PINGONE_USERNAME` environment variable. Must provide username and password, or access_token.
