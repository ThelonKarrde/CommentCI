package github

import (
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"log"
	"net/http"
)

func createOauthClient(accessToken string) (context.Context, *http.Client) {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	return ctx, oauth2.NewClient(ctx, ts)
}

func Comment(commentText *string, accessToken string, owner string, repo string, issue int) {
	ctx, oauthClient := createOauthClient(accessToken)
	client := github.NewClient(oauthClient)
	comment := github.IssueComment{Body: commentText}
	_, _, err := client.Issues.CreateComment(ctx, owner, repo, issue, &comment)
	if err != nil {
		log.Fatalf(err.Error())
	}
}
