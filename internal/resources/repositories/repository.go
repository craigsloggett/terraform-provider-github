package repositories

import (
	"context"
	"fmt"

	"github.com/craigsloggett/terraform-provider-github/internal/common"

	"github.com/google/go-github/v60/github"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &GitHubRepository{}
var _ resource.ResourceWithImportState = &GitHubRepository{}

type GitHubRepository struct {
	client *github.Client
	owner  string
}

type GitHubRepositoryModel struct {
	// Arguments
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
	GitignoreTemplate        types.String `tfsdk:"gitignore_template"`
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

	// Attributes
	ID     types.Int64  `tfsdk:"id"`
	NodeID types.String `tfsdk:"node_id"`
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
			// Arguments
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
				Description:         "Indicates if the repository has wiki enabled.",
				MarkdownDescription: "Indicates if the repository has wiki enabled.",
				Optional:            true,
			},
			"has_discussions": schema.BoolAttribute{
				Description:         "Indicates if the repository has discussions enabled.",
				MarkdownDescription: "Indicates if the repository has discussions enabled.",
				Optional:            true,
			},
			"team_id": schema.Int64Attribute{
				Description:         "The ID of the team associated with the repository.",
				MarkdownDescription: "The ID of the team associated with the repository.",
				Optional:            true,
			},
			"auto_init": schema.BoolAttribute{
				Description:         "Indicates if the repository is initialized with a README.",
				MarkdownDescription: "Indicates if the repository is initialized with a README.",
				Optional:            true,
			},
			"gitignore_template": schema.StringAttribute{
				Description:         "The .gitignore template used by the repository.",
				MarkdownDescription: "The .gitignore template used by the repository.",
				Optional:            true,
			},
			"license_template": schema.StringAttribute{
				Description:         "The license template used by the repository.",
				MarkdownDescription: "The license template used by the repository.",
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
				Description:         "Indicates if the repository is a template repository.",
				MarkdownDescription: "Indicates if the repository is a template repository.",
				Optional:            true,
			},
			// Attributes
			"id": schema.Int64Attribute{
				Description:         "GitHub ID for the repository.",
				MarkdownDescription: "GitHub ID for the repository.",
				Computed:            true,
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"node_id": schema.StringAttribute{
				Description:         "The node ID of the repository.",
				MarkdownDescription: "The node ID of the repository.",
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
		Description:         "This resource allows you to create and manage repositories within your GitHub organization or personal account.",
		MarkdownDescription: "This resource allows you to create and manage repositories within your GitHub organization or personal account.",
	}
}

func (r *GitHubRepository) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	config, ok := req.ProviderData.(*common.ClientConfiguration)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Type from ProviderData",
			fmt.Sprintf("Expected *common.GitHubProviderConfiguration, got: %T", req.ProviderData),
		)
		return
	}

	r.client = config.Client
	r.owner = config.Owner
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

	client := r.client

	// Read Terraform plan data into the model.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)

	if resp.Diagnostics.HasError() {
		return
	}

	repository := &github.Repository{
		Name:                     github.String(types.String.ValueString(model.Name)),
		Description:              github.String(types.String.ValueString(model.Description)),
		Homepage:                 github.String(types.String.ValueString(model.Homepage)),
		Private:                  github.Bool(types.Bool.ValueBool(model.Private)),
		HasIssues:                github.Bool(types.Bool.ValueBool(model.HasIssues)),
		HasProjects:              github.Bool(types.Bool.ValueBool(model.HasProjects)),
		HasWiki:                  github.Bool(types.Bool.ValueBool(model.HasWiki)),
		HasDiscussions:           github.Bool(types.Bool.ValueBool(model.HasDiscussions)),
		TeamID:                   github.Int64(types.Int64.ValueInt64(model.TeamID)),
		AutoInit:                 github.Bool(types.Bool.ValueBool(model.AutoInit)),
		GitignoreTemplate:        github.String(types.String.ValueString(model.GitignoreTemplate)),
		LicenseTemplate:          github.String(types.String.ValueString(model.LicenseTemplate)),
		AllowSquashMerge:         github.Bool(types.Bool.ValueBool(model.AllowSquashMerge)),
		AllowMergeCommit:         github.Bool(types.Bool.ValueBool(model.AllowMergeCommit)),
		AllowRebaseMerge:         github.Bool(types.Bool.ValueBool(model.AllowRebaseMerge)),
		AllowAutoMerge:           github.Bool(types.Bool.ValueBool(model.AllowAutoMerge)),
		DeleteBranchOnMerge:      github.Bool(types.Bool.ValueBool(model.DeleteBranchOnMerge)),
		SquashMergeCommitTitle:   github.String(types.String.ValueString(model.SquashMergeCommitTitle)),
		SquashMergeCommitMessage: github.String(types.String.ValueString(model.SquashMergeCommitMessage)),
		MergeCommitTitle:         github.String(types.String.ValueString(model.MergeCommitTitle)),
		MergeCommitMessage:       github.String(types.String.ValueString(model.MergeCommitMessage)),
		HasDownloads:             github.Bool(types.Bool.ValueBool(model.HasDownloads)),
		IsTemplate:               github.Bool(types.Bool.ValueBool(model.IsTemplate)),
	}

	repo, _, err := client.Repositories.Create(ctx, "", repository)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Communicating with the GitHub API",
			fmt.Sprintf("Unable to create the repository, got error: %s", err),
		)
		return
	}

	model.ID = types.Int64Value(repo.GetID())
	model.NodeID = types.StringValue(repo.GetNodeID())

	// Save updated data into Terraform state.
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *GitHubRepository) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var model GitHubRepositoryModel

	client := r.client
	owner := r.owner

	// Read Terraform plan data into the model.
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)

	if resp.Diagnostics.HasError() {
		return
	}

	repository := &github.Repository{
		Name:                     github.String(types.String.ValueString(model.Name)),
		Description:              github.String(types.String.ValueString(model.Description)),
		Homepage:                 github.String(types.String.ValueString(model.Homepage)),
		Private:                  github.Bool(types.Bool.ValueBool(model.Private)),
		HasIssues:                github.Bool(types.Bool.ValueBool(model.HasIssues)),
		HasProjects:              github.Bool(types.Bool.ValueBool(model.HasProjects)),
		HasWiki:                  github.Bool(types.Bool.ValueBool(model.HasWiki)),
		HasDiscussions:           github.Bool(types.Bool.ValueBool(model.HasDiscussions)),
		TeamID:                   github.Int64(types.Int64.ValueInt64(model.TeamID)),
		AutoInit:                 github.Bool(types.Bool.ValueBool(model.AutoInit)),
		GitignoreTemplate:        github.String(types.String.ValueString(model.GitignoreTemplate)),
		LicenseTemplate:          github.String(types.String.ValueString(model.LicenseTemplate)),
		AllowSquashMerge:         github.Bool(types.Bool.ValueBool(model.AllowSquashMerge)),
		AllowMergeCommit:         github.Bool(types.Bool.ValueBool(model.AllowMergeCommit)),
		AllowRebaseMerge:         github.Bool(types.Bool.ValueBool(model.AllowRebaseMerge)),
		AllowAutoMerge:           github.Bool(types.Bool.ValueBool(model.AllowAutoMerge)),
		DeleteBranchOnMerge:      github.Bool(types.Bool.ValueBool(model.DeleteBranchOnMerge)),
		SquashMergeCommitTitle:   github.String(types.String.ValueString(model.SquashMergeCommitTitle)),
		SquashMergeCommitMessage: github.String(types.String.ValueString(model.SquashMergeCommitMessage)),
		MergeCommitTitle:         github.String(types.String.ValueString(model.MergeCommitTitle)),
		MergeCommitMessage:       github.String(types.String.ValueString(model.MergeCommitMessage)),
		HasDownloads:             github.Bool(types.Bool.ValueBool(model.HasDownloads)),
		IsTemplate:               github.Bool(types.Bool.ValueBool(model.IsTemplate)),
	}

	repo, _, err := client.Repositories.Edit(ctx, owner, model.Name.ValueString(), repository)

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Communicating with the GitHub API",
			fmt.Sprintf("Unable to update the repository, got error: %s", err),
		)
		return
	}

	model.ID = types.Int64Value(repo.GetID())
	model.NodeID = types.StringValue(repo.GetNodeID())

	// Save updated data into Terraform state.
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *GitHubRepository) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var model GitHubRepositoryModel

	client := r.client
	owner := r.owner

	// Read Terraform prior state data into the model.
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)

	if resp.Diagnostics.HasError() {
		return
	}

	_, err := client.Repositories.Delete(ctx, owner, model.Name.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Communicating with the GitHub API",
			fmt.Sprintf("Unable to delete the repository, got error: %s", err),
		)
		return
	}

	// Save updated data into Terraform state.
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *GitHubRepository) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
