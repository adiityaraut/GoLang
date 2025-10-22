package basics

import "fmt"

func main() {
	// Simple iteration over range
	for i := 1; i < 5; i++ {
		fmt.Println("Iteration:", i)
	}
	//Iterate over collections
	numbers := []int{10, 20, 30, 40, 50}
	for index, numbers := range numbers {
		//As we are using placeholders we need to use fmt.Printf because fmt.Println does not support formatted strings
		fmt.Printf("Index: %d, Value: %d\n", index, numbers)
	}
	//Break : breaks the loop and the control goes to the statement after the loop
	//Continue : skips the current iteration and goes to the next iteration of the loop
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue // Skip even numbers
		}
		fmt.Println("Odd Number:", i)
		if i == 7 {
			break // Exit loop when i is 7
		}
	}
	rows := 5
	for i := 1; i <= rows; i++ {
		//inner loop for spaces bfore stars
		for j := 1; j <= rows-i; j++ {
			fmt.Print(" ")
		}
		//inner loop for stars
		for k := 1; k <= (2*i - 1); k++ {
			fmt.Print("*")
		}
		fmt.Println() // Move to the next line after each row
	}
	//WITH GO 1.22
	for i := range 10 {
		fmt.Println(10 - i)
	}
}
