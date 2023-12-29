terraform {
  required_providers {
    davinci = {
      source  = "pingidentity/davinci"
      version = "~> 0.3"
    }
  }
}

provider "davinci" {
}
