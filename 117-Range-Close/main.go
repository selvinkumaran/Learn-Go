package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i <= 10; i++ {
			c <- i
		}
		// If you use range loop we have to close
		close(c) // we must close channel else deadlock will occur
	}()

	for v := range c {
		fmt.Println(v)
	}
	fmt.Println("About to close")

}
