package website

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Login string `json:"login"`
}

type Issue struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Locked bool   `json:"locked"`
	User   User   `json:"user"`
}

func FetchIssues(owner, repo string) ([]Issue, error) {

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API Error: %s", res.Status)
	}

	var issues []Issue
	err = json.NewDecoder(res.Body).Decode(&issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
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
