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

// send
func bar(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
}

// receive
func foo(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
	}
}
