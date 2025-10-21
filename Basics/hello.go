package basics // Package is a way to organise and reuse code in Go
//Inorder to run go code both package main and main function are required
// Import is used to include other packages
import "fmt" // Importing the fmt package for formatted I/O
//Main function is the entry point of the program
func main() {
	fmt.Println("Hello World") //Output text to the console
	//Go supports unicode natively to handle multiple languages
	// go build creates an persistent executable file
	// go run compiles and runs the code without creating an executable file
}

/*
Standard Library Packages:
- fmt: for formatted I/O
- math: for mathematical functions
This modular approach promotes clarity,efficiency,and maintainability in Go programs.
*/

/*
Though import allows all funcions and types from a package to be used,but during compilation only the used ones are included in the final executable,optimizing its size and performance.Ghost compiler removes unused code from the final binary.And linker combines the necessary parts into the executable.
Tree shaking is the process of removing unused code during compilation to optimize the final executable size.
*/
