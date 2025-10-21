package basics

import "fmt"

var firstname string = "Aditya" // Package-level variable declaration with type inference
func main() {
	var age int // Uninitialized int variable, defaults to 0
	var name string = "Aditya"
	count := 10 // Short variable declaration, type inferred as int (Type Inference)

	//VAriable in go have block scope
	fmt.Println("Name:", firstname)
}
