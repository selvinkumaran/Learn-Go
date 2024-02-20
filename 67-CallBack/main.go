package main

import "fmt"

func main() {

	fmt.Println(printSquare(square, 3))

}

func printSquare(f func(int) int, a int) string {
	x := square(a)
	return fmt.Sprintf("the value is %v and the squared value vale is %v", a, x)
}

func square(n int) int {
	return n * n
}
