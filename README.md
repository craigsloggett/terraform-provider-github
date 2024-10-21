# Terraform GitHub Provider

The GitHub provider is used to manage and configure resources offered by GitHub. The provider needs to be configured with the proper credentials before it can be used.

## Authentication

The GitHub provider currently only supports the use of a personal access token to authenticate with the GitHub API.

### Personal Access Token

To authenticate using a fine-grained personal access token, ensure that the `token` argument or the `GITHUB_TOKEN` environment variable is set.

```terraform
provider "github" {
  token = var.github_token # Or the GITHUB_TOKEN environment variable.
}
```

## Contribution

A `Makefile` has been created for local development of this provider. To run the checks done in CI locally, simply run `make` before pushing your changes. The `Makefile` has been written such that tests are done hermetically and do not depend on tooling installed on your development machine.

Running `make` with no arguments will lint, build, generate documentation, and then test the provider.

### Authentication

In order to test the GitHub provider in this repository, a personal access token must be available to create, update, and delete entities in GitHub.

The provider looks for the `GITHUB_TOKEN` environment variable when configuring a GitHub client:

```shell
$ export GITHUB_TOKEN="github_pat_xxxxxxxxxxxxxxxxxxxxxxxxx"
```

### Release Automation

As part of the release pipeline, the GitHub Action workflow will sign the provider before uploading it to the public registry. To support this, the following secrets have been added to the GitHub repository:

| Secret Name                  | Description                                                                                     |
| ---------------------------- | ----------------------------------------------------------------------------------------------- |
| `GPG_PRIVATE_KEY`            | The GPG private key used to sign provider releases before publishing to the Terraform registry. |
| `GPG_PRIVATE_KEY_PASSPHRASE` | The passphrase for the GPG private signing key.                                                 |

The release pipeline will check commits made to the `main` branch since the last release and determine how to bump the version number. This is based on conventional commits.
