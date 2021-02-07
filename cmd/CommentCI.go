package main

import (
	cmt "github.com/ThelonKarrde/CommentCI/pkg/comments"
	"github.com/ThelonKarrde/CommentCI/pkg/config"
	ghi "github.com/ThelonKarrde/CommentCI/pkg/github"
	"github.com/ThelonKarrde/CommentCI/pkg/utils"
	"log"
)

func main() {
	data := config.ReadConfig()
	if data.CommentText != "" {
		if data.CommentFiles != nil {
			log.Println("Warning! Both comment and file-comments args are specified! Priority over single comment flag.")
		}
		ghi.CommentIssue(&data.CommentText, data.GitHubCommentToken, data.GitHubRepoOwner, data.GitHubRepoName, data.IssueNumber)
	} else {
		if data.MultiCommentMode == true {
			if data.FileList != nil {
				comment := cmt.MakeSingleComment(utils.FilesToStrings(data.FileList), data.CommentFiles, data.CodeStyleMode)
				ghi.CommentIssue(&comment, data.GitHubCommentToken, data.GitHubRepoOwner, data.GitHubRepoName, data.IssueNumber)
			} else {
				log.Fatalf("No files specified!")
			}
		}
	}
}
