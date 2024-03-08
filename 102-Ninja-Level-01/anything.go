package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Start CPU", runtime.NumCPU())
	fmt.Println("Start Goroutine", runtime.NumGoroutine())

	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		fmt.Println("I am fron thing one")
		wg.Done()
	}()

	go func() {
		fmt.Println("I am fron thing two")
		wg.Done()
	}()

	fmt.Println("Mid CPU", runtime.NumCPU())
	fmt.Println("Mid Goroutine", runtime.NumGoroutine())

	wg.Wait()

	fmt.Println("End CPU", runtime.NumCPU())
	fmt.Println("End Goroutine", runtime.NumGoroutine())

	fmt.Println("About to exit")

}
