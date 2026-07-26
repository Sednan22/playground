package main

import (
	"fmt"
)

func main() {

	err := redisTest()
	if err != nil {
		fmt.Print(err)
	}

}
