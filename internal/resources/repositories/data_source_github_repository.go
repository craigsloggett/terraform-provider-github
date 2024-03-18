package repositories

import (
	"context"
	"fmt"
	"time"

	"github.com/google/go-github/v60/github"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &GitHubRepository{}

type GitHubRepository struct {
	client any
}

type GitHubRepositoryModel struct {
	Owner                     types.String      `tfsdk:"owner"`
	Repo                      types.String      `tfsdk:"repo"`
	Id                        types.Int64       `tfsdk:"id"`
	NodeID                    types.String      `tfsdk:"node_id"`
	Name                      types.String      `tfsdk:"name"`
	FullName                  types.String      `tfsdk:"full_name"`
	Description               types.String      `tfsdk:"description"`
	Homepage                  types.String      `tfsdk:"homepage"`
	DefaultBranch             types.String      `tfsdk:"default_branch"`
	MasterBranch              types.String      `tfsdk:"master_branch"`
	CreatedAt                 timetypes.RFC3339 `tfsdk:"created_at"`
	PushedAt                  timetypes.RFC3339 `tfsdk:"pushed_at"`
	UpdatedAt                 timetypes.RFC3339 `tfsdk:"updated_at"`
	HTMLURL                   types.String      `tfsdk:"html_url"`
	CloneURL                  types.String      `tfsdk:"clone_url"`
	GitURL                    types.String      `tfsdk:"git_url"`
	MirrorURL                 types.String      `tfsdk:"mirror_url"`
	SSHURL                    types.String      `tfsdk:"ssh_url"`
	SVNURL                    types.String      `tfsdk:"svn_url"`
	Language                  types.String      `tfsdk:"language"`
	Fork                      types.Bool        `tfsdk:"fork"`
	ForksCount                types.Int64       `tfsdk:"forks_count"`
	NetworkCount              types.Int64       `tfsdk:"network_count"`
	OpenIssuesCount           types.Int64       `tfsdk:"open_issues_count"`
	StargazersCount           types.Int64       `tfsdk:"stargazers_count"`
	SubscribersCount          types.Int64       `tfsdk:"subscribers_count"`
	Size                      types.Int64       `tfsdk:"size"`
	AutoInit                  types.Bool        `tfsdk:"auto_init"`
	AllowRebaseMerge          types.Bool        `tfsdk:"allow_rebase_merge"`
	AllowUpdateBranch         types.Bool        `tfsdk:"allow_update_branch"`
	AllowSquashMerge          types.Bool        `tfsdk:"allow_squash_merge"`
	AllowMergeCommit          types.Bool        `tfsdk:"allow_merge_commit"`
	AllowAutoMerge            types.Bool        `tfsdk:"allow_auto_merge"`
	AllowForking              types.Bool        `tfsdk:"allow_forking"`
	WebCommitSignoffRequired  types.Bool        `tfsdk:"web_commit_signoff_required"`
	DeleteBranchOnMerge       types.Bool        `tfsdk:"delete_branch_on_merge"`
	UseSquashPRTitleAsDefault types.Bool        `tfsdk:"use_squash_pr_title_as_default"`
	SquashMergeCommitTitle    types.String      `tfsdk:"squash_merge_commit_title"`
	SquashMergeCommitMessage  types.String      `tfsdk:"squash_merge_commit_message"`
	MergeCommitTitle          types.String      `tfsdk:"merge_commit_title"`
	MergeCommitMessage        types.String      `tfsdk:"merge_commit_message"`
	Archived                  types.Bool        `tfsdk:"archived"`
	Disabled                  types.Bool        `tfsdk:"disabled"`
}

func NewGitHubRepository() datasource.DataSource {
	return &GitHubRepository{}
}

func TimeToFramework(_ context.Context, v *time.Time) timetypes.RFC3339 {
	return timetypes.NewRFC3339TimePointerValue(v)
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
			"node_id": schema.StringAttribute{
				Description:         "The node ID of the repository.",
				MarkdownDescription: "The node ID of the repository.",
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
			"description": schema.StringAttribute{
				Description:         "The description of the repository.",
				MarkdownDescription: "The description of the repository.",
				Computed:            true,
			},
			"homepage": schema.StringAttribute{
				Description:         "The homepage of the repository.",
				MarkdownDescription: "The homepage of the repository.",
				Computed:            true,
			},
			"default_branch": schema.StringAttribute{
				Description:         "The repository's default branch.",
				MarkdownDescription: "The repository's default branch.",
				Computed:            true,
			},
			"master_branch": schema.StringAttribute{
				Description:         "The repository's master branch.",
				MarkdownDescription: "The repository's master branch.",
				Computed:            true,
			},
			"created_at": schema.StringAttribute{
				CustomType:          timetypes.RFC3339Type{},
				Description:         "The timestamp of when the repository was created on GitHub.",
				MarkdownDescription: "The timestamp of when the repository was created on GitHub.",
				Computed:            true,
			},
			"pushed_at": schema.StringAttribute{
				CustomType:          timetypes.RFC3339Type{},
				Description:         "The timestamp of the last push to the repository.",
				MarkdownDescription: "The timestamp of the last push to the repository.",
				Computed:            true,
			},
			"updated_at": schema.StringAttribute{
				CustomType:          timetypes.RFC3339Type{},
				Description:         "The timestamp of when the repository was last updated.",
				MarkdownDescription: "The timestamp of when the repository was last updated.",
				Computed:            true,
			},
			"html_url": schema.StringAttribute{
				Description:         "The HTML URL of the repository.",
				MarkdownDescription: "The HTML URL of the repository.",
				Computed:            true,
			},
			"clone_url": schema.StringAttribute{
				Description:         "The URL used for cloning the repository.",
				MarkdownDescription: "The URL used for cloning the repository.",
				Computed:            true,
			},
			"git_url": schema.StringAttribute{
				Description:         "The git URL of the repository.",
				MarkdownDescription: "The git URL of the repository.",
				Computed:            true,
			},
			"mirror_url": schema.StringAttribute{
				Description:         "The mirror URL of the repository.",
				MarkdownDescription: "The mirror URL of the repository.",
				Computed:            true,
			},
			"ssh_url": schema.StringAttribute{
				Description:         "The SSH URL of the repository.",
				MarkdownDescription: "The SSH URL of the repository.",
				Computed:            true,
			},
			"svn_url": schema.StringAttribute{
				Description:         "The SVN URL of the repository.",
				MarkdownDescription: "The SVN URL of the repository.",
				Computed:            true,
			},
			"language": schema.StringAttribute{
				Description:         "The primary language of the repository.",
				MarkdownDescription: "The primary language of the repository.",
				Computed:            true,
			},
			"fork": schema.BoolAttribute{
				Description:         "Indicates if the repository is a fork.",
				MarkdownDescription: "Indicates if the repository is a fork.",
				Computed:            true,
			},
			"forks_count": schema.Int64Attribute{
				Description:         "The number of forks of the repository.",
				MarkdownDescription: "The number of forks of the repository.",
				Computed:            true,
			},
			"network_count": schema.Int64Attribute{
				Description:         "The number of repositories in the network.",
				MarkdownDescription: "The number of repositories in the network.",
				Computed:            true,
			},
			"open_issues_count": schema.Int64Attribute{
				Description:         "The number of open issues in the repository.",
				MarkdownDescription: "The number of open issues in the repository.",
				Computed:            true,
			},
			"stargazers_count": schema.Int64Attribute{
				Description:         "The number of users who have starred the repository.",
				MarkdownDescription: "The number of users who have starred the repository.",
				Computed:            true,
			},
			"subscribers_count": schema.Int64Attribute{
				Description:         "The number of users subscribed to the repository.",
				MarkdownDescription: "The number of users subscribed to the repository.",
				Computed:            true,
			},
			"size": schema.Int64Attribute{
				Description:         "The size of the repository, in kilobytes.",
				MarkdownDescription: "The size of the repository, in kilobytes.",
				Computed:            true,
			},
			"auto_init": schema.BoolAttribute{
				Description:         "Indicates if the repository is initialized with a README.",
				MarkdownDescription: "Indicates if the repository is initialized with a README.",
				Computed:            true,
			},
			"allow_rebase_merge": schema.BoolAttribute{
				Description:         "Indicates if rebase merging is allowed in the repository.",
				MarkdownDescription: "Indicates if rebase merging is allowed in the repository.",
				Computed:            true,
			},
			"allow_update_branch": schema.BoolAttribute{
				Description:         "Indicates if updating a pull request head branch that is behind its base branch is allowed.",
				MarkdownDescription: "Indicates if updating a pull request head branch that is behind its base branch is allowed.",
				Computed:            true,
			},
			"allow_squash_merge": schema.BoolAttribute{
				Description:         "Indicates if squash merging is allowed in the repository.",
				MarkdownDescription: "Indicates if squash merging is allowed in the repository.",
				Computed:            true,
			},
			"allow_merge_commit": schema.BoolAttribute{
				Description:         "Indicates if merge commits are allowed in the repository.",
				MarkdownDescription: "Indicates if merge commits are allowed in the repository.",
				Computed:            true,
			},
			"allow_auto_merge": schema.BoolAttribute{
				Description:         "Indicates if auto-merging is allowed in the repository.",
				MarkdownDescription: "Indicates if auto-merging is allowed in the repository.",
				Computed:            true,
			},
			"allow_forking": schema.BoolAttribute{
				Description:         "Indicates if forking is allowed for the repository.",
				MarkdownDescription: "Indicates if forking is allowed for the repository.",
				Computed:            true,
			},
			"web_commit_signoff_required": schema.BoolAttribute{
				Description:         "Indicates if commit signoff is required for web-based commits.",
				MarkdownDescription: "Indicates if commit signoff is required for web-based commits.",
				Computed:            true,
			},
			"delete_branch_on_merge": schema.BoolAttribute{
				Description:         "Indicates if branches are automatically deleted when pull requests are merged.",
				MarkdownDescription: "Indicates if branches are automatically deleted when pull requests are merged.",
				Computed:            true,
			},
			"use_squash_pr_title_as_default": schema.BoolAttribute{
				Description:         "Indicates if the squash PR title is used as the default commit message.",
				MarkdownDescription: "Indicates if the squash PR title is used as the default commit message.",
				Computed:            true,
			},
			"squash_merge_commit_title": schema.StringAttribute{
				Description:         "The title of squash merge commits for pull requests.",
				MarkdownDescription: "The title of squash merge commits for pull requests.",
				Computed:            true,
			},
			"squash_merge_commit_message": schema.StringAttribute{
				Description:         "The message of squash merge commits for pull requests.",
				MarkdownDescription: "The message of squash merge commits for pull requests.",
				Computed:            true,
			},
			"merge_commit_title": schema.StringAttribute{
				Description:         "The title of merge commits for pull requests.",
				MarkdownDescription: "The title of merge commits for pull requests.",
				Computed:            true,
			},
			"merge_commit_message": schema.StringAttribute{
				Description:         "The message of merge commits for pull requests.",
				MarkdownDescription: "The message of merge commits for pull requests.",
				Computed:            true,
			},
			"archived": schema.BoolAttribute{
				Description:         "Indicates if the repository is archived.",
				MarkdownDescription: "Indicates if the repository is archived.",
				Computed:            true,
			},
			"disabled": schema.BoolAttribute{
				Description:         "Indicates if the repository is disabled.",
				MarkdownDescription: "Indicates if the repository is disabled.",
				Computed:            true,
			},
		},
		Description:         "Use this data source to retrieve a list of GitHub repositories.",
		MarkdownDescription: "Use this data source to retrieve a list of GitHub repositories.",
	}
}

