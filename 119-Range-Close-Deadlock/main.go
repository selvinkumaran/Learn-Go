package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	//send
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch) // Close the channel after sending all values
	}()

	//receive
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("About to exit")
}

//The issue in the first program arises because the main() function is waiting indefinitely for values to be sent on the channel ch. Since the sender goroutine is also in main() and runs concurrently, the main() function might reach the end before the sender goroutine can send all values. This results in a deadlock.
