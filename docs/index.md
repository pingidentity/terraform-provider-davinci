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
    // named davinci for now until merged to actual pingone provider
    davinci = {
      source = "pingidentity/davinci"
    }
    pingone = {
      source = "pingidentity/pingone"
    }
    time = {
      source = "hashicorp/time"
    }
  }
}

// Login with PingOne Admin Environment user
provider "davinci" {
  //Must be Identity Data Admin for Environment 
  // (typically PingOne Admin Environment User)
  username = var.pingone_username
  password = var.pingone_password
  // This base_url is required
  region = "NorthAmerica"
  // User will be _authenticated_ to this environment
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

// DaVinci dependincies are as such:
// - Connections
// - Flows depend on Connections
// - Apps depend on Flows
resource "davinci_flow" "initial_flow" {
  flow_json = file("default_flow.json")
  deploy    = true
  // NOTICE: this is NOT resource.pingone_environment. Dependency is on the role assignment, not environment.
  environment_id = resource.pingone_role_assignment_user.admin_sso.scope_environment_id
}

// All other data/resources can occur after the first DV read action
resource "davinci_application" "one" {
  company_id = resource.pingone_environment.tf_trial.id
  name       = "FromTF"
  oauth {
    enabled = false
    values {
      enabled = false
    }
  }
  policies {
    name = "Simple Flow"
    policy_flows {
      flow_id    = resource.davinci_flow.initial_flow.flow_id
      version_id = -1
      weight     = 100
    }
  }
  saml {
    values {
      enabled                = false
      enforce_signed_request = false
    }
  }
  depends_on = [
    resource.davinci_flow.initial_flow
  ]
}
```
