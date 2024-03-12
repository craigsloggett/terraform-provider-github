package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ provider.Provider = &GitHubProvider{}

type GitHubProvider struct {
	version string
}

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &GitHubProvider{
			version: version,
		}
	}
}

// Define the provider metadata.
func (p *GitHubProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "github"
	resp.Version = p.version
}

// Define the provider-level schema for configuration data.
func (p *GitHubProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

// Prepare an API client for data sources and resources.
func (p *GitHubProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
}

// Define the data sources implemented by the provider.
func (p *GitHubProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Define the resources implemented by the provider.
func (p *GitHubProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
