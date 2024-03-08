package main

import "fmt"

func main() {

	xi := []int{23, 45, 245, 24, 2342, 35, 6, 6, 7}

	fmt.Println(xi)
	fmt.Println("---------------------")

	fmt.Println(xi[2:5])
	fmt.Println("---------------------")

	fmt.Println(xi[:6])
	fmt.Println("---------------------")

	fmt.Println(xi[2:])
	fmt.Println("---------------------")
	fmt.Println(xi[:])
	fmt.Println("---------------------")

	fmt.Printf("%#v",xi)

}
