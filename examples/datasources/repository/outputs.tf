output "github_repository_example_id" {
  description = "GitHub ID for the repository."
  value       = data.github_repository.example.id
}

output "github_repository_example_node_id" {
  description = "The node ID of the repository."
  value       = data.github_repository.example.node_id
}

output "github_repository_example_full_name" {
  description = "The full name of the repository."
  value       = data.github_repository.example.full_name
}

output "github_repository_example_description" {
  description = "The description of the repository."
  value       = data.github_repository.example.description
}

output "github_repository_example_homepage" {
  description = "The homepage of the repository."
  value       = data.github_repository.example.homepage
}

output "github_repository_example_default_branch" {
  description = "The repository's default branch."
  value       = data.github_repository.example.default_branch
}

output "github_repository_example_master_branch" {
  description = "The repository's master branch."
  value       = data.github_repository.example.master_branch
}

output "github_repository_example_created_at" {
  description = "The timestamp of when the repository was created on GitHub."
  value       = data.github_repository.example.created_at
}

output "github_repository_example_pushed_at" {
  description = "The timestamp of the last push to the repository."
  value       = data.github_repository.example.pushed_at
}

output "github_repository_example_updated_at" {
  description = "The timestamp of when the repository was last updated."
  value       = data.github_repository.example.updated_at
}

output "github_repository_example_html_url" {
  description = "The HTML URL of the repository."
  value       = data.github_repository.example.html_url
}

output "github_repository_example_clone_url" {
  description = "The URL used for cloning the repository."
  value       = data.github_repository.example.clone_url
}

output "github_repository_example_git_url" {
  description = "The git URL of the repository."
  value       = data.github_repository.example.git_url
}

output "github_repository_example_mirror_url" {
  description = "The mirror URL of the repository."
  value       = data.github_repository.example.mirror_url
}

output "github_repository_example_ssh_url" {
  description = "The SSH URL of the repository."
  value       = data.github_repository.example.ssh_url
}

output "github_repository_example_svn_url" {
  description = "The SVN URL of the repository."
  value       = data.github_repository.example.svn_url
}

output "github_repository_example_language" {
  description = "The primary language of the repository."
  value       = data.github_repository.example.language
}

output "github_repository_example_fork" {
  description = "Indicates if the repository is a fork."
  value       = data.github_repository.example.fork
}

output "github_repository_example_forks_count" {
  description = "The number of forks of the repository."
  value       = data.github_repository.example.forks_count
}

output "github_repository_example_network_count" {
  description = "The number of repositories in the network."
  value       = data.github_repository.example.network_count
}

output "github_repository_example_open_issues_count" {
  description = "The number of open issues in the repository."
  value       = data.github_repository.example.open_issues_count
}

output "github_repository_example_stargazers_count" {
  description = "The number of users who have starred the repository."
  value       = data.github_repository.example.stargazers_count
}

output "github_repository_example_subscribers_count" {
  description = "The number of users subscribed to the repository."
  value       = data.github_repository.example.subscribers_count
}

output "github_repository_example_size" {
  description = "The size of the repository, in kilobytes."
  value       = data.github_repository.example.size
}

output "github_repository_example_auto_init" {
  description = "Indicates if the repository is initialized with a README."
  value       = data.github_repository.example.auto_init
}

output "github_repository_example_parent" {
  description = "Details of the parent repository."
  value       = data.github_repository.example.parent
}

output "github_repository_example_parent_owner" {
  description = "The owner of the parent repository."
  value       = data.github_repository.example.parent.owner
}

output "github_repository_example_parent_name" {
  description = "The name of the parent repository."
  value       = data.github_repository.example.parent.name
}

output "github_repository_example_parent_full_name" {
  description = "The full name of the parent repository."
  value       = data.github_repository.example.parent.full_name
}

output "github_repository_example_source" {
  description = "Details of the source repository."
  value       = data.github_repository.example.source
}

