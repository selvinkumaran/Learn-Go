package main

import (
	"fmt"
	"sort"
)

func main() {

	xi := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(xi)

	fmt.Println("**************************")

	// xi = append(xi[:4], xi...)

	fmt.Println(xi)

	fmt.Println("**************************")

	xi = append(xi[:4], xi[5:]...)

	fmt.Println(xi, "KNDIJDIJ")

	//sorting a slice
	sort.Ints(xi)

}
