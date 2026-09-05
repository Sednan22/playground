package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var loop, count int
	percentage := int64(50)

	for count < 10 {

		n := rand.Int63n(100)

		if n <= percentage {
			count++
		} else {
			count = 0
		}
		loop++
	}

	fmt.Println("Loops:", loop)

}
