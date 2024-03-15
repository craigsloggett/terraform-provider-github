package provider

import (
	"context"
	"os"

	"github.com/craigsloggett/terraform-provider-github/internal/resources/repositories"
	"github.com/google/go-github/v60/github"
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
	Token types.String `tfsdk:"token"`
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
	var model GitHubProviderModel

	token := os.Getenv("GITHUB_TOKEN")
	resp.Diagnostics.Append(req.Config.Get(ctx, &model)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Prioritize a token configured in the provider over the GITHUB_TOKEN environment variable.
	if model.Token.ValueString() != "" {
		token = model.Token.ValueString()
	}

	if token == "" {
		resp.Diagnostics.AddError(
			"Missing Personal Access Token Configuration",
			"While configuring the provider, a GitHub token was not found in "+
				"the GITHUB_TOKEN environment variable or provider configuration "+
				"block token attribute.",
		)
	}

	resp.DataSourceData = github.NewClient(nil).WithAuthToken(token)
}

func (p *GitHubProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		repositories.NewGitHubRepository,
	}
}

func (p *GitHubProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}
