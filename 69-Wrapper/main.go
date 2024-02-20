package main

import (
	"fmt"
	"time"
)

func main() {
	timeFunc(doWork)
}

func timeFunc(f func()) {

	start := time.Now()
	fmt.Println(start, "start")
	f()
	elapsed := time.Since(start)
	fmt.Println(elapsed, "end")
}

func doWork() {
	for i := 0; i < 20000; i++ {
		fmt.Print(i, " ")
	}
}
