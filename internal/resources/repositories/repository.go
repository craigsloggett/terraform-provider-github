package repositories

import (
	"context"
	"fmt"

	"github.com/google/go-github/v60/github"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &GitHubRepository{}
var _ resource.ResourceWithImportState = &GitHubRepository{}

type GitHubRepository struct {
	client any
}

type GitHubRepositoryModel struct {
	Id   types.Int64  `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

func NewGitHubRepository() resource.Resource {
	return &GitHubRepository{}
}

func (r *GitHubRepository) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_repository"
}

func (r *GitHubRepository) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description:         "GitHub ID for the repository.",
				MarkdownDescription: "GitHub ID for the repository.",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				Description:         "The name of the repository.",
				MarkdownDescription: "The name of the repository.",
				Required:            true,
			},
		},
		Description:         "Use this resource to create a GitHub repository for the authenticated user.",
		MarkdownDescription: "Use this resource to create a GitHub repository for the authenticated user.",
	}
}

func (r *GitHubRepository) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*github.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Type from ProviderData",
			fmt.Sprintf("Expected *github.Client, got: %T", req.ProviderData),
		)
		return
	}

	r.client = client
}

func (r *GitHubRepository) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var model GitHubRepositoryModel

	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)

	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *GitHubRepository) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var model GitHubRepositoryModel

	client, ok := r.client.(*github.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Type from GitHubProvider's Client Field",
			fmt.Sprintf("Expected *github.Client, got: %T", r.client),
		)
		return
	}

	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)

	if resp.Diagnostics.HasError() {
		return
	}

	repository := &github.Repository{
		Name: github.String(types.String.ValueString(model.Name)),
	}

	repo, _, err := client.Repositories.Create(ctx, "", repository)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Communicating with the GitHub API",
			fmt.Sprintf("Unable to create the repository, got error: %s", err),
		)
		return
	}

	model.Id = types.Int64Value(repo.GetID())

	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *GitHubRepository) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *GitHubRepository) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *GitHubRepository) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
