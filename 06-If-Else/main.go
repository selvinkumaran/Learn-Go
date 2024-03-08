package main

import (
	"fmt"
	"math/rand"
)

func main() {

	fmt.Println("if statement")

	x := rand.Intn(30)
	if x > 20 {
		fmt.Printf("%v is greater than 20", x)
	} else {
		fmt.Printf("%v is lesser than 20", x)
	}

	fmt.Println()

	y := 20
	if y > 20 {
		fmt.Printf("%v is greater than 20", y)
	} else if y == 20 {
		fmt.Printf("%v is equals to 20", y)
	} else {
		fmt.Printf("%v is lesser than 20", y)
	}
}
