package config

import (
	"github.com/akamensky/argparse"
	"log"
	"os"
)

type Data struct {
	GitHubRepoOwner string
	GitHubRepoName string
	GitHubCommentUser string
	GitHubCommentToken string
}

func readEnvConfig() *Data {
	ghro, _ := os.LookupEnv("GITHUB_REPO_OWNER")
	ghrn, _ := os.LookupEnv("GITHUB_REPO_NAME")
	ghcu, _ := os.LookupEnv("GITHUB_COMMENT_USER")
	ghct, _ := os.LookupEnv("GITHUB_COMMENT_TOKEN")
	return &Data{
		GitHubRepoOwner:    ghro,
		GitHubRepoName:     ghrn,
		GitHubCommentUser:  ghcu,
		GitHubCommentToken: ghct,
	}
}

func readArgConfig() *Data {
	parser := argparse.NewParser("CommentCI", "Sent a comment to GitHub PR or Issue from your CI")
	ghro := parser.String("o", "github-owner", &argparse.Options{
		Required: true,
		Help:     "Name of the owner of repository",
	})
	ghrn := parser.String("r", "github-repository", &argparse.Options{
		Required: true,
		Help:     "Name of the github repository",
	})
	ghcu := parser.String("u", "github-user", &argparse.Options{
		Required: true,
		Help:     "User which will comment PR",
	})
	ghct := parser.String("t", "github-token", &argparse.Options{
		Required: false,
		Help:     "Access token to comment PR",
	})
	err := parser.Parse(os.Args)
	if err != nil {
		log.Println(err.Error())
	}
	return &Data{
		GitHubRepoOwner:    *ghro,
		GitHubRepoName:     *ghrn,
		GitHubCommentUser:  *ghcu,
		GitHubCommentToken: *ghct,
	}
}

func ReadConfig() Data {
	envData := readEnvConfig()
	argData := readArgConfig()
	data := Data{}
	if envData.GitHubRepoName == "" {
		data = *argData
	} else {
		data = *envData
	}
	return data
}