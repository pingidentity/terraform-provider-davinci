terraform {
  required_providers {
    davinci = {
      source  = "pingidentity/davinci"
      version = "~> 0.5"
    }
    pingone = {
      source  = "pingidentity/pingone"
      version = ">= 1.0, < 2.0"
    }
  }
}

provider "davinci" {
  username       = var.pingone_dv_admin_username
  password       = var.pingone_dv_admin_password
  environment_id = var.pingone_admin_environment_id
  region         = var.pingone_region
}

provider "pingone" {
  client_id      = var.pingone_admin_client_id
  client_secret  = var.pingone_admin_client_secret
  environment_id = var.pingone_admin_environment_id
  region_code    = var.pingone_region_code
}

resource "pingone_environment" "dv_example" {
  name        = "DaVinci Terraform Example"
  description = "A new trial environment for DaVinci Terraform configuration-as-code."
  license_id  = var.license_id

  services = [
    {
      type = "SSO"
    },
    {
      type = "MFA"
    },
    {
      type = "DaVinci"
      tags = ["DAVINCI_MINIMAL"]
    }
  ]
}

resource "davinci_flow" "mainflow" {
  environment_id = pingone_environment.dv_example.id

  # ...
}