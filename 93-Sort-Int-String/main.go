package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{3, 6, 8, 9, 56, 2, 1, 456, 7, 98, 23}

	xs := []string{"f", "e", "s", "w", "a"}

	fmt.Println(xi)
	fmt.Println(xs)

	fmt.Println("******************************")
	sort.Ints(xi)
	sort.Strings(xs)

	fmt.Println(xi)
	fmt.Println(xs)
}
