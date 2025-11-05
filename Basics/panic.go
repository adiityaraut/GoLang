package basics

import "fmt"

func main() {

	//Panic is a built-in function that stops normal execution of a function immediately.When a function encounters a panic, it stops executing and begins to unwind the stack, running any deferred functions along the way.
	//Panic is typically used to indicate a serious error that the program cannot recover from, such as an out-of-bounds array access or a nil pointer dereference.
	//When a panic occurs, the program prints a stack trace to the standard error output, which can help diagnose the cause of the panic.
	//After all deferred functions have been executed, the program terminates with a non-zero exit code.
	//Panic can be triggered explicitly by calling the panic function, or it can occur implicitly when certain runtime errors are encountered.
	//The syn\tax of panic function is called with an optional argument of any type, which represents the value associated with the panic.
	//Anything after panic will not be reachable to the runtime.
	process(-1)
}

func process(input int) {
	fmt.Println("Deferred 1")
	fmt.Println("Deferred 2")
	if input < 0 {
		panic("Negative input not allowed")
	}
	fmt.Println("Processing input:", input)
}
