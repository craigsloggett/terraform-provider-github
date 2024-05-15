package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const providerConfig = `
terraform {
  required_providers {
    github = {
      source  = "craigsloggett/github"
    }
  }
}

provider "github" {}
`

var testAccProtoV6ProviderFactories = map[string]func() (tfprotov6.ProviderServer, error){
	"github": providerserver.NewProtocol6WithError(NewGitHubProvider()()),
}

func testAccPreCheck(t *testing.T) {}
