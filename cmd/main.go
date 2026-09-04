package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Account struct {
	id   int
	name string
	api  string
}

func main() {

	accountID := flag.Int("account", 0, "account id")
	flag.Parse()

	if *accountID == 0 {
		log.Fatalf("usage: %s --account=x", os.Args[0])
	}

	accountLoaded := loadAccount(*accountID)
	if accountLoaded.api == "" {
		log.Fatal("failed loading account or account doesn't exist")
	}

	fmt.Println("ID:", accountLoaded.id, "| Name:", accountLoaded.name)

}

func loadAccount(accountID int) Account {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error: %v", err)
	}

	return Account{
		id:   accountID,
		name: fmt.Sprintf("test%d", accountID),
		api:  os.Getenv(fmt.Sprintf("TEST%d", accountID)),
	}
}
