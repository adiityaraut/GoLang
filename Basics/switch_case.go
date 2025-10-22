package basics

import "fmt"

func main() {

	// 	switch expression { // (switch case default but no break needed)
	// 	case value1:
	// 		// code to be executed if expression equals value1
	// 	case value2:
	// 		// code to be executed if expression equals value2
	// 	default:
	// 		// code to be executed if expression doesn't match any case
	// 	}
	fruit := "apple"

	switch fruit {
	case "apple":
		fmt.Println("This is an apple.")
	case "banana":
		fmt.Println("This is a banana.")
	default:
		fmt.Println("Unknown fruit.")
	}

	day := "Monday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Start of the work week.")
	default:
		fmt.Println("It's a weekday.")
	}

	//fallthrough is used to continue execution to the next case
	number := 2
	switch {
	case number > 1:
		fmt.Println("Greater than One")
		fallthrough
	case number == 2:
		fmt.Println("Two")
	default:
		fmt.Println("Other number")
	}
	checkType(10)
	checkType("Hello")
	checkType(true)
	checkType(3.14)
}

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Println("Unknown type")
	}
}
