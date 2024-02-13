package main

import (
	"fmt"
	"math/rand"
)

func main() {
	x := rand.Intn(20)
	y := rand.Intn(20)
	fmt.Println("x is", x)
	fmt.Println("y is", y)
	if x < 4 && y < 4 {
		fmt.Println("Both are lesser than 4")
	} else if x > 6 && y > 6 {
		fmt.Println("Both are greater than 6")
	} else if x >= 4 && x <= 6 {
		fmt.Println("x is from 4 to 6 inclusive of both numbers")
	} else if y != 5 {
		fmt.Println("y is not 5")
	} else {
		fmt.Println("none of the previous were met")
	}
}
