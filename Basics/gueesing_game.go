package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano()) // Seed the random number generator with the current time in nanoseconds
	random := rand.New(source)                      // Create a new random number generator

	//Generate a random number between 1 and 100
	target := random.Intn(100) + 1

	//Welcome message
	fmt.Println("Welcome to the Guessing Game!")
	fmt.Println("I have selected a random number between 1 and 100.")
	fmt.Println("Can you guess what it is?")
	var guess int
	for {
		fmt.Print("Enter your guess: ")
		//&guess is the address of the guess variable where the input will be stored as normal guess would just pass the value
		fmt.Scanln(&guess)

		//Check the guess
		if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Println("Congratulations! You guessed the correct number:", target)
			break // Exit the loop when the correct guess is made
		}
	}
}
