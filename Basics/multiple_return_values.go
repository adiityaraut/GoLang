package basics

import (
	"errors"
	"fmt"
)

func main() {

	// func functionName(parameter1 type1, parameter2 type2 ......) (returnType1, returnType2) {
	// 	Code Block
	// 	return value1, value2
	// }
	q, r := divide(10, 3)
	fmt.Println("Quotient:", q)
	fmt.Println("Remainder:", r)

	result, err := compare(5, 10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Comparison Result:", result)
	}

}
func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func compare(a, b int) (string, error) {
	if a > b {
		return "a is greater than b", nil
	} else if a < b {
		return "a is less than b", nil
	} else {
		return "", errors.New("Unable to compare equal values")
	}
}