output "github_repository_example_source_owner" {
  description = "The owner of the source repository."
  value       = data.github_repository.example.source.owner
}

output "github_repository_example_source_name" {
  description = "The name of the source repository."
  value       = data.github_repository.example.source.name
}

output "github_repository_example_source_full_name" {
  description = "The full name of the source repository."
  value       = data.github_repository.example.source.full_name
}

output "github_repository_example_template_repository" {
  description = "Details of the template repository."
  value       = data.github_repository.example.template_repository
}

output "github_repository_example_template_repository_owner" {
  description = "The owner of the template repository."
  value       = data.github_repository.example.template_repository.owner
}

output "github_repository_example_template_repository_name" {
  description = "The name of the template repository."
  value       = data.github_repository.example.template_repository.name
}

output "github_repository_example_template_repository_full_name" {
  description = "The full name of the template repository."
  value       = data.github_repository.example.template_repository.full_name
}

output "github_repository_example_organization" {
  description = "Details of the organization the repository is a part of."
  value       = data.github_repository.example.organization
}

output "github_repository_example_organization_name" {
  description = "The name of the organization the repository is a part of."
  value       = data.github_repository.example.organization.name
}

output "github_repository_example_permissions" {
  description = "A map of permissions for the repository."
  value       = data.github_repository.example.permissions
}

output "github_repository_example_security_and_analysis" {
  description = "Advanced security and secret scanning options."
  value       = data.github_repository.example.security_and_analysis
}

output "github_repository_example_security_and_analysis_advanced_security" {
  description = "Advanced security on the repository."
  value       = data.github_repository.example.security_and_analysis.advanced_security
}

output "github_repository_example_security_and_analysis_advanced_security_status" {
  description = "The state of advanced security on the repository."
  value       = data.github_repository.example.security_and_analysis.advanced_security.status
}

output "github_repository_example_security_and_analysis_secret_scanning" {
  description = "Secret scanning on the repository."
  value       = data.github_repository.example.security_and_analysis.secret_scanning
}

output "github_repository_example_security_and_analysis_secret_scanning_status" {
  description = "The state of secret scanning on the repository."
  value       = data.github_repository.example.security_and_analysis.secret_scanning.status
}

output "github_repository_example_security_and_analysis_secret_scanning_push_protection" {
  description = "Secret scanning push protection on the repository."
  value       = data.github_repository.example.security_and_analysis.secret_scanning_push_protection
}

output "github_repository_example_security_and_analysis_secret_scanning_push_protection_status" {
  description = "The state of secret scanning push protection on the repository."
  value       = data.github_repository.example.security_and_analysis.secret_scanning_push_protection.status
}

output "github_repository_example_security_and_analysis_secret_scanning_validity_checks" {
  description = "Secret scanning validity checks on the repository."
  value       = data.github_repository.example.security_and_analysis.secret_scanning_validity_checks
}

output "github_repository_example_security_and_analysis_secret_scanning_validity_checks_status" {
  description = "The state of secret scanning validity checks on the repository."
  value       = data.github_repository.example.security_and_analysis.secret_scanning_validity_checks.status
}

output "github_repository_example_security_and_analysis_dependabot_security_updates" {
  description = "Dependabot security updates on the repository."
  value       = data.github_repository.example.security_and_analysis.dependabot_security_updates
}

output "github_repository_example_security_and_analysis_dependabot_security_updates_status" {
  description = "The state of dependabot security updates on the repository."
  value       = data.github_repository.example.security_and_analysis.dependabot_security_updates.status
}

output "github_repository_example_team_id" {
  description = "The ID of the team associated with the repository."
  value       = data.github_repository.example.team_id
}

output "github_repository_example_url" {
  description = "The URL of the repository."
  value       = data.github_repository.example.url
}

output "github_repository_example_archive_url" {
  description = "The archive URL of the repository."
  value       = data.github_repository.example.archive_url
}

