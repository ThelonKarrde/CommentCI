package gitlab

import (
	gogitlab "github.com/xanzy/go-gitlab"
	"log"
)

func createClient(apiToken string) *gogitlab.Client {
	client, err := gogitlab.NewClient(apiToken)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return client
}

func makePid(owner string, repo string) string {
	return owner + "/" + repo
}

func Comment(cmType string, comment *string, apiToken string, owner string, repo string, id int) {
	git := createClient(apiToken)
	pid := makePid(owner, repo)
	var err error
	switch cmType {
	case "issue":
		cmtOpts := gogitlab.CreateIssueNoteOptions{Body: comment}
		_, _, err = git.Notes.CreateIssueNote(pid, id, &cmtOpts)
	case "merge-request":
		cmtOpts := gogitlab.CreateMergeRequestNoteOptions{Body: comment}
		_, _, err = git.Notes.CreateMergeRequestNote(pid, id, &cmtOpts)
	default:
		log.Fatalf("No target type sepcified for GitLab mode!")
	}
	if err != nil {
		log.Fatalf("Failt to create %s comment: %v", cmType, err)
	}
}
