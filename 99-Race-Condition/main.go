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
	const numGoroutines = 100
	var ws sync.WaitGroup
	ws.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			v := counter
			// time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter = v
			ws.Done()
			fmt.Println("Gorountine:", runtime.NumGoroutine())
		}()
	}
	ws.Wait()
	fmt.Println("COUNT:", counter)
	fmt.Println("Gorountine:", runtime.NumGoroutine())
}

// Another example
/*
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

var count = 0

func main() {
	wg.Add(1001)
	for i := 0; i <= 1000; i++ {
		go func() {
			mu.Lock()
			count++
			fmt.Println("count", count)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("Finish")
}

*/
