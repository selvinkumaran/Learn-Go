package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GOrountine:", runtime.NumGoroutine())

	var counter int64
	const gs = 100
	var ws sync.WaitGroup
	ws.Add(gs)	
	for i := 0; i < gs; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter \t", atomic.LoadInt64(&counter))
			ws.Done()
		}()
		fmt.Println("Gorountine:", runtime.NumGoroutine())
	}
	ws.Wait()
	fmt.Println("COUNT:", counter)
	fmt.Println("Gorountine:", runtime.NumGoroutine())
}
