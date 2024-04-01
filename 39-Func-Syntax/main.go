package main

import "fmt"

func main() {
	foo() //arguments
	foo1("selviin")
	fmt.Println(foo2("delllll"))
	fmt.Println(foo3("selvin", 22))
}

//func (receiver) identifier (p parameter(s)) (return(s)){...code...}

// No parameter no returns
func foo() { //parameters
	fmt.Println("i am foo")
}

// 1 parameter no returns
func foo1(s string) {
	fmt.Println("My name is", s)
}

// 1 parameter 1 returns
func foo2(s string) string {
	return s
}

// 2 parameter 2 returns
func foo3(name string, age int) (string, int) {
	return name, age
}
