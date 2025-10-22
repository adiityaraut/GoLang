package basics

import "fmt"

func main() {
	message := "Hello World"
	//Range iterates over each index and value of the string in an unspecified order, channel till it is closed, slice or map
	for i, v := range message {
		fmt.Println(i, v)
		fmt.Printf("Index: %d, Character: %c\n", i, v)
	}
}