output "github_repository_example_assignees_url" {
  description = "The assignees URL of the repository."
  value       = data.github_repository.example.assignees_url
}

output "github_repository_example_blobs_url" {
  description = "The blobs URL of the repository."
  value       = data.github_repository.example.blobs_url
}

output "github_repository_example_branches_url" {
  description = "The branches URL of the repository."
  value       = data.github_repository.example.branches_url
}

output "github_repository_example_collaborators_url" {
  description = "The collaborators URL of the repository."
  value       = data.github_repository.example.collaborators_url
}

output "github_repository_example_comments_url" {
  description = "The comments URL of the repository."
  value       = data.github_repository.example.comments_url
}

output "github_repository_example_commits_url" {
  description = "The commits URL of the repository."
  value       = data.github_repository.example.commits_url
}

output "github_repository_example_compare_url" {
  description = "The compare URL of the repository."
  value       = data.github_repository.example.compare_url
}

output "github_repository_example_contents_url" {
  description = "The contents URL of the repository."
  value       = data.github_repository.example.contents_url
}

output "github_repository_example_contributors_url" {
  description = "The contributors URL of the repository."
  value       = data.github_repository.example.contributors_url
}

output "github_repository_example_deployments_url" {
  description = "The deployments URL of the repository."
  value       = data.github_repository.example.deployments_url
}

output "github_repository_example_downloads_url" {
  description = "The downloads URL of the repository."
  value       = data.github_repository.example.downloads_url
}

output "github_repository_example_events_url" {
  description = "The events URL of the repository."
  value       = data.github_repository.example.events_url
}

output "github_repository_example_forks_url" {
  description = "The forks URL of the repository."
  value       = data.github_repository.example.forks_url
}

output "github_repository_example_git_commits_url" {
  description = "The Git commits URL of the repository."
  value       = data.github_repository.example.git_commits_url
}

output "github_repository_example_git_refs_url" {
  description = "The Git references URL of the repository."
  value       = data.github_repository.example.git_refs_url
}

output "github_repository_example_git_tags_url" {
  description = "The Git tags URL of the repository."
  value       = data.github_repository.example.git_tags_url
}

output "github_repository_example_hooks_url" {
  description = "The hooks URL of the repository."
  value       = data.github_repository.example.hooks_url
}

output "github_repository_example_issue_comment_url" {
  description = "The issue comment URL of the repository."
  value       = data.github_repository.example.issue_comment_url
}

output "github_repository_example_issue_events_url" {
  description = "The issue events URL of the repository."
  value       = data.github_repository.example.issue_events_url
}

output "github_repository_example_issues_url" {
  description = "The issues URL of the repository."
  value       = data.github_repository.example.issues_url
}

output "github_repository_example_keys_url" {
  description = "The keys URL of the repository."
  value       = data.github_repository.example.keys_url
}

output "github_repository_example_labels_url" {
  description = "The labels URL of the repository."
  value       = data.github_repository.example.labels_url
}

output "github_repository_example_languages_url" {
  description = "The languages URL of the repository."
  value       = data.github_repository.example.languages_url
}

output "github_repository_example_merges_url" {
  description = "The merges URL of the repository."
  value       = data.github_repository.example.merges_url
}

output "github_repository_example_milestones_url" {
  description = "The milestones URL of the repository."
  value       = data.github_repository.example.milestones_url
}

output "github_repository_example_notifications_url" {
  description = "The notifications URL of the repository."
  value       = data.github_repository.example.notifications_url
}

output "github_repository_example_pulls_url" {
  description = "The pull requests URL of the repository."
  value       = data.github_repository.example.pulls_url
}

output "github_repository_example_releases_url" {
  description = "The releases URL of the repository."
  value       = data.github_repository.example.releases_url
}

output "github_repository_example_stargazers_url" {
  description = "The stargazers URL of the repository."
  value       = data.github_repository.example.stargazers_url
}

output "github_repository_example_statuses_url" {
  description = "The statuses URL of the repository."
  value       = data.github_repository.example.statuses_url
}

