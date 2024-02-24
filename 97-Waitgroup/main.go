package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\n", runtime.GOOS)
	fmt.Println("ARCH\n", runtime.GOARCH)
	fmt.Println("CPUs\n", runtime.NumCPU())
	fmt.Println("Goroutine\n", runtime.NumGoroutine())

	wg.Add(1)
	go foo()
	bar()
	fmt.Println("CPUs\n", runtime.NumCPU())
	fmt.Println("Goroutine\n", runtime.NumGoroutine())
	wg.Wait()
	// time.Sleep(2000)

}

func foo() {
	for i := 0; i <= 10; i++ {
		fmt.Println("foo", i)
	}
	wg.Done()
}
func bar() {
	for i := 0; i <= 10; i++ {
		fmt.Println("bar", i)
	}
}
