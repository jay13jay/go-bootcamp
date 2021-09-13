package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	target := 10
	var count int
	// u_time := time.Now().UnixNano()
	rand.Seed(time.Now().UnixNano())

	fmt.Printf("Target: %d\nGuesses: ", target)
	for n := 0; n != target; {
		n = rand.Intn(target + 1)
		if n == target {
			fmt.Printf("\nGuess is correct!\n")
			fmt.Printf("Answer took %d tries\n", count+1)
		} else {
			fmt.Printf("%d ", n)
			count++
		}
	}
}
