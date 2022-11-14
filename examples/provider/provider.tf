terraform {
  required_providers {
    // named davinci for now until merged to actual pingone provider
    davinci = {
      source = "pingidentity/davinci"
    }
  }
}

// Login with DaVinci Team Member
provider "davinci" {
  username       = var.pingone_username
  password       = var.pingone_password
  environment_id = var.pingone_environment_id
  region         = var.pingone_region
}
