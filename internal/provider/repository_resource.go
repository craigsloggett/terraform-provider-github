package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/go-github/v79/github"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ resource.Resource = &GitHubRepositoryResource{}
var _ resource.ResourceWithImportState = &GitHubRepositoryResource{}

// Types

type GitHubRepositoryResource struct {
	client       *github.Client
	owner        string
	organization string
}

type GitHubRepositoryResourceModel struct {
	// Arguments
	Name                     types.String `tfsdk:"name"`
	Description              types.String `tfsdk:"description"`
	Homepage                 types.String `tfsdk:"homepage"`
	Private                  types.Bool   `tfsdk:"private"`
	HasIssues                types.Bool   `tfsdk:"has_issues"`
	HasProjects              types.Bool   `tfsdk:"has_projects"`
	HasWiki                  types.Bool   `tfsdk:"has_wiki"`
	HasDiscussions           types.Bool   `tfsdk:"has_discussions"`
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
	IsTemplate               types.Bool   `tfsdk:"is_template"`

	// Template Arguments
	TemplateRepository types.String `tfsdk:"template_repository"`
	TemplateOwner      types.String `tfsdk:"template_owner"`

	// Attributes
	ID     types.Int64  `tfsdk:"id"`
	NodeID types.String `tfsdk:"node_id"`
}

// Constructor

func NewGitHubRepositoryResource() resource.Resource {
	return &GitHubRepositoryResource{}
}

// Helpers

type expansionMode int

const (
	expandForCreate expansionMode = iota
	expandForUpdate
)

func expandRepository(model GitHubRepositoryResourceModel, mode expansionMode) *github.Repository {
	repo := &github.Repository{
		Name:                github.Ptr(model.Name.ValueString()),
		Description:         github.Ptr(model.Description.ValueString()),
		Homepage:            github.Ptr(model.Homepage.ValueString()),
		Private:             github.Ptr(model.Private.ValueBool()),
		HasIssues:           github.Ptr(model.HasIssues.ValueBool()),
		HasProjects:         github.Ptr(model.HasProjects.ValueBool()),
		HasWiki:             github.Ptr(model.HasWiki.ValueBool()),
		HasDiscussions:      github.Ptr(model.HasDiscussions.ValueBool()),
		AllowSquashMerge:    github.Ptr(model.AllowSquashMerge.ValueBool()),
		AllowMergeCommit:    github.Ptr(model.AllowMergeCommit.ValueBool()),
		AllowRebaseMerge:    github.Ptr(model.AllowRebaseMerge.ValueBool()),
		AllowAutoMerge:      github.Ptr(model.AllowAutoMerge.ValueBool()),
		DeleteBranchOnMerge: github.Ptr(model.DeleteBranchOnMerge.ValueBool()),
		IsTemplate:          github.Ptr(model.IsTemplate.ValueBool()),
	}

	// Create is the only time you can successfully pass these parameters into the GitHub API.
	if mode == expandForCreate {
		repo.AutoInit = github.Ptr(model.AutoInit.ValueBool())
		repo.GitignoreTemplate = github.Ptr(model.GitignoreTemplate.ValueString())
		repo.LicenseTemplate = github.Ptr(model.LicenseTemplate.ValueString())
	}

	if model.AllowSquashMerge.ValueBool() {
		repo.SquashMergeCommitTitle = github.Ptr(model.SquashMergeCommitTitle.ValueString())
		repo.SquashMergeCommitMessage = github.Ptr(model.SquashMergeCommitMessage.ValueString())
	}

	if model.AllowMergeCommit.ValueBool() {
		repo.MergeCommitTitle = github.Ptr(model.MergeCommitTitle.ValueString())
		repo.MergeCommitMessage = github.Ptr(model.MergeCommitMessage.ValueString())
	}

	return repo
}

