data "github_repository" "example" {
  owner = "craigsloggett"
  name  = "terraform-provider-github"
}

output "github_repository_example_owner" {
  description = "The owner of the repository."
  value       = provider::github::get_repository_owner(data.github_repository.example.full_name)
}
