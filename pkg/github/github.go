package github

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func createGHToken(token string) (context.Context, *http.Client) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	return ctx, oauth2.NewClient(ctx, ts)
}

func CommentIssue(cmt *string, token string, owner string, repo string, issue int) {
	ctx, tkn := createGHToken(token)
	ghClient := github.NewClient(tkn)
	ghCmt := github.IssueComment{Body: cmt}
	_, _, err := ghClient.Issues.CreateComment(ctx, owner, repo, issue, &ghCmt)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