func flattenRepository(model *GitHubRepositoryResourceModel, repo *github.Repository) {
	// IDs
	model.ID = types.Int64Value(repo.GetID())
	model.NodeID = types.StringValue(repo.GetNodeID())
	// Arguments
	model.Name = types.StringValue(repo.GetName())
	model.Description = types.StringValue(repo.GetDescription())
	model.Homepage = types.StringValue(repo.GetHomepage())
	model.Private = types.BoolValue(repo.GetPrivate())
	model.HasIssues = types.BoolValue(repo.GetHasIssues())
	model.HasProjects = types.BoolValue(repo.GetHasProjects())
	model.HasWiki = types.BoolValue(repo.GetHasWiki())
	model.HasDiscussions = types.BoolValue(repo.GetHasDiscussions())
	model.AllowSquashMerge = types.BoolValue(repo.GetAllowSquashMerge())
	model.AllowMergeCommit = types.BoolValue(repo.GetAllowMergeCommit())
	model.AllowRebaseMerge = types.BoolValue(repo.GetAllowRebaseMerge())
	model.AllowAutoMerge = types.BoolValue(repo.GetAllowAutoMerge())
	model.DeleteBranchOnMerge = types.BoolValue(repo.GetDeleteBranchOnMerge())
	model.SquashMergeCommitTitle = types.StringValue(repo.GetSquashMergeCommitTitle())
	model.SquashMergeCommitMessage = types.StringValue(repo.GetSquashMergeCommitMessage())
	model.MergeCommitTitle = types.StringValue(repo.GetMergeCommitTitle())
	model.MergeCommitMessage = types.StringValue(repo.GetMergeCommitMessage())
	model.IsTemplate = types.BoolValue(repo.GetIsTemplate())
	// Attributes
	model.Private = types.BoolValue(repo.GetPrivate())
	model.HasIssues = types.BoolValue(repo.GetHasIssues())
	model.HasProjects = types.BoolValue(repo.GetHasProjects())
	model.HasWiki = types.BoolValue(repo.GetHasWiki())
	model.HasDiscussions = types.BoolValue(repo.GetHasDiscussions())
	model.AllowSquashMerge = types.BoolValue(repo.GetAllowSquashMerge())
	model.AllowMergeCommit = types.BoolValue(repo.GetAllowMergeCommit())
	model.AllowRebaseMerge = types.BoolValue(repo.GetAllowRebaseMerge())
	model.AllowAutoMerge = types.BoolValue(repo.GetAllowAutoMerge())
	model.DeleteBranchOnMerge = types.BoolValue(repo.GetDeleteBranchOnMerge())
	model.SquashMergeCommitTitle = types.StringValue(repo.GetSquashMergeCommitTitle())
	model.SquashMergeCommitMessage = types.StringValue(repo.GetSquashMergeCommitMessage())
	model.MergeCommitTitle = types.StringValue(repo.GetMergeCommitTitle())
	model.MergeCommitMessage = types.StringValue(repo.GetMergeCommitMessage())
	model.IsTemplate = types.BoolValue(repo.GetIsTemplate())
}

// Resource Definition

func (r *GitHubRepositoryResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_repository"
}

