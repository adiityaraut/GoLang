package basics

import "fmt"

func main() {

	// For Loop as While Loop
	var i int = 1
	for i <= 5 {
		fmt.Println("Count:", i)
		i++
	}
	//for as while with break
	sum := 0
	for {
		if sum >= 20 {
			break
		}
		fmt.Println("Sum:", sum)
		sum += 3
	}
}
