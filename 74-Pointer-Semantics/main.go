package main

import "fmt"

func main() {

	// pointer semantics
	fmt.Println("Pointer semantics")
	b := 1
	fmt.Println(b)
	addOneP(&b)
	fmt.Println(b)

}

// pointer semantics
func addOneP(a *int) {
	*a++
}
