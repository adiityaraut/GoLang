package main

import "fmt"

func main() {
	// Declare two integer variables
	var a, b int = 15, 4

	// Perform arithmetic operations
	fmt.Println("Arithmetic Operations in Go:")
	fmt.Println("-----------------------------")
	fmt.Printf("a = %d, b = %d\n\n", a, b)

	fmt.Printf("Addition: a + b = %d\n", a+b)
	fmt.Printf("Subtraction: a - b = %d\n", a-b)
	fmt.Printf("Multiplication: a * b = %d\n", a*b)
	fmt.Printf("Division: a / b = %d\n", a/b)
	fmt.Printf("Modulus (Remainder): a %% b = %d\n", a%b)
}
