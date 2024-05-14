package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccRepositoryResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: providerConfig + testAccRepositoryResourceConfig("testing-repository"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("github_repository.test", "name", "testing-repository"),
				),
			},
			{
				ResourceName:      "github_repository.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccRepositoryResourceConfig(name string) string {
	return fmt.Sprintf(`
resource "github_repository" "test" {
  name               = %[1]q
}
`, name)
}
