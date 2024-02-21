package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{4, 24, 5, 6, 8, 8, 5, 32, 2, 4, 5, 7, 7}
	xs := []string{"zdjw", "ieie", "wieiie"}

	fmt.Println("Before Sorting")
	fmt.Println(xi)
	fmt.Println(xs)

	sort.Ints(xi)
	sort.Strings(xs)
	
	fmt.Println("After Sorting")
	fmt.Println(xi)
	fmt.Println(xs)
}
