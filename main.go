package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Sednan22/tests/website"
)


func main() {

	req, err := http.NewRequest("GET", "https://api.github.com/repos/golang/go/issues", nil)
	if err != nil {
		fmt.Println("Error making request: %v\n", err)
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error getting response: %v\n", err)
		return
	}
	defer res.Body.Close()

	var issues []website.Issue
	err = json.NewDecoder(res.Body).Decode(&issues)
	if err != nil {
		fmt.Println("ERROR reading body: %v\n", err)
		return
	}

	for i, issue := range issues {
		fmt.Printf("Issue: %d - %s - Author: %s\n", i, issue.Title, issue.User.Login)
	}

	fmt.Println("\nDONE...\n")
}


/*
// FOR LOOP TO FIND HOW DATA IS COMING FROM GET REQUEST | GOOD TO BUILD STRUCT
for _, obj := range test {
		for key, val := range obj {
			if key == "user" {
				userMap, ok := val.(map[string]interface{})
				if ok {
					for uKey, uVal := range userMap {
						fmt.Printf("UserKey: %-15s | Val: %-20T\n", uKey, uVal)
					}
				} 
			} else {
				fmt.Printf("Key: %-15s | Val: %-20T\n", key, val)
			}
		}
	}
*/