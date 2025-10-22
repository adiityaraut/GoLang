package basics

import (
	"fmt"
)

/*
=============================
      SLICES IN GO
=============================

DEFINITION:
A slice in Go is a dynamically-sized, flexible view into the elements of an array.
Unlike arrays, slices do not store data directly — they are *descriptors* that reference
an underlying array.

SLICE HEADER CONTAINS:
1. A pointer to the first element in the underlying array (that the slice can access)
2. The length (len) — number of accessible elements
3. The capacity (cap) — number of elements between the first element of the slice
   and the end of the underlying array.

Syntax to declare a slice:
    var sliceName []Type
Example:
    var nums []int  // Declares a nil slice of integers
*/

func main() {

	/*
		=============================
		   NIL vs EMPTY SLICE
		=============================
		- A nil slice has no underlying array (default zero value of []T)
		- An empty slice has length 0 but points to an existing (zero-length) array
	*/

	var nilSlice []int // nil slice (not yet initialized)
	fmt.Printf("nilSlice == nil : %v, len=%d, cap=%d\n", nilSlice == nil, len(nilSlice), cap(nilSlice))

	emptySlice1 := []int{} // empty literal slice
	fmt.Printf("emptySlice1 == nil : %v, len=%d, cap=%d\n", emptySlice1 == nil, len(emptySlice1), cap(emptySlice1))

	emptySlice2 := make([]int, 0) // empty slice using make()
	fmt.Printf("emptySlice2 == nil : %v, len=%d, cap=%d\n\n", emptySlice2 == nil, len(emptySlice2), cap(emptySlice2))

	/*
		=============================
		   CREATING SLICES
		=============================
		1. From an array
		2. Using make()
		3. From another slice
	*/

	// 1️.Creating from an array
	array := [6]int{10, 20, 30, 40, 50, 60}
	sliceA := array[1:4] // elements from index 1 to 3 (20, 30, 40)
	fmt.Printf("sliceA from array: %v, len=%d, cap=%d\n", sliceA, len(sliceA), cap(sliceA))
	// cap = len(array) - startIndex = 6 - 1 = 5

	// 2️.Creating with make(length)
	sliceB := make([]int, 5) // len=5, cap=5 (default capacity = length)
	fmt.Printf("sliceB (make len=5): len=%d, cap=%d\n", len(sliceB), cap(sliceB))

	// 3️.Creating with make(length, capacity)
	sliceC := make([]int, 3, 10) // len=3, cap=10
	fmt.Printf("sliceC (make len=3, cap=10): len=%d, cap=%d\n\n", len(sliceC), cap(sliceC))

	/*
		=============================
		   LENGTH vs CAPACITY
		=============================
		- len(s): current number of elements in slice
		- cap(s): total number of elements available in underlying array from start index
	*/

	s := []int{1, 2, 3, 4}
	fmt.Printf("s=%v, len=%d, cap=%d\n\n", s, len(s), cap(s))

	/*
		=============================
		   APPENDING ELEMENTS
		=============================
		- append() adds elements at the end of a slice.
		- If capacity is enough, the new elements are added in-place.
		- If capacity is insufficient, Go allocates a new underlying array (capacity usually doubles)
		  and returns a new slice referencing it.
	*/

	appendSlice := []int{}
	fmt.Printf("Initial appendSlice: len=%d cap=%d\n", len(appendSlice), cap(appendSlice))
	for i := 1; i <= 8; i++ {
		appendSlice = append(appendSlice, i)
		fmt.Printf("After append %d → len=%d cap=%d\n", i, len(appendSlice), cap(appendSlice))
	}
	fmt.Println()

	/*
		=============================
		   RESLICING
		=============================
		Reslicing means creating a new slice from an existing one.
		Syntax: newSlice := oldSlice[start:end]

		Rules:
		- start and end must be within the capacity of the slice (0 <= start <= end <= cap(oldSlice))
		- len(newSlice) = end - start
		- cap(newSlice) = cap(oldSlice) - start
	*/

	reslice := []int{10, 20, 30, 40, 50}
	fmt.Println("Original slice:", reslice)
	sub1 := reslice[1:4] // 20, 30, 40
	fmt.Printf("Reslice sub1=%v len=%d cap=%d\n", sub1, len(sub1), cap(sub1))

	sub2 := reslice[:3] // 10, 20, 30
	fmt.Printf("Reslice sub2=%v len=%d cap=%d\n\n", sub2, len(sub2), cap(sub2))

	/*
		=============================
		   UNDERLYING ARRAY SHARING
		=============================
		Multiple slices can share the same underlying array.
		Changing one slice may affect another if they overlap in memory.
	*/

	shared := []int{100, 200, 300, 400, 500}
	part1 := shared[1:4] // [200, 300, 400]
	part2 := shared[2:]  // [300, 400, 500]
	part1[1] = 999       // modifies shared[2]
	fmt.Println("After modifying part1[1]:")
	fmt.Println("shared array:", shared)
	fmt.Println("part1:", part1)
	fmt.Println("part2:", part2, "\n")

	/*
		=============================
		   COPY FUNCTION
		=============================
		copy(dst, src) copies min(len(dst), len(src)) elements from src to dst.
		Returns number of elements copied.
	*/

	source := []int{1, 2, 3, 4}
	dest := make([]int, 2) // destination shorter than source
	copied := copy(dest, source)
	fmt.Printf("Source=%v, Dest=%v, Elements Copied=%d\n\n", source, dest, copied)

	/*
		=============================
		   ITERATING OVER SLICES
		=============================
		We can iterate over slices using:
		- Classic for loop
		- Range loop
	*/

	rangeSlice := []int{10, 20, 30}
	fmt.Println("Iteration using range:")
	for i, v := range rangeSlice {
		fmt.Printf("Index=%d Value=%d\n", i, v)
	}
	fmt.Println()

	/*
		=============================
		   MULTIDIMENSIONAL SLICES
		=============================
		A 2D slice is a slice of slices (rows can have different lengths).
	*/

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println("matrix:", matrix)
	fmt.Println("matrix[1][2] =", matrix[1][2])

	// Creating 2D slice using make()
	rows, cols := 3, 4
	twoD := make([][]int, rows)
	for i := range twoD {
		twoD[i] = make([]int, cols)
	}
	twoD[0][0] = 99
	fmt.Printf("twoD[0][0] = %d\n\n", twoD[0][0])

	/*
		=============================
		   COMPARING SLICES
		=============================
		- You cannot directly compare two slices using == (except with nil).
		- To check equality, compare elements manually or use reflect.DeepEqual().
	*/

	sliceX := []int{1, 2, 3}
	sliceY := []int{1, 2, 3}
	fmt.Println("sliceX == nil:", sliceX == nil)
	fmt.Println("sliceY == nil:", sliceY == nil)
	// The below line would cause compile error:
	// fmt.Println(sliceX == sliceY)
	fmt.Println("To compare slices use reflect.DeepEqual(sliceX, sliceY)\n")

	/*
		=============================
		   COPYING vs SHARING
		=============================
		- Use copy() when you want to duplicate slice data to avoid modifying the original.
		- Simple assignment like slice2 = slice1 makes both share the same underlying array.
	*/

	original := []int{1, 2, 3}
	ref := original // shares underlying array
	copyVer := make([]int, len(original))
	copy(copyVer, original) // makes independent copy
	ref[0] = 99
	fmt.Println("original:", original)
	fmt.Println("ref (shared):", ref)
	fmt.Println("copyVer (independent):", copyVer, "\n")

	/*
		=============================
		   APPEND BEHAVIOR WITH SHARED SLICES
		=============================
		If two slices share the same array and one of them appends beyond its capacity,
		it will allocate a new array, breaking the link.
	*/

	base := []int{1, 2, 3}
	alias := base
	base = append(base, 4, 5, 6, 7)
	base[0] = 100
	fmt.Println("base after append:", base)
	fmt.Println("alias (original):", alias, "\n")

	/*
			=============================
			   BEST PRACTICES & NOTES
			=============================
			Use slices instead of arrays for flexibility.
		    Remember: appending may reallocate underlying array.
			Use copy() when you need independent data.
		    Be cautious when multiple slices share the same array.
			len() and cap() are O(1) operations.
			Avoid comparing slices directly (use reflect.DeepEqual if needed).
			nil slice and empty slice behave the same in range loops and len() checks.
	*/

}
