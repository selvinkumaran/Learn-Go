package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 23
	fmt.Println(ch)
}
