terraform {
  required_version = "~> 1.1"

  required_providers {
    github = {
      source  = "craigsloggett/github"
      version = "dev"
    }
  }
}
