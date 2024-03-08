package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Print("I'm init function")
}

func main() {
	fmt.Println("Switch Statement")

	x := rand.Intn(40)

	switch {
	case x > 20:
		fmt.Printf("%v is greater than 20", x)
	case x < 20:
		fmt.Printf("%v is lesser than 20", x)
	case x == 20:
		fmt.Printf("%v is equal to 20", x)
	default:
		fmt.Print("default")
	}
}
