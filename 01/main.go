package main

import (
	"fmt"
)

var x=20
const y=30

func main() {
	z:=40
	fmt.Printf("The value of x is %v \nThe type of x is %T \n",x,x)
	fmt.Printf("The value of y is %v \nThe type of y is %T \n",y,y)
	fmt.Printf("The value of z is %v \nThe type of z is %T",z,z)
}
