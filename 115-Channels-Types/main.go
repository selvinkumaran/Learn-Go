package main

import "fmt"

func main() {
	ch := make(chan int)
	chSend := make(chan<- int)    //send
	chReceive := make(<-chan int) //receive

	fmt.Printf("ch \t\t %T\n", ch)
	fmt.Printf("ch_send \t %T\n", chSend)
	fmt.Printf("ch_receive \t %T\n", chReceive)

	// //specific to generic doesn't assign
	// ch = ch_send
	// ch = ch_receive

	//generic to specific assigns
	chSend = ch
	chReceive = ch

}
