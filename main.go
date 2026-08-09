package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	opt, _ := redis.ParseURL(os.Getenv("REDIS_URL"))
	rdb := redis.NewClient(opt)
	defer rdb.Close()

	res2, err := rdb.JSONGet(ctx, "test", "$.foo.bar").Result()
	if err != nil {
		panic(err)
	}

	// fmt.Println(res2) // slice

	var result []string
	err = json.Unmarshal([]byte(res2), &result)
	if err != nil {
		panic(err)
	}

	if len(result) != 1 {
		fmt.Println("Got more or less than 1 val")
		return
	}

	val := result[0] // string

	// fmt.Println(val)

	if val == "poo" {
		fmt.Println("Works")
	}

}
