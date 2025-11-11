package intermediate

import "fmt"

func main() {
	/*
		-> A Pointer is a variable that stores the memory address of another variable.
		-> USE CASE
		  - Modify the value of a variable indirectly
		  - Pass large data structures efficiently between functions
		  - Manage memory directly for performance reasons
		-> Pointer Declaration And Initialisation
			-> Declaration Syntax:
				var ptr *int
				'ptr' is a pointer to an integer('*int')

			->Initialization
				var a int = 10
				ptr = &a //ptr now points to a's memory address

		->Pointer Operations: Limited to referencing('&') and dereferencing('*')
		-> Nil Pointers
		-> Go does not support pointer arithmetic like c/cpp
		-> Passing Pointers to Functions
		-> Pointers to Structs
		-> Use Pointers when a function needs to modify an arguments value
		-> 'Unsafe.Pointer(&x)' converts the address of 'x' to an 'unsafe.Pointer'
	*/
	var ptr *int
	var a int = 10
	ptr = &a //Referencing

	fmt.Println(a)
	fmt.Println(ptr)
	fmt.Println(*ptr) //Dereferencing

	modifyValue(ptr)
	fmt.Println(a)

}

func modifyValue(ptr *int) {
	*ptr++
}
