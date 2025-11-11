package intermediate

import "fmt"

func main() {
	//Recursion is the process of a function calling itself.It breaks down a problem into smaller subproblems of the same type until they become simple enough to solve directly.
	//In Every Recursive func there is a base case which the termination point of the recursion.Without base case, func would continue indefinitely leading to stack overflow.
	//Recursive solutions can be optimized by caching results of expensive function calls
	fmt.Println(factorial(10))
	fmt.Println(factorial(5))

	fmt.Println(sumOfIndex(9))
	fmt.Println(sumOfIndex(12))
	fmt.Println(sumOfIndex(12345))

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func sumOfIndex(n int) int {
	if n < 10 {
		return n
	}
	return n%10 + sumOfIndex(n/10)
}

/*
	PRACTICAL USE CASES
		- Mathematical Algo
		- Tree And Graph Traversal
		- Divide And Conquer

	Benefits Of Recursion
		- Simplicity
		- Clarity
		- Flexibility

	Considerations
		- Performance
		- Base Case

	Best Practices
		-Testing
		-Optimisations
		-Recursive call
*/
