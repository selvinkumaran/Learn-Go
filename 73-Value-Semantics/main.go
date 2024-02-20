package main

import "fmt"

func main() {

	// value semantics
	fmt.Println("value semantics")
	a := 1
	fmt.Println(a)         //1
	fmt.Println(addOne(a)) //2
	fmt.Println(a)         //1

}

// value semantics
func addOne(a int) int {
	return a + 1
}
