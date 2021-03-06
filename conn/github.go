package conn

import (
	"context"

	"github.com/google/go-github/v29/github"
	"golang.org/x/oauth2"
)

func InitClient(token string, ctx context.Context) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}
