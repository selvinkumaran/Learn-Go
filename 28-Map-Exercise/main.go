package main

import "fmt"

func main() {

	m := map[string][]string{
		"james":  {"cricket", "kabddi"},
		"philip": {"volleyball", "tennis", "kabddi"},
		"sam":    {"football", "kabddi"},
	}
	fmt.Println(m)

	fmt.Println(m["sam"])

	fmt.Println("*********************")

	for k, v := range m {
		fmt.Println(k, v)
	}

	m["jabes"] = []string{"kabaddi", "batmiton"}
	m["vinoth"] = []string{"throwbal", "batmiton"}

	fmt.Println("new added elements")

	for k, v := range m {
		fmt.Println(k, v)
	}
}
