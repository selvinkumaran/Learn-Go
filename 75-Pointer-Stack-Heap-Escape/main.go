package main

import "fmt"

func main() {

	a := 1
	// value semantics
	fmt.Println(a)
	// pointer semantics
	fmt.Println(&a)
}

// This is a commang to see escape
// go run -gcflags -m main.go
