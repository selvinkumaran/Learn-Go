package main

import "fmt"

type hotdog int

func main() {

	var y int

	var x hotdog = 42

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	y = int(x) //conversion

	fmt.Println(y)
	fmt.Printf("%T", y)
}

// fmt.Println("foo ran -", e, "\n", e.(customErr).info)
