package config

import (
	"github.com/akamensky/argparse"
	"log"
	"os"
)

type Data struct {
	RepoOwner        string
	RepoName         string
	ApiUser          string
	ApiToken         string
	Platform         string
	TargetType       string
	CommentText      string
	CommentFiles     []string
	FileList         []string
	CodeStyleMode    bool
	IssueNumber      int
	MultiCommentMode bool
}

func readEnvConfig(data *Data) *Data {
	apiUser, ok := os.LookupEnv("API_USER")
	if !ok {
		apiUser, ok = os.LookupEnv("GITHUB_COMMENT_USER")
		if !ok {
			log.Println("Warning! No API user specified! Check API_USER env.")
		}
		log.Println("GITHUB_COMMENT_USER environment variable deprecated, use API_USER instead!")
	}
	data.ApiUser = apiUser
	apiToken, ok := os.LookupEnv("API_TOKEN")
	if !ok {
		apiToken, ok = os.LookupEnv("GITHUB_COMMENT_TOKEN")
		if !ok {
			log.Fatalf("No API token specified! Check API_TOKEN env.")
		}
		log.Println("GITHUB_COMMENT_TOKEN environment variable deprecated, use API_TOKEN instead!")
	}
	data.ApiToken = apiToken
	return data
}

func readArgConfig(data *Data) *Data {
	parser := argparse.NewParser("CommentCI", "Sent a comment to GitHub PR or Issue from your CI")
	ghro := parser.String("o", "owner", &argparse.Options{
		Required: true,
		Help:     "Owner of the repository. User/Organisations.",
	})
	ghrn := parser.String("r", "repository", &argparse.Options{
		Required: true,
		Help:     "Name of the repository.",
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
	platform := parser.Selector("p", "platform", []string{"github", "gitlab"}, &argparse.Options{
		Required: true,
		Help:     "Select platform where to send comments",
	})
	glbType := parser.Selector("g", "target-type", []string{"issue", "merge-request"}, &argparse.Options{
		Required: false,
		Help:     "Select type of comment target (GitLab only)",
	})
	err := parser.Parse(os.Args)
	if err != nil {
		log.Println(err.Error())
	}
	data.RepoOwner = *ghro
	data.RepoName = *ghrn
	data.CommentText = *cmt
	data.CodeStyleMode = *csm
	data.FileList = *fList
	data.CommentFiles = *fCmt
	data.IssueNumber = *isn
	data.MultiCommentMode = *mcm
	data.Platform = *platform
	data.TargetType = *glbType
	return data
}

func ReadConfig() Data {
	data := Data{}
	data = *readArgConfig(&data)
	data = *readEnvConfig(&data)
	return data
}
