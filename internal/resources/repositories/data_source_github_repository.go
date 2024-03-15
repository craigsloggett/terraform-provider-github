package repositories

import (
	"context"
	"fmt"

	"github.com/google/go-github/v60/github"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &GitHubRepository{}

type GitHubRepository struct {
	client any
}

type GitHubRepositoryModel struct {
	Owner    types.String `tfsdk:"owner"`
	Repo     types.String `tfsdk:"repo"`
	Id       types.Int64  `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	FullName types.String `tfsdk:"full_name"`
}

func NewGitHubRepository() datasource.DataSource {
	return &GitHubRepository{}
}

func (d *GitHubRepository) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_repository"
}

func (d *GitHubRepository) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"owner": schema.StringAttribute{
				Description:         "The owner of the repository.",
				MarkdownDescription: "The owner of the repository.",
				Required:            true,
			},
			"repo": schema.StringAttribute{
				Description:         "The name of the repository.",
				MarkdownDescription: "The name of the repository.",
				Required:            true,
			},
			"id": schema.Int64Attribute{
				Description:         "GitHub ID for the repository.",
				MarkdownDescription: "GitHub ID for the repository.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				Description:         "The name of the repository.",
				MarkdownDescription: "The name of the repository.",
				Computed:            true,
			},
			"full_name": schema.StringAttribute{
				Description:         "The full name of the repository.",
				MarkdownDescription: "The full name of the repository.",
				Computed:            true,
			},
		},
		Description:         "Use this data source to retrieve a list of GitHub repositories.",
		MarkdownDescription: "Use this data source to retrieve a list of GitHub repositories.",
	}
}

func (d *GitHubRepository) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// ProviderData is nil until the ConfigureProvider RPC is called.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*github.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *github.Client, got: %T", req.ProviderData),
		)
	}

	d.client = client
}

func (d *GitHubRepository) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var model GitHubRepositoryModel

	client := d.client.(*github.Client)
	resp.Diagnostics.Append(req.Config.Get(ctx, &model)...)

	if resp.Diagnostics.HasError() {
		return
	}

	repo, _, err := client.Repositories.Get(ctx, model.Owner.ValueString(), model.Repo.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Communicating with the GitHub API",
			fmt.Sprintf("Unable to get repository, got error: %s", err),
		)
		return
	}

	model.Id = types.Int64Value(repo.GetID())
	model.Name = types.StringValue(repo.GetName())
	model.FullName = types.StringValue(repo.GetFullName())

	resp.State.Set(ctx, &model)
}
