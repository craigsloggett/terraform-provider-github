resource "github_repository" "example" {
  name = "terraform-aws-module"

  template_repository = "terraform-module-template"
  template_owner      = "craigsloggett-lab"
}
