package main

import "fmt"

func main() {

	xi := []int{23, 45, 66, 67}

	fmt.Println(xi)
	fmt.Println("***************************")
	xi = append(xi, 23, 23, 45, 24, 54)

	fmt.Println(xi)

	fmt.Println("***************************")

	xs := []string{"ssd", "sdijid", "sjndj"}

	fmt.Println(xs)
	fmt.Println("***************************")
	xs = append(xs, "sdni")

	fmt.Println(xs)

}
