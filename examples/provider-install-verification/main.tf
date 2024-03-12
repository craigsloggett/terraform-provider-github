terraform {
  required_providers {
    github = {
      source = "craigsloggett/github"
    }
  }
}


provider "github" {}

data "github_repos" "example" {}
