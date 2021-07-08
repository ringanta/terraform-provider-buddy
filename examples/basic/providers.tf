terraform {
  required_version = "= 1.0.2"

  required_providers {
    buddy = {
      source  = "ringanta.id/ringanta/buddy"
      version = "0.1"
    }
  }
}

provider "buddy" {}
