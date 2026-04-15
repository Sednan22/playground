package website

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

const cacheFileName = "website/cache_issues.json"

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

	CacheIssues(issues)

	return issues, nil
}

func CacheIssues(issues []Issue) {

	dat, err := json.MarshalIndent(issues, "", "  ")
	if err != nil {
		fmt.Printf("Error marshal issues: %v", err)
		return
	}

	err = os.WriteFile(cacheFileName, dat, 0666)
	if err != nil {
		fmt.Printf("Error writing to file: %v", err)
	}
}

func ReadCacheIssues() ([]Issue, error) {

	dat, err := os.ReadFile(cacheFileName)
	if err != nil {
		return nil, err
	}

	var issues []Issue
	err = json.Unmarshal(dat, &issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}
