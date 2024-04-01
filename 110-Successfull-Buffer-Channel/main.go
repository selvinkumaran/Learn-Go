package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 23

	fmt.Println(<-ch)
}
