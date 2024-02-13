package main

import (
	"fmt"
)

func main() {

	i := 0

	for i := 0; i <= 20; i++ {
		if i%2 == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	for i <= 20 {
		fmt.Print(i, " ")
		i++
	}
	//Infinite Loop
	for {
		fmt.Println("selvin")
	}
}
