package basics

import "fmt"

// Arrays in Go are fixed-size collections of elements of the same type.
func main() {
	//var arrayName [size]dataType
	var numbers [5]int // Declare an array of 5 integers
	fmt.Println("Numbers:", numbers)
	//In go when you assign an array to new variable or pass it as argument to function a copy of the array is made and modifications to the new variable or parameter do not affect the original array.
	originalArray := [3]int{1, 2, 3}
	copiedArray := originalArray                  // A copy of originalArray is made
	copiedArray[0] = 10                           // Modify the copy
	fmt.Println("Original Array:", originalArray) // Output: [1 2 3]
	fmt.Println("Copied Array:", copiedArray)     // Output: [10 2 3]
	numbers[0] = 10
	fmt.Println(numbers)                                // Output: [0 0 0 0 0] (default zero values)
	var fruits = [3]string{"Apple", "Banana", "Cherry"} // Declare and initialize an array of strings
	fmt.Println("Fruits:", fruits)

	for _, v := range numbers { // _ here is a blank identifier used to ignore the index value
		fmt.Printf("Value %d", v)
		fmt.Println()
	}
	n := 2
	_ = n // To avoid unused variable error
	//length of an array is fixed and can be obtained using len() function determined at compile time
	// Go also supports multidimensional arrays
	var matrix [2][3]int // 2 rows and 3 columns
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3
	matrix[1][0] = 4
	matrix[1][1] = 5
	matrix[1][2] = 6
	fmt.Println("Matrix:", matrix)
}
