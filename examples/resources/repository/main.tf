resource "github_repository" "example" {
  name       = "testing-repository"
  private    = true
  visibility = "public"
}
