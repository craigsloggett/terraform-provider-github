package common

import (
	"github.com/google/go-github/v60/github"
)

type ClientConfiguration struct {
	Client *github.Client
	Owner  string
}
