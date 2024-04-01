package main

import "fmt"

func main() {
	//Send Only Channel
	ch := make(<-chan int, 2)

	// Un Comment to see error
	// ch <- 23
	// ch <- 24

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("--------------------")
	fmt.Printf("%T\n", ch)
}
