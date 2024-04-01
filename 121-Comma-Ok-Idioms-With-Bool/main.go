package main

import "fmt"

func main() {

	oddd := make(chan int)
	even := make(chan int)
	quit := make(chan bool)

	//send
	go send(oddd, even, quit)

	//receive
	receive(oddd, even, quit)

	fmt.Println("About to close")
}

func send(oddd, even chan<- int, quit chan<- bool) {

	for i := 0; i < 5; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			oddd <- i
		}
	}
	close(quit)
}

func receive(oddd, even <-chan int, quit <-chan bool) {
	for {
		select {
		case v := <-even:
			fmt.Println("From the even channel", v)
		case v := <-oddd:
			fmt.Println("From the oddd channel", v)
		case i, ok := <-quit:
			if !ok {
				fmt.Println("From comma ok", i, ok)
			} else {
				fmt.Println("From comma ok", i)
			}
			return
		}
	}
}
