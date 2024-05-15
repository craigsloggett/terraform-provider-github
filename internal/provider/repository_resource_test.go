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
						knownvalue.StringExact(""),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("homepage"),
						knownvalue.StringExact(""),
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
						knownvalue.Bool(true),
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
						knownvalue.Null(),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("gitignore_template"),
						knownvalue.Null(),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("license_template"),
						knownvalue.Null(),
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
						knownvalue.Bool(false),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("delete_branch_on_merge"),
						knownvalue.Bool(false),
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
						tfjsonpath.New("merge_commit_title"),
						knownvalue.StringExact("MERGE_MESSAGE"),
					),
					statecheck.ExpectKnownValue(
						"github_repository.test",
						tfjsonpath.New("merge_commit_message"),
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

func testAccRepositoryResourceConfig(name string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
  name = %[1]q
}
`, name)
}
