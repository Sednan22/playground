package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	start := time.Now()
	loop := 0
	randNumber := int64(10000)
	randomness := int64(0)
	for {
		randNumber += rand.Int63n(10)
		randomness += rand.Int63n(10)

		if randomness > randNumber {
			fmt.Printf("RandNumber number is: %d\n", randNumber)
			fmt.Printf("DoNothing number is: %d\n", randomness)
			fmt.Printf("It took %v and %d loops\n", time.Since(start), loop)
			break
		}
		loop++
	}
}
