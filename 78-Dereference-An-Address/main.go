package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	p := "Drop by drop,the bucket gets filled."
	q := "Persistenly, patiently, you are bound to succeed."
	r := "The meaning of the life ..."
	s := 22
	a = &p
	b = &q
	c = &r
	d = &s
}

func main() {

	fmt.Printf("%T \t %v\n", a, a)
	fmt.Printf("%T \t %v\n", b, b)
	fmt.Printf("%T \t %v\n", c, c)
	fmt.Printf("%T \t\t %v\n", d, d)

	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)

}
