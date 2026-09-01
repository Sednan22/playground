package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Account struct {
	Name string
	API string
}

func main() {

	accountID := flag.Int("account", 0, "account id")
	flag.Parse()
	
	if *accountID == 0 {
		log.Fatal("--account=x needed")
	}

	accountLoaded, err := loadConfig(*accountID)
	if err != nil {
		log.Fatal("accountLoaded failed")
	}
	
	fmt.Println(accountLoaded)
}

func loadConfig(accountID int) (Account, error) {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error: %v", err)
	}

	if accountID == 1 {
		return Account{
			Name: "test1",
			API: os.Getenv("TEST1")
		},nil
	}
	if accountID == 2 {
		return Account{
			Name: "test2",
			API: os.Getenv("TEST2")
		},nil
	}
	if accountID == 3 {
		return Account{
			Name: "test3",
			API: os.Getenv("TEST3")
		},nil
	}

	return Account{}, fmt.Error("account not found")
}
