package main

import (
	"fmt"
	"math/rand"
	"time"
)


 func getGuess(secret int) int {
	var guess int
	attempts := 0
	
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)
		attempts++

		if guess < secret {
			fmt.Println("Too low! Try again.")
		} else if guess > secret {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Println("Correct!")
			fmt.Println("Attempts: ", attempts)
			return guess
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	for {
		fmt.Println("\n--- New Game Started ---")
		secret := rand.Intn(100) + 1
		fmt.Println("Welcome to the Number Guessing Game!")
		getGuess(secret)

		var choice string
		fmt.Print("Do you want to play again? (yes/no): ")
		fmt.Scan(&choice)
		if choice != "yes" {
			fmt.Println("Thanks for playing!")
			break
		}
	}
}