package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("GOrountine:", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var ws sync.WaitGroup
	var mu sync.Mutex
	ws.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			ws.Done()
			mu.Unlock()
			fmt.Println("Gorountine:", runtime.NumGoroutine())
		}()
	}
	ws.Wait()
	fmt.Println("COUNT:", counter)
	fmt.Println("Gorountine:", runtime.NumGoroutine())
}
