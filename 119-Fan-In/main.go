package main

import (
	"fmt"
	"sync"
)

func main() {
	eve := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(eve, odd)
	go receive(eve, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}
	fmt.Println("About to close")

}

// Send Channel
func send(eve, odd chan<- int) {

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			eve <- i
		} else {
			odd <- i
		}
	}
	close(eve)
	close(odd)
}

// Receive Channel
func receive(eve, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range eve {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v := range odd {
			fanin <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fanin)
}
