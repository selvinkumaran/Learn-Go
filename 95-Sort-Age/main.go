package main

import (
	"fmt"
	"sort"
)

type user struct {
	Name string
	Age  int
}

type ByAge []user

func (x ByAge) Len() int           { return len(x) }
func (x ByAge) Less(i, j int) bool { return x[i].Age < x[j].Age }
func (x ByAge) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {

	u1 := user{
		"selvin", 22,
	}
	u2 := user{
		"mani", 32,
	}
	u3 := user{
		"sharuk", 42,
	}
	u4 := user{
		"nagaraj", 12,
	}

	user := []user{u1, u2, u3, u4}
	fmt.Println(user)
	sort.Sort(ByAge(user))
	fmt.Println(user)

}
