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
  environment_id = var.pingone_admin_environment_id
  region         = var.pingone_region
}
