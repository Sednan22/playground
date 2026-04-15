package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	
	type responseParams struct {
		Title string	`json:"title"`
		Number int		`json:"number"`
	}


	req, err := http.NewRequest("GET", "https://api.github.com/repos/golang/go/issues", nil)
	if err != nil {
		fmt.Println("Error making request")
		return
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Error getting response")
		return
	}

	defer res.Body.Close()

	var test []responseParams
	err = json.NewDecoder(res.Body).Decode(&test)
	if err != nil {
		fmt.Println("ERROR reading body")
		return
	}

	for _, obj := range test {
		fmt.Println(obj.Title)
	}

	fmt.Println("\nDONE...\n")

}