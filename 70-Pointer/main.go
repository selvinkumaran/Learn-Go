package main

import "fmt"

func main() {
	x := 43

	fmt.Println(x)
	fmt.Println(&x)
	// type and value for pointers
	fmt.Printf("%T and %v\n", &x, &x)
	s := "selvin"
	fmt.Printf("%T and %v\n", &s, &s)
}
