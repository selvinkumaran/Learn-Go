package main

import "fmt"

func main() {

	//If not passing any arguments also it will work
	// result := sum() this also will work
	result := sum(1, 2, 3, 4, 5)
	fmt.Println("The sum is ", result)
}

// The variatic parameter have to come last
func sum(nums ...int /*unlimited type*/) int {
	fmt.Println(nums)
	fmt.Printf("%T\n", nums)
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}
