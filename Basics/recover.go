package basics

import "fmt"

func main() {
	//Recover is a built-in function that allows you to regain control of a panicking goroutine.It is used in deferred functions to catch and handle panics, preventing the program from crashing.
	//When a panic occurs, the normal execution of the program is interrupted, and the control is transferred to the nearest deferred function that calls recover.
	//If recover is called within a deferred function and there is an active panic, it stops the panic and returns the value passed to panic.
	//If there is no active panic, recover returns nil.
	//By using recover in a deferred function, you can gracefully handle panics and allow the program to continue executing instead of terminating abruptly.
	process()
	fmt.Println("Program continues after recovery from panic")
}

func process() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
		//Dont be confused
		//r:= recover()
		//if r != nil {
	}()
	fmt.Println("Starting process")
	panic("Something went wrong!")
	fmt.Println("Process completed") // This line will not be executed
}
