package main

import "fmt"

func main() {

	xs := []string{
		"the", "is", "and", "the", "is", "and", "the", "is", "and", "the", "is", "and", "the", "is", "and", "the", "is", "and", "after", "before", "before",
	}

	m := map[string]int{}

	for _, v := range xs {
		m[v]++
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
