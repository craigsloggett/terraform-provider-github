resource "github_repository" "example" {
  name                        = "testing-repository"
  description                 = "This is a description."
  homepage                    = "https://github.com"
  private                     = true
  has_issues                  = true
  has_projects                = true
  has_wiki                    = true
  has_discussions             = true
  auto_init                   = true
  gitignore_template          = "Terraform"
  license_template            = "mpl-2.0"
  allow_squash_merge          = true
  allow_merge_commit          = true
  allow_rebase_merge          = false
  allow_auto_merge            = true
  delete_branch_on_merge      = true
  squash_merge_commit_title   = "PR_TITLE"
  squash_merge_commit_message = "COMMIT_MESSAGES"
  merge_commit_title          = "PR_TITLE"
  merge_commit_message        = "PR_BODY"
  is_template                 = true
}
