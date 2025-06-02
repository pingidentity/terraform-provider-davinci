terraform {
  required_providers {
    davinci = {
      source  = "pingidentity/davinci"
      version = "~> 0.5"
    }
  }
}

provider "davinci" {
}
