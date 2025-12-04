package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/knownvalue"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/tfjsonpath"
)

func testAccRepositoryResourceConfig(name string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
  name                     = %[1]q

  allow_auto_merge            = true
  allow_merge_commit          = true
  allow_rebase_merge          = true
  allow_squash_merge          = true
  auto_init                   = true
  delete_branch_on_merge      = true
  description                 = "This is a description."
  gitignore_template          = "Terraform"
  has_discussions             = false
  has_issues                  = true
  has_projects                = false
  has_wiki                    = false
  homepage                    = "https://github.com"
  is_template                 = false
  license_template            = "mpl-2.0"
  private                     = false
  squash_merge_commit_message = "COMMIT_MESSAGES"
  squash_merge_commit_title   = "COMMIT_OR_PR_TITLE"
  merge_commit_message        = "PR_BODY"
  merge_commit_title          = "PR_TITLE"

  template_repository = "terraform-module-template"
  template_owner      = "craigsloggett-lab"
}
`, name)
}

func TestAccRepositoryResource(t *testing.T) {
	repoName := "testing-repository-" + acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + testAccRepositoryResourceConfig(repoName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("name"),
						knownvalue.StringExact(repoName),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("description"),
						knownvalue.StringExact("This is a description."),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("homepage"),
						knownvalue.StringExact("https://github.com"),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("private"),
						knownvalue.Bool(false),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("has_issues"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("has_projects"),
						knownvalue.Bool(false),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("has_wiki"),
						knownvalue.Bool(false),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("has_discussions"),
						knownvalue.Bool(false),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("auto_init"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("gitignore_template"),
						knownvalue.StringExact("Terraform"),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("license_template"),
						knownvalue.StringExact("mpl-2.0"),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("allow_squash_merge"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("allow_merge_commit"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("allow_rebase_merge"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("allow_auto_merge"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("delete_branch_on_merge"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("squash_merge_commit_title"),
						knownvalue.StringExact("COMMIT_OR_PR_TITLE"),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("squash_merge_commit_message"),
						knownvalue.StringExact("COMMIT_MESSAGES"),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("merge_commit_message"),
						knownvalue.StringExact("PR_BODY"),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("merge_commit_title"),
						knownvalue.StringExact("PR_TITLE"),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("is_template"),
						knownvalue.Bool(false),
					),
				},
			},
			{
				ResourceName:      "github_repository.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"auto_init",
					"gitignore_template",
					"license_template",
				},
			},
			{
				Config: providerConfig + testAccRepositoryResourceConfig(repoName+"-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("github_repository.test", "name", repoName+"-updated"),
				),
			},
		},
	})
}

// Disable Merge Commits

func testAccRepositoryResourceNoMergeCommitConfig(name string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
  name                     = %[1]q

  allow_auto_merge            = true
  allow_merge_commit          = false
  allow_rebase_merge          = true
  allow_squash_merge          = true
  auto_init                   = true
  delete_branch_on_merge      = true
  description                 = "This is a description."
  gitignore_template          = "Terraform"
  has_discussions             = false
  has_issues                  = true
  has_projects                = false
  has_wiki                    = false
  homepage                    = "https://github.com"
  is_template                 = false
  license_template            = "mpl-2.0"
  private                     = false
  squash_merge_commit_message = "COMMIT_MESSAGES"
  squash_merge_commit_title   = "COMMIT_OR_PR_TITLE"

  template_repository = "terraform-module-template"
  template_owner      = "craigsloggett-lab"
}
`, name)
}

func TestAccRepositoryResourceNoMergeCommit(t *testing.T) {
	repoName := "testing-repository-" + acctest.RandString(8)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + testAccRepositoryResourceNoMergeCommitConfig(repoName),
				ConfigStateChecks: []statecheck.StateCheck{
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("name"),
						knownvalue.StringExact(repoName),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("description"),
						knownvalue.StringExact("This is a description."),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("homepage"),
						knownvalue.StringExact("https://github.com"),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("private"),
						knownvalue.Bool(false),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("has_issues"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("has_projects"),
						knownvalue.Bool(false),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("has_wiki"),
						knownvalue.Bool(false),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("has_discussions"),
						knownvalue.Bool(false),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("auto_init"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("gitignore_template"),
						knownvalue.StringExact("Terraform"),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("license_template"),
						knownvalue.StringExact("mpl-2.0"),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("allow_squash_merge"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("allow_merge_commit"),
						knownvalue.Bool(false),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("allow_rebase_merge"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("allow_auto_merge"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("delete_branch_on_merge"),
						knownvalue.Bool(true),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("squash_merge_commit_title"),
						knownvalue.StringExact("COMMIT_OR_PR_TITLE"),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("squash_merge_commit_message"),
						knownvalue.StringExact("COMMIT_MESSAGES"),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("is_template"),
						knownvalue.Bool(false),
					),
				},
			},
			{
				ResourceName:      "github_repository.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"auto_init",
					"gitignore_template",
					"license_template",
				},
			},
			{
				Config: providerConfig + testAccRepositoryResourceNoMergeCommitConfig(repoName+"-updated"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("github_repository.test", "name", repoName+"-updated"),
				),
			},
		},
	})
}