func (d *GitHubRepository) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.client = client
}

func (d *GitHubRepository) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var model GitHubRepositoryModel

	client, ok := d.client.(*github.Client)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Type from GitHubProvider's Client Field",
			fmt.Sprintf("Expected *github.Client, got: %T", d.client),
		)
		return
	}

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

	model.NodeID = types.StringValue(repo.GetNodeID())
	model.Name = types.StringValue(repo.GetName())
	model.FullName = types.StringValue(repo.GetFullName())
	model.Description = types.StringValue(repo.GetDescription())
	model.Homepage = types.StringValue(repo.GetHomepage())
	model.DefaultBranch = types.StringValue(repo.GetDefaultBranch())
	model.MasterBranch = types.StringValue(repo.GetMasterBranch())

	createdAt := repo.GetCreatedAt()
	pushedAt := repo.GetPushedAt()
	updatedAt := repo.GetUpdatedAt()
	model.CreatedAt = timetypes.NewRFC3339TimePointerValue(createdAt.GetTime())
	model.PushedAt = timetypes.NewRFC3339TimePointerValue(pushedAt.GetTime())
	model.UpdatedAt = timetypes.NewRFC3339TimePointerValue(updatedAt.GetTime())

	model.HTMLURL = types.StringValue(repo.GetHTMLURL())
	model.CloneURL = types.StringValue(repo.GetCloneURL())
	model.GitURL = types.StringValue(repo.GetGitURL())
	model.MirrorURL = types.StringValue(repo.GetMirrorURL())
	model.SSHURL = types.StringValue(repo.GetSSHURL())
	model.SVNURL = types.StringValue(repo.GetSVNURL())
	model.Language = types.StringValue(repo.GetLanguage())

	model.Fork = types.BoolValue(repo.GetFork())

	model.ForksCount = types.Int64Value(int64(repo.GetForksCount()))
	model.NetworkCount = types.Int64Value(int64(repo.GetNetworkCount()))
	model.OpenIssuesCount = types.Int64Value(int64(repo.GetOpenIssuesCount()))
	model.StargazersCount = types.Int64Value(int64(repo.GetStargazersCount()))
	model.SubscribersCount = types.Int64Value(int64(repo.GetSubscribersCount()))
	model.Size = types.Int64Value(int64(repo.GetSize()))

	model.AutoInit = types.BoolValue(repo.GetAutoInit())
	model.AllowRebaseMerge = types.BoolValue(repo.GetAllowRebaseMerge())
	model.AllowUpdateBranch = types.BoolValue(repo.GetAllowUpdateBranch())
	model.AllowSquashMerge = types.BoolValue(repo.GetAllowSquashMerge())
	model.AllowMergeCommit = types.BoolValue(repo.GetAllowMergeCommit())
	model.AllowAutoMerge = types.BoolValue(repo.GetAllowAutoMerge())
	model.AllowForking = types.BoolValue(repo.GetAllowForking())
	model.WebCommitSignoffRequired = types.BoolValue(repo.GetWebCommitSignoffRequired())
	model.DeleteBranchOnMerge = types.BoolValue(repo.GetDeleteBranchOnMerge())
	model.UseSquashPRTitleAsDefault = types.BoolValue(repo.GetUseSquashPRTitleAsDefault())

	model.SquashMergeCommitTitle = types.StringValue(repo.GetSquashMergeCommitTitle())
	model.SquashMergeCommitMessage = types.StringValue(repo.GetSquashMergeCommitMessage())
	model.MergeCommitTitle = types.StringValue(repo.GetMergeCommitTitle())
	model.MergeCommitMessage = types.StringValue(repo.GetMergeCommitMessage())

	model.Archived = types.BoolValue(repo.GetArchived())
	model.Disabled = types.BoolValue(repo.GetDisabled())

	resp.State.Set(ctx, &model)
}
