package provider

import (
	"context"
	"fmt"

	"github.com/google/go-github/v60/github"
	"github.com/hashicorp/terraform-plugin-framework-timetypes/timetypes"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = &GitHubRepositoryDataSource{}

type GitHubRepositoryDataSource struct {
	client *github.Client
}

func NewGitHubRepositoryDataSource() datasource.DataSource {
	return &GitHubRepositoryDataSource{}
}

type GitHubRepositoryDataSourceModel struct {
	// Arguments
	Owner types.String `tfsdk:"owner"`
	Name  types.String `tfsdk:"name"`

	// Attributes
	ID                       types.Int64               `tfsdk:"id"`
	NodeID                   types.String              `tfsdk:"node_id"`
	FullName                 types.String              `tfsdk:"full_name"`
	Description              types.String              `tfsdk:"description"`
	Homepage                 types.String              `tfsdk:"homepage"`
	DefaultBranch            types.String              `tfsdk:"default_branch"`
	MasterBranch             types.String              `tfsdk:"master_branch"`
	CreatedAt                timetypes.RFC3339         `tfsdk:"created_at"`
	PushedAt                 timetypes.RFC3339         `tfsdk:"pushed_at"`
	UpdatedAt                timetypes.RFC3339         `tfsdk:"updated_at"`
	HTMLURL                  types.String              `tfsdk:"html_url"`
	CloneURL                 types.String              `tfsdk:"clone_url"`
	GitURL                   types.String              `tfsdk:"git_url"`
	MirrorURL                types.String              `tfsdk:"mirror_url"`
	SSHURL                   types.String              `tfsdk:"ssh_url"`
	SVNURL                   types.String              `tfsdk:"svn_url"`
	Language                 types.String              `tfsdk:"language"`
	Fork                     types.Bool                `tfsdk:"fork"`
	ForksCount               types.Int64               `tfsdk:"forks_count"`
	NetworkCount             types.Int64               `tfsdk:"network_count"`
	OpenIssuesCount          types.Int64               `tfsdk:"open_issues_count"`
	StargazersCount          types.Int64               `tfsdk:"stargazers_count"`
	SubscribersCount         types.Int64               `tfsdk:"subscribers_count"`
	Size                     types.Int64               `tfsdk:"size"`
	AutoInit                 types.Bool                `tfsdk:"auto_init"`
	Parent                   *linkedRepositoryModel    `tfsdk:"parent"`
	Source                   *linkedRepositoryModel    `tfsdk:"source"`
	TemplateRepository       *linkedRepositoryModel    `tfsdk:"template_repository"`
	Organization             *organizationModel        `tfsdk:"organization"`
	Permissions              *permissionsModel         `tfsdk:"permissions"`
	AllowRebaseMerge         types.Bool                `tfsdk:"allow_rebase_merge"`
	AllowUpdateBranch        types.Bool                `tfsdk:"allow_update_branch"`
	AllowSquashMerge         types.Bool                `tfsdk:"allow_squash_merge"`
	AllowMergeCommit         types.Bool                `tfsdk:"allow_merge_commit"`
	AllowAutoMerge           types.Bool                `tfsdk:"allow_auto_merge"`
	AllowForking             types.Bool                `tfsdk:"allow_forking"`
	WebCommitSignoffRequired types.Bool                `tfsdk:"web_commit_signoff_required"`
	DeleteBranchOnMerge      types.Bool                `tfsdk:"delete_branch_on_merge"`
	SquashMergeCommitTitle   types.String              `tfsdk:"squash_merge_commit_title"`
	SquashMergeCommitMessage types.String              `tfsdk:"squash_merge_commit_message"`
	MergeCommitTitle         types.String              `tfsdk:"merge_commit_title"`
	MergeCommitMessage       types.String              `tfsdk:"merge_commit_message"`
	Topics                   types.List                `tfsdk:"topics"`
	Archived                 types.Bool                `tfsdk:"archived"`
	Disabled                 types.Bool                `tfsdk:"disabled"`
	Private                  types.Bool                `tfsdk:"private"`
	HasIssues                types.Bool                `tfsdk:"has_issues"`
	HasWiki                  types.Bool                `tfsdk:"has_wiki"`
	HasPages                 types.Bool                `tfsdk:"has_pages"`
	HasProjects              types.Bool                `tfsdk:"has_projects"`
	HasDownloads             types.Bool                `tfsdk:"has_downloads"`
	HasDiscussions           types.Bool                `tfsdk:"has_discussions"`
	IsTemplate               types.Bool                `tfsdk:"is_template"`
	LicenseTemplate          types.String              `tfsdk:"license_template"`
	GitignoreTemplate        types.String              `tfsdk:"gitignore_template"`
	SecurityAndAnalysis      *securityAndAnalysisModel `tfsdk:"security_and_analysis"`
	TeamID                   types.Int64               `tfsdk:"team_id"`
	URL                      types.String              `tfsdk:"url"`
	ArchiveURL               types.String              `tfsdk:"archive_url"`
	AssigneesURL             types.String              `tfsdk:"assignees_url"`
	BlobsURL                 types.String              `tfsdk:"blobs_url"`
	BranchesURL              types.String              `tfsdk:"branches_url"`
	CollaboratorsURL         types.String              `tfsdk:"collaborators_url"`
	CommentsURL              types.String              `tfsdk:"comments_url"`
	CommitsURL               types.String              `tfsdk:"commits_url"`
	CompareURL               types.String              `tfsdk:"compare_url"`
	ContentsURL              types.String              `tfsdk:"contents_url"`
	ContributorsURL          types.String              `tfsdk:"contributors_url"`
	DeploymentsURL           types.String              `tfsdk:"deployments_url"`
	DownloadsURL             types.String              `tfsdk:"downloads_url"`
	EventsURL                types.String              `tfsdk:"events_url"`
	ForksURL                 types.String              `tfsdk:"forks_url"`
	GitCommitsURL            types.String              `tfsdk:"git_commits_url"`
	GitRefsURL               types.String              `tfsdk:"git_refs_url"`
	GitTagsURL               types.String              `tfsdk:"git_tags_url"`
	HooksURL                 types.String              `tfsdk:"hooks_url"`
	IssueCommentURL          types.String              `tfsdk:"issue_comment_url"`
	IssueEventsURL           types.String              `tfsdk:"issue_events_url"`
	IssuesURL                types.String              `tfsdk:"issues_url"`
	KeysURL                  types.String              `tfsdk:"keys_url"`
	LabelsURL                types.String              `tfsdk:"labels_url"`
	LanguagesURL             types.String              `tfsdk:"languages_url"`
	MergesURL                types.String              `tfsdk:"merges_url"`
	MilestonesURL            types.String              `tfsdk:"milestones_url"`
	NotificationsURL         types.String              `tfsdk:"notifications_url"`
	PullsURL                 types.String              `tfsdk:"pulls_url"`
	ReleasesURL              types.String              `tfsdk:"releases_url"`
	StargazersURL            types.String              `tfsdk:"stargazers_url"`
	StatusesURL              types.String              `tfsdk:"statuses_url"`
	SubscribersURL           types.String              `tfsdk:"subscribers_url"`
	SubscriptionURL          types.String              `tfsdk:"subscription_url"`
	TagsURL                  types.String              `tfsdk:"tags_url"`
	TreesURL                 types.String              `tfsdk:"trees_url"`
	TeamsURL                 types.String              `tfsdk:"teams_url"`
	Visibility               types.String              `tfsdk:"visibility"`
}

type linkedRepositoryModel struct {
	Owner    types.String `tfsdk:"owner"`
	Name     types.String `tfsdk:"name"`
	FullName types.String `tfsdk:"full_name"`
}

type organizationModel struct {
	Name types.String `tfsdk:"name"`
}

type permissionsModel struct {
	Admin    types.Bool `tfsdk:"admin"`
	Pull     types.Bool `tfsdk:"pull"`
	Triage   types.Bool `tfsdk:"triage"`
	Push     types.Bool `tfsdk:"push"`
	Maintain types.Bool `tfsdk:"maintain"`
}

type securityAndAnalysisModel struct {
	AdvancedSecurity             *advancedSecurityModel             `tfsdk:"advanced_security"`
	SecretScanning               *secretScanningModel               `tfsdk:"secret_scanning"`
	SecretScanningPushProtection *secretScanningPushProtectionModel `tfsdk:"secret_scanning_push_protection"`
	SecretScanningValidityChecks *secretScanningValidityChecksModel `tfsdk:"secret_scanning_validity_checks"`
	DependabotSecurityUpdates    *dependabotSecurityUpdatesModel    `tfsdk:"dependabot_security_updates"`
}

type advancedSecurityModel struct {
	Status types.String `tfsdk:"status"`
}

type secretScanningModel struct {
	Status types.String `tfsdk:"status"`
}

type secretScanningPushProtectionModel struct {
	Status types.String `tfsdk:"status"`
}

type secretScanningValidityChecksModel struct {
	Status types.String `tfsdk:"status"`
}

type dependabotSecurityUpdatesModel struct {
	Status types.String `tfsdk:"status"`
}

func (d *GitHubRepositoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_repository"
}

func (d *GitHubRepositoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"owner": schema.StringAttribute{
				Description:         "The owner of the repository.",
				MarkdownDescription: "The owner of the repository.",
				Required:            true,
			},
			"name": schema.StringAttribute{
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
			"parent": schema.SingleNestedAttribute{
				Description:         "Details of the parent repository.",
				MarkdownDescription: "Details of the parent repository.",
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"owner": schema.StringAttribute{
						Description:         "The owner of the parent repository.",
						MarkdownDescription: "The owner of the parent repository.",
						Computed:            true,
					},
					"name": schema.StringAttribute{
						Description:         "The name of the parent repository.",
						MarkdownDescription: "The name of the parent repository.",
						Computed:            true,
					},
					"full_name": schema.StringAttribute{
						Description:         "The full name of the parent repository.",
						MarkdownDescription: "The full name of the parent repository.",
						Computed:            true,
					},
				},
			},
			"source": schema.SingleNestedAttribute{
				Description:         "Details of the source repository.",
				MarkdownDescription: "Details of the source repository.",
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"owner": schema.StringAttribute{
						Description:         "The owner of the source repository.",
						MarkdownDescription: "The owner of the source repository.",
						Computed:            true,
					},
					"name": schema.StringAttribute{
						Description:         "The name of the source repository.",
						MarkdownDescription: "The name of the source repository.",
						Computed:            true,
					},
					"full_name": schema.StringAttribute{
						Description:         "The full name of the source repository.",
						MarkdownDescription: "The full name of the source repository.",
						Computed:            true,
					},
				},
			},
			"template_repository": schema.SingleNestedAttribute{
				Description:         "Details of the template repository.",
				MarkdownDescription: "Details of the template repository.",
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"owner": schema.StringAttribute{
						Description:         "The owner of the template repository.",
						MarkdownDescription: "The owner of the template repository.",
						Computed:            true,
					},
					"name": schema.StringAttribute{
						Description:         "The name of the template repository.",
						MarkdownDescription: "The name of the template repository.",
						Computed:            true,
					},
					"full_name": schema.StringAttribute{
						Description:         "The full name of the template repository.",
						MarkdownDescription: "The full name of the template repository.",
						Computed:            true,
					},
				},
			},
			"organization": schema.SingleNestedAttribute{
				Description:         "Details of the organization the repository is a part of.",
				MarkdownDescription: "Details of the organization the repository is a part of.",
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"name": schema.StringAttribute{
						Description:         "The name of the organization the repository is a part of.",
						MarkdownDescription: "The name of the organization the repository is a part of.",
						Computed:            true,
					},
				},
			},
			"permissions": schema.SingleNestedAttribute{
				Description:         "A map of permissions for the repository.",
				MarkdownDescription: "A map of permissions for the repository.",
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"admin": schema.BoolAttribute{
						Description:         "Has Admin permissions.",
						MarkdownDescription: "Has Admin permissions.",
						Computed:            true,
					},
					"pull": schema.BoolAttribute{
						Description:         "Has Pull permissions.",
						MarkdownDescription: "Has Pull permissions.",
						Computed:            true,
					},
					"triage": schema.BoolAttribute{
						Description:         "Has Triage permissions.",
						MarkdownDescription: "Has Triage permissions.",
						Computed:            true,
					},
					"push": schema.BoolAttribute{
						Description:         "Has Push permissions.",
						MarkdownDescription: "Has Push permissions.",
						Computed:            true,
					},
					"maintain": schema.BoolAttribute{
						Description:         "Has Maintain permissions.",
						MarkdownDescription: "Has Maintain permissions.",
						Computed:            true,
					},
				},
			},
			"allow_rebase_merge": schema.BoolAttribute{
				Description:         "Indicates if rebase merging is allowed in the repository.",
				MarkdownDescription: "Indicates if rebase merging is allowed in the repository.",
				Computed:            true,
			},
			"allow_update_branch": schema.BoolAttribute{
				Description:         "Indicates if updating a pull request head branch is allowed.",
				MarkdownDescription: "Indicates if updating a pull request head branch is allowed.",
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
			"topics": schema.ListAttribute{
				ElementType:         types.StringType,
				Description:         "The list of topics associated with the repository.",
				MarkdownDescription: "The list of topics associated with the repository.",
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
			"private": schema.BoolAttribute{
				Description:         "Indicates if the repository is private.",
				MarkdownDescription: "Indicates if the repository is private.",
				Computed:            true,
			},
			"has_issues": schema.BoolAttribute{
				Description:         "Indicates if the repository has issues enabled.",
				MarkdownDescription: "Indicates if the repository has issues enabled.",
				Computed:            true,
			},
			"has_wiki": schema.BoolAttribute{
				Description:         "Indicates if the repository has wiki enabled.",
				MarkdownDescription: "Indicates if the repository has wiki enabled.",
				Computed:            true,
			},
			"has_pages": schema.BoolAttribute{
				Description:         "Indicates if the repository has GitHub Pages enabled.",
				MarkdownDescription: "Indicates if the repository has GitHub Pages enabled.",
				Computed:            true,
			},
			"has_projects": schema.BoolAttribute{
				Description:         "Indicates if the repository has projects enabled.",
				MarkdownDescription: "Indicates if the repository has projects enabled.",
				Computed:            true,
			},
			"has_downloads": schema.BoolAttribute{
				Description:         "Indicates if the repository has downloads enabled.",
				MarkdownDescription: "Indicates if the repository has downloads enabled.",
				Computed:            true,
			},
			"has_discussions": schema.BoolAttribute{
				Description:         "Indicates if the repository has discussions enabled.",
				MarkdownDescription: "Indicates if the repository has discussions enabled.",
				Computed:            true,
			},
			"is_template": schema.BoolAttribute{
				Description:         "Indicates if the repository is a template repository.",
				MarkdownDescription: "Indicates if the repository is a template repository.",
				Computed:            true,
			},
			"license_template": schema.StringAttribute{
				Description:         "The license template used by the repository.",
				MarkdownDescription: "The license template used by the repository.",
				Computed:            true,
			},
			"gitignore_template": schema.StringAttribute{
				Description:         "The .gitignore template used by the repository.",
				MarkdownDescription: "The .gitignore template used by the repository.",
				Computed:            true,
			},
			"security_and_analysis": schema.SingleNestedAttribute{
				Description:         "Advanced security and secret scanning options.",
				MarkdownDescription: "Advanced security and secret scanning options.",
				Computed:            true,
				Attributes: map[string]schema.Attribute{
					"advanced_security": schema.SingleNestedAttribute{
						Description:         "Advanced security on the repository.",
						MarkdownDescription: "Advanced security on the repository.",
						Computed:            true,
						Attributes: map[string]schema.Attribute{
							"status": schema.StringAttribute{
								Description:         "The state of advanced security on the repository.",
								MarkdownDescription: "The state of advanced security on the repository.",
								Computed:            true,
							},
						},
					},
					"secret_scanning": schema.SingleNestedAttribute{
						Description:         "Secret scanning on the repository.",
						MarkdownDescription: "Secret scanning on the repository.",
						Computed:            true,
						Attributes: map[string]schema.Attribute{
							"status": schema.StringAttribute{
								Description:         "The state of secret scanning on the repository.",
								MarkdownDescription: "The state of secret scanning on the repository.",
								Computed:            true,
							},
						},
					},
					"secret_scanning_push_protection": schema.SingleNestedAttribute{
						Description:         "Secret scanning push protection on the repository.",
						MarkdownDescription: "Secret scanning push protection on the repository.",
						Computed:            true,
						Attributes: map[string]schema.Attribute{
							"status": schema.StringAttribute{
								Description:         "The state of secret scanning push protection on the repository.",
								MarkdownDescription: "The state of secret scanning push protection on the repository.",
								Computed:            true,
							},
						},
					},
					"secret_scanning_validity_checks": schema.SingleNestedAttribute{
						Description:         "Secret scanning validity checks on the repository.",
						MarkdownDescription: "Secret scanning validity checks on the repository.",
						Computed:            true,
						Attributes: map[string]schema.Attribute{
							"status": schema.StringAttribute{
								Description:         "The state of secret scanning validity checks on the repository.",
								MarkdownDescription: "The state of secret scanning validity checks on the repository.",
								Computed:            true,
							},
						},
					},
					"dependabot_security_updates": schema.SingleNestedAttribute{
						Description:         "Dependabot security updates on the repository.",
						MarkdownDescription: "Dependabot security updates on the repository.",
						Computed:            true,
						Attributes: map[string]schema.Attribute{
							"status": schema.StringAttribute{
								Description:         "The state of dependabot security updates on the repository.",
								MarkdownDescription: "The state of dependabot security updates on the repository.",
								Computed:            true,
							},
						},
					},
				},
			},
			"team_id": schema.Int64Attribute{
				Description:         "The ID of the team associated with the repository.",
				MarkdownDescription: "The ID of the team associated with the repository.",
				Computed:            true,
			},
			"url": schema.StringAttribute{
				Description:         "The URL of the repository.",
				MarkdownDescription: "The URL of the repository.",
				Computed:            true,
			},
			"archive_url": schema.StringAttribute{
				Description:         "The archive URL of the repository.",
				MarkdownDescription: "The archive URL of the repository.",
				Computed:            true,
			},
			"assignees_url": schema.StringAttribute{
				Description:         "The assignees URL of the repository.",
				MarkdownDescription: "The assignees URL of the repository.",
				Computed:            true,
			},
			"blobs_url": schema.StringAttribute{
				Description:         "The blobs URL of the repository.",
				MarkdownDescription: "The blobs URL of the repository.",
				Computed:            true,
			},
			"branches_url": schema.StringAttribute{
				Description:         "The branches URL of the repository.",
				MarkdownDescription: "The branches URL of the repository.",
				Computed:            true,
			},
			"collaborators_url": schema.StringAttribute{
				Description:         "The collaborators URL of the repository.",
				MarkdownDescription: "The collaborators URL of the repository.",
				Computed:            true,
			},
			"comments_url": schema.StringAttribute{
				Description:         "The comments URL of the repository.",
				MarkdownDescription: "The comments URL of the repository.",
				Computed:            true,
			},
			"commits_url": schema.StringAttribute{
				Description:         "The commits URL of the repository.",
				MarkdownDescription: "The commits URL of the repository.",
				Computed:            true,
			},
			"compare_url": schema.StringAttribute{
				Description:         "The compare URL of the repository.",
				MarkdownDescription: "The compare URL of the repository.",
				Computed:            true,
			},
			"contents_url": schema.StringAttribute{
				Description:         "The contents URL of the repository.",
				MarkdownDescription: "The contents URL of the repository.",
				Computed:            true,
			},
			"contributors_url": schema.StringAttribute{
				Description:         "The contributors URL of the repository.",
				MarkdownDescription: "The contributors URL of the repository.",
				Computed:            true,
			},
			"deployments_url": schema.StringAttribute{
				Description:         "The deployments URL of the repository.",
				MarkdownDescription: "The deployments URL of the repository.",
				Computed:            true,
			},
			"downloads_url": schema.StringAttribute{
				Description:         "The downloads URL of the repository.",
				MarkdownDescription: "The downloads URL of the repository.",
				Computed:            true,
			},
			"events_url": schema.StringAttribute{
				Description:         "The events URL of the repository.",
				MarkdownDescription: "The events URL of the repository.",
				Computed:            true,
			},
			"forks_url": schema.StringAttribute{
				Description:         "The forks URL of the repository.",
				MarkdownDescription: "The forks URL of the repository.",
				Computed:            true,
			},
			"git_commits_url": schema.StringAttribute{
				Description:         "The Git commits URL of the repository.",
				MarkdownDescription: "The Git commits URL of the repository.",
				Computed:            true,
			},
			"git_refs_url": schema.StringAttribute{
				Description:         "The Git references URL of the repository.",
				MarkdownDescription: "The Git references URL of the repository.",
				Computed:            true,
			},
			"git_tags_url": schema.StringAttribute{
				Description:         "The Git tags URL of the repository.",
				MarkdownDescription: "The Git tags URL of the repository.",
				Computed:            true,
			},
			"hooks_url": schema.StringAttribute{
				Description:         "The hooks URL of the repository.",
				MarkdownDescription: "The hooks URL of the repository.",
				Computed:            true,
			},
			"issue_comment_url": schema.StringAttribute{
				Description:         "The issue comment URL of the repository.",
				MarkdownDescription: "The issue comment URL of the repository.",
				Computed:            true,
			},
			"issue_events_url": schema.StringAttribute{
				Description:         "The issue events URL of the repository.",
				MarkdownDescription: "The issue events URL of the repository.",
				Computed:            true,
			},
			"issues_url": schema.StringAttribute{
				Description:         "The issues URL of the repository.",
				MarkdownDescription: "The issues URL of the repository.",
				Computed:            true,
			},
			"keys_url": schema.StringAttribute{
				Description:         "The keys URL of the repository.",
				MarkdownDescription: "The keys URL of the repository.",
				Computed:            true,
			},
			"labels_url": schema.StringAttribute{
				Description:         "The labels URL of the repository.",
				MarkdownDescription: "The labels URL of the repository.",
				Computed:            true,
			},
			"languages_url": schema.StringAttribute{
				Description:         "The languages URL of the repository.",
				MarkdownDescription: "The languages URL of the repository.",
				Computed:            true,
			},
			"merges_url": schema.StringAttribute{
				Description:         "The merges URL of the repository.",
				MarkdownDescription: "The merges URL of the repository.",
				Computed:            true,
			},
			"milestones_url": schema.StringAttribute{
				Description:         "The milestones URL of the repository.",
				MarkdownDescription: "The milestones URL of the repository.",
				Computed:            true,
			},
			"notifications_url": schema.StringAttribute{
				Description:         "The notifications URL of the repository.",
				MarkdownDescription: "The notifications URL of the repository.",
				Computed:            true,
			},
			"pulls_url": schema.StringAttribute{
				Description:         "The pull requests URL of the repository.",
				MarkdownDescription: "The pull requests URL of the repository.",
				Computed:            true,
			},
			"releases_url": schema.StringAttribute{
				Description:         "The releases URL of the repository.",
				MarkdownDescription: "The releases URL of the repository.",
				Computed:            true,
			},
			"stargazers_url": schema.StringAttribute{
				Description:         "The stargazers URL of the repository.",
				MarkdownDescription: "The stargazers URL of the repository.",
				Computed:            true,
			},
			"statuses_url": schema.StringAttribute{
				Description:         "The statuses URL of the repository.",
				MarkdownDescription: "The statuses URL of the repository.",
				Computed:            true,
			},
			"subscribers_url": schema.StringAttribute{
				Description:         "The subscribers URL of the repository.",
				MarkdownDescription: "The subscribers URL of the repository.",
				Computed:            true,
			},
			"subscription_url": schema.StringAttribute{
				Description:         "The subscription URL of the repository.",
				MarkdownDescription: "The subscription URL of the repository.",
				Computed:            true,
			},
			"tags_url": schema.StringAttribute{
				Description:         "The tags URL of the repository.",
				MarkdownDescription: "The tags URL of the repository.",
				Computed:            true,
			},
			"trees_url": schema.StringAttribute{
				Description:         "The trees URL of the repository.",
				MarkdownDescription: "The trees URL of the repository.",
				Computed:            true,
			},
			"teams_url": schema.StringAttribute{
				Description:         "The teams URL of the repository.",
				MarkdownDescription: "The teams URL of the repository.",
				Computed:            true,
			},
			"visibility": schema.StringAttribute{
				Description:         "The visibility of the repository (public or private).",
				MarkdownDescription: "The visibility of the repository (public or private).",
				Computed:            true,
			},
		},
		Description:         "Use this data source to retrieve information about a GitHub repository.",
		MarkdownDescription: "Use this data source to retrieve information about a GitHub repository.",
	}
}

