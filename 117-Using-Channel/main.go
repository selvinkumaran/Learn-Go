package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	//send
	go foo(ch)

	//receive
	bar(ch)

	fmt.Println("About to exit")
}

func bar(ch chan<- int) {
	ch <- 43
}

func foo(ch <-chan int) {
	fmt.Println(<-ch)
}
