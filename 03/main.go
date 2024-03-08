package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	fmt.Println("start", runtime.NumGoroutine())
	go func() {
		fmt.Println("sum 1")
		wg.Done()
	}()
	go func() {
		fmt.Println("sum 2")
		wg.Done()
	}()
	go func() {
		fmt.Println("sum 3")
		wg.Done()
	}()
	go func() {
		fmt.Println("sum 4")
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("end", runtime.NumGoroutine())
}
