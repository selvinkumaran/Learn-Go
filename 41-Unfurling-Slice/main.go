package main

import "fmt"

func main() {

	xi := []int{1, 2, 3, 4, 5}
	result := sum(xi...)
	fmt.Println("The sum is ", result)
}

//The variatic parameter have to come last
func sum(nums ...int) int {
	fmt.Println(nums)
	fmt.Printf("%T\n", nums)
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}
