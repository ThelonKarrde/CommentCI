package main

import (
	cmt "github.com/ThelonKarrde/CommentCI/internal/comments"
	"github.com/ThelonKarrde/CommentCI/internal/config"
	ghi "github.com/ThelonKarrde/CommentCI/internal/github"
	glb "github.com/ThelonKarrde/CommentCI/internal/gitlab"
	"github.com/ThelonKarrde/CommentCI/internal/utils"
	"log"
)

func main() {
	data := config.ReadConfig()
	if data.CommentText != "" {
		if len(data.FileList) != 0 {
			log.Println("Warning! Both single-comment and file-comment args are specified! Priority over single comment flag.")
		}
		comment := cmt.MakeComment(data.CommentText, "", data.CodeStyleMode)
		if data.Platform == "github" {
			ghi.CommentIssue(&comment, data.ApiToken, data.RepoOwner, data.RepoName, data.IssueNumber)
		} else {
			if data.TargetType == "issue" {
				glb.CommentIssue(&comment, data.ApiToken, data.RepoOwner, data.RepoName, data.IssueNumber)
			}
			if data.TargetType == "merge-request" {
				glb.CommentMergeRequest(&comment, data.ApiToken, data.RepoOwner, data.RepoName, data.IssueNumber)
			} else {
				log.Fatalf("No target type sepcified for GitLab mode!")
			}
		}
	} else {
		if data.MultiCommentMode == false {
			if data.FileList != nil {
				comment := cmt.MakeSingleComment(utils.ConvertFilesToStrings(data.FileList), data.CommentFiles, data.CodeStyleMode)
				if data.Platform == "github" {
					ghi.CommentIssue(&comment, data.ApiToken, data.RepoOwner, data.RepoName, data.IssueNumber)
				} else {
					if data.TargetType == "issue" {
						glb.CommentIssue(&comment, data.ApiToken, data.RepoOwner, data.RepoName, data.IssueNumber)
					}
					if data.TargetType == "merge-request" {
						glb.CommentMergeRequest(&comment, data.ApiToken, data.RepoOwner, data.RepoName, data.IssueNumber)
					} else {
						log.Fatalf("No target type sepcified for GitLab mode!")
					}
				}
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
				if data.Platform == "github" {
					ghi.CommentIssue(&comment, data.ApiToken, data.RepoOwner, data.RepoName, data.IssueNumber)
				} else {
					if data.TargetType == "issue" {
						glb.CommentIssue(&comment, data.ApiToken, data.RepoOwner, data.RepoName, data.IssueNumber)
					}
					if data.TargetType == "merge-request" {
						glb.CommentMergeRequest(&comment, data.ApiToken, data.RepoOwner, data.RepoName, data.IssueNumber)
					} else {
						log.Fatalf("No target type sepcified for GitLab mode!")
					}
				}
			}
		}
	}
}