func (d *GitHubRepositoryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

	d.client = config.Client
}

func (d *GitHubRepositoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var model GitHubRepositoryDataSourceModel

	client := d.client

	// Read Terraform configuration data into the model.
	resp.Diagnostics.Append(req.Config.Get(ctx, &model)...)

	if resp.Diagnostics.HasError() {
		return
	}

	repo, _, err := client.Repositories.Get(ctx, model.Owner.ValueString(), model.Name.ValueString())

	if err != nil {
		resp.Diagnostics.AddError(
			"Error Communicating with the GitHub API",
			fmt.Sprintf("Unable to get repository, got error: %s", err),
		)
		return
	}

	createdAt := repo.GetCreatedAt()
	pushedAt := repo.GetPushedAt()
	updatedAt := repo.GetUpdatedAt()
	permissions := repo.GetPermissions()
	parent := repo.GetParent()
	source := repo.GetSource()
	templateRepository := repo.GetTemplateRepository()
	organization := repo.GetOrganization()
	securityAndAnalysis := repo.GetSecurityAndAnalysis()
	advancedSecurity := securityAndAnalysis.GetAdvancedSecurity()
	secretScanning := securityAndAnalysis.GetSecretScanning()
	secretScanningPushProtection := securityAndAnalysis.GetSecretScanningPushProtection()
	secretScanningValidityChecks := securityAndAnalysis.GetSecretScanningValidityChecks()
	dependabotSecurityUpdates := securityAndAnalysis.GetDependabotSecurityUpdates()

	model.ID = types.Int64Value(repo.GetID())
	model.NodeID = types.StringValue(repo.GetNodeID())
	model.Name = types.StringValue(repo.GetName())
	model.FullName = types.StringValue(repo.GetFullName())
	model.Description = types.StringValue(repo.GetDescription())
	model.Homepage = types.StringValue(repo.GetHomepage())
	model.DefaultBranch = types.StringValue(repo.GetDefaultBranch())
	model.MasterBranch = types.StringValue(repo.GetMasterBranch())
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
	model.Parent = &linkedRepositoryModel{}
	model.Parent.Owner = types.StringValue(parent.GetOwner().GetLogin())
	model.Parent.Name = types.StringValue(parent.GetName())
	model.Parent.FullName = types.StringValue(parent.GetFullName())
	model.Source = &linkedRepositoryModel{}
	model.Source.Owner = types.StringValue(source.GetOwner().GetLogin())
	model.Source.Name = types.StringValue(source.GetName())
	model.Source.FullName = types.StringValue(source.GetFullName())
	model.TemplateRepository = &linkedRepositoryModel{}
	model.TemplateRepository.Owner = types.StringValue(templateRepository.GetOwner().GetLogin())
	model.TemplateRepository.Name = types.StringValue(templateRepository.GetName())
	model.TemplateRepository.FullName = types.StringValue(templateRepository.GetFullName())
	model.Organization = &organizationModel{}
	model.Organization.Name = types.StringValue(organization.GetLogin())
	model.Permissions = &permissionsModel{}
	model.Permissions.Admin = types.BoolValue(permissions["admin"])
	model.Permissions.Pull = types.BoolValue(permissions["pull"])
	model.Permissions.Triage = types.BoolValue(permissions["triage"])
	model.Permissions.Push = types.BoolValue(permissions["push"])
	model.Permissions.Maintain = types.BoolValue(permissions["maintain"])
	model.AllowRebaseMerge = types.BoolValue(repo.GetAllowRebaseMerge())
	model.AllowUpdateBranch = types.BoolValue(repo.GetAllowUpdateBranch())
	model.AllowSquashMerge = types.BoolValue(repo.GetAllowSquashMerge())
	model.AllowMergeCommit = types.BoolValue(repo.GetAllowMergeCommit())
	model.AllowAutoMerge = types.BoolValue(repo.GetAllowAutoMerge())
	model.AllowForking = types.BoolValue(repo.GetAllowForking())
	model.WebCommitSignoffRequired = types.BoolValue(repo.GetWebCommitSignoffRequired())
	model.DeleteBranchOnMerge = types.BoolValue(repo.GetDeleteBranchOnMerge())
	model.SquashMergeCommitTitle = types.StringValue(repo.GetSquashMergeCommitTitle())
	model.SquashMergeCommitMessage = types.StringValue(repo.GetSquashMergeCommitMessage())
	model.MergeCommitTitle = types.StringValue(repo.GetMergeCommitTitle())
	model.MergeCommitMessage = types.StringValue(repo.GetMergeCommitMessage())
	model.Topics, _ = types.ListValueFrom(ctx, types.StringType, repo.Topics)
	model.Archived = types.BoolValue(repo.GetArchived())
	model.Disabled = types.BoolValue(repo.GetDisabled())
	model.Private = types.BoolValue(repo.GetPrivate())
	model.HasIssues = types.BoolValue(repo.GetHasIssues())
	model.HasWiki = types.BoolValue(repo.GetHasWiki())
	model.HasPages = types.BoolValue(repo.GetHasPages())
	model.HasProjects = types.BoolValue(repo.GetHasProjects())
	model.HasDownloads = types.BoolValue(repo.GetHasDownloads())
	model.HasDiscussions = types.BoolValue(repo.GetHasDiscussions())
	model.IsTemplate = types.BoolValue(repo.GetIsTemplate())
	model.LicenseTemplate = types.StringValue(repo.GetLicenseTemplate())
	model.GitignoreTemplate = types.StringValue(repo.GetGitignoreTemplate())
	model.SecurityAndAnalysis = &securityAndAnalysisModel{}
	model.SecurityAndAnalysis.AdvancedSecurity = &advancedSecurityModel{}
	model.SecurityAndAnalysis.AdvancedSecurity.Status = types.StringValue(advancedSecurity.GetStatus())
	model.SecurityAndAnalysis.SecretScanning = &secretScanningModel{}
	model.SecurityAndAnalysis.SecretScanning.Status = types.StringValue(secretScanning.GetStatus())
	model.SecurityAndAnalysis.SecretScanningPushProtection = &secretScanningPushProtectionModel{}
	model.SecurityAndAnalysis.SecretScanningPushProtection.Status = types.StringValue(secretScanningPushProtection.GetStatus())
	model.SecurityAndAnalysis.SecretScanningValidityChecks = &secretScanningValidityChecksModel{}
	model.SecurityAndAnalysis.SecretScanningValidityChecks.Status = types.StringValue(secretScanningValidityChecks.GetStatus())
	model.SecurityAndAnalysis.DependabotSecurityUpdates = &dependabotSecurityUpdatesModel{}
	model.SecurityAndAnalysis.DependabotSecurityUpdates.Status = types.StringValue(dependabotSecurityUpdates.GetStatus())
	model.TeamID = types.Int64Value(repo.GetTeamID())
	model.URL = types.StringValue(repo.GetURL())
	model.ArchiveURL = types.StringValue(repo.GetArchiveURL())
	model.AssigneesURL = types.StringValue(repo.GetAssigneesURL())
	model.BlobsURL = types.StringValue(repo.GetBlobsURL())
	model.BranchesURL = types.StringValue(repo.GetBranchesURL())
	model.CollaboratorsURL = types.StringValue(repo.GetCollaboratorsURL())
	model.CommentsURL = types.StringValue(repo.GetCommentsURL())
	model.CommitsURL = types.StringValue(repo.GetCommitsURL())
	model.CompareURL = types.StringValue(repo.GetCompareURL())
	model.ContentsURL = types.StringValue(repo.GetContentsURL())
	model.ContributorsURL = types.StringValue(repo.GetContributorsURL())
	model.DeploymentsURL = types.StringValue(repo.GetDeploymentsURL())
	model.DownloadsURL = types.StringValue(repo.GetDownloadsURL())
	model.EventsURL = types.StringValue(repo.GetEventsURL())
	model.ForksURL = types.StringValue(repo.GetForksURL())
	model.GitCommitsURL = types.StringValue(repo.GetGitCommitsURL())
	model.GitRefsURL = types.StringValue(repo.GetGitRefsURL())
	model.GitTagsURL = types.StringValue(repo.GetGitTagsURL())
	model.HooksURL = types.StringValue(repo.GetHooksURL())
	model.IssueCommentURL = types.StringValue(repo.GetIssueCommentURL())
	model.IssueEventsURL = types.StringValue(repo.GetIssueEventsURL())
	model.IssuesURL = types.StringValue(repo.GetIssuesURL())
	model.KeysURL = types.StringValue(repo.GetKeysURL())
	model.LabelsURL = types.StringValue(repo.GetLabelsURL())
	model.LanguagesURL = types.StringValue(repo.GetLanguagesURL())
	model.MergesURL = types.StringValue(repo.GetMergesURL())
	model.MilestonesURL = types.StringValue(repo.GetMilestonesURL())
	model.NotificationsURL = types.StringValue(repo.GetNotificationsURL())
	model.PullsURL = types.StringValue(repo.GetPullsURL())
	model.ReleasesURL = types.StringValue(repo.GetReleasesURL())
	model.StargazersURL = types.StringValue(repo.GetStargazersURL())
	model.StatusesURL = types.StringValue(repo.GetStatusesURL())
	model.SubscribersURL = types.StringValue(repo.GetSubscribersURL())
	model.SubscriptionURL = types.StringValue(repo.GetSubscriptionURL())
	model.TagsURL = types.StringValue(repo.GetTagsURL())
	model.TreesURL = types.StringValue(repo.GetTreesURL())
	model.TeamsURL = types.StringValue(repo.GetTeamsURL())
	model.Visibility = types.StringValue(repo.GetVisibility())

	// Save updated data into Terraform state.
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
