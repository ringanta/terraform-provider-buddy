terraform {
  required_version = "= 1.0.2"

  required_providers {
    buddy = {
      source  = "ringanta.id/ringanta/buddy"
      version = "0.2"
    }
  }
}

provider "buddy" {}
