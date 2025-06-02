terraform {
  required_providers {
    davinci = {
      source  = "pingidentity/davinci"
      version = "~> 0.5"
    }
  }
}

provider "davinci" {
  username       = var.pingone_dv_admin_username
  password       = var.pingone_dv_admin_password
  environment_id = var.pingone_environment_id
  region         = var.pingone_region
}
