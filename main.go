package main

import (
	"context"
	"log"

	"github.com/craigsloggett/terraform-provider-github/internal/provider"
	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

func main() {
	opts := providerserver.ServeOpts{
		Address: "registry.terraform.io/craigsloggett/github",
		Debug:   false,
	}

	err := providerserver.Serve(context.Background(), provider.NewGitHubProvider(), opts)

	if err != nil {
		log.Fatal(err.Error())
	}
}
