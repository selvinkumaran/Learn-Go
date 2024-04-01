package main

import "fmt"

func main() {

	oddd := make(chan int)
	even := make(chan int)
	quit := make(chan int)

	//send
	go send(oddd, even, quit)

	//receive
	receive(oddd, even, quit)

	fmt.Println("About to close")
}

func send(oddd, even, quit chan<- int) {

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			oddd <- i
		}
	}
	close(even)
	close(oddd)
	quit <- 0
}

func receive(oddd, even, quit <-chan int) {
	for {
		select {
		case v := <-even:
			fmt.Println("From the even channel", v)
		case v := <-oddd:
			fmt.Println("From the oddd channel", v)
		case v := <-quit:
			fmt.Println("From the quit  channel", v)
			return
		}
	}
}
