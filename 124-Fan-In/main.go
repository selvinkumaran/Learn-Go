package main

import (
	"fmt"
	"sync"
)

// Fanin
// Taking values from many channels and putting those values onto one channel

func main() {

	oddd := make(chan int)
	even := make(chan int)
	fanin := make(chan int)

	//send
	go send(oddd, even)

	//receive
	go receive(oddd, even, fanin)

	for v := range fanin {
		fmt.Println(v)
	}

	fmt.Println("About to close")
}

func send(oddd, even chan<- int) {

	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			oddd <- i
		}
	}
	close(even)
	close(oddd)
}

func receive(oddd, even <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range oddd {
			fanin <- v
		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
