terraform {
  required_providers {
    davinci = {
      source  = "pingidentity/davinci"
      version = "~> 0.3"
    }
    pingone = {
      source  = "pingidentity/pingone"
      version = "~> 0.24"
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
  name        = "DaVinci Terraform Example"
  description = "A new trial environment for DaVinci Terraform configuration-as-code."
  license_id  = var.license_id

  service {
    type = "SSO"
  }

  service {
    type = "MFA"
  }

  service {
    type = "DaVinci"
    tags = ["DAVINCI_MINIMAL"]
  }
}

resource "davinci_flow" "mainflow" {
  environment_id = pingone_environment.dv_example.id

  # ...
}