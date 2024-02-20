package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := foo(xi...)
	fmt.Println(sum)
	res := bar(xi)
	fmt.Println(res)

}

func foo(vi ...int) int {
	sum := 0
	for _, v := range vi {
		sum += v
	}
	return sum
}

func bar(vi []int) int {
	sum := 0
	for _, v := range vi {
		sum += v
	}
	return sum
}
