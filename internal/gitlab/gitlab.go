package gitlab

import (
	gogitlab "github.com/xanzy/go-gitlab"
	"log"
)

func createGitlabClient(apiToken string) *gogitlab.Client {
	git, err := gogitlab.NewClient(apiToken)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return git
}

func makePid(owner string, repo string) string {
	return owner + "/" + repo
}

func CommentIssue(comment *string, apiToken string, owner string, repo string, issue int) {
	git := createGitlabClient(apiToken)
	cmtOpts := gogitlab.CreateIssueNoteOptions{Body: comment}
	pid := makePid(owner, repo)
	_, _, err := git.Notes.CreateIssueNote(pid, issue, &cmtOpts)
	if err != nil {
		log.Fatalf("Failed to create Issue comment: %v", err)
	}
}

func CommentMergeRequest(comment *string, apiToken string, owner string, repo string, mr int) {
	git := createGitlabClient(apiToken)
	cmtOpts := gogitlab.CreateMergeRequestNoteOptions{Body: comment}
	pid := makePid(owner, repo)
	_, _, err := git.Notes.CreateMergeRequestNote(pid, mr, &cmtOpts)
	if err != nil {
		log.Fatalf("Failed to create Merge Request comment: %v", err)
	}
}
