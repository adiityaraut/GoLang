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
}
func sum(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}
