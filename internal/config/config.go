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
	IssueNumber        int
	MultiCommentMode   bool
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
		Required: true,
		Help:     "Owner of the repository. User/Organisations.",
	})
	ghrn := parser.String("r", "github-repository", &argparse.Options{
		Required: true,
		Help:     "Name of the github repository.",
	})
	cmt := parser.String("s", "single-comment", &argparse.Options{
		Required: false,
		Help:     "Single comment string to sent to GitHub.",
	})
	csm := parser.Flag("c", "codify", &argparse.Options{
		Required: false,
		Help:     "Put comments to the Markdown code block.",
	})
	fList := parser.StringList("f", "file", &argparse.Options{
		Required: false,
		Help:     "By repeating this flag you can specify multiple files which content will be sent to comment.",
	})
	fCmt := parser.StringList("l", "file-comment", &argparse.Options{
		Required: false,
		Help:     "By repeating this flag you can specify comments for provided files in according order.",
	})
	isn := parser.Int("i", "issue-number", &argparse.Options{
		Required: true,
		Help:     "Number(id) of the Issue/PR to sent a comment.",
	})
	mcm := parser.Flag("m", "multi-comment", &argparse.Options{
		Required: false,
		Default:  false,
		Help:     "Put each file into a separate comment in GitHub.",
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
	data.IssueNumber = *isn
	data.MultiCommentMode = *mcm
	return data
}

func ReadConfig() Data {
	data := Data{}
	data = *readArgConfig(&data)
	data = *readEnvConfig(&data)
	return data
}
