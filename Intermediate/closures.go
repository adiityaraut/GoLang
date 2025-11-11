package intermediate

import "fmt"

func main() {
	//Clodures are powerful concept that allows functions to capture and manipute variables from their surrounding lexical scope.
	//A closure is a function value that references variables from outside its body.The function may access and assign to the referenced variables; in this sense, the function is "bound" to the variables.These variables are called free variables because they are not defined within the function but are used by it and these variables persist as long as the closure itself is referenced .
	//Closures are often used to create function factories, implement data hiding, and manage state in a functional way.
	//Closures work with lexical scoping, meaning they capture variables based on their position in the source code, not based on the call stack.Thid allows closures to retain access to variables even after the outer function has completed execution.
	/*
		In Go, functions are first-class citizens, meaning they can be assigned to variables, passed as arguments to other functions, and returned from functions.Closures leverage this feature by allowing functions to capture and retain access to variables from their surrounding scope.
	*/
	/*
		i:=0
		return func() int {
			i++
			return i
		}
	*/
	//Added func returns a variable to sequence
	sequence := adder()
	//We see value is incrementing and its stored in the form of state which is incrementing and i is captured by closure its value is persisting.
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())
	fmt.Println(sequence())

	//Creating a new instance of adder function
	newSequence := adder()
	fmt.Println(newSequence())
	fmt.Println(newSequence())

	subtraction := func() func(int) int {
		countDown := 10
		return func(x int) int {
			countDown -= x
			return countDown
		}
	}()
	fmt.Println(subtraction(1))
	fmt.Println(subtraction(2))
	fmt.Println(subtraction(3))

}

// Func name is added which returns another function that returns an integer.
func adder() func() int {
	//gofer captures the variable i from its surrounding lexical scope.
	i := 0
	fmt.Println("Previous Value of i", i)
	return func() int {
		i++
		fmt.Println("Current Value of i", i)
		return i
	}
}

//Func returned is an anonymous function which is a function without a name.

/*
	PRACTICAL USE CASES
	 - Stateful Functions
	 - Encapsulation
	 - Callback

	USEFUL OF CLOSURES
	 - Encapsulation
	 - Flexibility
	 - Readability

	 CONSIDERATIONS
	 - Memory usage(CLosures can keep variables alive longer than expected if they hold references to large objects or resources)
	 - Concurency(Race conds and unncessary)

	 BEST PRACTICES
	 - Limit Scope
	 - Avoid Overuse
*/
