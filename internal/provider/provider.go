package provider

import (
	"context"
	"fmt"
	"os"

	"github.com/craigsloggett/terraform-provider-github/internal/functions"

	"github.com/google/go-github/v60/github"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ provider.Provider = &GitHubProvider{}

type GitHubProvider struct{}

type GitHubProviderModel struct {
	Owner types.String `tfsdk:"owner"`
	Token types.String `tfsdk:"token"`
}

type GitHubClientConfiguration struct {
	Client *github.Client
	Owner  string
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
			"owner": schema.StringAttribute{
				MarkdownDescription: "The target GitHub organization or individual user account to manage. Alternatively, can be configured using the `GITHUB_OWNER` environment variable.",
				Optional:            true,
			},
			"token": schema.StringAttribute{
				MarkdownDescription: "The GitHub fine-grained personal access token used to authenticate with the API. Alternatively, can be configured using the `GITHUB_TOKEN` environment variable.",
				Optional:            true,
			},
		},
	}
}

func (p *GitHubProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var model GitHubProviderModel

	owner := os.Getenv("GITHUB_OWNER")
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

	client := github.NewClient(nil).WithAuthToken(token)

	// Prioritize an owner configured in the provider over the GITHUB_OWNER environment variable.
	if model.Owner.ValueString() != "" {
		owner = model.Owner.ValueString()
	}

	// If an owner has not been configured, default to the individual user account owning the token.
	if owner == "" {

		user, _, err := client.Users.Get(ctx, "")

		if err != nil {
			resp.Diagnostics.AddError(
				"Error Communicating with the GitHub API",
				fmt.Sprintf("Unable to get user, got error: %s", err),
			)
			return
		}

		owner = user.GetLogin()
	}

	config := &GitHubClientConfiguration{
		Client: client,
		Owner:  owner,
	}

	resp.DataSourceData = config
	resp.ResourceData = config
}

func (p *GitHubProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewGitHubRepositoryDataSource,
	}
}

func (p *GitHubProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewGitHubRepositoryResource,
	}
}

func (p *GitHubProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		functions.NewGetRepositoryName,
		functions.NewGetRepositoryOwner,
	}
}
