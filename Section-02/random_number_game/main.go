package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	guess := 10
	var count int
	rand.Seed(time.Now().Unix())

	for n := 0; n != guess; {
		n = rand.Intn(guess + 1)
		if n == guess {
			fmt.Printf("Guess: %d\nTarget: %d\nGuess is correct!\n", guess, n)
			fmt.Printf("Answer took %d tries\n", count+1)
		} else {
			fmt.Printf("Guess: %d\nTarget: %d\nWrong answer... Try again\n", guess, n)
			count++
		}
	}
}
