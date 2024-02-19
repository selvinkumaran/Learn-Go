package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Hello selvin")
	}()

	func(s string) {
		fmt.Println("Hi This is", s)
	}("selvin")
}

func foo() {
	fmt.Println("foo")
}

// a named function with an identifier
// func (r reciever) identifier(p p arameter(s)) (r return(s)) { code }

// an anonymous function
// func (p parameter(s)) (r return(s)) { code }
