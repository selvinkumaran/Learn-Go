package main

import "fmt"

func main() {
	//Bi Directional Channel
	ch := make(chan int, 2)

	ch <- 23
	ch <- 24

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println("--------------------")
	fmt.Printf("%T\n", ch)
}
