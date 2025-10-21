package basics

//During build process,tree shaking statically analyzes the code base to determine which imports are actually used in the code.
import (
	"fmt"
	foo "net/http" //foo is named import
)

func main() {
	fmt.Println("Hello, Go Standard Library!")
	resp, err := foo.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println("Error fetching URL:", err)
	}
	defer resp.Body.Close()
	fmt.Println("HTTP Response Status:", resp.Status)
}

//Pointers, Slices,maps,functions and structs are initialized to nil by default
