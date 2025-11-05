package basics

import "fmt"

func main() {
	//In go defer is a mechanism that allows you to postpone the execution of a function until the surrounding function returns.
	//The deferred function calls are pushed onto a stack and executed in last-in-first-out (LIFO) order when the surrounding function completes.
	//This is particularly useful for resource management tasks such as closing files, releasing locks, or cleaning up resources, ensuring that these actions are performed regardless of how the function exits (whether normally or due to an error).
	//Goroutine will execute the deferred functions even if a panic occurs in the surrounding function.Goroutines are basically lightweight threads managed by the Go runtime.These run in background and are used to perform concurrent tasks.
	//Defer function is followed by a function or method call.The function is evaluated immediately, but its execution is deferred until the surrounding function returns.
	//Defer function is always part of surrounding function and cannot be used outside of it.The surrounding function means the function in which the defer statement is present.
	//We can have multiple defer statements in a single function.The deferred functions will be executed in reverse order of their appearance (LIFO order) when the surrounding function returns.

	process(10)
}

func process(i int) {
	defer fmt.Println("Deferred value of i:", i)
	defer fmt.Println("First Deferred statement executed")
	defer fmt.Println("Second deferred statement executed")
	defer fmt.Println("Third Deferred statement executed")
	fmt.Println("Normal statement executed")
	i++
	fmt.Println("Value of i in process function:", i)
	//Practical use cases of defer:
	//1. Resource Management: Defer is commonly used to ensure that resources such as files, network connections, or database connections are properly closed after their use.
	//2. Mutex Unlocking: When working with concurrent programming and mutexes, defer can be used to ensure that a mutex is always unlocked after it has been locked, even if an error occurs.
	//3. Logging and Tracing: Defer can be used to log function entry and exit points, making it easier to trace the flow of execution in complex applications.
	//4. Error Handling: Defer can be used to handle cleanup tasks in case of errors, ensuring that necessary actions are taken before exiting a function.
	//Best Practices:
	//1.Keep deferred actions short
	//2.Understand evaluation timing
	//3.Avoid complex control flow
}
