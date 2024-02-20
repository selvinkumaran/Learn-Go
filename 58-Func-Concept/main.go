package main

import "fmt"

func main() {
	sum := sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9})
	fmt.Println(sum)
}

func sum(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

// Named return
/*func sum(xi []int) (total int) {
	total = 0
	for _, v := range xi {
		total += v
	}
	return
}
*/
