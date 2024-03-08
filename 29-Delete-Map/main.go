package main

import "fmt"

func main() {

	m := map[string][]string{
		"james":  {"cricket", "kabddi"},
		"philip": {"volleyball", "tennis", "kabddi"},
		"sam":    {"football", "kabddi"},
	}

	m["jabes"] = []string{"kabaddi", "batmiton"}
	m["vinoth"] = []string{"throwbal", "batmiton"}

	fmt.Println("*********************")

	for k, v := range m {
		fmt.Println(k, v)
	}

	delete(m, "jabes")
	
	fmt.Println("********AFTER DELETING*************")

	for k, v := range m {
		fmt.Println(k, v)
	}
}
