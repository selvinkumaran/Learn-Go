package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx := context.Background()

	fmt.Println("context\t\t", ctx)
	fmt.Println("context Error:\t", ctx.Err())
	fmt.Printf("context Type:\t%T\n", ctx)
	fmt.Println("----------------------------")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context\t\t", ctx)
	fmt.Println("context Error:\t", ctx.Err())
	fmt.Printf("context Type:\t%T\n", ctx)
	fmt.Println("cancel\t\t", cancel)
	fmt.Printf("cancel Type:\t%T\n", cancel)
	fmt.Println("----------------------------")

	cancel()

	fmt.Println("context\t\t", ctx)
	fmt.Println("context Error:\t", ctx.Err())
	fmt.Printf("context Type:\t%T\n", ctx)
	fmt.Println("cancel\t\t", cancel)
	fmt.Printf("cancel Type:\t%T\n", cancel)
	fmt.Println("----------------------------")

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num gortins 1:", runtime.NumGoroutine())

	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num gortins 2:", runtime.NumGoroutine())

	fmt.Println("about to cancel context")
	cancel()
	fmt.Println("cancelled context")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 3:", ctx.Err())
	fmt.Println("num gortins 3:", runtime.NumGoroutine())

}
