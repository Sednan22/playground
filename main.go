package main

import (
	"fmt"

	"github.com/Sednan22/tests/website"
)

func main() {

	issues, err := website.FetchIssues("golang", "go")
	if err != nil {
		fmt.Printf("Error fetching issues: %v\n", err)
		return
	}

	for i, issue := range issues {
		fmt.Printf("Issue: %d - %s - Author: %s\n", i, issue.Title, issue.User.Login)
	}

	fmt.Println("\nDONE...")
}
