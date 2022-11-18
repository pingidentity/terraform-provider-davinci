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

// DaVinci dependincies are as such and should be de:
// - Connections/Variables
// - Flows depends_on Connections/Variables
// - Apps depends_on Flows
resource "davinci_flow" "initial_flow" {
  flow_json = file("default_flow.json")
  deploy    = true
  // NOTICE: this is NOT resource.pingone_environment. Dependency is on the role assignment, not environment.
  environment_id = resource.pingone_role_assignment_user.admin_sso.scope_environment_id
}

// All other data/resources can occur after the first DV read action
resource "davinci_application" "one" {
  environment_id = resource.pingone_role_assignment_user.admin_sso.scope_environment_id
  name           = "FromTF"
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
