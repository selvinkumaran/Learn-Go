package main

import "fmt"

func main() {

	x := doMath(42, 16, add)
	y := doMath(42, 16, sub)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", add)
	fmt.Printf("%T\n", sub)
	fmt.Printf("%T\n", doMath)
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func add(a int, b int) int {
	return a + b
}

func sub(a int, b int) int {
	return a - b
}
