package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {

	accountID := flag.Int("account", 0, "account id")
	flag.Parse()

	accountLoaded := loadConfig(*accountID)
	if *accountID == 0 {
		log.Fatal("--account=x needed")
	}

	fmt.Println(accountLoaded)
}

func loadConfig(accountID int) string {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error: %v", err)
	}

	if accountID == 1 {
		return os.Getenv("TEST1")
	}
	if accountID == 2 {
		return os.Getenv("TEST2")
	}
	if accountID == 3 {
		return os.Getenv("TEST3")
	}

	return "0"
}
