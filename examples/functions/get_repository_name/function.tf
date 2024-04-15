data "github_repository" "example" {
  owner = "craigsloggett"
  name  = "terraform-provider-github"
}

output "github_repository_example_name" {
  description = "The name of the repository."
  value       = provider::github::get_repository_name(data.github_repository.example.full_name)
}
