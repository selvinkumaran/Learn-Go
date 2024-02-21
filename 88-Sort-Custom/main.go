package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

// // By Age
// type ByAge []person

// func (x ByAge) Len() int           { return len(x) }
// func (x ByAge) Less(i, j int) bool { return x[i].Age < x[j].Age }
// func (x ByAge) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

// func main() {

// 	p4 := person{"nandha", 33}
// 	p1 := person{"selvin", 232}
// 	p2 := person{"narkunan", 22}
// 	p3 := person{"venket", 23}

// 	people := []person{p1, p2, p3, p4}
// 	fmt.Println(people)
// 	sort.Sort(ByAge(people))
// 	fmt.Println(people)
// }

// By Name
type ByName []person

func (x ByName) Len() int           { return len(x) }
func (x ByName) Less(i, j int) bool { return x[i].Name < x[j].Name }
func (x ByName) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {

	p4 := person{"nandha", 33}
	p1 := person{"selvin", 232}
	p2 := person{"narkunan", 22}
	p3 := person{"venket", 23}

	people := []person{p1, p2, p3, p4}
	fmt.Println(people)
	sort.Sort(ByName(people))
	fmt.Println(people)

}
