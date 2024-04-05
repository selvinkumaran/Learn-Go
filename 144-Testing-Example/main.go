package main

func main() {
}

func AddSum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}