output "github_repository_example_subscribers_url" {
  description = "The subscribers URL of the repository."
  value       = data.github_repository.example.subscribers_url
}

output "github_repository_example_subscription_url" {
  description = "The subscription URL of the repository."
  value       = data.github_repository.example.subscription_url
}

output "github_repository_example_tags_url" {
  description = "The tags URL of the repository."
  value       = data.github_repository.example.tags_url
}

output "github_repository_example_trees_url" {
  description = "The trees URL of the repository."
  value       = data.github_repository.example.trees_url
}

output "github_repository_example_teams_url" {
  description = "The teams URL of the repository."
  value       = data.github_repository.example.teams_url
}

output "github_repository_example_visibility" {
  description = "The visibility of the repository (public or private)."
  value       = data.github_repository.example.visibility
}

#output "github_repository_example_id" {
#  description = "This is an example."
#  value       = data.github_repository.example.id
#}
#
#output "github_repository_example_node_id" {
#  description = "This is an example."
#  value       = data.github_repository.example.node_id
#}
#
#output "github_repository_example_name" {
#  description = "This is an example."
#  value       = data.github_repository.example.name
#}
#
#output "github_repository_example_full_name" {
#  description = "This is an example."
#  value       = data.github_repository.example.full_name
#}
#
#output "github_repository_example_description" {
#  description = "This is an example."
#  value       = data.github_repository.example.description
#}
#
#output "github_repository_example_homepage" {
#  description = "This is an example."
#  value       = data.github_repository.example.homepage
#}
#
#output "github_repository_example_default_branch" {
#  description = "This is an example."
#  value       = data.github_repository.example.default_branch
#}
#
#output "github_repository_example_master_branch" {
#  description = "This is an example."
#  value       = data.github_repository.example.master_branch
#}
#
#output "github_repository_example_created_at" {
#  description = "This is an example."
#  value       = data.github_repository.example.created_at
#}
#
#output "github_repository_example_pushed_at" {
#  description = "This is an example."
#  value       = data.github_repository.example.pushed_at
#}
#
#output "github_repository_example_updated_at" {
#  description = "This is an example."
#  value       = data.github_repository.example.updated_at
#}
#
#output "github_repository_example_html_url" {
#  description = "This is an example."
#  value       = data.github_repository.example.html_url
#}
#
#output "github_repository_example_clone_url" {
#  description = "This is an example."
#  value       = data.github_repository.example.clone_url
#}
#
#output "github_repository_example_git_url" {
#  description = "This is an example."
#  value       = data.github_repository.example.git_url
#}
#
#output "github_repository_example_mirror_url" {
#  description = "This is an example."
#  value       = data.github_repository.example.mirror_url
#}
#
#output "github_repository_example_ssh_url" {
#  description = "This is an example."
#  value       = data.github_repository.example.ssh_url
#}
#
#output "github_repository_example_svn_url" {
#  description = "This is an example."
#  value       = data.github_repository.example.svn_url
#}
#
#output "github_repository_example_language" {
#  description = "This is an example."
#  value       = data.github_repository.example.language
#}
#
#output "github_repository_example_fork" {
#  description = "This is an example."
#  value       = data.github_repository.example.fork
#}
#
#output "github_repository_example_forks_count" {
#  description = "The number of forks of the repository."
#  value       = data.github_repository.example.forks_count
#}
#
#output "github_repository_example_network_count" {
#  description = "The count of the repository network."
#  value       = data.github_repository.example.network_count
#}
#
#output "github_repository_example_open_issues_count" {
#  description = "The number of open issues in the repository."
#  value       = data.github_repository.example.open_issues_count
#}
#
#output "github_repository_example_stargazers_count" {
#  description = "The number of users who have starred the repository."
#  value       = data.github_repository.example.stargazers_count
#}
#
#output "github_repository_example_subscribers_count" {
#  description = "The number of users subscribed to the repository."
#  value       = data.github_repository.example.subscribers_count
#}
#
#output "github_repository_example_size" {
#  description = "The size of the repository, in kilobytes."
#  value       = data.github_repository.example.size
#}
#
#output "github_repository_example_auto_init" {
#  description = "This is an example."
#  value       = data.github_repository.example.auto_init
#}
#
#output "github_repository_example_parent" {
#  description = "This is an example."
#  value       = data.github_repository.example.parent
#}
#
#output "github_repository_example_source" {
#  description = "This is an example."
#  value       = data.github_repository.example.source
#}
#
#output "github_repository_example_template_repository" {
#  description = "This is an example."
#  value       = data.github_repository.example.template_repository
#}
#
#output "github_repository_example_organization" {
#  description = "This is an example."
#  value       = data.github_repository.example.organization
#}
#
#output "github_repository_example_permissions_admin" {
#  description = "This is an example."
#  value       = data.github_repository.example.permissions.admin
#}
#
#output "github_repository_example_permissions_pull" {
#  description = "This is an example."
#  value       = data.github_repository.example.permissions.pull
#}
#
#output "github_repository_example_permissions_triage" {
#  description = "This is an example."
#  value       = data.github_repository.example.permissions.triage
#}
#
#output "github_repository_example_permissions_push" {
#  description = "This is an example."
#  value       = data.github_repository.example.permissions.push
#}
#
#output "github_repository_example_permissions_maintain" {
#  description = "This is an example."
#  value       = data.github_repository.example.permissions.maintain
#}
#
#output "github_repository_example_allow_rebase_merge" {
#  description = "This is an example."
#  value       = data.github_repository.example.allow_rebase_merge
#}
#
#output "github_repository_example_allow_update_branch" {
#  description = "This is an example."
#  value       = data.github_repository.example.allow_update_branch
#}
#
#output "github_repository_example_allow_squash_merge" {
#  description = "This is an example."
#  value       = data.github_repository.example.allow_squash_merge
#}
#
#output "github_repository_example_allow_merge_commit" {
#  description = "This is an example."
#  value       = data.github_repository.example.allow_merge_commit
#}
#
#output "github_repository_example_allow_auto_merge" {
#  description = "This is an example."
#  value       = data.github_repository.example.allow_auto_merge
#}
#
#output "github_repository_example_allow_forking" {
#  description = "This is an example."
#  value       = data.github_repository.example.allow_forking
#}
#
#output "github_repository_example_web_commit_signoff_required" {
#  description = "This is an example."
#  value       = data.github_repository.example.web_commit_signoff_required
#}
#
#output "github_repository_example_delete_branch_on_merge" {
#  description = "This is an example."
#  value       = data.github_repository.example.delete_branch_on_merge
#}
#
#output "github_repository_example_use_squash_pr_title_as_default" {
#  description = "This is an example."
#  value       = data.github_repository.example.use_squash_pr_title_as_default
#}
#
#output "github_repository_example_squash_merge_commit_title" {
#  description = "This is an example."
#  value       = data.github_repository.example.squash_merge_commit_title
#}
#
#output "github_repository_example_squash_merge_commit_message" {
#  description = "This is an example."
#  value       = data.github_repository.example.squash_merge_commit_message
#}
#
#output "github_repository_example_merge_commit_title" {
#  description = "This is an example."
#  value       = data.github_repository.example.merge_commit_title
#}
#
#output "github_repository_example_merge_commit_message" {
#  description = "This is an example."
#  value       = data.github_repository.example.merge_commit_message
#}
#
#output "github_repository_example_topics" {
#  description = "This is an example."
#  value       = data.github_repository.example.topics
#}
#
#output "github_repository_example_archived" {
#  description = "This is an example."
#  value       = data.github_repository.example.archived
#}
#
#output "github_repository_example_disabled" {
#  description = "This is an example."
#  value       = data.github_repository.example.disabled
#}
#
