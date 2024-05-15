resource "github_repository" "example" {
  name = "testing-repository"

  allow_auto_merge            = true
  allow_merge_commit          = true
  allow_rebase_merge          = false
  allow_squash_merge          = true
  auto_init                   = true
  delete_branch_on_merge      = true
  description                 = "This is a description."
  gitignore_template          = "Terraform"
  has_discussions             = true
  has_issues                  = true
  has_projects                = true
  has_wiki                    = true
  homepage                    = "https://github.com"
  is_template                 = true
  license_template            = "mpl-2.0"
  merge_commit_message        = "PR_BODY"
  merge_commit_title          = "PR_TITLE"
  private                     = true
  squash_merge_commit_message = "COMMIT_MESSAGES"
  squash_merge_commit_title   = "PR_TITLE"
}
