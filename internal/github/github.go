package github

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func createGithubToken(accessToken string) (context.Context, *http.Client) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	return ctx, oauth2.NewClient(ctx, ts)
}

func CommentIssue(comment *string, accessToken string, owner string, repo string, issue int) {
	ctx, token := createGithubToken(accessToken)
	ghClient := github.NewClient(token)
	ghCmt := github.IssueComment{Body: comment}
	_, _, err := ghClient.Issues.CreateComment(ctx, owner, repo, issue, &ghCmt)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
