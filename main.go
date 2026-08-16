package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env")
	}

	cfgTokens, err := NewConfig()
	if err != nil {
		fmt.Println("error getting new config for jwt tokens")
		return
	}

	// using a new create token
	token, err := CreateToken(cfgTokens, "1")
	if err != nil {
		fmt.Println("error creating new access token")
		return
	}

	// customToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWUsImlhdCI6MTUxNjIzOTAyMn0.KMUFsIDTnFmyG3nMiGM6H9FNFUROf3wh7SmqJp-QV30"
	_, err = ValidateToken(token, cfgTokens.TokenSecret)
	if err != nil {
		fmt.Println("error validating token", err)
		return
	}

	fmt.Println("valid token")

}
