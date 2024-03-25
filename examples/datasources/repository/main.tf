data "github_repository" "example" {
  owner = "craigsloggett"
  repo  = "terraform-provider-github"
}

resource "github_repository" "example" {}
