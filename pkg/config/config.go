package config

import (
	"fmt"
	"log"
	"os"
)

type Data struct {
	GitHubRepoOwner string
	GitHubRepoName string
	GitHubCommentUser string
	GitHubCommentToken string
}

func ReadEnvConfig() {
	ghro, ok := os.LookupEnv("GITHUB_REPO_OWNER")
	if !ok {
		log.Printf("No GITHUB_REPO_OWNER value")
	}
	fmt.Println(ghro)
}