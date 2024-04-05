package main

import "fmt"

func main() {
	s := check()
	fmt.Println(s)
}

func check() string {
	i := 0
	defer fmt.Println(i)
	defer fmt.Println("3")
	i++
	return "selvin"
}
