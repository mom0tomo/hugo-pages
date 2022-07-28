package main

import (
	"fmt"
	"log"
	"time"

	"github.com/go-git/go-git/v5"
)

func main() {
	r, err := git.PlainOpen(".")
	if err != nil {
		log.Fatal(err)
	}
	commits, err := r.Log(&git.LogOptions{})
	if err != nil {
		log.Fatal(err)
	}
	commit, err := commits.Next()
	if err != nil {
		log.Fatal(err)
	}

	lastCommitDate := commit.Author.When
	sinceLastCommit := time.Since(lastCommitDate)
	fmt.Printf("%d", int(sinceLastCommit.Hours()/24))
}
