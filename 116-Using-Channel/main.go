package main

import "fmt"

func main() {

	c := make(chan int)

	go send(c)
	receive(c)
	fmt.Println("About To Exit")
}

func send(c chan<- int) {
	c <- 22
}
func receive(c <-chan int) {
	fmt.Println(<-c)
}
