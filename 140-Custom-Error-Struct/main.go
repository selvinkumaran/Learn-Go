package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("Here is the error: %v", ce.info)
}

func main() {
	c1 := customErr{
		info: "need more errors to learn",
	}

	fmt.Println(c1.Error())

	foo(c1)
}

func foo(e error) {
	fmt.Println("foo ran -", e)
	fmt.Println("foo ran -", e.(customErr).info) //type assertion
}
