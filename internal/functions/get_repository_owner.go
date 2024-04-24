package functions

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = GetRepositoryOwner{}

func NewGetRepositoryOwner() function.Function {
	return GetRepositoryOwner{}
}

type GetRepositoryOwner struct{}

func (r GetRepositoryOwner) Metadata(_ context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "get_repository_owner"
}

func (r GetRepositoryOwner) Definition(_ context.Context, _ function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:             "Get the owner of a repository from the full name.",
		Description:         "Get the owner of a repository from the full name.",
		MarkdownDescription: "Get the owner of a repository from the full name.",
		Parameters: []function.Parameter{
			function.StringParameter{
				Name:                "full_name",
				Description:         "The full name of the repository.",
				MarkdownDescription: "The full name of the repository.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (r GetRepositoryOwner) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var data string

	resp.Error = function.ConcatFuncErrors(req.Arguments.Get(ctx, &data))

	if resp.Error != nil {
		return
	}

	parts := strings.SplitN(data, "/", 2)
	if len(parts) > 0 {
		data = parts[0]
	}

	resp.Error = function.ConcatFuncErrors(resp.Result.Set(ctx, data))
}
