package basics

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

	const p float64 = 22 / 7
	fmt.Printf("Approximation of Pi using integer division 22/7 = %f\n", p)
	//Overflow occurs when a calculation produces a result that exceeds the maximum limit of the data type.Overflow results in wrap-around behavior, leading to incorrect and unexpected values.
	var maxInt int64 = 9223372036854775807
	fmt.Println(maxInt)
	maxInt += 1
	fmt.Println("Overflowed Value:", maxInt)
}
