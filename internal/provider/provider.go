package provider

import (
	"context"
	"os"

	"github.com/craigsloggett/terraform-provider-github/internal/resources/repositories"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &GitHubProvider{}

type GitHubProvider struct {
}

type GitHubProviderModel struct {
	ApiToken types.String `tfsdk:"token"`
}

func NewGitHubProvider() func() provider.Provider {
	return func() provider.Provider {
		return &GitHubProvider{}
	}
}

func (p *GitHubProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "github"
}

func (p *GitHubProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"token": schema.StringAttribute{
				MarkdownDescription: "The GitHub fine-grained personal access token used to authenticate with the API.",
				Optional:            true,
			},
		},
	}
}

func (p *GitHubProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	apiToken := os.Getenv("GITHUB_TOKEN")

	var data GitHubProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if data.ApiToken.ValueString() != "" {
		apiToken = data.ApiToken.ValueString()
	}

	if apiToken == "" {
		resp.Diagnostics.AddError(
			"Missing API Token Configuration",
			"While configuring the provider, a GitHub API token was not found in "+
				"the GITHUB_TOKEN environment variable or provider configuration "+
				"block token attribute.",
		)
	}
}

func (p *GitHubProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		repositories.NewGitHubRepository,
	}
}

func (p *GitHubProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
