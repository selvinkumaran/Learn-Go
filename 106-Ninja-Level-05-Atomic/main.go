package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var wg sync.WaitGroup

	var increment int64

	gs := 100

	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&increment, 1)
			fmt.Println(atomic.LoadInt64(&increment))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value", increment)

}
