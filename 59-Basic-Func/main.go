package main

import "fmt"

func main() {

	i := foo()
	x, y := bar()
	fmt.Println(i)
	fmt.Println(x, y)
}

func foo() int {
	return 45
}

func bar() (int, string) {
	return 3, "selvin"
}
