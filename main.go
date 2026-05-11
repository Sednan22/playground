package main

import "fmt"

func main() {

	err := FetchWebSite()
	if err != nil {
		fmt.Printf("Error in fetchWebSite: %v", err)
		return
	}

	fmt.Println("Done")
}
