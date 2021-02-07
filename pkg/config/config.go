package config

import (
	"github.com/akamensky/argparse"
	"log"
	"os"
)

type Data struct {
	GitHubRepoOwner    string
	GitHubRepoName     string
	GitHubCommentUser  string
	GitHubCommentToken string
	CommentText        string
	CommentFiles       []string
	FileList           []string
	CodeStyleMode      bool
}

func readEnvConfig(data *Data) *Data {
	ghcu, ok := os.LookupEnv("GITHUB_COMMENT_USER")
	if !ok {
		log.Fatalf("No GitHub user specified! Check GITHUB_COMMENT_USER env.")
	}
	data.GitHubCommentUser = ghcu
	ghct, ok := os.LookupEnv("GITHUB_COMMENT_TOKEN")
	if !ok {
		log.Fatalf("No GitHub token specified! Check GITHUB_COMMENT_TOKEN env.")
	}
	data.GitHubCommentToken = ghct
	return data
}

func readArgConfig(data *Data) *Data {
	parser := argparse.NewParser("CommentCI", "Sent a comment to GitHub PR or Issue from your CI")
	ghro := parser.String("o", "github-owner", &argparse.Options{
		Required: false,
		Help:     "Name of the owner of repository",
	})
	ghrn := parser.String("r", "github-repository", &argparse.Options{
		Required: false,
		Help:     "Name of the github repository",
	})
	cmt := parser.String("s", "comment", &argparse.Options{
		Required: false,
		Help:     "Single comment string to sent to PR",
	})
	csm := parser.Flag("c", "code", &argparse.Options{
		Help: "Set this flag if you want to put your file outputs to Markdown code block",
	})
	fList := parser.StringList("f", "file", &argparse.Options{
		Required: false,
		Help:     "By repeating this flag you can specify multiply files which needs to be send to comment.",
	})
	fCmt := parser.StringList("l", "file-comments", &argparse.Options{
		Required: false,
		Help:     "By repeating this flag you can specify comments for provided files in according order",
	})
	err := parser.Parse(os.Args)
	if err != nil {
		log.Println(err.Error())
	}
	data.GitHubRepoOwner = *ghro
	data.GitHubRepoName = *ghrn
	data.CommentText = *cmt
	data.CodeStyleMode = *csm
	data.FileList = *fList
	data.CommentFiles = *fCmt
	return data
}

func ReadConfig() Data {
	data := Data{}
	data = *readEnvConfig(&data)
	data = *readArgConfig(&data)
	return data
}
