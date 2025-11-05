package basics

import "fmt"

func main() {
	// func <name> (parameter list) return type {
	// function body
	// return value
	// }
	fmt.Println(add(5, 10))

	greet := func() {
		fmt.Println("This is an anonymous function.")
	}
	greet()
	// Function as First-Class Citizens that is functions can be assigned to variables, passed as arguments to other functions, and returned from other functions.First claas objects are functions that can be treated like any other variable.
	operation := add
	result := operation(7, 3)
	fmt.Println("Result of operation:", result)
	double := createMultiplier(2) // captures factor = 2
	triple := createMultiplier(3) // captures factor = 3

	fmt.Println(double(5))  // 10
	fmt.Println(triple(5))  // 15
	fmt.Println(double(10)) // 20

}

func add(a int, b int) int {
	return a + b
}

//Another type of function is anonymous function, which is defined without a name and can be assigned to a variable or passed as an argument to other functions.Its also called closure when it captures variables from its surrounding scope.Its also called funcion literal.

// Go also supports higher-order functions, which are functions that can take other functions as arguments or return them as results.

//Function that takes a function as an argument
func applyOperation(x int, y int, operation func(int, int) int) int {
	return operation(x, y)
}

// Function that returns another function
func createMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// In Go, functions are first-class citizens, meaning they can:

// be passed as arguments,

// be returned from other functions,

// be assigned to variables.

//Multiple return values: Go functions can return multiple values, which is useful for returning both a result and an error status.

//Named return values: Go allows you to name the return values of a function, which can improve code readability.
