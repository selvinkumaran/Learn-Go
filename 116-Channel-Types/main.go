package main

import "fmt"

func main() {
	bch := make(chan int)   // Bi directional channel
	rch := make(chan<- int) // Receive only channel
	sch := make(<-chan int) // Send only channel

	fmt.Printf("%T\n", bch)
	fmt.Printf("%T\n", rch)
	fmt.Printf("%T\n", sch)

	//Specific to general deosn't work
	// bch = rch
	// bch = sch

	//General to Specific works
	// rch = bch
	// sch = bch

	//Specific to Specific deosn't work
	// rch = sch

}
