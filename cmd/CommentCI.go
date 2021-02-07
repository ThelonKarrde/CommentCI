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
		if data.FileList != nil {
			log.Println("Warning! Both single-comment and file-comment args are specified! Priority over single comment flag.")
		}
		ghi.CommentIssue(&data.CommentText, data.GitHubCommentToken, data.GitHubRepoOwner, data.GitHubRepoName, data.IssueNumber)
	} else {
		if data.MultiCommentMode == false {
			if data.FileList != nil {
				comment := cmt.MakeSingleComment(utils.ConvertFilesToStrings(data.FileList), data.CommentFiles, data.CodeStyleMode)
				ghi.CommentIssue(&comment, data.GitHubCommentToken, data.GitHubRepoOwner, data.GitHubRepoName, data.IssueNumber)
			} else {
				log.Fatalf("No files specified!")
			}
		} else {
			for i, p := range data.FileList {
				var comment string
				if i >= len(data.CommentFiles) {
					comment = cmt.MakeComment(utils.ReadFileToString(p), "", data.CodeStyleMode)
				} else {
					comment = cmt.MakeComment(utils.ReadFileToString(p), data.CommentFiles[i], data.CodeStyleMode)
				}
				ghi.CommentIssue(&comment, data.GitHubCommentToken, data.GitHubRepoOwner, data.GitHubRepoName, data.IssueNumber)
			}
		}
	}
}
