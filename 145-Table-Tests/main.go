package main

import "fmt"

func main() {
	fmt.Println(AddSum(1, 2, 3, 2))
}

func AddSum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
