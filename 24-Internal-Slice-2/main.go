package main

import (
	"fmt"
	"sort"
)

func main() {

	a := []int{1979, 2, 3, 5, 6, 8, 9}

	// b := []int{66}

	// b:=make([]int ,len(a))

	// copy(b, a)

	// // b=a

	// fmt.Println("a is", a)
	// fmt.Println("b is", b)

	// fmt.Println("***********************")

	// a[0] = 2222

	// fmt.Println("a is", a)
	// fmt.Println("b is", b)

	// x := make([]int, len(a))

	// copy(x, a)

	// sort.Ints(a)
	// sort.Ints(x)

	// fmt.Println(x, "x value")

	b := []int{1979, 2, 3, 5, 6, 8, 9}

	fmt.Println("*******************")

	fmt.Println(a, "a value")

	fmt.Println("*******************")

	fmt.Println(sample(a))
	fmt.Println(a, "a fun 1")

	fmt.Println(sample2(b))
	fmt.Println(b, "a fun 2")
}

func sample(e []int) int {
	sort.Ints(e)
	return 2
}

func sample2(e []int) int {
	x := make([]int, len(e))
	copy(x, e)
	sort.Ints(x)
	return 2
}
