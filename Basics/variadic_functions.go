package basics

import "fmt"

func main() {
	//Variadic functions in go allow you to create functions that can accept a variable number of arguments.
	//In go,variadic functions are defined using an ellipsis (...) before the type of the last parameter.
	//Difference between regular and variadic parameter is that variadic parameter can accept zero or more arguments of the specified type, whereas regular parameter accepts a fixed number of arguments.
	// func functionName(param1 type1, param2 type2, ... , paramN ... typeN) returnType {
	// 	// Code Block
	// 	return value
	// }

	fmt.Println("Sum of 1,2,3:", sum(1, 2, 3))
	fmt.Println("Sum of 10,20,30,40,50:", sum(10, 20, 30, 40, 50))
	fmt.Println("Sum of no numbers:", sum())
	result, statement := addz("Adding numbers:", 5, 10, 15)
	fmt.Println(statement, result)
}
func sum(nums ...int) int {
	total := 0
	//variadic parameters is like a slice inside the function
	for _, num := range nums {
		total += num
	}
	return total
}

func addz(returnString string, nums ...int) (string, int) {
	total := 0
	//variadic parameters is like a slice inside the function
	for _, num := range nums {
		total += num
	}
	return total
}

//We can use slices as variadic arguments by unpacking the slice using the ellipsis (...) operator when calling the function.For example:
// numbers := []int{1, 2, 3, 4, 5}
// result := sum(numbers...) // Unpacking the slice

// Difference between returnString and sequence is that returnString is a regular parameter of type string, while nums is a variadic parameter of type int.
// This means that when calling the addz function, you must provide a single string argument for returnString, followed by zero or more integer arguments for nums.
