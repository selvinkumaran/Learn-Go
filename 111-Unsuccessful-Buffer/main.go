package main

import "fmt"

func main() {
	ch := make(chan int, 1)

	ch <- 42
	ch <- 43

	fmt.Println(<-ch)
}
