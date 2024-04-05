// Package Example_test ask you to ready to rock
package Example_Test

// AddSum  add unlimited number of values of type int
func AddSums(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}
