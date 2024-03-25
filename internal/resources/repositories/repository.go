package repositories

import (
	"context"
	"fmt"

	"github.com/google/go-github/v60/github"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var _ resource.Resource = &GitHubRepository{}
var _ resource.ResourceWithImportState = &GitHubRepository{}

type GitHubRepository struct {
	client any
}

type GitHubRepositoryModel struct {
}

func NewGitHubRepository() resource.Resource {
	return &GitHubRepository{}
}

func (r *GitHubRepository) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_repository"
}

func (r *GitHubRepository) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
}

func (r *GitHubRepository) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *GitHubRepository) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *GitHubRepository) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
