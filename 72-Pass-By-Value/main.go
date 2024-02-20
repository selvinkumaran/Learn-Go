package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	intDelta(&a)
	fmt.Println(a)

	xi := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(xi)
	sliceDelta(xi)
	fmt.Println(xi)

	m := make(map[string]int)
	m["selvin"] = 22
	fmt.Println(m["selvin"])
	mapDelta(m, "selvin")
	fmt.Println(m["selvin"])
}

func intDelta(n *int) {
	*n = 43
}

func sliceDelta(xi []int) {
	xi[0] = 99
}
func mapDelta(x map[string]int, s string) {
	x[s] = 43
}
