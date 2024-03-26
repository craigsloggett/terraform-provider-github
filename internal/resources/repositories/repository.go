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
	Name                     types.String `tfsdk:"name"`
	Description              types.String `tfsdk:"description"`
	Homepage                 types.String `tfsdk:"homepage"`
	Private                  types.Bool   `tfsdk:"private"`
	HasIssues                types.Bool   `tfsdk:"has_issues"`
	HasProjects              types.Bool   `tfsdk:"has_projects"`
	HasWiki                  types.Bool   `tfsdk:"has_wiki"`
	HasDiscussions           types.Bool   `tfsdk:"has_discussions"`
	TeamID                   types.Int64  `tfsdk:"team_id"`
	AutoInit                 types.Bool   `tfsdk:"auto_init"`
	GitIgnoreTemplate        types.String `tfsdk:"git_ignore_template"`
	LicenseTemplate          types.String `tfsdk:"license_template"`
	AllowSquashMerge         types.Bool   `tfsdk:"allow_squash_merge"`
	AllowMergeCommit         types.Bool   `tfsdk:"allow_merge_commit"`
	AllowRebaseMerge         types.Bool   `tfsdk:"allow_rebase_merge"`
	AllowAutoMerge           types.Bool   `tfsdk:"allow_auto_merge"`
	DeleteBranchOnMerge      types.Bool   `tfsdk:"delete_branch_on_merge"`
	SquashMergeCommitTitle   types.String `tfsdk:"squash_merge_commit_title"`
	SquashMergeCommitMessage types.String `tfsdk:"squash_merge_commit_message"`
	MergeCommitTitle         types.String `tfsdk:"merge_commit_title"`
	MergeCommitMessage       types.String `tfsdk:"merge_commit_message"`
	HasDownloads             types.Bool   `tfsdk:"has_downloads"`
	IsTemplate               types.Bool   `tfsdk:"is_template"`
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
			"name": schema.StringAttribute{
				Description:         "The name of the repository.",
				MarkdownDescription: "The name of the repository.",
				Required:            true,
			},
			"description": schema.StringAttribute{
				Description:         "The description of the repository.",
				MarkdownDescription: "The description of the repository.",
				Optional:            true,
			},
			"homepage": schema.StringAttribute{
				Description:         "The homepage of the repository.",
				MarkdownDescription: "The homepage of the repository.",
				Optional:            true,
			},
			"private": schema.BoolAttribute{
				Description:         "Indicates if the repository is private.",
				MarkdownDescription: "Indicates if the repository is private.",
				Optional:            true,
			},
			"has_issues": schema.BoolAttribute{
				Description:         "Indicates if the repository has issues enabled.",
				MarkdownDescription: "Indicates if the repository has issues enabled.",
				Optional:            true,
			},
			"has_projects": schema.BoolAttribute{
				Description:         "Indicates if the repository has projects enabled.",
				MarkdownDescription: "Indicates if the repository has projects enabled.",
				Optional:            true,
			},
			"has_wiki": schema.BoolAttribute{
				Description:         "Indicates if the repository has the wiki feature enabled.",
				MarkdownDescription: "Indicates if the repository has the wiki feature enabled.",
				Optional:            true,
			},
			"has_discussions": schema.BoolAttribute{
				Description:         "Indicates if the repository has discussions enabled.",
				MarkdownDescription: "Indicates if the repository has discussions enabled.",
				Optional:            true,
			},
			"team_id": schema.Int64Attribute{
				Description:         "The ID of the team that owns the repository. Only applicable to organizations.",
				MarkdownDescription: "The ID of the team that owns the repository. Only applicable to organizations.",
				Optional:            true,
			},
			"auto_init": schema.BoolAttribute{
				Description:         "Indicates if the repository is initialized with a README.",
				MarkdownDescription: "Indicates if the repository is initialized with a README.",
				Optional:            true,
			},
			"git_ignore_template": schema.StringAttribute{
				Description:         "The .gitignore template to apply to the repository upon creation.",
				MarkdownDescription: "The `.gitignore` template to apply to the repository upon creation.",
				Optional:            true,
			},
			"license_template": schema.StringAttribute{
				Description:         "The license template to apply to the repository upon creation.",
				MarkdownDescription: "The license template to apply to the repository upon creation.",
				Optional:            true,
			},
			"allow_squash_merge": schema.BoolAttribute{
				Description:         "Indicates if squash merging is allowed in the repository.",
				MarkdownDescription: "Indicates if squash merging is allowed in the repository.",
				Optional:            true,
			},
			"allow_merge_commit": schema.BoolAttribute{
				Description:         "Indicates if merge commits are allowed in the repository.",
				MarkdownDescription: "Indicates if merge commits are allowed in the repository.",
				Optional:            true,
			},
			"allow_rebase_merge": schema.BoolAttribute{
				Description:         "Indicates if rebase merging is allowed in the repository.",
				MarkdownDescription: "Indicates if rebase merging is allowed in the repository.",
				Optional:            true,
			},
			"allow_auto_merge": schema.BoolAttribute{
				Description:         "Indicates if auto-merging is allowed in the repository.",
				MarkdownDescription: "Indicates if auto-merging is allowed in the repository.",
				Optional:            true,
			},
			"delete_branch_on_merge": schema.BoolAttribute{
				Description:         "Indicates if branches are automatically deleted when pull requests are merged.",
				MarkdownDescription: "Indicates if branches are automatically deleted when pull requests are merged.",
				Optional:            true,
			},
			"squash_merge_commit_title": schema.StringAttribute{
				Description:         "The title of squash merge commits for pull requests.",
				MarkdownDescription: "The title of squash merge commits for pull requests.",
				Optional:            true,
			},
			"squash_merge_commit_message": schema.StringAttribute{
				Description:         "The message of squash merge commits for pull requests.",
				MarkdownDescription: "The message of squash merge commits for pull requests.",
				Optional:            true,
			},
			"merge_commit_title": schema.StringAttribute{
				Description:         "The title of merge commits for pull requests.",
				MarkdownDescription: "The title of merge commits for pull requests.",
				Optional:            true,
			},
			"merge_commit_message": schema.StringAttribute{
				Description:         "The message of merge commits for pull requests.",
				MarkdownDescription: "The message of merge commits for pull requests.",
				Optional:            true,
			},
			"has_downloads": schema.BoolAttribute{
				Description:         "Indicates if the repository has downloads enabled.",
				MarkdownDescription: "Indicates if the repository has downloads enabled.",
				Optional:            true,
			},
			"is_template": schema.BoolAttribute{
				Description:         "Indicates if the repository is marked as a template repository.",
				MarkdownDescription: "Indicates if the repository is marked as a template repository.",
				Optional:            true,
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

	// Read Terraform prior state data into the model.
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state.
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

	// Read Terraform plan data into the model.
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

	// Save updated data into Terraform state.
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *GitHubRepository) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

func (r *GitHubRepository) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

func (r *GitHubRepository) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
