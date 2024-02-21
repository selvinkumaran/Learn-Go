package main

import "fmt"

func addI(a, b int) int {
	return a + b
}

func addF(a, b float64) float64 {
	return a + b
}

type myNumbers interface {
	~int | ~float64
}

func addT[T myNumbers](a, b T) T {
	return a + b
}

type myAlias int

func main() {

	var x myAlias = 45
	fmt.Println(addI(2, 2))
	fmt.Println(addF(2.2, 2.2))

	fmt.Println(addT(x, 2))
	fmt.Println(addT(2.2, 2.2))
}