func (r *GitHubRepositoryResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
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
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"homepage": schema.StringAttribute{
				Description:         "The homepage of the repository.",
				MarkdownDescription: "The homepage of the repository.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"private": schema.BoolAttribute{
				Description:         "Indicates if the repository is private.",
				MarkdownDescription: "Indicates if the repository is private.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"has_issues": schema.BoolAttribute{
				Description:         "Indicates if the repository has issues enabled.",
				MarkdownDescription: "Indicates if the repository has issues enabled.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"has_projects": schema.BoolAttribute{
				Description:         "Indicates if the repository has projects enabled.",
				MarkdownDescription: "Indicates if the repository has projects enabled.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"has_wiki": schema.BoolAttribute{
				Description:         "Indicates if the repository has wiki enabled.",
				MarkdownDescription: "Indicates if the repository has wiki enabled.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"has_discussions": schema.BoolAttribute{
				Description:         "Indicates if the repository has discussions enabled.",
				MarkdownDescription: "Indicates if the repository has discussions enabled.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
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
				Description:         "Indicates if squash merging is allowed in the repository. Defaults to 'true'.",
				MarkdownDescription: "Indicates if squash merging is allowed in the repository. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"allow_merge_commit": schema.BoolAttribute{
				Description:         "Indicates if merge commits are allowed in the repository. Defaults to `true`.",
				MarkdownDescription: "Indicates if merge commits are allowed in the repository. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"allow_rebase_merge": schema.BoolAttribute{
				Description:         "Indicates if rebase merging is allowed in the repository. Defaults to `true`.",
				MarkdownDescription: "Indicates if rebase merging is allowed in the repository. Defaults to `true`.",
				Optional:            true,
				Computed:            true,
				Default:             booldefault.StaticBool(true),
			},
			"allow_auto_merge": schema.BoolAttribute{
				Description:         "Indicates if auto-merging is allowed in the repository.",
				MarkdownDescription: "Indicates if auto-merging is allowed in the repository.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"delete_branch_on_merge": schema.BoolAttribute{
				Description:         "Indicates if branches are automatically deleted when pull requests are merged.",
				MarkdownDescription: "Indicates if branches are automatically deleted when pull requests are merged.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"squash_merge_commit_title": schema.StringAttribute{
				Description: `The default value for a squash merge commit title.

	Must be one of:

	'PR_TITLE' defaults to the pull request's title.

	'COMMIT_OR_PR_TITLE' defaults to the commit's title (if only one commit) or the pull request's title (when more than one commit).`,
				MarkdownDescription: `The default value for a squash merge commit title.

	Must be one of:

	` + "`" + `PR_TITLE` + "`" + ` defaults to the pull request's title.

	` + "`" + `COMMIT_OR_PR_TITLE` + "`" + ` defaults to the commit's title (if only one commit) or the pull request's title (when more than one commit).`,
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("COMMIT_OR_PR_TITLE"),
				Validators: []validator.String{
					stringvalidator.OneOf("PR_TITLE", "COMMIT_OR_PR_TITLE"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"squash_merge_commit_message": schema.StringAttribute{
				Description: `The default value for a squash merge commit message.

	Must be one of:

	'PR_BODY' defaults to the pull request's body.

	'COMMIT_MESSAGES' defaults to the branch's commit messages.

	'BLANK' defaults to a blank commit message.`,
				MarkdownDescription: `The default value for a squash merge commit message.

	Must be one of:

	` + "`" + `PR_BODY` + "`" + ` defaults to the pull request's body.

	` + "`" + `COMMIT_MESSAGES` + "`" + ` defaults to the branch's commit messages.

	` + "`" + `BLANK` + "`" + ` defaults to a blank commit message.`,
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("COMMIT_MESSAGES"),
				Validators: []validator.String{
					stringvalidator.OneOf("PR_BODY", "COMMIT_MESSAGES", "BLANK"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"merge_commit_title": schema.StringAttribute{
				Description: `The default value for a merge commit title.

	Must be one of:

	'PR_TITLE' defaults to the pull request's title.

	'MERGE_MESSAGE' defaults to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).`,
				MarkdownDescription: `The default value for a merge commit title.

	Must be one of:

	` + "`" + `PR_TITLE` + "`" + ` defaults to the pull request's title.

	` + "`" + `MERGE_MESSAGE` + "`" + ` defaults to the classic title for a merge message (e.g., Merge pull request #123 from branch-name).`,
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("MERGE_MESSAGE"),
				Validators: []validator.String{
					stringvalidator.OneOf("PR_TITLE", "MERGE_MESSAGE"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"merge_commit_message": schema.StringAttribute{
				Description: `The default value for a merge commit message.

	Must be one of:

	'PR_BODY' defaults to the pull request's body.

	'PR_TITLE' defaults to the pull request's body.

	'BLANK' defaults to a blank commit message.`,
				MarkdownDescription: `The default value for a merge commit message.

	Must be one of:

	` + "`" + `PR_BODY` + "`" + ` defaults to the pull request's body.

	` + "`" + `PR_TITLE` + "`" + ` defaults to the pull request's body.

	` + "`" + `BLANK` + "`" + ` defaults to a blank commit message.`,
				Optional: true,
				Computed: true,
				Default:  stringdefault.StaticString("PR_TITLE"),
				Validators: []validator.String{
					stringvalidator.OneOf("PR_BODY", "PR_TITLE", "BLANK"),
				},
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"is_template": schema.BoolAttribute{
				Description:         "Indicates if the repository is a template repository.",
				MarkdownDescription: "Indicates if the repository is a template repository.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			// Template Arguments
			"template_repository": schema.StringAttribute{
				Description:         "The name of the template repository to use.",
				MarkdownDescription: "The name of the template repository to use.",
				Optional:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
			},
			"template_owner": schema.StringAttribute{
				Description:         "The owner of the template repository.",
				MarkdownDescription: "The owner of the template repository.",
				Optional:            true,
				Computed:            true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
				},
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

func (r *GitHubRepositoryResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	config, ok := req.ProviderData.(*GitHubClientConfiguration)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Type from ProviderData",
			fmt.Sprintf("Expected *common.GitHubProviderConfiguration, got: %T", req.ProviderData),
		)
		return
	}

	r.client = config.Client
	r.owner = config.Owner
	r.organization = config.Organization
}

// Resource Lifecycle

func (r *GitHubRepositoryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var model GitHubRepositoryResourceModel

	client := r.client

	// Read Terraform State
	resp.Diagnostics.Append(req.State.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	repo, _, err := client.Repositories.GetByID(ctx, model.ID.ValueInt64())
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Communicating with the GitHub API",
			fmt.Sprintf("Unable to get repository, got error: %s", err),
		)
		return
	}

	flattenRepository(&model, repo)

	// Save updated data into Terraform state.
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *GitHubRepositoryResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var model GitHubRepositoryResourceModel

	// Read Plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	client := r.client
	owner := r.owner
	organization := r.organization

	var repo *github.Repository
	var err error

	if !model.TemplateRepository.IsNull() {
		// Create from a Template
		templateRepo := model.TemplateRepository.ValueString()
		templateOwner := model.TemplateOwner.ValueString()

		// Use the Provider Configured `owner` for the API Call
		if templateOwner == "" {
			templateOwner = owner
		}

		templateReq := &github.TemplateRepoRequest{
			Name:        github.Ptr(model.Name.ValueString()),
			Owner:       github.Ptr(owner), // The owner (org) where the new repository will be created.
			Description: github.Ptr(model.Description.ValueString()),
			Private:     github.Ptr(model.Private.ValueBool()),
		}

		_, _, err = client.Repositories.CreateFromTemplate(ctx, templateOwner, templateRepo, templateReq)
		if err != nil {
			resp.Diagnostics.AddError("Error Creating Repository from Template", err.Error())
			return
		}

		repository := expandRepository(model, expandForUpdate)
		repo, _, err = client.Repositories.Edit(ctx, owner, model.Name.ValueString(), repository)
		if err != nil {
			resp.Diagnostics.AddError("Error Applying Settings after Template Creation", err.Error())
			return
		}
	} else {
		// Standard Creation
		repository := expandRepository(model, expandForCreate)
		repo, _, err = client.Repositories.Create(ctx, organization, repository)
		if err != nil {
			resp.Diagnostics.AddError("Error Creating Repository", err.Error())
			return
		}
	}

	flattenRepository(&model, repo)

	if model.TemplateRepository.IsNull() {
		// If not using a template, ensure these are null in state.
		model.TemplateRepository = types.StringNull()
		model.TemplateOwner = types.StringNull()
	} else {
		// If using a template, ensure TemplateOwner is set.
		if model.TemplateOwner.IsUnknown() || model.TemplateOwner.IsNull() {
			// If owner is unknown at plan time, default to 'owner'.
			model.TemplateOwner = types.StringValue(owner)
		}
	}

	// Save updated data into Terraform state.
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *GitHubRepositoryResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var model GitHubRepositoryResourceModel
	var state GitHubRepositoryResourceModel

	client := r.client
	owner := r.owner

	// Read Plan
	resp.Diagnostics.Append(req.Plan.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read Terraform State
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	repository := expandRepository(model, expandForUpdate)
	repo, _, err := client.Repositories.Edit(ctx, owner, state.Name.ValueString(), repository)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Communicating with the GitHub API",
			fmt.Sprintf("Unable to update the repository, got error: %s", err),
		)
		return
	}

	flattenRepository(&model, repo)

	// Save updated data into Terraform state.
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}

func (r *GitHubRepositoryResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var model GitHubRepositoryResourceModel

	client := r.client
	owner := r.owner

	// Read Terraform State
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

func (r *GitHubRepositoryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error importing item",
			"Could not import the repository, the ID should be an integer: "+err.Error(),
		)
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}
