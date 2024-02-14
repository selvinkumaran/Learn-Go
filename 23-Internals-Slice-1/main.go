package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 5, 6, 8, 9}

	b := a

	fmt.Println("a is", a)
	fmt.Println("b is", b)

	fmt.Println("***********************")

	a[0] = 2222

	fmt.Println("a is", a)
	fmt.Println("b is", b)
}
