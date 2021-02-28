package main

import (
	"github.com/ThelonKarrde/CommentCI/internal/config"
	"github.com/ThelonKarrde/CommentCI/internal/format"
	ghi "github.com/ThelonKarrde/CommentCI/internal/github"
	glb "github.com/ThelonKarrde/CommentCI/internal/gitlab"
	"github.com/ThelonKarrde/CommentCI/internal/utils"
	"log"
)

func main() {
	cfg := config.ReadConfig()
	var comments []string
	if cfg.CommentText != "" {
		comments = append(comments, format.Comment(cfg.CommentText, "", cfg.CodeStyleMode))
	} else {
		if cfg.MultiCommentMode {
			for i, p := range cfg.FileList {
				cmt := ""
				if i < len(cfg.CommentFiles) {
					cmt = cfg.CommentFiles[i]
				}
				comments = append(comments, format.Comment(utils.ReadFileToString(p), cmt, cfg.CodeStyleMode))
			}
		} else {
			if len(cfg.FileList) > 0 {
				comments = append(comments, format.SingleComment(utils.ConvertFilesToStrings(cfg.FileList), cfg.CommentFiles, cfg.CodeStyleMode))
			} else {
				log.Fatalf("No files specified!")
			}
		}
	}
	for _, c := range comments {
		if cfg.Platform == "github" {
			ghi.Comment(&c, cfg.ApiToken, cfg.RepoOwner, cfg.RepoName, cfg.IssueNumber)
		} else {
			glb.Comment(cfg.TargetType, &c, cfg.ApiToken, cfg.RepoOwner, cfg.RepoName, cfg.IssueNumber)
		}
	}
}
