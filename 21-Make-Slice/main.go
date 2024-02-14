package main

import (
	"fmt"
	"sort"
)

func main() {

	// xi := []int{1, 2, 3, 4, 5, 6, 7, 8}

	// fmt.Println(xi)

	xi := make([]int, 0, 10)
	fmt.Println(xi)
	fmt.Println("length", len(xi))
	fmt.Println("capacity", cap(xi))

	xi = append(xi, 1, 2, 3, 4, 5, 6, 7, 8)
	fmt.Println(xi)

	fmt.Println("*************************")

	xi = append(xi, 9, 10, 11, 12, 13)
	fmt.Println(xi)
	fmt.Println("length", len(xi))
	fmt.Println("capacity", cap(xi))

	fmt.Println("*************************")

	xi = append(xi, 14, 15, 16, 17, 18, 19, 20, 21)
	fmt.Println(xi)
	fmt.Println("length", len(xi))
	fmt.Println("capacity", cap(xi))

	fmt.Println("*************************")

	xi = append(xi, 14, 15, 16, 17, 18, 19, 20, 21)
	xi = append(xi, 14, 15, 16, 17, 18, 19, 20, 21)
	xi = append(xi, 14, 15, 16, 17, 18, 19, 20, 21)
	fmt.Println(xi, "BEFORE")
	sort.Ints(xi)
	fmt.Println(xi, "AFTER")
	fmt.Println("length", len(xi))
	fmt.Println("capacity", cap(xi))

}
