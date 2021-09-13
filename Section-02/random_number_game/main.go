package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// define our game parameters
	var tries int = 5     // amount of tries before Game Over
	var maxGuess int = 10 // number the user has to guess
	var guesses []int     // array of wrong answers tried

	// Seed random using unix nanosecond timestamp and initialize game
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(maxGuess + 1)
	initGame(tries, target)

	// Loop repeatedly until tries reaches 0
	for guess := 0; len(guesses) < tries; {
		// guess = rand.Intn(target + 1) // logic is currently backwords to simulate a user guessing
		fmt.Printf("Enter Your Guess (number can be between 0 and %d): ", maxGuess)
		fmt.Scanln(&guess)

		if checkNum(guess, target) {
			guesses = append(guesses, guess) // add final guess to array

			// Game Over - Victory
			gameOver(true, guesses, target)

			return
		} else {
			fmt.Println("\nIncorrect answer! Please try again")
			fmt.Printf("Number of tries left: %d\n\n", tries-len(guesses))
			guesses = append(guesses, guess)
		}
	}
	// Game Over - Defeat
	gameOver(false, guesses, target)

}

func checkNum(guess int, target int) bool {
	if guess == target {
		return (true)
	} else {
		return (false)
	}
}

func gameOver(win bool, guesses []int, target int) {
	if win {
		fmt.Printf("\nYou Win!")

		if len(guesses) == 1 {
			fmt.Printf("\nGuessed on first try! What a BEAST!\n")
		}
	} else {
		fmt.Println("GAME OVER!!!")
	}
	fmt.Printf("\nGuesses: %d\n", guesses)
	fmt.Printf("Target was: %d\n", target)
	fmt.Printf("Number of Attempts: %d\n", len(guesses))
}

func initGame(tries int, target int) {
	fmt.Printf("You have %d tries to guess the number!\n", tries)
	fmt.Printf("Target: %d\n", target)
}
