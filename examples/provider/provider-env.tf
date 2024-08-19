terraform {
  required_providers {
    davinci = {
      source  = "pingidentity/davinci"
      version = "~> 0.4"
    }
  }
}

provider "davinci" {
}
